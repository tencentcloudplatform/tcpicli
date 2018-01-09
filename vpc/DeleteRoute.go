// 2018-01-09 09:23:43.250266 -0800 PST m=+104.576124825
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteRouteResp struct {
	Code     int64         `json:"code"`
	CodeDesc string        `json:"codeDesc"`
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5689
func (c *VpcClient) DeleteRoute(options ...string) (*DeleteRouteResp, error) {
	resp, err := c.DoAction("DeleteRoute", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteRouteResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DeleteRoute(options ...string) (*DeleteRouteResp, error) {
	return DefaultClient.DeleteRoute(options...)
}

func (r *DeleteRouteResp) String(args ...interface{}) (string, error) {
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
