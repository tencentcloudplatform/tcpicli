package vpc

import (
	"encoding/json"
)

func DeleteRouteTable(options ...string) (*RespDeleteRouteTable, error) {
	b, err := DoAction("DeleteRouteTable", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteRouteTable)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
