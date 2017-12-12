package eip

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["eip"]

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return core.DoAction("eip", action, options...)
}
