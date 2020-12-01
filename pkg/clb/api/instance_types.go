package api

type DescribeInstancesLabelsAndNodeNameArgs struct {
	ClusterID string `json:"cluster_id"`
	NodeID    string `json:"node_id"`
}
type DescribeInstancesLabelsAndNodeNameResponse struct {
	Response
	Data struct {
		Labels      []map[string]string `json:"labels"`
		NodeName    string              `json:"node_name"`
		Taints      []map[string]string `json:"taints"`
		InternalIPs []string            `json:"internal_ips"`
		ExternalIPs []string            `json:"external_ips"`
	} `json:"Data"`
}

type DescribeInstanceExistsByProviderIDArgs struct {
	ProviderID string `json:"provider_id"`
}

type DescribeInstanceExistsByProviderIDResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"Data"`
}
