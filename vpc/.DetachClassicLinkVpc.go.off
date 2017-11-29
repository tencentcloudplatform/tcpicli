package vpc

import (
	"encoding/json"
)

func DetachClassicLinkVpc(options ...string) (*RespDetachClassicLinkVpc, error) {
	b, err := DoAction("DetachClassicLinkVpc", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDetachClassicLinkVpc)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
