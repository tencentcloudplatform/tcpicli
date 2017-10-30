package cvm

import (
	"encoding/json"
)

type TerminateInstancesResp struct {
	Response struct {
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

func TerminateInstances(options ...string) (*TerminateInstancesResp, error) {
	resp, err := DoAction("TerminateInstances", options...)
	if err != nil {
		return nil, err
	}
	var s TerminateInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
