package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func DoAction(service, action string, options ...string) ([]byte, error) {
	client := NewClient(Endpoint[service], true)
	method := "POST"

	params := make(map[string]interface{})
	params["Action"] = action
	AssignParams(params, options...)

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
	if val, ok := m["code"]; ok {
		if val.(float64) != 0 {
			return b, errors.New(string(b))
		}
	}
	return b, nil
}
