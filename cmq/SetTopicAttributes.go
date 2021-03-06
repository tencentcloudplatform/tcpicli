// 2018-02-08 13:05:53.53239 -0800 PST m=+26.995628724
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type SetTopicAttributesResp struct {
	Code      float64 `json:"code"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/7406
func (c *CmqClient) SetTopicAttributes(options ...string) (*SetTopicAttributesResp, error) {
	resp, err := c.DoAction("SetTopicAttributes", options...)
	if err != nil {
		return nil, err
	}
	var s SetTopicAttributesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func SetTopicAttributes(options ...string) (*SetTopicAttributesResp, error) {
	return DefaultClient.SetTopicAttributes(options...)
}

func (r *SetTopicAttributesResp) String(args ...interface{}) (string, error) {
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
