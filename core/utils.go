package core

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

// used to determine cmq endpoint
func CmqEndPoint(action string, options ...string) string {
	queryType := ""
	name := ""
	region := ""
	queueRe := regexp.MustCompile("(?i)((send|receive|delete)message|batch.*message|queue)")
	if queueAction := queueRe.MatchString(action); queueAction {
		queryType = "queue"
	}
	topicRe := regexp.MustCompile("(?i)(topic|subsc|publish)")
	if topicAction := topicRe.MatchString(action); topicAction {
		queryType = "topic"
	}
	if Internal() {
		name = "tencentyun"
	} else if !Internal() {
		name = "qcloud"
	}

	for _, v := range options {
		v = strings.ToLower(v)
		if strings.HasPrefix(v, "region=") {
			region = strings.Split(v, "=")[1]
		}
	}
	return fmt.Sprintf("cmq-%s-%s.api.%s.com/v2/index.php", queryType, region, name)
}
