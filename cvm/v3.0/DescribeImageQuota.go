// 2018-05-23 22:48:12.697417381 +0000 UTC m=+3.978832571
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeImageQuotaResp struct {
	Response struct {
		ImageNumQuota int64  `json:"ImageNumQuota"`
		RequestID     string `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/15719
func (c *CvmClient) DescribeImageQuota(options ...string) (*DescribeImageQuotaResp, error) {
	resp, err := c.DoAction("DescribeImageQuota", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeImageQuotaResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeImageQuota(options ...string) (*DescribeImageQuotaResp, error) {
	return DefaultClient.DescribeImageQuota(options...)
}

func (r *DescribeImageQuotaResp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
