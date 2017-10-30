package cvm

import (
	"encoding/json"
)

type StartInstancesResp struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Detail   interface{} `json:"detail"`
	Message  string      `json:"message"`
}

func StartInstances(options ...string) (*StartInstancesResp, error) {
	resp, err := DoAction("StartInstances", options...)
	if err != nil {
		return nil, err
	}
	var s StartInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
