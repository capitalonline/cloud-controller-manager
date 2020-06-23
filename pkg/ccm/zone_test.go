package ccm

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"reflect"
	"strings"
	"testing"

	cloudprovider "k8s.io/cloud-provider"

)

func Test_GetZoneByProviderID (t *testing.T) {
	// params
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("newCloud:: Failed to create kubernetes config: %v", err)
	}
	k8sClientSet, err := kubernetes.NewForConfig(config)

	fakeClusterID := ""
	fakeRegion := "test1"
	fakeProviderID := ""
	resources := newResources(fakeClusterID)
	fakeZones := newZones(resources, k8sClientSet, fakeRegion)

	// expected value
	expected := cloudprovider.Zone{Region: "test1"}

	// func
	actual, err := fakeZones.GetZoneByProviderID(context.TODO(), fakeProviderID)

	// verify region
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected region. got: %+v want: %+v", actual, expected)
	}

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %v", err)
	}

}

func Test_getZoneByProviderID (t *testing.T) {
	// params
	fakeClusterID := ""
	fakeProviderID := ""

	// expected value
	expected := ""

	// func
	res, err := getZoneByProviderID(fakeClusterID, fakeProviderID)

	// verify region
	actual := res.Data.Region
	if strings.Compare(actual, expected) != 0 {
		t.Errorf("unexpected region got: %s want: %s", actual, expected)
	}

	// verify
	if err != nil {
		t.Errorf("unexpected err, expected nil. got: %v", err)
	}

}