// 2018-01-09 09:22:00.463195 -0800 PST m=+1.788863307
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateVpcResp struct {
	Code          int64  `json:"code"`
	CodeDesc      string `json:"codeDesc"`
	Message       string `json:"message"`
	RouteTableSet []struct {
		RouteTableID   string `json:"routeTableId"`
		RouteTableName string `json:"routeTableName"`
		RouteTableType int64  `json:"routeTableType"`
	} `json:"routeTableSet"`
	SubnetSet     []interface{} `json:"subnetSet"`
	UniqVpcID     string        `json:"uniqVpcId"`
	VpcCreateTime string        `json:"vpcCreateTime"`
	VpcID         string        `json:"vpcId"`
}

// Implement https://cloud.tencent.com/document/api/215/1309
func (c *VpcClient) CreateVpc(options ...string) (*CreateVpcResp, error) {
	resp, err := c.DoAction("CreateVpc", options...)
	if err != nil {
		return nil, err
	}
	var s CreateVpcResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func CreateVpc(options ...string) (*CreateVpcResp, error) {
	return DefaultClient.CreateVpc(options...)
}

func (r *CreateVpcResp) String(args ...interface{}) (string, error) {
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
