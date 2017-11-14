package vpc

import (
	"encoding/json"
)

func AcceptVpcPeeringConnection(options ...string) (*RespAcceptVpcPeeringConnection, error) {
	b, err := DoAction("AcceptVpcPeeringConnection", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAcceptVpcPeeringConnection)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
