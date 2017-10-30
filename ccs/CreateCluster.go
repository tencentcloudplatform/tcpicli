package ccs

import (
	"encoding/json"
)

func CreateCluster(region ...string) (*RespCreateCluster, error) {
	b, err := DoAction("CreateCluster", region...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateCluster)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
