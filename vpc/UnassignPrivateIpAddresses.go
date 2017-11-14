package vpc

import (
	"encoding/json"
)

func UnassignPrivateIpAddresses(options ...string) (*RespUnassignPrivateIpAddresses, error) {
	b, err := DoAction("UnassignPrivateIpAddresses", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespUnassignPrivateIpAddresses)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
