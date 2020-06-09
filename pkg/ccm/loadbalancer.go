package ccm

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"

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
	log.Infof("GetLoadBalancer:: clusterName is: %s, service is: %+v", clusterName, service)
	loadBalancerName := cloudprovider.GetLoadBalancerName(service)
	log.Infof("GetLoadBalancer:: clusterName is: %s, loadBalancerName is: %s", clusterName, loadBalancerName)

	loadBalancer, err := getLoadBalancerByName(clusterName, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer  is not exist")
			return nil, false, nil
		}
		log.Errorf("GetLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
		return nil, false, err
	}
	log.Infof("GetLoadBalancer:: cloud.getLoadBalancerByName, res is: %s", loadBalancer)

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
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	res, err := getLoadBalancerByName(clusterName, loadBalancerName)
	if err != nil {
		return res.Data.Name
	}
	return err.Error()
}
func (l *loadBalancers) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	log.Infof("EnsureLoadBalancer:: clusterName is: %s, service is: %+v, nodes is: %+v", clusterName, service, nodes)
	if service.Spec.SessionAffinity != v1.ServiceAffinityNone {
		log.Errorf("EnsureLoadBalancer:: SessionAffinity is not supported currently, only support 'None' type")
		return nil, errors.New("SessionAffinity is not supported currently, only support 'None' type")
	}

	// TODO check if kubernetes has already do validate
	loadBalancerExist := true
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	// step-1. get loadBalancer status
	loadBalancerGet, err := getLoadBalancerByName(clusterName, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Infof("EnsureLoadBalancer:: step-1 cloud.getLoadBalancerByName is succeed, loadBalancer is not exist")
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
			err := updateClassicLoadBalancer(ctx, clusterName, service, nodes, loadBalancerName)
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
			err := createClassicLoadBalancer(ctx, clusterName, service, nodes, loadBalancerName)
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
	loadBalancerVerify, err := getLoadBalancerByName(clusterName, loadBalancerName)
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
	log.Infof("UpdateLoadBalancer:: clusterName is: %s, service is: %+v, nodes is: %+v", clusterName, service, nodes)
	loadBalancerName := cloudprovider.GetLoadBalancerName(service)
	// only support classic yet
	switch 0 {
	case ClbLoadBalancerKindClassic:
		err := updateClassicLoadBalancer(ctx, clusterName, service, nodes, loadBalancerName)
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
	log.Infof("EnsureLoadBalancerDeleted:: clusterName is: %s, service is: %+v", clusterName, service)
	return deleteLoadBalancer(ctx, clusterName, service)
}

func getLoadBalancerByName(clusterName, loadBalancerName string) (*clb.DescribeLoadBalancersResponse, error) {
	// we don't need to check loadbalancer kind here because ensureLoadBalancerInstance will ensure the kind is right
	response, err := clb.DescribeLoadBalancers(&clb.DescribeLoadBalancersArgs{
		ClusterName: clusterName,
		LoadBalancerName: loadBalancerName,
	})
	// sdk with error
	if err != nil {
		return nil, err
	}
	// loadBalancer is not exist
	if len(response.Data.Vips) < 1 {
		return nil, ErrCloudLoadBalancerNotFound
	}

	return response, nil
}

func updateClassicLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node, loadBalancerName string) error {
	res, err := clb.UpdateLoadBalancers(&clb.UpdateLoadBalancersArgs{
		ClusterName:clusterName,
		LoadBalancerName: loadBalancerName,
		Service: service,
		Nodes: nodes,
	})
	if err != nil {
		return err
	}
	log.Infof("updateClassicLoadBalancer:: clb.UpdateLoadBalancers is succeed, res is: %+v", res)
	return nil
}

func createClassicLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node, loadBalancerName string) error {
	res, err := clb.CreateLoadBalancers(&clb.CreateLoadBalancersArgs{
		ClusterName:clusterName,
		LoadBalancerName: loadBalancerName,
		Service: service,
		Nodes: nodes,
	})
	if err != nil {
		return err
	}
	log.Infof("createClassicLoadBalancer:: clb.CreateLoadBalancers is succeed, res is: %+v", res)
	return nil
}

func deleteLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) error {
	// check the loadBalancer is exist
	loadBalancerName := cloudprovider.DefaultLoadBalancerName(service)
	res, err := getLoadBalancerByName(clusterName, loadBalancerName)
	if err != nil {
		if err == ErrCloudLoadBalancerNotFound {
			log.Warnf("deleteLoadBalancer:: cloud.getLoadBalancerByName, loadBalancer is not exist, so do not delete action, return nil")
			return nil
		}
		log.Errorf("deleteLoadBalancer:: cloud.getLoadBalancerByName is error, err is: %s", err)
		return err
	}
	log.Infof("deleteLoadBalancer:: cloud.getLoadBalancerByName, res is: %+v, then delete it", res)
	// delete the loadBalancer
	_, err2 := clb.DeleteLoadBalancers(&clb.DeleteLoadBalancersArgs{
		ClusterName: clusterName,
		LoadBalancerName: loadBalancerName,
	})
	if err2 != nil {
		log.Errorf("deleteLoadBalancer:: clb.DeleteLoadBalancers is error, err is: %s", err2)
		return err2
	}
	log.Infof("deleteLoadBalancer:: clb.DeleteLoadBalancers delete loadBalancer succeed!")
	return nil
}