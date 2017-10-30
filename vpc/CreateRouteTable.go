package vpc

import (
	"encoding/json"
)

func CreateRouteTable(options ...string) (*RespCreateRouteTable, error) {
	b, err := DoAction("CreateRouteTable", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateRouteTable)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
