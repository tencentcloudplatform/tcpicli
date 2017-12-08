// 2017-12-07 10:22:57.879276 -0800 PST m=+11.898639020
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyInstancesAttributeResp struct {
	Response struct {
		RequestID string      `json:"RequestId"`
		Error     interface{} `json:"Error,omitempty"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9381
func ModifyInstancesAttribute(options ...string) (*ModifyInstancesAttributeResp, error) {
	resp, err := DoAction("ModifyInstancesAttribute", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyInstancesAttributeResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *ModifyInstancesAttributeResp) String(args ...interface{}) (string, error) {
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