package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type AssociateInstancesKeyPairsResp struct {
	Response struct {
		RequestID string      `json:"RequestId"`
		Error     interface{} `json:"Error,omitempty"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9402
func AssociateInstancesKeyPairs(options ...string) (*AssociateInstancesKeyPairsResp, error) {
	resp, err := DoAction("AssociateInstancesKeyPairs", options...)
	if err != nil {
		return nil, err
	}
	var s AssociateInstancesKeyPairsResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *AssociateInstancesKeyPairsResp) String(args ...interface{}) (string, error) {
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
