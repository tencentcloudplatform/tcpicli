package vpc

import (
	"encoding/json"
)

func ModifyVpcAttribute(options ...string) (*RespModifyVpcAttribute, error) {
	b, err := DoAction("ModifyVpcAttribute", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyVpcAttribute)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}