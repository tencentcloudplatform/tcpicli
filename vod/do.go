package vod

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["vod"]

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return core.DoAction("vod", action, options...)
}
