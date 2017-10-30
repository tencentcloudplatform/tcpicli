package vpc

import (
	"encoding/json"
)

func CreateSubnetAclRule(options ...string) (*RespCreateSubnetAclRule, error) {
	b, err := DoAction("CreateSubnetAclRule", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateSubnetAclRule)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
