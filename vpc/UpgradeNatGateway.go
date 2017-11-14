package vpc

import (
	"encoding/json"
)

func UpgradeNatGateway(options ...string) (*RespUpgradeNatGateway, error) {
	b, err := DoAction("UpgradeNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespUpgradeNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
