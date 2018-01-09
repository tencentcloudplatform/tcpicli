// 2018-01-09 09:22:25.414564 -0800 PST m=+26.740278073
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeVpcPeeringConnectionsResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     []struct {
		AppID                 string `json:"appId"`
		Bandwidth             int64  `json:"bandwidth"`
		CreateTime            string `json:"createTime"`
		PeerAppID             string `json:"peerAppId"`
		PeerRegion            string `json:"peerRegion"`
		PeerUin               int64  `json:"peerUin"`
		PeerVpcID             string `json:"peerVpcId"`
		PeeringConnectionID   string `json:"peeringConnectionId"`
		PeeringConnectionName string `json:"peeringConnectionName"`
		Region                string `json:"region"`
		State                 int64  `json:"state"`
		Type                  int64  `json:"type"`
		Uin                   int64  `json:"uin"`
		UnPeerVpcID           string `json:"unPeerVpcId"`
		UnVpcID               string `json:"unVpcId"`
		UniqPeerVpcID         string `json:"uniqPeerVpcId"`
		UniqVpcID             string `json:"uniqVpcId"`
		VpcID                 string `json:"vpcId"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

// Implement https://cloud.tencent.com/document/api/215/2101
func (c *VpcClient) DescribeVpcPeeringConnections(options ...string) (*DescribeVpcPeeringConnectionsResp, error) {
	resp, err := c.DoAction("DescribeVpcPeeringConnections", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeVpcPeeringConnectionsResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeVpcPeeringConnections(options ...string) (*DescribeVpcPeeringConnectionsResp, error) {
	return DefaultClient.DescribeVpcPeeringConnections(options...)
}

func (r *DescribeVpcPeeringConnectionsResp) String(args ...interface{}) (string, error) {
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
