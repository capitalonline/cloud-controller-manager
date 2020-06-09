package ccm

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cloud-provider"
)

type instances struct {
	region    string
	resources *resources
}

func newInstances(resources *resources, region string) cloudprovider.Instances {
	return &instances{
		resources: resources,
		region:    region,
	}
}

// NodeAddresses returns all the valid addresses of the droplet identified by
// nodeName. Only the public/private IPv4 addresses are considered for now.
//
// When nodeName identifies more than one droplet, only the first will be
// considered.
func (i *instances) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
	log.Infof("not support yet")

	return nil, nil
}

// NodeAddressesByProviderID returns all the valid addresses of the droplet
// identified by providerID. Only the public/private IPv4 addresses will be
// considered for now.
func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	log.Infof("not support yet")

	return nil, nil
}

// ExternalID returns the cloud provider ID of the droplet identified by
// nodeName. If the droplet does not exist or is no longer running, the
// returned error will be cloudprovider.InstanceNotFound.
//
// When nodeName identifies more than one droplet, only the first will be
// considered.
func (i *instances) ExternalID(ctx context.Context, nodeName types.NodeName) (string, error) {
	return i.InstanceID(ctx, nodeName)
}

// InstanceID returns the cloud provider ID of the droplet identified by nodeName.
func (i *instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	log.Infof("not support yet")

	return "", nil
}

// InstanceType returns the type of the droplet identified by name.
func (i *instances) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	log.Infof("not support yet")

	return "", nil
}

// InstanceTypeByProviderID returns the type of the droplet identified by providerID.
func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	log.Infof("not support yet")

	return "", nil
}

// AddSSHKeyToAllInstances is not implemented; it always returns an error.
func (i *instances) AddSSHKeyToAllInstances(_ context.Context, _ string, _ []byte) error {
	return errors.New("not implemented")
}

// CurrentNodeName returns hostname as a NodeName value.
func (i *instances) CurrentNodeName(_ context.Context, hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}

// InstanceExistsByProviderID returns true if the droplet identified by
// providerID is running.
func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	log.Infof("not support yet")

	return false, nil
}

// InstanceShutdownByProviderID returns true if the droplet is turned off
func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	log.Infof("not support yet")

	return false, nil
}
