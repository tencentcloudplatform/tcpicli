package cvm

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CvmClient struct {
	core.Client
}

var DefaultClient = CvmClient{Client: *core.NewClient()}

func NewClient() *CvmClient {
	return &CvmClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("cvm", action, options...)
}

func (client *CvmClient) DoAction(action string, options ...string) ([]byte, error) {
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
	return client.Client.DoAction("cvm", action, options...)
}
