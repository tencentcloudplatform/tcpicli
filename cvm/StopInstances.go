package cvm

import (
	"encoding/json"
)

type StopInstancesResp struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Detail   interface{} `json:"detail"`
	Message  string      `json:"message"`
}

func StopInstances(options ...string) (*StopInstancesResp, error) {
	resp, err := DoAction("StopInstances", options...)
	if err != nil {
		return nil, err
	}
	var s StopInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
