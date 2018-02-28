// 2018-01-29 16:57:24.482042 -0800 PST m=+8.009254516
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ckafka

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type SetTopicAttributesResp struct {
	Code     int64         `json:"code"`
	CodeDesc string        `json:"codeDesc"`
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/597/10098
func (c *CkafkaClient) SetTopicAttributes(options ...string) (*SetTopicAttributesResp, error) {
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
