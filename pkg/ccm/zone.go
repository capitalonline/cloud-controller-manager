package ccm

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cloud-provider"
)

type zones struct {
	resources *resources
	region    string
}

func newZones(resources *resources, region string) cloudprovider.Zones {
	return zones{
		resources: resources,
		region:    region,
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
	return cloudprovider.Zone{Region: "not support"}, nil
}

// GetZoneByNodeName returns a cloudprovider.Zone from the droplet identified
// by nodeName. GetZoneByNodeName only sets the Region field of the returned
// cloudprovider.Zone.
func (z zones) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{Region: "not support yet"}, nil
}