// 2018-02-08 13:05:54.151316 -0800 PST m=+27.614547121
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type SubscribeResp struct {
	Code           float64 `json:"code"`
	Message        string  `json:"message"`
	RequestID      string  `json:"requestId"`
	SubscriptionID string  `json:"subscriptionId"`
}

// Implement https://cloud.tencent.com/document/api/406/7414
func (c *CmqClient) Subscribe(options ...string) (*SubscribeResp, error) {
	resp, err := c.DoAction("Subscribe", options...)
	if err != nil {
		return nil, err
	}
	var s SubscribeResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func Subscribe(options ...string) (*SubscribeResp, error) {
	return DefaultClient.Subscribe(options...)
}

func (r *SubscribeResp) String(args ...interface{}) (string, error) {
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
