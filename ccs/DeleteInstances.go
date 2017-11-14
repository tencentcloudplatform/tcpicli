// 2017-11-12 22:22:51.509253 +0800 CST m=+16.123388983
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ccs

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteInstancesResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		FailedMap   []interface{} `json:"failedMap"`
		SuccessList []string      `json:"successList"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9432
func DeleteInstances(options ...string) (*DeleteInstancesResp, error) {
	resp, err := DoAction("DeleteInstances", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DeleteInstancesResp) String(args ...interface{}) (string, error) {
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
