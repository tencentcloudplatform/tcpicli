package cvm

import (
	"encoding/json"
)

type DescribeInstancesStatusResp struct {
	Response struct {
		Error             interface{} `json:"Error,omitempty"`
		InstanceStatusSet []struct {
			InstanceID    string `json:"InstanceId"`
			InstanceState string `json:"InstanceState"`
		} `json:"InstanceStatusSet,omitempty"`
		RequestID  string `json:"RequestId"`
		TotalCount int    `json:"TotalCount"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9389
func DescribeInstancesStatus(options ...string) (*DescribeInstancesStatusResp, error) {
	resp, err := DoAction("DescribeInstancesStatus", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeInstancesStatusResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
