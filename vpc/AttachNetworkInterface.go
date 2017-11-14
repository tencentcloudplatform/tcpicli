package vpc

import (
	"encoding/json"
)

func AttachNetworkInterface(options ...string) (*RespAttachNetworkInterface, error) {
	b, err := DoAction("AttachNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAttachNetworkInterface)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
