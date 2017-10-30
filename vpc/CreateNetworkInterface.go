package vpc

import (
	"encoding/json"
)

func CreateNetworkInterface(options ...string) (*RespCreateNetworkInterface, error) {
	b, err := DoAction("CreateNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
