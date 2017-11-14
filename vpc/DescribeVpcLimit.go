package vpc

import (
	"encoding/json"
)

func DescribeVpcLimit(options ...string) (*RespDescribeVpcLimit, error) {
	b, err := DoAction("DescribeVpcLimit", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeVpcLimit)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
