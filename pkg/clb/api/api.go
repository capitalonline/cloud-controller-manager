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
	log.Infof("api:: DescribeLoadBalancers, args is: +v", args)
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
	log.Infof("api:: CreateLoadBalancers, args is: %+v", args)
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
	log.Infof("api:: UpdateLoadBalancers, args is: %+v", args)
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
	log.Infof("api:: DeleteLoadBalancers, args is: %+v", args)
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
