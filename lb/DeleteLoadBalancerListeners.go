// 2017-12-09 00:52:41.46197 -0800 PST m=+25.947722870
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package lb

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteLoadBalancerListenersResp struct {
	Code      int64  `json:"code"`
	CodeDesc  string `json:"codeDesc"`
	Message   string `json:"message"`
	RequestID int64  `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/214/1256
func DeleteLoadBalancerListeners(options ...string) (*DeleteLoadBalancerListenersResp, error) {
	resp, err := DoAction("DeleteLoadBalancerListeners", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteLoadBalancerListenersResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DeleteLoadBalancerListenersResp) String(args ...interface{}) (string, error) {
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
