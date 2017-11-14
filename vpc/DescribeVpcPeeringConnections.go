package vpc

import (
	"encoding/json"
)

func DescribeVpcPeeringConnections(options ...string) (*RespDescribeVpcPeeringConnections, error) {
	b, err := DoAction("DescribeVpcPeeringConnections", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeVpcPeeringConnections)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
