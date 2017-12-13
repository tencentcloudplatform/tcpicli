// 2017-12-12 20:49:28.938818845 +0000 UTC m=+86.934868653
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cdn

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type GetCdnStatusCodeResp struct {
	Code     int64       `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/228/3943
func (c *CdnClient) GetCdnStatusCode(options ...string) (*GetCdnStatusCodeResp, error) {
	resp, err := c.DoAction("GetCdnStatusCode", options...)
	if err != nil {
		return nil, err
	}
	var s GetCdnStatusCodeResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func GetCdnStatusCode(options ...string) (*GetCdnStatusCodeResp, error) {
	return DefaultClient.GetCdnStatusCode(options...)
}

func (r *GetCdnStatusCodeResp) String(args ...interface{}) (string, error) {
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
