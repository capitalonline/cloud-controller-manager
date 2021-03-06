package ccm

import (
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"os"

	"k8s.io/client-go/informers"
	cloudprovider "k8s.io/cloud-provider"
)

const (
	ProviderName string = "cdscloud"
	// At some point we should revisit how we start up our CCM implementation.
	// Having to look at env vars here instead of in the cmd itself is not ideal.
	// One option is to construct our own command that's specific to us.
	// Alibaba's ccm is an example how this is done.
	// https://github.com/kubernetes/cloud-provider-alibaba-cloud/blob/master/cmd/cloudprovider/app/ccm.go
	cdsClusterID      string = "CDS_CLUSTER_ID"
	cdsClusterRegionID  string = "CDS_CLUSTER_REGION_ID"
)

//var (
//	CloudInstanceNotFound = errors.New("cdscloud instance not found")
//)

type cloud struct {
	instances     cloudprovider.Instances
	zones         cloudprovider.Zones
	loadbalancers cloudprovider.LoadBalancer

	httpServer *http.Server
}

func newCloud() (cloudprovider.Interface, error) {
	// init k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("newCloud:: Failed to create kubernetes config: %v", err)
	}
	k8sClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("newCloud:: Failed to create kubernetes client: %v", err)
	}

	clusterID := os.Getenv(cdsClusterID)
	region := os.Getenv(cdsClusterRegionID)
	resources := newResources(clusterID)

	var httpServer *http.Server
	return &cloud{
		instances:     newInstances(resources, k8sClientSet, region),
		zones:         newZones(resources, k8sClientSet, region),
		loadbalancers: newLoadBalancers(resources, region),
		httpServer: httpServer,
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(io.Reader) (cloudprovider.Interface, error) {
		return newCloud()
	})
}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping activities within the cloud provider.
func (c *cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	clientset := clientBuilder.ClientOrDie("do-shared-informers")
	sharedInformer := informers.NewSharedInformerFactory(clientset, 0)
	// clusterID := os.Getenv(cdsClusterID)
	// res := NewResourcesController(clusterID, sharedInformer.Core().V1().Services(), clientset)
	sharedInformer.Start(nil)
	sharedInformer.WaitForCacheSync(nil)
	// go res.Run(stop)
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadbalancers, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) Instances() (cloudprovider.Instances, bool) {
	return c.instances, true
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (c *cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (c *cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (c *cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (c *cloud) ProviderName() string {
	return ProviderName
}

// HasClusterID returns true if a ClusterID is required and set
func (c *cloud) HasClusterID() bool {
	return false
}
