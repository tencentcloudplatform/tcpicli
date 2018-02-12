// 2018-02-08 13:05:59.666557 -0800 PST m=+33.129720573
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteTopicResp struct {
	Code      float64 `json:"code"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/7409
func (c *CmqClient) DeleteTopic(options ...string) (*DeleteTopicResp, error) {
	resp, err := c.DoAction("DeleteTopic", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteTopicResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DeleteTopic(options ...string) (*DeleteTopicResp, error) {
	return DefaultClient.DeleteTopic(options...)
}

func (r *DeleteTopicResp) String(args ...interface{}) (string, error) {
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
