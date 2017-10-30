package ccs

import (
	"encoding/json"
)

func DescribeClusterInstances(options ...string) (*RespDescribeClusterInstances, error) {
	b, err := DoAction("DescribeClusterInstances", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeClusterInstances)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
