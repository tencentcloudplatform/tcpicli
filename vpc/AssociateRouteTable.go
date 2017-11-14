package vpc

import (
	"encoding/json"
)

func AssociateRouteTable(options ...string) (*RespAssociateRouteTable, error) {
	b, err := DoAction("AssociateRouteTable", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAssociateRouteTable)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
