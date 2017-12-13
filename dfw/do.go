package dfw

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DfwClient struct {
	core.Client
}

var DefaultClient = DfwClient{Client: *core.NewClient()}

func NewClient() *DfwClient {
	return &DfwClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("dfw", action, options...)
}

func (client *DfwClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("dfw", action, options...)
}
