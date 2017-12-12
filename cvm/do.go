package cvm

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

var requesturl string = core.Endpoint["cvm"]

func DoAction(action string, options ...string) ([]byte, error) {
	version, ok := core.HasVersion(options...)
	if !ok {
		version = "2017-03-12"
		options = append(options, "Version="+version)
	}
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return core.DoAction("cvm", action, options...)
}
