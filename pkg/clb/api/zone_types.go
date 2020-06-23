package api


type DescribeZoneByProviderIDArgs struct {
	ClusterID  string`json:"cluster_id"`
	NodeID     string`json:"node_id"`
}
type DescribeZoneByProviderIDResponse struct {
	Response
	Data struct {
		Region string`json:"region"`
	}`json:"data"`
}