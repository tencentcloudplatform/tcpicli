package vpc

import (
	"encoding/json"
)

func DescribeNatGateway(options ...string) (*RespDescribeNatGateway, error) {
	b, err := DoAction("DescribeNatGateway", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeNatGateway)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
