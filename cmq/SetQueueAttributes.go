// 2018-02-08 13:05:28.643478 -0800 PST m=+2.107020591
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type SetQueueAttributesResp struct {
	ClientRequestID     float64 `json:"clientRequestId"`
	Code                float64 `json:"code"`
	MaxMsgHeapNum       float64 `json:"maxMsgHeapNum"`
	MaxMsgSize          float64 `json:"maxMsgSize"`
	Message             string  `json:"message"`
	MsgRetentionSeconds float64 `json:"msgRetentionSeconds"`
	PollingWaitSeconds  float64 `json:"pollingWaitSeconds"`
	RequestID           string  `json:"requestId"`
	VisibilityTimeout   float64 `json:"visibilityTimeout"`
}

// Implement https://cloud.tencent.com/document/api/406/5835
func (c *CmqClient) SetQueueAttributes(options ...string) (*SetQueueAttributesResp, error) {
	resp, err := c.DoAction("SetQueueAttributes", options...)
	if err != nil {
		return nil, err
	}
	var s SetQueueAttributesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func SetQueueAttributes(options ...string) (*SetQueueAttributesResp, error) {
	return DefaultClient.SetQueueAttributes(options...)
}

func (r *SetQueueAttributesResp) String(args ...interface{}) (string, error) {
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
