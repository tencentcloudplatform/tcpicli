package ccs

import (
	"encoding/json"
)

type ModifyClusterServiceResp struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

func ModifyClusterService(options ...string) (*ModifyClusterServiceResp, error) {
	resp, err := DoAction("ModifyClusterService", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyClusterServiceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
