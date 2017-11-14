package vpc

import (
	"encoding/json"
)

func DescribeNetworkAcl(options ...string) (*RespDescribeNetworkAcl, error) {
	b, err := DoAction("DescribeNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeNetworkAcl)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
