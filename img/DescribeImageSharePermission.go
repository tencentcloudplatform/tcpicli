package img

import (
	"encoding/json"
)

type DescribeImageSharePermissionResp struct {
	Response struct {
		Error               interface{}   `json:"Error,omitempty"`
		RequestID           string        `json:"RequestId"`
		SharedPermissionSet []interface{} `json:"SharedPermissionSet,omitempty"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9419
func DescribeImageSharePermission(options ...string) (*DescribeImageSharePermissionResp, error) {
	resp, err := DoAction("DescribeImageSharePermission", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeImageSharePermissionResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
