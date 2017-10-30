package vpc

import (
	"encoding/json"
)

func DeleteVpcPeeringConnectionEx(options ...string) (*RespDeleteVpcPeeringConnectionEx, error) {
	b, err := DoAction("DeleteVpcPeeringConnectionEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteVpcPeeringConnectionEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
