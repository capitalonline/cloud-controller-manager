package ccm

import (
	"sync"
	"time"

	v1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	v1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog"
)

const (
	controllerSyncTagsPeriod = 15 * time.Minute
	// syncTagsTimeout          = 1 * time.Minute
)

//type tagMissingError struct {
//	error
//}

type resources struct {
	clusterID    string

	kclient kubernetes.Interface

	mutex sync.RWMutex
}

// newResources initializes a new resources instance.
// kclient can only be set during the cloud. Initialize call since that is when
// the cloud provider framework provides us with a clientset. Fortunately, the
// initialization order guarantees that kclient won't be consumed prior to it
// being set.
func newResources(clusterID string) *resources {
	return &resources{
		clusterID:    clusterID,
	}
}

type syncer interface {
	Sync(name string, period time.Duration, stopCh <-chan struct{}, fn func() error)
}

type tickerSyncer struct{}

func (s *tickerSyncer) Sync(name string, period time.Duration, stopCh <-chan struct{}, fn func() error) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	// manually call to avoid initial tick delay
	if err := fn(); err != nil {
		klog.Errorf("%s failed: %s", name, err)
	}

	for {
		select {
		case <-ticker.C:
			if err := fn(); err != nil {
				klog.Errorf("%s failed: %s", name, err)
			}
		case <-stopCh:
			return
		}
	}
}

// ResourcesController is responsible for managing DigitalOcean cloud
// resources. It maintains a local state of the resources and
// synchronizes when needed.
type ResourcesController struct {
	kclient   kubernetes.Interface
	svcLister v1lister.ServiceLister

	clusterID string
	syncer    syncer
}

// NewResourcesController returns a new resource controller.
func NewResourcesController(clusterID string, inf v1informers.ServiceInformer, client kubernetes.Interface) *ResourcesController {
	r := &ResourcesController{}
	r.kclient = client

	return &ResourcesController{
		clusterID: clusterID,
		kclient:   client,
		svcLister: inf.Lister(),
		syncer:    &tickerSyncer{},
	}
}

// Run starts the resources controller loop.
func (r *ResourcesController) Run(stopCh <-chan struct{}) {
	if r.clusterID == "" {
		klog.Info("No cluster ID configured -- skipping cluster dependent syncers.")
		return
	}
	go r.syncer.Sync("tags syncer", controllerSyncTagsPeriod, stopCh, nil)
}


