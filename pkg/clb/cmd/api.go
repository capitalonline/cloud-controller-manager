package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cloud-controller-manager/pkg/clb/common"
	"io/ioutil"
	"net/http"
)

func DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDescribeLoadBalancers, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeLoadBalancersResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func CreateLoadBalancers(args *CreateLoadBalancersArgs) (*CreateLoadBalancerResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateLoadBalancers, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &CreateLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func UpdateLoadBalancers(args *UpdateLoadBalancersArgs) (*UpdateLoadBalancerResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateLoadBalancers, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &UpdateLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DeleteLoadBalancers(args *DeleteLoadBalancersArgs) (*DeleteLoadBalancersResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDeleteLoadBalancers, http.MethodDelete, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeleteLoadBalancersResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DescribeLoadBalancersTaskResult(args *DescribeLoadBalancersTaskResultArgs) (*DescribeLoadBalancersTaskResultResponse, error) {
	params := map[string]string {
		"task_id": args.RequestId,
	}
	req, err := common.NewCCKRequest(common.ActionDeleteLoadBalancers, http.MethodDelete, params, nil)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeLoadBalancersTaskResultResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DescribeInstances(args *DescribeInstancesArgs) (*DescribeInstancesResponse, error) {
	params := map[string]string {
		"task_id": args.RequesetId,
	}
	req, err := common.NewCCKRequest(common.ActionDeleteLoadBalancers, http.MethodDelete, params, nil)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeInstancesResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func RegisterInstancesWithLoadBalancer(args *RegisterInstancesWithLoadBalancerArgs)(*RegisterInstancesWithLoadBalancerResponse, error) {
	params := map[string]string {
		"cluster_name": args.ClusterName,
		"name": args.LoadBalancerName,
		"backends": args.Backends,
	}
	req, err := common.NewCCKRequest(common.ActionDeleteLoadBalancers, http.MethodDelete, params, nil)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &RegisterInstancesWithLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DeRegisterInstancesWithLoadBalancer(args *DeRegisterInstancesWithLoadBalancerArgs)(*DeRegisterInstancesWithLoadBalancerResponse, error) {
	params := map[string]string {
		"cluster_name": args.ClusterName,
		"name": args.LoadBalancerName,
		"backends": args.Backends,
	}
	req, err := common.NewCCKRequest(common.ActionDeleteLoadBalancers, http.MethodDelete, params, nil)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeRegisterInstancesWithLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}