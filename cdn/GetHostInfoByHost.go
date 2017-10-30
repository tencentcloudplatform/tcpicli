package cdn

import (
	"encoding/json"
	"fmt"
)

func GetHostInfoByHost(hostnames ...string) (*RespGetHostInfo, error) {
	var options []string
	for k, v := range hostnames {
		options = append(options, fmt.Sprintf("hosts.%v=%v", k, v))
	}
	b, err := DoAction("GetHostInfoByHost", options...)
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
