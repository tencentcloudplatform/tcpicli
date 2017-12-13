package vod

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type VodClient struct {
	core.Client
}

var DefaultClient = VodClient{Client: *core.NewClient()}

func NewClient() *VodClient {
	return &VodClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("vod", action, options...)
}

func (client *VodClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("vod", action, options...)
}
