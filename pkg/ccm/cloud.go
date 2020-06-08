package ccm

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"
)

const (
	ProviderName string = "cdstcloud"
)

var (
	CloudInstanceNotFound = errors.New("cdscloud instance not found")
)

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, NewCloud)
}

func NewCloud(config io.Reader) (cloudprovider.Interface, error) {
	var c Config
	if config != nil {
		cfg, err := ioutil.ReadAll(config)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(cfg, &c); err != nil {
			return nil, err
		}
	}

	if c.Region == "" {
		c.Region = os.Getenv("CDSCLOUD_CLOUD_CONTROLLER_MANAGER_REGION")
	}
	if c.SecretId == "" {
		c.SecretId = os.Getenv("CDSCLOUD_CLOUD_CONTROLLER_MANAGER_SECRET_ID")
	}
	if c.SecretKey == "" {
		c.SecretKey = os.Getenv("CDSCLOUD_CLOUD_CONTROLLER_MANAGER_SECRET_KEY")
	}

	return &Cloud{config: c}, nil
}

type Cloud struct {
	config Config

	kubeClient kubernetes.Interface
}

type Config struct {
	Region string `json:"region"`

	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`

}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping activities within the cloud provider.
func (cloud *Cloud) Initialize(clientBuilder controller.ControllerClientBuilder) {
	cloud.kubeClient = clientBuilder.ClientOrDie("cdstcloud-cloud-provider")
	return
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (cloud *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return cloud, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (cloud *Cloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (cloud *Cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (cloud *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (cloud *Cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (cloud *Cloud) ProviderName() string {
	return ProviderName
}

// HasClusterID returns true if a ClusterID is required and set
func (cloud *Cloud) HasClusterID() bool {
	return false
}
