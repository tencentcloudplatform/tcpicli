// 2018-01-31 15:15:20.234533 -0800 PST m=+1.251573107
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package trade

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeUserInfoResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	UserInfo struct {
		IsOwner    string `json:"isOwner"`
		Mail       string `json:"mail"`
		MailStatus string `json:"mailStatus"`
		Name       string `json:"name"`
		Phone      string `json:"phone"`
	} `json:"userInfo"`
}

// Implement https://cloud.tencent.com/document/api/378/4397
func (c *TradeClient) DescribeUserInfo(options ...string) (*DescribeUserInfoResp, error) {
	resp, err := c.DoAction("DescribeUserInfo", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeUserInfoResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeUserInfo(options ...string) (*DescribeUserInfoResp, error) {
	return DefaultClient.DescribeUserInfo(options...)
}

func (r *DescribeUserInfoResp) String(args ...interface{}) (string, error) {
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
