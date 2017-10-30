package vpc

import (
	"encoding/json"
)

func DeleteSubnet(options ...string) (*RespDeleteSubnet, error) {
	b, err := DoAction("DeleteSubnet", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteSubnet)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
