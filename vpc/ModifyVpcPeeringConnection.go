package vpc

import (
	"encoding/json"
)

func ModifyVpcPeeringConnection(options ...string) (*RespModifyVpcPeeringConnection, error) {
	b, err := DoAction("ModifyVpcPeeringConnection", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyVpcPeeringConnection)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
