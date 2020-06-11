package ccm

import (
	"context"
	"errors"
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
	ErrCloudLoadBalancerNotFound = errors.New("LoadBalancer not found")

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
	// testing code
	ingressesT := make([]v1.LoadBalancerIngress, 1)
	log.Infof("testing, return directly")
	return &v1.LoadBalancerStatus{
		Ingress: ingressesT,
	}, true, nil
	// business code
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	log.Infof("GetLoadBalancer:: clusterName is: %s, loadBalancerName is: %s", clusterName, loadBalancerName)

	loadBalancer, err := getLoadBalancerByName(clusterName, clusterID, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer  is not exist")
			return nil, false, nil
		}
		log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
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

func (l *loadBalancers) GetLoadBalancerName (ctx context.Context, clusterName string, service *v1.Service) string {
	clusterID := l.resources.clusterID
	log.Infof("GetLoadBalancerName:: clusterID is: %s, service.name is: %+v", clusterID, service.ObjectMeta.Name)
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	res, err := getLoadBalancerByName(clusterName, clusterID, loadBalancerName)
	if err != nil {
		log.Infof("GetLoadBalancerName:: getLoadBalancerByName succeed, return loadBalancerName is: %s", res.Data.Name)
		return res.Data.Name
	}
	log.Errorf("GetLoadBalancerName:: getLoadBalancerByName is error, err is: %s", err)
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
		return nil, errors.New("SessionAffinity is not supported currently, only support 'None' type")
	}

	// TODO check if kubernetes has already do validate
	loadBalancerExist := true
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	// step-1. get loadBalancer status
	loadBalancerGet, err := getLoadBalancerByName(clusterName, clusterID, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Infof("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is succeed, loadBalancer is not exist, then to create it")
			loadBalancerExist = false
		}
		log.Errorf("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is error, err is: %s", err)
		return nil, err
	} else {
		log.Infof("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is succeed, loadBalancer is exist, res is: %+v", loadBalancerGet)
	}

	// step-2. create or update loadBalancer by loadBalancerExist flag
	if loadBalancerExist {
		// loadBalancer is exist, then update it
		switch 0 {
		// only support classic yet
		case ClbLoadBalancerKindClassic:
			err := updateClassicLoadBalancer(ctx, clusterName, service, nodes, clusterID, loadBalancerName)
			if err != nil {
				log.Errorf("EnsureLoadBalancer:: step-2 cloud.updateClassicLoadBalancer is error, err is: %s", err)
				return nil, err
			}
			log.Infof("EnsureLoadBalancer:: step-2 cloud.updateClassicLoadBalancer succeed")
		default:
			log.Errorf("EnsureLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet")
			return nil, errors.New("Unsupported loadbalancer kind, only support [classic] yet")
		}
	} else {
		// loadBalancer is not exist, then create it
		switch 0 {
		// only support classic yet
		case ClbLoadBalancerKindClassic:
			err := createClassicLoadBalancer(ctx, clusterName, service, nodes, clusterID, loadBalancerName)
			if err != nil {
				log.Errorf("EnsureLoadBalancer:: step-2 cloud.createClassicLoadBalancer is error, err is: %s", err)
				return nil, err
			}
			log.Infof("EnsureLoadBalancer:: step-2 cloud.createClassicLoadBalancer succeed")
		default:
			log.Infof("EnsureLoadBalancer:: Unsupported loadbalancer kind, only support [classic] yet")
			return nil, errors.New("Unsupported loadbalancer kind, only support [classic] yet")
		}
	}

	// step-3. verify loadBalancer create or update successfully
	loadBalancerVerify, err := getLoadBalancerByName(clusterName, clusterID, loadBalancerName)
	if err != nil {
		log.Errorf("EnsureLoadBalancer:: step-3 cloud.getLoadBalancerByName is error, err is: %s", err)
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

	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	switch 0 {
	// only support classic yet
	case ClbLoadBalancerKindClassic:
		err := updateClassicLoadBalancer(ctx, clusterName, service, nodes, clusterID, loadBalancerName)
		if err != nil {
			log.Errorf("UpdateLoadBalancer:: cloud.updateClassicLoadBalancer is error, err is: %s", err)
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
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	log.Infof("EnsureLoadBalancerDeleted:: clusterID is: %s, service.name is: %+v", clusterID, service.ObjectMeta.Name)
	return deleteLoadBalancer(ctx, clusterName, clusterID, loadBalancerName)
}

func getLoadBalancerByName(clusterName, clusterID, loadBalancerName string) (*clb.DescribeLoadBalancersResponse, error) {
	// we don't need to check loadbalancer kind here because ensureLoadBalancerInstance will ensure the kind is right
	response, err := clb.DescribeLoadBalancers(&clb.DescribeLoadBalancersArgs{
		ClusterName: clusterName,
		CLusterID: clusterID,
		LoadBalancerName: loadBalancerName,
	})

	// api with error
	if err != nil {
		return nil, err
	}

	// loadBalancer is not exist
	if len(response.Data.Vips) < 1 {
		return nil, ErrCloudLoadBalancerNotFound
	}

	return response, nil
}

func updateClassicLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node, clusterID, loadBalancerName string) error {
	// get service ports info
	var portMapSlice []clb.PortMapping
	var portMapTmp clb.PortMapping
	for _, ports := range service.Spec.Ports {
		portMapTmp.Port = ports.Port
		portMapTmp.Nodeport = ports.NodePort
		portMapTmp.Protocol = ports.Protocol
	}
	portMapSlice = append(portMapSlice, portMapTmp)
	log.Infof("updateClassicLoadBalancer:: portMapSlice is: %s", portMapSlice)

	// get nodes providerID info
	var nodeIdSlice []string
	for _, node := range nodes {
		nodeIdSlice = append(nodeIdSlice, node.Spec.ProviderID)
	}
	log.Infof("updateClassicLoadBalancer:: nodeIdSlice is: %s", nodeIdSlice)

	// to create loadBalancer
	res, err := clb.UpdateLoadBalancers(&clb.UpdateLoadBalancersArgs{
		ClusterName:clusterName,
		CLusterID: clusterID,
		LoadBalancerName: loadBalancerName,
		// need to get from cluster
		NodeID: make([]string, 2),
		Annotations: service.ObjectMeta.Annotations,
		PortMap: portMapSlice,
	})

	if err != nil {
		return err
	}

	taskID := res.Data.TaskID
	log.Infof("updateClassicLoadBalancer: create task succeed, TaskID is: %+v", taskID)

	// to check loadBalancer update task result
	err2 := describeLoadBalancersTaskResult(taskID)

	if err2 != nil {
		return err2
	}
	return nil
}

func createClassicLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node, clusterID, loadBalancerName string) error {
	// get services ports info
	var portMapSlice []clb.PortMapping
	var portMapTmp clb.PortMapping
	for _, ports := range service.Spec.Ports {
		portMapTmp.Port = ports.Port
		portMapTmp.Nodeport = ports.NodePort
		portMapTmp.Protocol = ports.Protocol
	}
	portMapSlice = append(portMapSlice, portMapTmp)
	log.Infof("createClassicLoadBalancer:: portMapSlice is: %s", portMapSlice)

	// get nodes providerID info
	var nodeIdSlice []string
	for _, node := range nodes {
		nodeIdSlice = append(nodeIdSlice, node.Spec.ProviderID)
	}
	log.Infof("createClassicLoadBalancer:: nodeIdSlice is: %s", nodeIdSlice)

	// to create loadBalancer
	res, err := clb.CreateLoadBalancers(&clb.CreateLoadBalancersArgs{
		ClusterName:clusterName,
		CLusterID: clusterID,
		LoadBalancerName: loadBalancerName,
		// need to get from cluster
		NodeID: nodeIdSlice,
		Annotations: service.ObjectMeta.Annotations,
		PortMap: portMapSlice,
	})

	if err != nil {
		return err
	}

	taskID := res.Data.TaskID
	log.Infof("createClassicLoadBalancer: create task succeed, TaskID is: %+v", taskID)

	// to check loadBalancer create task result
	err2 := describeLoadBalancersTaskResult(taskID)

	if err2 != nil {
		return err2
	}
	return nil
}

func deleteLoadBalancer(ctx context.Context, clusterName, clusterID, loadBalancerName string) error {
	// check the loadBalancer is exist or not
	res, err := getLoadBalancerByName(clusterName, clusterID, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Warnf("deleteLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer is not exist, so do not delete action, return nil")
			return nil
		}
		log.Errorf("deleteLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
		return err
	}
	log.Infof("deleteLoadBalancer:: cloud.getLoadBalancerByName, res is: %+v, then to delete it", res)

	// loadBalancer is exist, then to delete it
	res2, err2 := clb.DeleteLoadBalancers(&clb.DeleteLoadBalancersArgs{
		ClusterName: clusterName,
		LoadBalancerName: loadBalancerName,
	})

	if err2 != nil {
		return err2
	}

	taskID := res2.Data.TaskID
	log.Infof("deleteLoadBalancer:: clb.DeleteLoadBalancers delete task_id is: %s", taskID)

	// to check loadBalancer delete task result
	err3 := describeLoadBalancersTaskResult(taskID)

	if err3 != nil {
		return err3
	}

	return nil
}

func describeLoadBalancersTaskResult (taskID string) error {
	for i:= 1; i < 120; i++ {
		res, err := clb.DescribeLoadBalancersTaskResult(&clb.DescribeLoadBalancersTaskResultArgs{
			TaskID: taskID,
		})
		if err != nil {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult is error, err is:%s", err)
		}
		if res.Data.Status == "running"{
			log.Infof("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult is running")
		} else if res.Data.Status == "ok" {
			log.Infof("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult succeed, status is: %s", res.Data.Status)
			return nil
		} else if res.Data.Status == "failed" {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult failed, status is: %s", res.Data.Status)
			return errors.New("clb.DescribeLoadBalancersTaskResult failed")
		} else if res.Data.Status == "error" {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult error, status is: %s", res.Data.Status)
			return errors.New("clb.DescribeLoadBalancersTaskResult error")
		} else {
			log.Errorf("DescribeLoadBalancersTaskResult:: clb.DescribeLoadBalancersTaskResult time out, running more than 10 minutes")
			return errors.New("clb.DescribeLoadBalancersTaskResult time out, running more than 20 minutes")
		}

		time.Sleep(time.Second * 10)
		log.Infof("DescribeLoadBalancersTaskResult:: time.sleep 30s")
	}
	return nil
}