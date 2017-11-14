package vpc

import (
	"encoding/json"
)

func AcceptVpcPeeringConnectionEx(options ...string) (*RespAcceptVpcPeeringConnectionEx, error) {
	b, err := DoAction("AcceptVpcPeeringConnectionEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAcceptVpcPeeringConnectionEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
