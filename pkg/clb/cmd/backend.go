package cmd

type LoadBalancerBackends struct {
	InstanceId     string   `json:"instanceId"`
	UnInstanceId   string   `json:"unInstanceId"`
	Weight         int      `json:"weight"`
	InstanceName   string   `json:"instanceName"`
	LanIp          string   `json:"lanIp"`
	WanIpSet       []string `json:"wanIpSet"`
	InstanceStatus int      `json:"instanceStatus"`
}

type DescribeLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name"`
}
type DescribeLoadBalancersResponse struct {
	Response
	TotalCount      int            `json:"totalCount"`
	LoadBalancerSet []LoadBalancer `json:"loadBalancerSet"`
}

type DeleteLoadBalancersArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name"`
}
type DeleteLoadBalancersResponse struct {
	Response
	Data struct {
		Status 	string `json:"status"`
		Name	string `json:"name"`
		Vips	[]string `json:"ha_ip"`
	} `json:"data"`
}

type DescribeLoadBalancerBackendsArgs struct {
	ClusterName 		string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name"`
	Offset				int`json:"offset"`
	Limit				int`json:"limit"`
}
type DescribeLoadBalancerBackendsResponse struct {
	Response
	TotalCount int                    `json:"totalCount"`
	BackendSet []LoadBalancerBackends `json:"backendSet"`
}

type DescribeLoadBalancersTaskResultArgs struct {
	RequestId string`json:"request_id"`
}
type DescribeLoadBalancersTaskResultResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}

type RegisterInstancesOpts struct {
	InstanceId string `qcloud_arg:"instanceId,required"`
	Weight     *int   `qcloud_arg:"weight"`
}

type RegisterInstancesWithLoadBalancerArgs struct {
	ClusterName			string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name,required"`
	Backends       		[]RegisterInstancesOpts`json:"backends,required"`
}

type RegisterInstancesWithLoadBalancerResponse struct {
	Response
	Data struct {
		RequestId string`json:"request_id"`
	} `json:"data"`
}

type DeRegisterInstancesWithLoadBalancerArgs struct {
	ClusterName			string`json:"cluster_name"`
	LoadBalancerName 	string`json:"name,required"`
	Backends       		[]RegisterInstancesOpts`json:"backends,required"`
}

type DeRegisterInstancesWithLoadBalancerResponse struct {
	Response
	Data struct {
		RequestId string`json:"request_id"`
	} `json:"data"`
}