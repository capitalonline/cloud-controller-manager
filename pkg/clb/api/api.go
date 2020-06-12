package api

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cloud-controller-manager/pkg/clb/common"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResponse, error) {
	log.Infof("api:: DescribeHaproxyLoadBalancerInstance")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDescribeHaproxyLoadBalancerInstance, http.MethodGet, nil, body)
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
	log.Infof("api:: CreateHaproxyLoadBalancerInstance")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
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
	log.Infof("api:: UpdateHaproxyLoadBalancerInstance")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionUpdateHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
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
	log.Infof("api:: DeleteHaproxyLoadBalancerInstance")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDeleteHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
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
	log.Infof("api:: CheckHaproxyLoadbalancerTaskStatus")
	params := map[string]string{
		"task_id": args.TaskID,
	}
	req, err := common.NewCCKRequest(common.ActionCheckHaproxyLoadbalancerTaskStatus, http.MethodGet, params, nil)
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

func DescribeInstancesLabelsAndNodeName(args *DescribeInstancesLabelsAndNodeNameArgs) (*DescribeInstancesLabelsAndNodeNameResponse, error) {
	log.Infof("api:: DescribeInstancesLabelsAndNodeName")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateHaproxyLoadBalancerInstance, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeInstancesLabelsAndNodeNameResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DescribeZoneByProviderID(args *DescribeZoneByProviderIDArgs) (*DescribeZoneByProviderIDResponse, error) {
	log.Infof("api:: DescribeZoneByProviderID")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateHaproxyLoadBalancerInstance, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeZoneByProviderIDResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}
