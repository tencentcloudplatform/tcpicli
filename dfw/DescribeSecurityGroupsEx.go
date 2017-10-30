package dfw

import (
	"encoding/json"
)

type DescribeSecurityGroupExResp struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Detail []struct {
			BeAssociateCount int    `json:"beAssociateCount"`
			CreateTime       string `json:"createTime"`
			ProjectID        string `json:"projectId"`
			SgID             string `json:"sgId"`
			SgName           string `json:"sgName"`
			SgRemark         string `json:"sgRemark"`
		} `json:"detail"`
		TotalNum int `json:"totalNum"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/1232
func DescribeSecurityGroupEx(options ...string) (*DescribeSecurityGroupExResp, error) {
	resp, err := DoAction("DescribeSecurityGroupEx", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeSecurityGroupExResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
