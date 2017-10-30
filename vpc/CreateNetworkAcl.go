package vpc

import (
	"encoding/json"
)

func CreateNetworkAcl(options ...string) (*RespCreateNetworkAcl, error) {
	b, err := DoAction("CreateNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateNetworkAcl)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
