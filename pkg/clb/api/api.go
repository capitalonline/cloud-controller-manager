package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/capitalonline/cloud-controller-manager/pkg/clb/common"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ErrCloudLoadBalancerNotFound = errors.New("LoadBalancer not found")
)

func DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResponse, error) {
	log.Infof("api:: DescribeLoadBalancers")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDescribeHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		if strings.Contains(string(content), "DataNotExists") {
			return nil, ErrCloudLoadBalancerNotFound
		}
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	log.Infof("api:: content is: %s", content)
	res := &DescribeLoadBalancersResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func CreateLoadBalancers(args *CreateLoadBalancersArgs) (*CreateLoadBalancerResponse, error) {
	log.Infof("api:: CreateLoadBalancers")
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

	log.Infof("api:: content is: %s", content)
	res := &CreateLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func UpdateLoadBalancers(args *UpdateLoadBalancersArgs) (*UpdateLoadBalancerResponse, error) {
	log.Infof("api:: UpdateLoadBalancers")
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

	log.Infof("api:: content is: %s", content)
	res := &UpdateLoadBalancerResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DeleteLoadBalancers(args *DeleteLoadBalancersArgs) (*DeleteLoadBalancersResponse, error) {
	log.Infof("api:: DeleteLoadBalancers")
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

	log.Infof("api:: content is: %s", content)
	res := &DeleteLoadBalancersResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DescribeLoadBalancersTaskResult(args *DescribeLoadBalancersTaskResultArgs) (*DescribeLoadBalancersTaskResultResponse, error) {
	log.Infof("api:: DescribeLoadBalancersTaskResult")
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCheckHaproxyLoadBalancerTaskStatus, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	log.Infof("api:: content is: %s", content)
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
	req, err := common.NewCCKRequest(common.ActionCreateHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	log.Infof("api:: content is: %s", content)
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
	req, err := common.NewCCKRequest(common.ActionCreateHaproxyLoadBalancerInstance, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	log.Infof("api:: content is: %s", content)
	res := &DescribeZoneByProviderIDResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}
