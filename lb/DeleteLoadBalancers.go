// 2018-02-08 13:11:00.397837 -0800 PST m=+48.749782540
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package lb

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteLoadBalancersResp struct {
	Code      float64 `json:"code"`
	CodeDesc  string  `json:"codeDesc"`
	Message   string  `json:"message"`
	RequestID float64 `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/214/1257
func (c *LbClient) DeleteLoadBalancers(options ...string) (*DeleteLoadBalancersResp, error) {
	resp, err := c.DoAction("DeleteLoadBalancers", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteLoadBalancersResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DeleteLoadBalancers(options ...string) (*DeleteLoadBalancersResp, error) {
	return DefaultClient.DeleteLoadBalancers(options...)
}

func (r *DeleteLoadBalancersResp) String(args ...interface{}) (string, error) {
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
