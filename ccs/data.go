package ccs

import ()

// DescribeClusterInstances
type RespDescribeClusterInstances struct {
	Code     int                          `json:"code"`
	CodeDesc string                       `json:"codeDesc"`
	Data     DescribeClusterInstancesData `json:"data"`
	Message  string                       `json:"message"`
}
type DescribeClusterInstancesData struct {
	Nodes      []DescribeClusterInstancesNodes `json:"message"`
	TotalCount int                             `json:"totalCount"`
}
type DescribeClusterInstancesNodes struct {
	AbnormalReason       string `json:"abnormalReason"`
	Cpu                  int    `json:"cpu"`
	CreatedAt            string `json:"createdAt"`
	CvmPayMode           int    `json:"cvmPayMode"`
	CvmState             int    `json:"cvmState"`
	InstanceCreateTime   string `json:"instanceCreateTime"`
	InstanceDeadlineTime string `json:"instanceDeadlineTime"`
	InstanceId           string `json:"instanceId"`
	InstanceName         string `json:"instanceName"`
	InstanceType         string `json:"instanceType"`
	IsNormal             int    `json:"isNormal"`
	KernelVersion        string `json:"kernelVersion"`
	LanIp                string `json:"lanIp"`
	Mem                  int    `json:"mem"`
	NetworkPayMode       int    `json:"networkPayMode"`
	OsImage              string `json:"osImage"`
	PodCidr              string `json:"podCidr"`
	Unschedulable        bool   `json:"unschedulable"`
	WanIp                string `json:"wanIp"`
	Zone                 string `json:"zone"`
	ZoneId               string `json:"zoneId"`
}

// CreateCluster
type RespCreateCluster struct {
	Code     int                   `json:"code"`
	CodeDesc string                `json:"codedesc"`
	Data     RespCreateClusterData `json:"data"`
	Message  string                `json:"message"`
}
type RespCreateClusterData struct {
	ClusterId string `json:"clusterId"`
	RequestId int    `json:"requestId"`
}

// DeleteCluster
type RespDeleteCluster struct {
	Code     int                   `json:"code"`
	CodeDesc string                `json:"codedesc"`
	Data     RespDeleteClusterData `json:"data"`
}
type RespDeleteClusterData struct {
	RequestId int `json:"requestId"`
}

// AddClusterInstances
type RespAddClusterInstances struct {
	Code     int                         `json:"code"`
	CodeDesc string                      `json:"codeDesc"`
	Data     RespAddClusterInstancesData `json:"data"`
	Message  string                      `json:"message"`
}
type RespAddClusterInstancesData struct {
	InstanceIds []string `json:"instanceIds"`
	RequestId   int      `json:"requestId"`
}
