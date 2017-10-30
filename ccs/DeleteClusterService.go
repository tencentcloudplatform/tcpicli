package ccs

import (
	"encoding/json"
)

type DeleteClusterServiceResp struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9437
func DeleteClusterService(options ...string) (*DeleteClusterServiceResp, error) {
	resp, err := DoAction("DeleteClusterService", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteClusterServiceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
