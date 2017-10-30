package cvm
import (
	"encoding/json"
)


type ResetInstancesResp struct {
	Code      int    `json:"code"`
	CodeDesc  string `json:"codeDesc"`
	Message   string `json:"message"`
	RequestID int    `json:"requestId"`
}


func ResetInstances(options ...string) (*ResetInstancesResp, error) {
	resp, err := DoAction("ResetInstances", options...)
	if err != nil {
		return nil, err
	}
	var s ResetInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
