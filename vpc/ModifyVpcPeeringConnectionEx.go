package vpc

import (
	"encoding/json"
)

func ModifyVpcPeeringConnectionEx(options ...string) (*RespModifyVpcPeeringConnectionEx, error) {
	b, err := DoAction("ModifyVpcPeeringConnectionEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyVpcPeeringConnectionEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
