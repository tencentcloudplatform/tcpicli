// 2018-05-23 14:32:48.571095967 -0700 PDT m=+1.651219862
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeRegionsResp struct {
	Response struct {
		RegionSet []struct {
			Region      string `json:"Region"`
			RegionName  string `json:"RegionName"`
			RegionState string `json:"RegionState"`
		} `json:"RegionSet"`
		RequestID  string `json:"RequestId"`
		TotalCount int64  `json:"TotalCount"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/15708
func (c *CvmClient) DescribeRegions(options ...string) (*DescribeRegionsResp, error) {
	resp, err := c.DoAction("DescribeRegions", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeRegionsResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeRegions(options ...string) (*DescribeRegionsResp, error) {
	return DefaultClient.DescribeRegions(options...)
}

func (r *DescribeRegionsResp) String(args ...interface{}) (string, error) {
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
