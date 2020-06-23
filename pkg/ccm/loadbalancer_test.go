package ccm

import (
	"context"
	"k8s.io/api/core/v1"
	"strings"
	"testing"
)


func Test_getLoadBalancerByName(t *testing.T) {
	// params
	fakeClusterName := "kubernetes"
	fakeClusterID := ""
	fakeLoadBalancerName := ""

	// expected value
	expected := ""

	// func
	res, err := getLoadBalancerByName(fakeClusterName, fakeClusterID, fakeLoadBalancerName)

	// verify loadBalancerName
	actual := res.Data.Name
	if strings.Compare(actual, expected) != 0 {
		t.Errorf("unexpected loadBalancerName got: %s want: %s", actual, expected)
	}

	// err
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %s", err)
	}

}

func Test_createClassicLoadBalancer(t *testing.T) {
	// params
	fakeClsuterName := "kubernetes"
	fakeClusterID := ""
	fakeLoadBalancerName := ""
	fakeProviderID := ""
	var service  *v1.Service
	var nodes 	 []*v1.Node
	var servicePorts v1.ServicePort
	var serviceSlice []v1.ServicePort
	portTmp := [] int32 {22, 23, 24}

	// init params
	for _, port := range portTmp {
		servicePorts.Port = port
		servicePorts.NodePort = 300 + port
		servicePorts.Protocol = "TCP"
		serviceSlice = append(serviceSlice, servicePorts)
	}
	service.Spec.Ports = serviceSlice

	for _, node := range nodes {
		node.Spec.ProviderID = fakeProviderID
	}

	// func
	err := updateClassicLoadBalancer(context.TODO(), fakeClsuterName, service, nodes, fakeClusterID, fakeLoadBalancerName)

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %s", err)
	}
}

func Test_updateClassicLoadBalancer(t *testing.T) {
	// params
	fakeClsuterName := "kubernetes"
	fakeClusterID := ""
	fakeLoadBalancerName := ""
	fakeProviderID := ""
	var service  *v1.Service
	var nodes 	 []*v1.Node
	var servicePorts v1.ServicePort
	var serviceSlice []v1.ServicePort
	portTmp := [] int32 {22, 23, 24}

	// init params
	for _, port := range portTmp {
		servicePorts.Port = port
		servicePorts.NodePort = 300 + port
		servicePorts.Protocol = "TCP"
		serviceSlice = append(serviceSlice, servicePorts)
	}
	service.Spec.Ports = serviceSlice

	for _, node := range nodes {
		node.Spec.ProviderID = fakeProviderID
	}

	// func
	err := updateClassicLoadBalancer(context.TODO(), fakeClsuterName, service, nodes, fakeClusterID, fakeLoadBalancerName)

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %s", err)
	}

}

func Test_deleteLoadBalancer(t *testing.T) {
	// params
	fakeClsuterName := "kubernetes"
	fakeClusterID := ""
	fakeLoadBalancerName := ""

	// func
	err := deleteLoadBalancer(context.TODO(), fakeClsuterName, fakeClusterID, fakeLoadBalancerName)

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %s", err)
	}
}

func Test_describeLoadBalancersTaskResult(t *testing.T) {
	// params
	fakeTaskID := ""

	// func
	err := describeLoadBalancersTaskResult(fakeTaskID)

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %s", err)
	}

}
