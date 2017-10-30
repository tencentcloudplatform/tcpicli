package cvm

import (
	"encoding/json"
)

type RunInstancesResp struct {
	Response struct {
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

func RunInstances(options ...string) (*RunInstancesResp, error) {
	resp, err := DoAction("RunInstances", options...)
	if err != nil {
		return nil, err
	}
	var s RunInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

//tcpicli -vv cvm RunInstances Region=gz Version=2017-03-12 InstanceChargeType=POSTPAID_BY_HOUR InstanceType=S2.MEDIUM4 InternetAccessible.InternetChargeType=TRAFFIC_POSTPAID_BY_HOUR InternetAccessible.InternetMaxBandwidthOut=100 Placement.Zone=ap-guangzhou-2 Placement.ProjectId=1059582 ImageId=img-dkwyg6sr VirtualPrivateCloud.VpcId=vpc-d49ixki1 VirtualPrivateCloud.SubnetId=subnet-kxjambf4 InstanceName="createdbyApi" LoginSettings.KeyIds.0=skey-8b6ih5f5 SecurityGroupIds.0=sg-0xgkqyh5 DataDisks.0.DiskType=CLOUD_BASIC DataDisks.0.DiskSize=100 SystemDisk.DiskType=CLOUD_BASIC
