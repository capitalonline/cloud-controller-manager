package api

import (
	"k8s.io/api/core/v1"
)

type LoadBalancer struct {
	Data struct {
		Status 	string`json:"status"`
		Name	string`json:"name"`
		Vips	[]string`json:"ha_ip"`
	} `json:"data"`
}

type DescribeLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
}
type DescribeLoadBalancersResponse struct {
	Response
	Data struct {
		Name 	string`json:"name"`
		Status 	string`json:"status"`
		Vips 	[]string`json:"ha_ip"`
	}`json:"data"`
}

type CreateLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
	Service 			*v1.Service`json:"service"`
	Nodes 				[]*v1.Node`json:"nodes"`
}
type CreateLoadBalancerResponse struct {
	Response
	Data struct {
		Name 	string`json:"name"`
		Status	string`json:"status"`
	}`json:"data"`
}

type UpdateLoadBalancersArgs struct {
	ClusterName			string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
	Service 			*v1.Service`json:"service"`
	Nodes 				[]*v1.Node`json:"nodes"`
}
type UpdateLoadBalancerResponse struct {
	Response
	Data struct {
		Name 	string`json:"name"`
		Status	string`json:"status"`
	}`json:"data"`
}

type DeleteLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"lb_name"`
}
type DeleteLoadBalancersResponse struct {
	Response
	Data struct { }`json:"data"`
}