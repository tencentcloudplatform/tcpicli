package vpc

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["vpc"]

func DoAction(action string, options ...string) ([]byte, error) {
	return core.DoAction("vpc", action, options...)
}
