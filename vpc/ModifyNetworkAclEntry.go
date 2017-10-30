package vpc

import (
	"encoding/json"
)

func ModifyNetworkAclEntry(options ...string) (*RespModifyNetworkAclEntry, error) {
	b, err := DoAction("ModifyNetworkAclEntry", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespModifyNetworkAclEntry)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
