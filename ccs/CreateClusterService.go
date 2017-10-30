package ccs

import (
	"encoding/json"
)

type CreateClusterServiceResp struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
}

func CreateClusterService(options ...string) (*CreateClusterServiceResp, error) {
	resp, err := DoAction("CreateClusterService", options...)
	if err != nil {
		return nil, err
	}
	var s CreateClusterServiceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
