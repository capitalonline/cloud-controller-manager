package cmd

import "k8s.io/api/core/v1"

type LoadBalancer struct {
	Data struct {
		Status 	string `json:"status"`
		Name	string `json:"name"`
		Vips	[]string `json:"ha_ip"`
	} `json:"data"`
}

type ModifyLoadBalancerAttributesArgs struct {
	LoadBalancerId   string  `qcloud_arg:"loadBalancerId,required"`
	LoadBalancerName *string `qcloud_arg:"loadBalancerName"`
	DomainPrefix     *string `qcloud_arg:"domainPrefix"`
}

type ModifyLoadBalancerAttributesResponse struct {
	Response
	RequestId int `json:"requestId"`
}

type InquiryLBPriceArgs struct {
	LoadBalancerType int `qcloud_arg:"loadBalancerType,required"`
}

type InquiryLBPriceResponse struct {
	Response
	Price int `json:"price"`
}

type CreateLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name"`
	Service 			*v1.Service`json:"service"`
	Nodes 				[]*v1.Node`json:"nodes"`
}
type CreateLoadBalancerResponse struct {
	Response
	TotalCount      int            `json:"totalCount"`
	LoadBalancerSet []LoadBalancer `json:"loadBalancerSet"`
}

type UpdateLoadBalancersArgs struct {
	ClusterName			string`json:"cluster_name"`
	LoadBalancerName 	string`json:"loadBalancer_name"`
	Service 			*v1.Service`json:"service"`
	Nodes 				[]*v1.Node`json:"nodes"`
}
type UpdateLoadBalancerResponse struct {
	Response
	TotalCount      int            `json:"totalCount"`
	LoadBalancerSet []LoadBalancer `json:"loadBalancerSet"`
}
