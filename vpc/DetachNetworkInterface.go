package vpc

import (
	"encoding/json"
)

func DetachNetworkInterface(options ...string) (*RespDetachNetworkInterface, error) {
	b, err := DoAction("DetachNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDetachNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
