package dfw

import (
	"encoding/json"
)

type CreateSecurityGroupResp struct {
}

// Implement https://cloud.tencent.com/document/api/213/9388
func CreateSecurityGroup(options ...string) (*CreateSecurityGroupResp, error) {
	resp, err := DoAction("CreateSecurityGroup", options...)
	if err != nil {
		return nil, err
	}
	var s CreateSecurityGroupResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
