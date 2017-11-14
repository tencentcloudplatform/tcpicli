package vpc

import (
	"encoding/json"
)

func DeteleSubnetAclRule(options ...string) (*RespDeteleSubnetAclRule, error) {
	b, err := DoAction("DeteleSubnetAclRule", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeteleSubnetAclRule)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
