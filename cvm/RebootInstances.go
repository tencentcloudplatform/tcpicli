// 2017-12-06 15:39:03.730105 -0800 PST m=+10.527196469
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type RebootInstancesResp struct {
	Response struct {
		RequestID string      `json:"RequestId"`
		Error     interface{} `json:"Error,omitempty"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9369
func RebootInstances(options ...string) (*RebootInstancesResp, error) {
	resp, err := DoAction("RebootInstances", options...)
	if err != nil {
		return nil, err
	}
	var s RebootInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *RebootInstancesResp) String(args ...interface{}) (string, error) {
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
