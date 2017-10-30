package vpc

import (
	"encoding/json"
)

func DescribeSubnetEx(options ...string) (*RespDescribeSubnetEx, error) {
	b, err := DoAction("DescribeSubnetEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeSubnetEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
