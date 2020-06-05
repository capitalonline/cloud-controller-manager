package common

import (
	"os"
)

const (
	defaultApiHost         = "http://cdsapi.capitalonline.net"
	apiHostLiteral         = "API_HOST"
	accessKeyIdLiteral     = "ACCESS_KEY_ID"
	accessKeySecretLiteral = "ACCESS_KEY_SECRET"
	cckProductType         = "cck"
	version                = "2019-08-08"
	signatureVersion       = "1.0"
	signatureMethod        = "HMAC-SHA1"
	timeStampFormat        = "2006-01-02T15:04:05Z"
)

const (
	// load balancer
	ActionDescribeLoadBalancers = "DescribeLoadBalancers"
	ActionCreateLoadBalancers = "CreateLoadBalancers"
	ActionDeleteLoadBalancers = "DeleteLoadBalancers"
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

	if APIHost == "" {
		APIHost = defaultApiHost
	}
}
