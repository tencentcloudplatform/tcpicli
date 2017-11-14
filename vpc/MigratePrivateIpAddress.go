package vpc

import (
	"encoding/json"
)

func MigratePrivateIpAddress(options ...string) (*RespMigratePrivateIpAddress, error) {
	b, err := DoAction("MigratePrivateIpAddress", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespMigratePrivateIpAddress)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
