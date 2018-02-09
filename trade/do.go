package trade

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type TradeClient struct {
	*core.Client
}

var DefaultClient = TradeClient{Client: core.DefaultClient}

func NewClient() *TradeClient {
	return &TradeClient{Client: core.DefaultClient}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("trade", action, options...)
}

func (client *TradeClient) DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return client.Client.DoAction("trade", action, options...)
}
