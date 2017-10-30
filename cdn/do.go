package cdn

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["cdn"]

func DoAction(action string, options ...string) ([]byte, error) {
	return core.DoAction("cdn", action, options...)
}
