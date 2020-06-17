package api

type DescribeLoadBalancersArgs struct {
	ClusterName      string `json:"cluster_name"`
	CLusterID        string `json:"cluster_id"`
	LoadBalancerName string `json:"loadbalancer_name"`
}
type DescribeLoadBalancersResponse struct {
	Response
	Data struct {
		Status string   `json:"status"`
		Name   string   `json:"name"`
		Vips   []string `json:"vips"`
	} `json:"Data"`
}

type PortMapping struct {
	// Protocol v1.Protocol `json:"protocol"`
	Port     int32       `json:"port"`
	NodePort int32       `json:"node_port"`
}
type CreateLoadBalancersArgs struct {
	ClusterName      string            `json:"cluster_name"`
	LoadBalancerName string            `json:"loadbalancer_name"`
	CLusterID        string            `json:"cluster_id"`
	NodeID           []string          `json:"node_id"`
	Annotations      []string			`json:"annotations"`
	PortMap          []PortMapping     `json:"port_map"`
}
type CreateLoadBalancerResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type UpdateLoadBalancersArgs struct {
	ClusterName      string            `json:"cluster_name"`
	LoadBalancerName string            `json:"loadbalancer_name"`
	CLusterID        string            `json:"cluster_id"`
	NodeID           []string          `json:"node_id"`
	Annotations      []string `json:"annotations"`
	PortMap          []PortMapping     `json:"port_map"`
}
type UpdateLoadBalancerResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DeleteLoadBalancersArgs struct {
	ClusterName      string `json:"cluster_name"`
	CLusterID        string `json:"cluster_id"`
	LoadBalancerName string `json:"loadbalancer_name"`
}
type DeleteLoadBalancersResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DescribeLoadBalancersTaskResultArgs struct {
	TaskID string`json:"task_id"`
}
type DescribeLoadBalancersTaskResultResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	}`json:"Data"`
}
