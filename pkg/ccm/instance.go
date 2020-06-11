package ccm

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/cloud-provider"

	clb "github.com/capitalonline/cloud-controller-manager/pkg/clb/api"
)

type instances struct {
	region    string
	resources *resources
	k8sClient *kubernetes.Clientset
}

func newInstances(resources *resources, k8sClientSet *kubernetes.Clientset, region string) cloudprovider.Instances {
	return &instances{
		resources: resources,
		region:    region,
		k8sClient: k8sClientSet,
	}
}

// NodeAddresses returns all the valid addresses of the droplet identified by
// nodeName. Only the public/private IPv4 addresses are considered for now.
//
// When nodeName identifies more than one droplet, only the first will be
// considered.
func (i *instances) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
	log.Infof("NodeAddresses:: nodeName is: %s", nodeName)
	log.Infof("not support yet")

	return nil, nil
}

// NodeAddressesByProviderID returns all the valid addresses of the droplet
// identified by providerID. Only the public/private IPv4 addresses will be
// considered for now.
func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	log.Infof("NodeAddressesByProviderID:: providerID is: %s", providerID)
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
	log.Infof("ExternalID:: nodeName is: %s", nodeName)
	return i.InstanceID(ctx, nodeName)
}

// InstanceID returns the cloud provider ID of the droplet identified by nodeName.
func (i *instances) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	log.Infof("InstanceID:: nodeName is: %s", nodeName)
	log.Infof("not support yet")

	return "", nil
}

// InstanceType returns the type of the droplet identified by name.
func (i *instances) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	log.Infof("InstanceType:: name is: %s", name)
	log.Infof("not support yet")

	return "", nil
}

// InstanceTypeByProviderID returns the type of the droplet identified by providerID.
func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	clusterID := i.resources.clusterID
	log.Infof("InstanceTypeByProviderID:: clusterID is: %s, providerID is: %s", clusterID, providerID)
	// get node labels and nodeName
	res, err := getNodeInstanceTypeAndNodeNameByProviderID(clusterID, providerID)
	if err != nil {
		log.Errorf("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID is error, err is: %s", err)
	}
	log.Infof("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, res is: %+v", res)

	if res.Data.NodeName == "" {
		log.Errorf("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
		return "", errors.New("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
	}

	// init cluster node labels exclude "node.kubernetes.io/instance-type"
	nodeName := res.Data.NodeName
	node, err := i.k8sClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Get to get node labels error, err is: %s", err)
		return "", errors.New("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Get to get node labels error")
	}

	returnInstanceTypeValue := ""
	for _, label := range res.Data.Labels {
		if label.Key == "node.kubernetes.io/instance-type" {
			returnInstanceTypeValue = label.Value
			continue
		}
		node.ObjectMeta.Labels[label.Key] = label.Value
	}

	log.Infof("InstanceTypeByProviderID:: succeed, returnInstanceTypeValue is: %s", returnInstanceTypeValue)
	// return node label which labels.key is "node.kubernetes.io/instance-type"
	return returnInstanceTypeValue, nil
}

// AddSSHKeyToAllInstances is not implemented; it always returns an error.
func (i *instances) AddSSHKeyToAllInstances(_ context.Context, _ string, _ []byte) error {
	log.Infof("AddSSHKeyToAllInstances:: none")
	return errors.New("not implemented")
}

// CurrentNodeName returns hostname as a NodeName value.
func (i *instances) CurrentNodeName(_ context.Context, hostname string) (types.NodeName, error) {
	log.Infof("CurrentNodeName:: hostname is: %s", hostname)
	return types.NodeName(hostname), nil
}

// InstanceExistsByProviderID returns true if the droplet identified by
// providerID is running.
func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	log.Infof("InstanceExistsByProviderID:: providerID is: %s", providerID)
	log.Infof("not support yet")

	return false, nil
}

// InstanceShutdownByProviderID returns true if the droplet is turned off
func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	log.Infof("InstanceShutdownByProviderID:: providerID is: %s", providerID)
	log.Infof("not support yet")

	return false, nil
}

func getNodeInstanceTypeAndNodeNameByProviderID(clusterID, providerID string)(*clb.DescribeInstancesLabelsAndNodeNameResponse, error){
	response, err := clb.DescribeInstancesLabelsAndNodeName(&clb.DescribeInstancesLabelsAndNodeNameArgs{
		ClusterID: clusterID,
		NodeID: providerID,
	})

	// api with error
	if err != nil {
		return nil, err
	}

	return response, nil
}