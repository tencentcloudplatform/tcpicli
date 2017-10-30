package ccs

import (
	"encoding/json"
)

func DeleteCluster(options ...string) (*RespDeleteCluster, error) {
	b, err := DoAction("DeleteCluster", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteCluster)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}