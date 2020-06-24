package common

import (
	"os"
)

const (
	defaultApiHost         = "http://cdsapi.capitalonline.net"
	defaultApiHostOversea  = "http://cdsapi-us.capitalonline.net"
	apiHostLiteral         = "CDS_API_HOST"
	accessKeyIdLiteral     = "CDS_ACCESS_KEY_ID"
	accessKeySecretLiteral = "CDS_ACCESS_KEY_SECRET"
	overseaFlag			   = "CDS_OVERSEA"
	cckProductType         = "cck"
	version                = "2019-08-08"
	signatureVersion       = "1.0"
	signatureMethod        = "HMAC-SHA1"
	timeStampFormat        = "2006-01-02T15:04:05Z"
)

const (
	// load balancer
	ActionDescribeHaproxyLoadBalancerInstance = "DescribeHaproxyLoadbalancerInstance"
	ActionCreateHaproxyLoadBalancerInstance   = "CreateHaproxyLoadbalancerInstance"
	ActionUpdateHaproxyLoadBalancerInstance   = "UpdateHaproxyLoadbalancerInstance"
	ActionDeleteHaproxyLoadBalancerInstance   = "DeleteHaproxyLoadbalancerInstance"
	ActionCheckHaproxyLoadBalancerTaskStatus  = "CheckHaproxyLoadbalancerTaskStatus"

	// instance
	ActionDescribeInstancesLabelsAndNodeName = "DescribeInstancesLabelsAndNodeName"
	ActionDescribeNodeInstanceStatus = "DescribeNodeInstanceStatus"
)

var (
	APIHost         string
	AccessKeyID     string
	AccessKeySecret string
)

func IsAccessKeySet() bool {
	return AccessKeyID != "" && AccessKeySecret != ""
}

func init() {
	if APIHost == "" {
		APIHost = os.Getenv(apiHostLiteral)
	}
	if AccessKeyID == "" {
		AccessKeyID = os.Getenv(accessKeyIdLiteral)
	}
	if AccessKeySecret == "" {
		AccessKeySecret = os.Getenv(accessKeySecretLiteral)
	}

	// True is oversea cluster; False is domestic cluster
	if os.Getenv(overseaFlag) == "True" &&  APIHost == "" {
		APIHost = defaultApiHostOversea
	} else if os.Getenv(overseaFlag) == "False" &&  APIHost == "" {
		APIHost = defaultApiHost
	}

}
