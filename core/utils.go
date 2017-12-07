package core

import (
	"io/ioutil"
	"strings"
)

func AssignParams(params map[string]interface{}, options ...string) error {
	for _, v := range options {
		p := strings.Split(v, "=")
		if len(p) == 1 {
			continue
		}
		if len(p[1]) > 0 && strings.HasPrefix(p[1], "@") {
			fname := strings.TrimLeft(p[1], "@")
			b, err := ioutil.ReadFile(fname)
			if err != nil {
				return err
			}
			params[p[0]] = string(b)
		} else {
			params[p[0]] = p[1]
		}
	}
	return nil
}

func HasRegion(options ...string) (string, bool) {
	for _, v := range options {
		v = strings.ToLower(v)
		if strings.HasPrefix(v, "region=") {
			return strings.Split(v, "=")[1], true
		}
	}
	return "", false
}

func HasVersion(options ...string) (string, bool) {
	for _, v := range options {
		v = strings.ToLower(v)
		if strings.HasPrefix(v, "version=") {
			return strings.Split(v, "=")[1], true
		}
	}
	return "", false
}
