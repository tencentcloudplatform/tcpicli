package vpc

import (
	"encoding/json"
)

func DescribeSubnet(options ...string) (*RespDescribeSubnet, error) {
	b, err := DoAction("DescribeSubnet", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeSubnet)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
