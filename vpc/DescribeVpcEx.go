// 2018-01-09 09:22:06.013079 -0800 PST m=+7.338756979
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeVpcExResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     []struct {
		CidrBlock      string `json:"cidrBlock"`
		ClassicLinkNum int64  `json:"classicLinkNum"`
		CreateTime     string `json:"createTime"`
		IsDefault      bool   `json:"isDefault"`
		IsMulticast    bool   `json:"isMulticast"`
		NatNum         int64  `json:"natNum"`
		RouteTableNum  int64  `json:"routeTableNum"`
		SubnetNum      int64  `json:"subnetNum"`
		UnVpcID        string `json:"unVpcId"`
		VpcDeviceNum   int64  `json:"vpcDeviceNum"`
		VpcID          string `json:"vpcId"`
		VpcName        string `json:"vpcName"`
		VpcPeerNum     int64  `json:"vpcPeerNum"`
		VpgNum         int64  `json:"vpgNum"`
		VpnGwNum       int64  `json:"vpnGwNum"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

// Implement https://cloud.tencent.com/document/api/215/1372
func (c *VpcClient) DescribeVpcEx(options ...string) (*DescribeVpcExResp, error) {
	resp, err := c.DoAction("DescribeVpcEx", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeVpcExResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeVpcEx(options ...string) (*DescribeVpcExResp, error) {
	return DefaultClient.DescribeVpcEx(options...)
}

func (r *DescribeVpcExResp) String(args ...interface{}) (string, error) {
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
