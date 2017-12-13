package cmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
	"io/ioutil"
)

//go:generate go run gen.go

func doAction(endpoint string, action string, options ...string) ([]byte, error) {
	client := core.NewClient()
	method := "POST"

	params := make(map[string]interface{})
	params["Action"] = action
	core.AssignParams(params, options...)

	resp, err := client.SendRequest(method, endpoint, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	b, err = json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}
	if val, ok := m["code"]; ok {
		if val.(float64) != 0 {
			return b, errors.New(string(b))
		}
	}
	return b, nil
}

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	name := "tencentyun"
	if !core.Internal() {
		name = "qcloud"
	}
	requesturl := fmt.Sprintf("cmq-queue-%s.api.%s.com/v2/index.php", region, name)
	return doAction(requesturl, action, options...)
}
