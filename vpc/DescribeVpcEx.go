package vpc

import (
	"encoding/json"
)

func DescribeVpcEx(options ...string) (*RespDescribeVpcEx, error) {
	b, err := DoAction("DescribeVpcEx", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeVpcEx)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
