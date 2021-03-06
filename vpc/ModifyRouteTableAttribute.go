// 2017-11-17 03:27:38.631831 -0800 PST m=+37.122327178
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyRouteTableAttributeResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/1417
func ModifyRouteTableAttribute(options ...string) (*ModifyRouteTableAttributeResp, error) {
	resp, err := DoAction("ModifyRouteTableAttribute", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyRouteTableAttributeResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *ModifyRouteTableAttributeResp) String(args ...interface{}) (string, error) {
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
