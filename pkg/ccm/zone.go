package ccm

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/cloud-provider"

	clb "github.com/capitalonline/cloud-controller-manager/pkg/clb/api"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type zones struct {
	resources *resources
	region    string
	k8sClient *kubernetes.Clientset
}

func newZones(resources *resources, k8sClientSet *kubernetes.Clientset, region string) cloudprovider.Zones {
	return zones{
		resources: resources,
		region:    region,
		k8sClient: k8sClientSet,
	}
}

// GetZone returns a cloudprovider.Zone from the region of z. GetZone only sets
// the Region field of the returned cloudprovider.Zone.
//
// Kuberenetes uses this method to get the region that the program is running in.
func (z zones) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: z.region}, nil
}

// GetZoneByProviderID returns a cloudprovider.Zone from the droplet identified
// by providerID. GetZoneByProviderID only sets the Region field of the
// returned cloudprovider.Zone.
func (z zones) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	clusterID := z.resources.clusterID
	log.Infof("GetZoneByProviderID:: clusterID is: %s, providerID is: %s", clusterID, providerID)

	res, err := getZoneByProviderID(clusterID, providerID)
	if err != nil {
		log.Errorf("GetZoneByProviderID:: getZoneByProviderID is error, err is: %s", err)
		SentrySendError(fmt.Errorf("GetZoneByProviderID:: getZoneByProviderID is error, err is: %s", err))
	}

	region := res.Data.Region
	return cloudprovider.Zone{Region: region}, nil
}

// GetZoneByNodeName returns a cloudprovider.Zone from the droplet identified
// by nodeName. GetZoneByNodeName only sets the Region field of the returned
// cloudprovider.Zone.
func (z zones) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	clusterID := z.resources.clusterID
	log.Infof("GetZoneByNodeName:: clusterID is: %s, nodeName is: %s", clusterID, string(nodeName))

	// get providerID from cluster
	res, err := z.k8sClient.CoreV1().Nodes().Get(string(nodeName), metav1.GetOptions{})
	if err != nil {
		log.Errorf("GetZoneByNodeName:: k8sClient.CoreV1().Nodes().Get to get node's providerID is error, err is: %s", err)
		SentrySendError(fmt.Errorf("GetZoneByNodeName:: k8sClient.CoreV1().Nodes().Get to get node's providerID is error, err is: %s", err))
		return  cloudprovider.Zone{Region: ""}, err
	}
	providerID := res.Spec.ProviderID

	// get zone by providerID
	res2, err2 := getZoneByProviderID(clusterID, providerID)
	if err2 != nil {
		log.Errorf("GetZoneByProviderID:: getZoneByProviderID is error, err is: %s", err)
		SentrySendError(fmt.Errorf("GetZoneByProviderID:: getZoneByProviderID is error, err is: %s", err))
		return  cloudprovider.Zone{Region: ""}, err2
	}

	region := res2.Data.Region
	log.Infof("GetZoneByNodeName:: succeed, region is: %s", region)
	return cloudprovider.Zone{Region: region}, nil
}

func getZoneByProviderID(clusterID, providerID string) (*clb.DescribeZoneByProviderIDResponse, error) {
	response, err := clb.DescribeZoneByProviderID(&clb.DescribeZoneByProviderIDArgs{
		ClusterID: clusterID,
		NodeID: providerID,
	})

	// api with error
	if err != nil {
		return nil, err
	}

	return response, nil
}