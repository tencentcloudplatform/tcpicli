package vpc

import (
	"encoding/json"
)

func DeleteNetworkAcl(options ...string) (*RespDeleteNetworkAcl, error) {
	b, err := DoAction("DeleteNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteNetworkAcl)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
