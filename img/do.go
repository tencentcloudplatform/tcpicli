package img

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
	"io/ioutil"
)

//go:generate go run gen.go
func doAction(endpoint string, action string, options ...string) ([]byte, error) {
	client := core.NewClient(endpoint, true)
	method := "POST"

	params := make(map[string]interface{})
	params["Action"] = action
	core.AssignParams(params, options...)

	resp, err := client.SendRequest(method, params)
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
	// img api uses multiple standards for JSON error responses. valN reflects error conditions
	if valC, ok := m["code"]; ok {
		if valC.(float64) != 0 {
			return b, errors.New(string(b))
		}
	}
	if valE, ok := m["Error"]; ok {
		if valE.(interface{}) != nil {
			return b, errors.New(string(b))
		}
	}
	if valEA, ok := m["Error"]; ok {
		if valEA.([]interface{}) != nil {
			return b, errors.New(string(b))
		}
	}
	return b, nil
}

func appendVersion(action string, options ...string) ([]string, error) {
	oldActions := []string{
		"DescribeImages",
		"SyncImages",
	}
	version := "Version=2017-03-12"

	for _, value := range oldActions {
		if value == action {
			options = append(options, version)
		}
	}
	return options, nil
}

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	requesturl := fmt.Sprintf(core.Endpoint["img"])
	options, err := appendVersion(action, options...)
	if err != nil {
		return nil, err
	}
	return doAction(requesturl, action, options...)
}
