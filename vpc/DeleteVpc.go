package vpc

import (
	"encoding/json"
)

func DeleteVpc(vpcid ...string) (*RespDeleteVpc, error) {
	b, err := DoAction("DeleteVpc", vpcid...)
	if err != nil {
		return nil, err
	}
	m := new(RespDeleteVpc)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
