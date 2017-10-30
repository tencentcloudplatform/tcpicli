package ccs

import (
	"encoding/json"
)

func AddClusterInstances(region ...string) (*RespAddClusterInstances, error) {
	b, err := DoAction("AddClusterInstances", region...)
	if err != nil {
		return nil, err
	}
	m := new(RespAddClusterInstances)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
