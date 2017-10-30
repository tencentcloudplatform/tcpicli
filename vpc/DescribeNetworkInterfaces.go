package vpc

import (
	"encoding/json"
)

func DescribeNetworkInterfaces(options ...string) (*RespDescribeNetworkInterfaces, error) {
	b, err := DoAction("DescribeNetworkInterfaces", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeNetworkInterfaces)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
