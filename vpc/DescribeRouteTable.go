package vpc

import (
	"encoding/json"
)

func DescribeRouteTable(options ...string) (*RespDescribeRouteTable, error) {
	b, err := DoAction("DescribeRouteTable", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeRouteTable)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
