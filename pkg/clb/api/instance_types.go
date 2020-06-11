package api


type DescribeInstancesLabelsAndNodeNameArgs struct {
	ClusterID  string`json:"cluster_id"`
	NodeID     string`json:"node_id"`
}
type DescribeInstancesLabelsAndNodeNameResponse struct {
	Response
	Data struct {
		Labels  	[]LabelMapping`json:"labels"`
		NodeName 	string`json:"node_name"`
	}`json:"data"`
}
type LabelMapping struct {
	Key 	string`json:"key"`
	Value 	string`json:"value"`
}
