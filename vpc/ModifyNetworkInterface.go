package vpc

import (
	"encoding/json"
)

func ModifyNetworkInterface(options ...string) (*RespModifyNetworkInterface, error) {
	b, err := DoAction("ModifyNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
