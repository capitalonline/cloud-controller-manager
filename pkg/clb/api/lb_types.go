package api

import (
	"k8s.io/api/core/v1"
)

type DescribeLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	CLusterID 	 		string`json:"cluster_id"`
	LoadBalancerName 	string`json:"lb_name"`
}
type DescribeLoadBalancersResponse struct {
	Response
	Data struct {
		Status 	string`json:"status"`
		Name	string`json:"name"`
		Vips	[]string`json:"ha_ip"`
	} `json:"data"`
}

type PortMapping struct {
	Protocol v1.Protocol
	Port     int32
	Nodeport int32
}
type CreateLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
	CLusterID 	 		string`json:"cluster_id"`
	NodeID 				[]string`json:"node_id"`
	Annotations			map[string]string`json:"annotations"`
	PortMap 			[]PortMapping`json:"port_map"`
}
type CreateLoadBalancerResponse struct {
	Response
	Data struct {
		TaskID 	string`json:"task_id"`
	}`json:"data"`
}

type UpdateLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
	CLusterID 	 		string`json:"cluster_id"`
	NodeID 				[]string`json:"node_id"`
	Annotations			map[string]string`json:"annotations"`
	PortMap 			[]PortMapping`json:"port_map"`
}
type UpdateLoadBalancerResponse struct {
	Response
	Data struct {
		TaskID 	string`json:"task_id"`
	}`json:"data"`
}

type DeleteLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	CLusterID 	 		string`json:"cluster_id"`
	LoadBalancerName 	string`json:"lb_name"`
}
type DeleteLoadBalancersResponse struct {
	Response
	Data struct {
		TaskID 	string`json:"task_id"`
	}`json:"data"`
}

type DescribeLoadBalancersTaskResultArgs struct {
	TaskID 	string`json:"task_id"`
}
type DescribeLoadBalancersTaskResultResponse struct {
	Response
	Data struct {
		Status 	string`json:"status"`
	}
}