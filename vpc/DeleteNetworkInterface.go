package vpc

import (
	"encoding/json"
)

func DeleteNetworkInterface(options ...string) (*RespDeleteNetworkInterface, error) {
	b, err := DoAction("DeleteNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
