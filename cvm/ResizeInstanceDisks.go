package cvm

import (
	"encoding/json"
)

type ResizeInstanceDisksResp struct {
	Response struct {
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

func ResizeInstanceDisks(options ...string) (*ResizeInstanceDisksResp, error) {
	resp, err := DoAction("ResizeInstanceDisks", options...)
	if err != nil {
		return nil, err
	}
	var s ResizeInstanceDisksResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
