package vpc

import (
	"encoding/json"
)

func EipBindNatGateway(options ...string) (*RespEipBindNatGateway, error) {
	b, err := DoAction("EipBindNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespEipBindNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
