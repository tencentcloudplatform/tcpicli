package vpc

import (
	"encoding/json"
)

func EipUnBindNatGateway(options ...string) (*RespEipUnBindNatGateway, error) {
	b, err := DoAction("EipUnBindNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespEipUnBindNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
