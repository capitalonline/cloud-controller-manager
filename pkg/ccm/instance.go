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
func (i *instances) NodeAddressesByProviderID(ctx context.Context, providerID string, nodeName string) ([]v1.NodeAddress, error) {
	clusterID := i.resources.clusterID
	snatIp := ""
	nodeAnnotations, err := i.k8sClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	log.Infof("NodeAddressesByProviderID:: node.ObjectMeta.Annotations are: %+v", nodeAnnotations.ObjectMeta.Annotations)
	for key, value := range nodeAnnotations.ObjectMeta.Annotations {
		log.Infof("annotation:::::::%s====%s", key, value)
	}
	log.Infof("snatIp:: %s", snatIp)
	
	log.Infof("NodeAddressesByProviderID:: providerID is: %s", providerID)
	// get node nodeName
	res, err := getNodeInstanceTypeAndNodeNameByProviderID(clusterID, providerID)
	if err != nil {
		log.Errorf("NodeAddressesByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID is error, err is: %s", err)
		return nil, err
	}
	log.Infof("NodeAddressesByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, res is: %+v", res)

	//if res.Data.NodeName == "" {
	//	log.Errorf("NodeAddressesByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
	//	return nil, errors.New("NodeAddressesByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
	//}
	//
	//// get cluster node address
	//nodeAddress, err := i.k8sClient.CoreV1().Nodes().Get(res.Data.NodeName, metav1.GetOptions{})
	//if err != nil {
	//	log.Errorf("NodeAddressesByProviderID:: k8sClient.CoreV1().Nodes().Get to get node address error, err is: %s", err)
	//	return nil, errors.New("NodeAddressesByProviderID:: k8sClient.CoreV1().Nodes().Get to get node address error")
	//}

	// get node's address by providerID

	var nodeAddressStruct v1.NodeAddress
	var nodeAddressSlice []v1.NodeAddress
	// init internal ip
	if len(res.Data.InternalIPs) != 0 {
		for _, internalIP := range res.Data.InternalIPs {
			nodeAddressStruct.Type = v1.NodeAddressType("InternalIP")
			nodeAddressStruct.Address = internalIP
			nodeAddressSlice = append(nodeAddressSlice, nodeAddressStruct)
		}
	}

	// init external ip
	if len(res.Data.ExternalIPs) != 0 {
		for _, externalIP := range res.Data.ExternalIPs {
			nodeAddressStruct.Type = v1.NodeAddressType("ExternalIP")
			nodeAddressStruct.Address = externalIP
			nodeAddressSlice = append(nodeAddressSlice, nodeAddressStruct)
		}
	}

	// init hostname
	if res.Data.NodeName != "" {
		nodeAddressStruct.Type = v1.NodeAddressType("Hostname")
		nodeAddressStruct.Address = res.Data.NodeName
		nodeAddressSlice = append(nodeAddressSlice, nodeAddressStruct)
	}
	//nodeAddressStructTmp.Type = v1.NodeAddressType("ExternalIP")
	//nodeAddressStructTmp.Address = "117.168.192.110"

	// update node status
	log.Infof("NodeAddressesByProviderID: nodeAddressSlice is: %+v", nodeAddressSlice)
	//nodeAddress.Status.Addresses = nodeSliceTmp
	//_, err = i.k8sClient.CoreV1().Nodes().Update(nodeAddress)
	//if err != nil {
	//	log.Errorf("NodeAddressesByProviderID:: k8sClient.CoreV1().Nodes().Update(node) error, err is: %s", err)
	//	return nil, err
	//}

	log.Infof("NodeAddressesByProviderID: succeed!")
	return nodeAddressSlice, nil
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
		return "", err
	}
	log.Infof("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, res is: %+v", res)

	if res.Data.NodeName == "" {
		log.Errorf("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
		return "", errors.New("InstanceTypeByProviderID:: getNodeInstanceTypeAndNodeNameByProviderID, nodeName is empty")
	}

	// init cluster node labels exclude "node.kubernetes.io/instance-type"
	nodeName := res.Data.NodeName
	nodeLabels, err := i.k8sClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Get to get node labels error, err is: %s", err)
		return "", errors.New("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Get to get node labels error")
	}

	// set labels
	returnInstanceTypeValue := ""
	if len(res.Data.Labels) != 0 {
		for _, label := range res.Data.Labels {
			for key, value := range label {
				if key == "node.kubernetes.io/instance-type" {
					returnInstanceTypeValue = value
					continue
				}
				nodeLabels.ObjectMeta.Labels[key] = value
				log.Infof("InstanceTypeByProviderID:: set node.ObjectMeta.Labels is: %s", label)
			}
		}
		// update nodes
		_, err = i.k8sClient.CoreV1().Nodes().Update(nodeLabels)
		if err != nil {
			log.Errorf("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Update(node) error, err is: %s", err)
			return "", err
		}
		log.Infof("InstanceTypeByProviderID: update node's label succeed!")
		log.Infof("InstanceTypeByProviderID:: node.ObjectMeta.Labels are: %+v", nodeLabels.ObjectMeta.Labels)
	}
	
	// set Annotations
	nodeAnnotations, err := i.k8sClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if len(res.Data.Annotations) != 0 {
		for _, annotation := range res.Data.Annotations {
			for key, value := range annotation {
				nodeAnnotations.ObjectMeta.Annotations[key] = value
				log.Infof("InstanceTypeByProviderID:: nodeLabels.ObjectMeta.Annotations: %s", annotation)
			}
		}
		// update nodes
		_, err = i.k8sClient.CoreV1().Nodes().Update(nodeAnnotations)
		if err != nil {
			log.Errorf("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Update(node) error, err is: %s", err)
			return "", err
		}
		log.Infof("InstanceTypeByProviderID: update node's annotation succeed!")
		log.Infof("InstanceTypeByProviderID:: node.ObjectMeta.Annotations are: %+v", nodeAnnotations.ObjectMeta.Annotations)
	}

	// set taints
	// to fix "the object has been modified; please apply your changes to the latest version and try again" issue
	nodeTaints, err := i.k8sClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	var taintStructTmp v1.Taint
	// var taintTmpSlice []v1.Taint
	taintSliceTmp := nodeTaints.Spec.Taints
	if len(res.Data.Taints) != 0 {
		for _, taint := range res.Data.Taints {
			for key, value := range taint {
				taintStructTmp.Key = key
				taintStructTmp.Effect = v1.TaintEffect(value)
			}
			taintSliceTmp = append(taintSliceTmp, taintStructTmp)
		}

		log.Infof("InstanceTypeByProviderID:: taintSliceTmp is: %+v", taintSliceTmp)
		nodeTaints.Spec.Taints = taintSliceTmp

		// update nodes
		_, err = i.k8sClient.CoreV1().Nodes().Update(nodeTaints)
		if err != nil {
			log.Errorf("InstanceTypeByProviderID:: k8sClient.CoreV1().Nodes().Update(node) error, err is: %s", err)
			return "", err
		}
		log.Infof("InstanceTypeByProviderID: update node's taints succeed!")
		log.Infof("InstanceTypeByProviderID:: node.Spec.Taints are: %+v", nodeTaints.Spec.Taints)
	}

	// returnInstanceTypeValue := "cds.vm.8c.8g"
	log.Infof("InstanceTypeByProviderID:: succeed, returnInstanceTypeValue is: %s", returnInstanceTypeValue)
	// return node label which label.key is "node.kubernetes.io/instance-type"
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

// InstanceExistsByProviderID returns true if the instance identified by
// providerID is running.
func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	log.Infof("InstanceExistsByProviderID:: providerID is: %s", providerID)
	ok, err := describeInstanceExistsByProviderID(providerID)

	if !ok {
		if err != nil {
			log.Errorf("InstanceExistsByProviderID:: instance with unknown error")
			return false, err
		}
		log.Errorf("InstanceExistsByProviderID:: instance is not exist")
		return false, nil
	}
	log.Infof("InstanceExistsByProviderID:: instance is exist")
	return true, nil
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

func describeInstanceExistsByProviderID(providerID string) (bool, error) {
	response, err := clb.DescribeInstanceExistsByProviderID(&clb.DescribeInstanceExistsByProviderIDArgs{
		ProviderID: providerID,
	})

	// api with error
	if err != nil {
		return false, err
	}

	if response.Data.Status == "true" {
		return true, nil
	} else if response.Data.Status == "false" {
		return  false, nil
	} else {
		return false, errors.New("unknown error")
	}
}
