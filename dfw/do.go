package dfw

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["dfw"]

func DoAction(action string, options ...string) ([]byte, error) {
	return core.DoAction("dfw", action, options...)
}
