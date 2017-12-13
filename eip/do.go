package eip

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type EipClient struct {
	core.Client
}

var DefaultClient = EipClient{Client: *core.NewClient()}

func NewClient() *EipClient {
	return &EipClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("eip", action, options...)
}

func (client *EipClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("eip", action, options...)
}
