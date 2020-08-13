package ccm

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/cloud-provider"
	"time"

	clb "github.com/capitalonline/cloud-controller-manager/pkg/clb/api"
)

const (
	// defaultActiveTimeout is the number of seconds to wait for a load balancer to
	// reach the active state.
	defaultActiveTimeout = 90

	// defaultActiveCheckTick is the number of seconds between load balancer
	// status checks when waiting for activation.
	defaultActiveCheckTick = 5
)

var (
	ClbLoadBalancerKindClassic = 0
	//ClbLoadBalancerKindApplication = 1
)

type loadBalancers struct {
	resources         *resources
	region            string
	clusterID         string
	lbActiveTimeout   int
	lbActiveCheckTick int
}

// newLoadbalancers returns a cloudprovider.LoadBalancer whose concrete type is a *loadbalancer.
func newLoadBalancers(resources *resources, region string) cloudprovider.LoadBalancer {
	return &loadBalancers{
		resources:         resources,
		region:            region,
		lbActiveTimeout:   defaultActiveTimeout,
		lbActiveCheckTick: defaultActiveCheckTick,
	}
}

func (l *loadBalancers) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
	clusterID := l.resources.clusterID
	log.Infof("GetLoadBalancer:: clusterID is: %s, service.name is: %s", clusterID, service.ObjectMeta.Name)

	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	log.Infof("GetLoadBalancer:: serviceName is: %s, serviceNameSpace is: %s, serviceUid is: %s", serviceName, serviceNameSpace, serviceUid)

	loadBalancer, err := getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid)
	if err != nil {
		if err == clb.ErrCloudLoadBalancerNotFound {
			log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer  is not exist")
			SentrySendError(fmt.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer  is not exist"))
			return nil, false, nil
		}
		log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
		SentrySendError(fmt.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err))
		return nil, false, err
	}
	log.Infof("GetLoadBalancer:: cloud.getLoadBalancerByName, res is: %+v", loadBalancer)

	ingresses := make([]v1.LoadBalancerIngress, len(loadBalancer.Data.Vips))

	for i, vip := range loadBalancer.Data.Vips {
		ingresses[i] = v1.LoadBalancerIngress{IP: vip}
	}

	log.Infof("GetLoadBalancer:: successfully, ingresses are: %s", ingresses)
	return &v1.LoadBalancerStatus{
		Ingress: ingresses,
	}, true, nil
}

func (l *loadBalancers) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	clusterID := l.resources.clusterID
	log.Infof("GetLoadBalancerName:: clusterID is: %s, service.name is: %+v", clusterID, service.ObjectMeta.Name)

	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	log.Infof("GetLoadBalancer:: serviceName is: %s, serviceNameSpace is: %s, serviceUid is: %s", serviceName, serviceNameSpace, serviceUid)

	res, err := getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid)

	if err != nil {
		log.Infof("GetLoadBalancerName:: getLoadBalancerByName succeed, return loadBalancerName is: %s", res.Data.Name)
		return res.Data.Name
	}
	log.Errorf("GetLoadBalancerName:: getLoadBalancerByName is error, err is: %s", err)
	SentrySendError(fmt.Errorf("GetLoadBalancerName:: getLoadBalancerByName is error, err is: %s", err))
	return err.Error()
}

func (l *loadBalancers) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	clusterID := l.resources.clusterID
	log.Infof("EnsureLoadBalancer:: clusterID is: %s, service.name is: %s", clusterID, service.ObjectMeta.Name)
	log.Infof("EnsureLoadBalancer:: nodes totally num is: %d", len(nodes))

	for _, node := range nodes {
		log.Infof("EnsureLoadBalancer:: node.name is: %s", node.ObjectMeta.Name)
	}
	if service.Spec.SessionAffinity != v1.ServiceAffinityNone {
		log.Errorf("EnsureLoadBalancer:: SessionAffinity is not supported currently, only support 'None' type")
		SentrySendError(fmt.Errorf("EnsureLoadBalancer:: SessionAffinity is not supported currently, only support 'None' type"))
		return nil, errors.New("SessionAffinity is not supported currently, only support 'None' type")
	}

	// TODO check if kubernetes has already do validate
	loadBalancerExist := true
	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	log.Infof("GetLoadBalancer:: serviceName is: %s, serviceNameSpace is: %s, serviceUid is: %s", serviceName, serviceNameSpace, serviceUid)

	// step-1. get loadBalancer status
	loadBalancerGet, err := getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid)
	if err != nil {
		if err == clb.ErrCloudLoadBalancerNotFound {
			log.Infof("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is succeed, loadBalancer is not exist, then to create it")
			loadBalancerExist = false
		} else {
			log.Errorf("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is error, err is: %s", err)
			SentrySendError(fmt.Errorf("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is error, err is: %s", err))
			return nil, err
		}

	} else {
		log.Infof("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is succeed, loadBalancer is exist, res is: %+v", loadBalancerGet)
	}

	// step-2. create or update loadBalancer by loadBalancerExist flag
	if loadBalancerExist {
		// loadBalancer is exist, then update it
		switch 0 {
		// only support classic yet
		case ClbLoadBalancerKindClassic:
			err := updateClassicLoadBalancer(ctx, service, nodes, clusterID)
			if err != nil {
				log.Errorf("EnsureLoadBalancer:: step-2 cloud.updateClassicLoadBalancer is error, err is: %s", err)
				SentrySendError(fmt.Errorf("EnsureLoadBalancer:: step-2 cloud.updateClassicLoadBalancer is error, err is: %s", err))
				return nil, err
			}
			log.Infof("EnsureLoadBalancer:: step-2 cloud.updateClassicLoadBalancer succeed")
		default:
			log.Errorf("EnsureLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet")
			SentrySendError(fmt.Errorf("EnsureLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet"))
			return nil, errors.New("Unsupported loadbalancer kind, only support [classic] yet")
		}
	} else {
		// loadBalancer is not exist, then create it
		switch 0 {
		// only support classic yet
		case ClbLoadBalancerKindClassic:
			err := createClassicLoadBalancer(ctx, service, nodes, clusterID)
			if err != nil {
				log.Errorf("EnsureLoadBalancer:: step-2 cloud.createClassicLoadBalancer is error, err is: %s", err)
				SentrySendError(fmt.Errorf("EnsureLoadBalancer:: step-2 cloud.createClassicLoadBalancer is error, err is: %s", err))
				return nil, err
			}
			log.Infof("EnsureLoadBalancer:: step-2 cloud.createClassicLoadBalancer succeed")
		default:
			log.Infof("EnsureLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet")
			return nil, errors.New("Unsupported loadbalancer kind, only support [classic] yet")
		}
	}

	// step-3. verify loadBalancer create or update successfully
	loadBalancerVerify, err := getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid)
	if err != nil {
		log.Errorf("EnsureLoadBalancer:: step-3 cloud.getLoadBalancerByName is error, err is: %s", err)
		SentrySendError(fmt.Errorf("EnsureLoadBalancer:: step-3 cloud.getLoadBalancerByName is error, err is: %s", err))
		return nil, err
	}
	log.Infof("EnsureLoadBalancer:: step-3 cloud.getLoadBalancerByName, res is: %+v", loadBalancerVerify)

	ingresses := make([]v1.LoadBalancerIngress, len(loadBalancerVerify.Data.Vips))
	for i, vip := range loadBalancerVerify.Data.Vips {
		ingresses[i] = v1.LoadBalancerIngress{IP: vip}
	}

	log.Infof("EnsureLoadBalancer:: successfully, ingresses are: %s", ingresses)
	return &v1.LoadBalancerStatus{
		Ingress: ingresses,
	}, nil
}

func (l *loadBalancers) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	clusterID := l.resources.clusterID
	log.Infof("UpdateLoadBalancer:: clusterID is: %s, service.name is: %+v", clusterID, service.ObjectMeta.Name)
	log.Infof("UpdateLoadBalancer:: nodes totally num is: %d", len(nodes))
	for _, node := range nodes {
		log.Infof("EnsureLoadBalancer:: node.name is: %s", node.ObjectMeta.Name)
	}

	switch 0 {
	// only support classic yet
	case ClbLoadBalancerKindClassic:
		err := updateClassicLoadBalancer(ctx, service, nodes, clusterID)
		if err != nil {
			log.Errorf("UpdateLoadBalancer:: cloud.updateClassicLoadBalancer is error, err is: %s", err)
			SentrySendError(fmt.Errorf("UpdateLoadBalancer:: cloud.updateClassicLoadBalancer is error, err is: %s", err))
			return err
		}
	default:
		log.Infof("UpdateLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet")
		return errors.New("Unsupported loadbalancer kind, only support [classic] yet")
	}
	log.Infof("UpdateLoadBalancer:: succeed!")
	return nil
}

func (l *loadBalancers) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	clusterID := l.resources.clusterID
	log.Infof("EnsureLoadBalancerDeleted:: clusterID is: %s, service.name is: %+v", clusterID, service.ObjectMeta.Name)

	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	log.Infof("GetLoadBalancer:: serviceName is: %s, serviceNameSpace is: %s, serviceUid is: %s", serviceName, serviceNameSpace, serviceUid)

	return deleteLoadBalancer(ctx, clusterID, serviceName, serviceNameSpace, serviceUid)
}

func getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid string) (*clb.DescribeLoadBalancersResponse, error) {
	// we don't need to check loadbalancer kind here because ensureLoadBalancerInstance will ensure the kind is right
	response, err := clb.DescribeLoadBalancers(&clb.DescribeLoadBalancersArgs{
		ClusterID:clusterID,
		ServiceName: serviceName,
		ServiceNameSpace: serviceNameSpace,
		ServiceUid: serviceUid,
	})

	// api with error
	if err != nil {
		return nil, err
	}

	// loadBalancer is not exist
	if len(response.Data.Vips) < 1 {
		return nil, clb.ErrCloudLoadBalancerNotFound
	}

	return response, nil
}

func updateClassicLoadBalancer(ctx context.Context, service *v1.Service, nodes []*v1.Node, clusterID string) error {
	// loadBalacer params
	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	// get service ports info
	var portMapSlice []clb.PortMapping
	var portMapTmp clb.PortMapping
	for _, ports := range service.Spec.Ports {
		portMapTmp.Port = ports.Port
		portMapTmp.NodePort = ports.NodePort
		// portMapTmp.Protocol = ports.Protocol
		// add to slice
		portMapSlice = append(portMapSlice, portMapTmp)
	}
	log.Infof("updateClassicLoadBalancer:: portMapSlice is: %+v", portMapSlice)

	// get nodes providerID info
	var nodeIdSlice []string
	for _, node := range nodes {
		nodeIdSlice = append(nodeIdSlice, node.Spec.ProviderID)
	}
	log.Infof("updateClassicLoadBalancer:: nodeIdSlice is: %s", nodeIdSlice)

	// get service annotations
	var annotationsSliceTmp []string
	if len(service.ObjectMeta.Annotations) != 0 {
		for key, value := range service.ObjectMeta.Annotations {
			stringTmp := key + ":" + value
			annotationsSliceTmp = append(annotationsSliceTmp, stringTmp)
		}

	} else {
		annotationsSliceTmp = append(annotationsSliceTmp, "")
	}
	log.Infof("updateClassicLoadBalancer:: annotationsSliceTmp is: %s", annotationsSliceTmp)

	// to create loadBalancer
	res, err := clb.UpdateLoadBalancers(&clb.UpdateLoadBalancersArgs{
		ClusterID:        clusterID,
		NodeID: nodeIdSlice,
		Annotations: annotationsSliceTmp,
		PortMap:     portMapSlice,
		ServiceName: serviceName,
		ServiceNameSpace: serviceNameSpace,
		ServiceUid: serviceUid,
	})

	if err != nil {
		return err
	}

	taskID := res.TaskID
	log.Infof("updateClassicLoadBalancer: create task succeed, TaskID is: %+v", taskID)

	// to check loadBalancer update task result
	err2 := describeLoadBalancersTaskResult(taskID)

	if err2 != nil {
		return err2
	}
	return nil
}

func createClassicLoadBalancer(ctx context.Context, service *v1.Service, nodes []*v1.Node, clusterID string) error {
	// loadBalacer params
	serviceName := service.ObjectMeta.Name
	serviceNameSpace := service.ObjectMeta.Namespace
	serviceUid := string(service.UID)

	// get services ports info
	var portMapSlice []clb.PortMapping
	var portMapTmp clb.PortMapping
	for _, ports := range service.Spec.Ports {
		portMapTmp.Port = ports.Port
		portMapTmp.NodePort = ports.NodePort
		// portMapTmp.Protocol = ports.Protocol
		// append to slice
		portMapSlice = append(portMapSlice, portMapTmp)
	}
	log.Infof("createClassicLoadBalancer:: portMapSlice is: %s", portMapSlice)

	// get nodes providerID info
	var nodeIdSlice []string
	for _, node := range nodes {
		nodeIdSlice = append(nodeIdSlice, node.Spec.ProviderID)
	}
	log.Infof("createClassicLoadBalancer:: nodeIdSlice is: %s", nodeIdSlice)

	// get service annotations
	var annotationsSliceTmp []string

	if len(service.ObjectMeta.Annotations) != 0 {
		for key, value := range service.ObjectMeta.Annotations {
			stringTmp := key + ":" + value
			annotationsSliceTmp = append(annotationsSliceTmp, stringTmp)
		}

	} else {
		annotationsSliceTmp = append(annotationsSliceTmp, "")
	}
	log.Infof("updateClassicLoadBalancer:: annotationsSliceTmp is: %s", annotationsSliceTmp)

	// to create loadBalancer
	res, err := clb.CreateLoadBalancers(&clb.CreateLoadBalancersArgs{
		ClusterID:        clusterID,
		NodeID:      nodeIdSlice,
		Annotations: annotationsSliceTmp,
		PortMap:     portMapSlice,
		ServiceName: serviceName,
		ServiceNameSpace: serviceNameSpace,
		ServiceUid: serviceUid,
	})

	if err != nil {
		return err
	}

	taskID := res.TaskID
	log.Infof("createClassicLoadBalancer: create task succeed, TaskID is: %+v", taskID)

	// to check loadBalancer create task result
	err2 := describeLoadBalancersTaskResult(taskID)

	if err2 != nil {
		return err2
	}
	return nil
}

func deleteLoadBalancer(ctx context.Context, clusterID, serviceName, serviceNameSpace, serviceUid string) error {
	// check the loadBalancer is exist or not
	res, err :=  getLoadBalancerByName(clusterID, serviceName, serviceNameSpace, serviceUid)
	if err != nil {
		if err == clb.ErrCloudLoadBalancerNotFound {
			log.Warnf("deleteLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer is not exist, so do not delete action, return nil")
			return nil
		}
		log.Errorf("deleteLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
		SentrySendError(fmt.Errorf("deleteLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err))
		return err
	}
	log.Infof("deleteLoadBalancer:: cloud.getLoadBalancerByName, res is: %+v, then to delete it", res)

	// loadBalancer is exist, then to delete it
	res2, err2 := clb.DeleteLoadBalancers(&clb.DeleteLoadBalancersArgs{
		ClusterID:        clusterID,
		ServiceName: serviceName,
		ServiceNameSpace: serviceNameSpace,
		ServiceUid: serviceUid,
	})

	if err2 != nil {
		return err2
	}

	taskID := res2.TaskID
	log.Infof("deleteLoadBalancer:: clb.DeleteLoadBalancers delete task_id is: %s", taskID)

	// to check loadBalancer delete task result
	err3 := describeLoadBalancersTaskResult(taskID)

	if err3 != nil {
		return err3
	}

	return nil
}

func describeLoadBalancersTaskResult(taskID string) error {
	for i := 1; i < 120; i++ {
		res, err := clb.DescribeLoadBalancersTaskResult(&clb.DescribeLoadBalancersTaskResultArgs{
			TaskID: taskID,
		})
		if err != nil {
			SentrySendError(fmt.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult is error, err is:%s", err))
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult is error, err is:%s", err)
		}
		if res.Data.Status == "doing" {
			log.Infof("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult is doing")
		} else if res.Data.Status == "finish" {
			log.Infof("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult succeed, status is: %s", res.Data.Status)
			return nil
		} else if res.Data.Status == "error" {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult error, status is: %s", res.Data.Status)
			SentrySendError(fmt.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult error, status is: %s", res.Data.Status))
			return errors.New("clb.DescribeLoadBalancersTaskResult error")
		} else {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult time out, running more than 10 minutes")
			SentrySendError(fmt.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult time out, running more than 10 minutes"))
			return errors.New("clb.DescribeLoadBalancersTaskResult time out, running more than 20 minutes")
		}

		log.Infof("DescribeLoadBalancersTaskResult:: time.sleep 30s")
		time.Sleep(time.Second * 10)
	}
	return nil
}
