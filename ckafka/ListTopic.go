// 2018-01-29 16:54:26.475658 -0800 PST m=+7.434896530
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ckafka

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ListTopicResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TopicList []struct {
			TopicID   string `json:"topicId"`
			TopicName string `json:"topicName"`
		} `json:"topicList"`
		TotalCount int64 `json:"totalCount"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/597/10101
func (c *CkafkaClient) ListTopic(options ...string) (*ListTopicResp, error) {
	resp, err := c.DoAction("ListTopic", options...)
	if err != nil {
		return nil, err
	}
	var s ListTopicResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func ListTopic(options ...string) (*ListTopicResp, error) {
	return DefaultClient.ListTopic(options...)
}

func (r *ListTopicResp) String(args ...interface{}) (string, error) {
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
