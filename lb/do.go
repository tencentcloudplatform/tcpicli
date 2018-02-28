package lb

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type LbClient struct {
	core.Client
}

var DefaultClient = LbClient{Client: *core.NewClient()}

func NewClient() *LbClient {
	return &LbClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}

	return DefaultClient.Client.DoAction("lb", action, options...)
}

func (client *LbClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("lb", action, options...)
}
