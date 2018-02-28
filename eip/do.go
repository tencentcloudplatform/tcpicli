package eip

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type EipClient struct {
	*core.Client
}

var DefaultClient = EipClient{Client: core.DefaultClient}

func NewClient() *EipClient {
	return &EipClient{Client: core.DefaultClient}
}

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return DefaultClient.Client.DoAction("eip", action, options...)
}

func (client *EipClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("eip", action, options...)
}
