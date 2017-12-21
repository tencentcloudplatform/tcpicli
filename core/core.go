package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"text/template"
)

const PWD = `github.com/tencentcloudplatform/tcpicli/core`

func DoAction(service, action string, options ...string) ([]byte, error) {
	return DefaultClient.DoAction(service, action, options...)
}

func (client *Client) DoAction(service, action string, options ...string) ([]byte, error) {
	// 	client := NewClient(Endpoint[service], true)
	method := "POST"
	requesturl := Endpoint[service]

	params := make(map[string]interface{})
	params["Action"] = action
	AssignParams(params, options...)

	resp, err := client.SendRequest(method, requesturl, params)
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
	if _, ok := HasVersion(options...); ok {
		for _, val := range m {
			if view, ok := val.(map[string]interface{}); ok {
				for k, _ := range view {
					if k == "Error" {
						return b, errors.New(string(b))
					}
				}
			}
		}
	}
	return b, nil
}

func FmtOutput(view string, intf interface{}) ([]byte, error) {
	if _, err := os.Stat(view); !os.IsNotExist(err) {
		b, err := ioutil.ReadFile(view)
		if err != nil {
			return nil, err
		}
		view = string(b)
	}
	return execTemplate(view, intf)
}

func execTemplate(tmpl string, body interface{}) ([]byte, error) {
	j := func(i interface{}) string {
		b, err := json.MarshalIndent(i, "", "  ")
		if err != nil {
			return err.Error()
		}
		return string(b)
	}
	m := template.FuncMap{
		"json": j,
	}
	t, err := template.New("tmpl").Funcs(m).Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
