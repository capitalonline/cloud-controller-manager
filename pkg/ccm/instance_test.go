package ccm

import (
	"strings"
	"testing"
)

func Test_getNodeInstanceTypeAndNodeNameByProviderID (t *testing.T) {

	// params
	fakeClusterID := ""
	fakeProviderID := ""

	// expected value
	expectedNodeName := ""

	// func
	res, err := getNodeInstanceTypeAndNodeNameByProviderID(fakeClusterID, fakeProviderID)

	// verify nodeName
	actualNodeName := res.Data.NodeName
	if strings.Compare(actualNodeName, expectedNodeName) != 0 {
		t.Errorf("unexpected NodeName got: %s want: %s", actualNodeName, expectedNodeName)
	}

	// verify labels
	actualLabels := res.Data.Labels
	t.Logf("actualLabels is: %s", actualLabels)

	// verify taints
	actualTaints := res.Data.Taints
	t.Logf("actualTaints is: %s", actualTaints)

	// err
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %v", err)
	}

}

func Test_describeInstanceExistsByProviderID (t *testing.T) {
	// params
	fakeProviderID := ""

	// func
	actualStatus, err := describeInstanceExistsByProviderID(fakeProviderID)

	// verify
	if actualStatus != true {
		t.Errorf("unexpected status got: %v want: true", actualStatus)
	}

	// err
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %v", err)
	}
}
