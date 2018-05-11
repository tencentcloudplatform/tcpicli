// 2018-02-06 11:33:05.087329 -0800 PST m=+1.323175655
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateQueueResp struct {
	ClientRequestID int64  `json:"clientRequestId"`
	Code            int64  `json:"code"`
	Message         string `json:"message"`
	QueueID         string `json:"queueId"`
	RequestID       string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/5832
func (c *CmqClient) CreateQueue(options ...string) (*CreateQueueResp, error) {
	resp, err := c.DoAction("CreateQueue", options...)
	if err != nil {
		return nil, err
	}
	var s CreateQueueResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func CreateQueue(options ...string) (*CreateQueueResp, error) {
	return DefaultClient.CreateQueue(options...)
}

func (r *CreateQueueResp) String(args ...interface{}) (string, error) {
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
