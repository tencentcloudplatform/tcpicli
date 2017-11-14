package vpc

import (
	"encoding/json"
)

func RejectVpcPeeringConnectionEx(options ...string) (*RespRejectVpcPeeringConnectionEx, error) {
	b, err := DoAction("RejectVpcPeeringConnectionEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespRejectVpcPeeringConnectionEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
