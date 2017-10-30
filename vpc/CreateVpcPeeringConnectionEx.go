package vpc

import (
	"encoding/json"
)

func CreateVpcPeeringConnectionEx(options ...string) (*RespCreateVpcPeeringConnectionEx, error) {
	b, err := DoAction("CreateVpcPeeringConnectionEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateVpcPeeringConnectionEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
