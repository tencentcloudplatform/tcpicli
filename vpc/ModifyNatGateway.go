package vpc

import (
	"encoding/json"
)

func ModifyNatGateway(options ...string) (*RespModifyNatGateway, error) {
	b, err := DoAction("ModifyNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
