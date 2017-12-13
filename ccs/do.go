package ccs

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CcsClient struct {
	core.Client
}

var DefaultClient = CcsClient{Client: *core.NewClient()}

func NewClient() *CcsClient {
	return &CcsClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("ccs", action, options...)
}

func (client *CcsClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("ccs", action, options...)
}
