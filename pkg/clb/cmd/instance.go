package cmd

import "time"

type DescribeInstancesArgs struct {
	Version     string    `qcloud_arg:"Version,required"`
	InstanceIds *[]string `qcloud_arg:"InstanceIds"`
	//LanIps      *[]string `qcloud_arg:"lanIps"`
	Filters     *[]Filter `qcloud_arg:"Filters"`
	Offset      *int      `qcloud_arg:"Offset"`
	Limit       *int      `qcloud_arg:"Limit"`
}

type Filter struct {
	Name   string        `qcloud_arg:"Name"`
	Values []interface{} `qcloud_arg:"Values"`
}

type CvmResponse struct {
	Response interface{} `json:"Response"`
}

type DescribeInstancesResponse struct {
	TotalCount  int            `json:"TotalCount"`
	InstanceSet []InstanceInfo `json:"InstanceSet"`
	RequestID   string         `json:"RequestId"`
}

type Placement struct {
	Zone      string      `json:"Zone"`
	HostID    interface{} `json:"HostId"`
	ProjectID int         `json:"ProjectId"`
}

type Disk struct {
	DiskType string `json:"DiskType"`
	DiskID   string `json:"DiskId"`
	DiskSize int    `json:"DiskSize"`
}

type InternetAccessible struct {
	InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut"`
	InternetChargeType      string `json:"InternetChargeType"`
}

type VirtualPrivateCloud struct {
	VpcID        string `json:"VpcId"`
	SubnetID     string `json:"SubnetId"`
	AsVpcGateway bool   `json:"AsVpcGateway"`
}

type InstanceInfo struct {
	InstanceID         string   `json:"InstanceId"`
	InstanceType       string   `json:"InstanceType"`
	CPU                int      `json:"CPU"`
	Memory             int      `json:"Memory"`
	InstanceName       string   `json:"InstanceName"`
	InstanceChargeType string   `json:"InstanceChargeType"`
	PrivateIPAddresses []string `json:"PrivateIpAddresses"`
	PublicIPAddresses  []string `json:"PublicIpAddresses"`
	ImageID            string   `json:"ImageId"`
	OsName             string   `json:"OsName"`
	RenewFlag          string   `json:"RenewFlag"`

	Placement           Placement           `json:"Placement"`
	SystemDisk          Disk                `json:"SystemDisk"`
	DataDisks           []Disk              `json:"DataDisks"`
	InternetAccessible  InternetAccessible  `json:"InternetAccessible"`
	VirtualPrivateCloud VirtualPrivateCloud `json:"VirtualPrivateCloud"`

	CreatedTime time.Time `json:"CreatedTime"`
	ExpiredTime time.Time `json:"ExpiredTime"`
}
