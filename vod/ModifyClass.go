// 2017-12-01 12:07:54.213306 -0800 PST m=+5.682849474
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vod

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyClassResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/266/7815
func ModifyClass(options ...string) (*ModifyClassResp, error) {
	resp, err := DoAction("ModifyClass", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyClassResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *ModifyClassResp) String(args ...interface{}) (string, error) {
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
