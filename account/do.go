package main

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

func DoAction(action string, options ...string) ([]byte, error) {
	return core.DoAction("account", action, options...)
}
