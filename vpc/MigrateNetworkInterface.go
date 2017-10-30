package vpc

import (
	"encoding/json"
)

func MigrateNetworkInterface(options ...string) (*RespMigrateNetworkInterface, error) {
	b, err := DoAction("MigrateNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespMigrateNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
