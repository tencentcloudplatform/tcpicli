package vpc

import (
	"encoding/json"
)

func CreateVpcPeeringConnection(options ...string) (*RespCreateVpcPeeringConnection, error) {
	b, err := DoAction("CreateVpcPeeringConnection", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateVpcPeeringConnection)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
