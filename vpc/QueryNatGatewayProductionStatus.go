package vpc

import (
	"encoding/json"
)

func QueryNatGatewayProductionStatus(options ...string) (*RespQueryNatGatewayProductionStatus, error) {
	b, err := DoAction("QueryNatGatewayProductionStatus", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespQueryNatGatewayProductionStatus)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
