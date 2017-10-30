package cvm

import (
	"encoding/json"
)

type RebootInstancesResp struct {
	Response struct {
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

func RebootInstances(options ...string) (*RebootInstancesResp, error) {
	resp, err := DoAction("RebootInstances", options...)
	if err != nil {
		return nil, err
	}
	var s RebootInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
