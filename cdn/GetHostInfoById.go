package cdn

import (
	"encoding/json"
	"fmt"
)

func GetHostInfoById(ids ...string) (*RespGetHostInfo, error) {
	var options []string
	for k, v := range ids {
		options = append(options, fmt.Sprintf("ids.%v=%v", k, v))
	}
	b, err := DoAction("GetHostInfoById", options...)
	if err != nil {
		return nil, err
	}
	m := new(RespGetHostInfo)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
