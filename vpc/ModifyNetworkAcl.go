package vpc

import (
	"encoding/json"
)

func ModifyNetworkAcl(options ...string) (*RespModifyNetworkAcl, error) {
	b, err := DoAction("ModifyNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyNetworkAcl)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
