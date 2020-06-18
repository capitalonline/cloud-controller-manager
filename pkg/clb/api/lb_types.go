package api

type DescribeLoadBalancersArgs struct {
	ClusterID        string `json:"cluster_id"`
	ServiceName 	 string `json:"service_name"`
	ServiceNameSpace string `json:"service_name_space"`
	ServiceUid       string `json:"service_uid"`
}
type DescribeLoadBalancersResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
		Name   string `json:"name"`
		Vips   []string `json:"vips"`
	} `json:"Data"`
}

type PortMapping struct {
	// Protocol v1.Protocol `json:"protocol"`
	Port     int32 `json:"port"`
	NodePort int32 `json:"node_port"`
}
type CreateLoadBalancersArgs struct {
	ClusterID        string `json:"cluster_id"`
	NodeID           []string `json:"node_id"`
	Annotations      []string `json:"annotations"`
	PortMap          []PortMapping `json:"port_map"`
	ServiceName 	 string `json:"service_name"`
	ServiceNameSpace string `json:"service_name_space"`
	ServiceUid       string `json:"service_uid"`
}
type CreateLoadBalancerResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type UpdateLoadBalancersArgs struct {
	ClusterID        string `json:"cluster_id"`
	NodeID           []string `json:"node_id"`
	Annotations      []string `json:"annotations"`
	PortMap          []PortMapping `json:"port_map"`
	ServiceName 	 string `json:"service_name"`
	ServiceNameSpace string `json:"service_name_space"`
	ServiceUid       string `json:"service_uid"`
}
type UpdateLoadBalancerResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DeleteLoadBalancersArgs struct {
	ClusterID        string `json:"cluster_id"`
	ServiceName 	 string `json:"service_name"`
	ServiceNameSpace string `json:"service_name_space"`
	ServiceUid       string `json:"service_uid"`
}
type DeleteLoadBalancersResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DescribeLoadBalancersTaskResultArgs struct {
	TaskID string `json:"task_id"`
}
type DescribeLoadBalancersTaskResultResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	}`json:"Data"`
}
