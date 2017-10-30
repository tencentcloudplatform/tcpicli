package vpc

import (
	"encoding/json"
)

func AttachClassicLinkVpc(options ...string) (*RespAttachClassicLinkVpc, error) {
	b, err := DoAction("AttachClassicLinkVpc", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespAttachClassicLinkVpc)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
