package vpc

import (
	"encoding/json"
)

func DeleteNatGateway(options ...string) (*RespDeleteNatGateway, error) {
	b, err := DoAction("DeleteNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
