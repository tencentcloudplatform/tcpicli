package vpc

import (
	"encoding/json"
)

func AssignPrivateIpAddresses(options ...string) (*RespAssignPrivateIpAddresses, error) {
	b, err := DoAction("AssignPrivateIpAddresses", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAssignPrivateIpAddresses)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
