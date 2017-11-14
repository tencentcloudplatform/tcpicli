package vpc

import (
	"encoding/json"
)

func ModifySubnetAttribute(options ...string) (*RespModifySubnetAttribute, error) {
	b, err := DoAction("ModifySubnetAttribute", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifySubnetAttribute)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
