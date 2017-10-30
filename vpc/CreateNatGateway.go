package vpc

import (
	"encoding/json"
)

func CreateNatGateway(options ...string) (*RespCreateNatGateway, error) {
	b, err := DoAction("CreateNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
