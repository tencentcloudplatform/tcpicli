package vpc

import (
	"encoding/json"
)

func DescribeVpcClassicLink(options ...string) (*RespDescribeVpcClassicLink, error) {
	b, err := DoAction("DescribeVpcClassicLink", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeVpcClassicLink)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
