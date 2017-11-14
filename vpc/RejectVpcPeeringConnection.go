package vpc

import (
	"encoding/json"
)

func RejectVpcPeeringConnection(options ...string) (*RespRejectVpcPeeringConnection, error) {
	b, err := DoAction("RejectVpcPeeringConnection", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespRejectVpcPeeringConnection)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
