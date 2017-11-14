package vpc

import (
	"encoding/json"
)

func CreateSubnet(options ...string) (*RespCreateSubnet, error) {
	b, err := DoAction("CreateSubnet", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespCreateSubnet)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
