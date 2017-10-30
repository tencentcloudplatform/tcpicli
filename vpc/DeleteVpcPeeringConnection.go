package vpc

import (
	"encoding/json"
)

func DeleteVpcPeeringConnection(options ...string) (*RespDeleteVpcPeeringConnection, error) {
	b, err := DoAction("DeleteVpcPeeringConnection", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteVpcPeeringConnection)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
