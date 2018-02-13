package feecenter

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type FeecenterClient struct {
	*core.Client
}

var DefaultClient = FeecenterClient{Client: core.DefaultClient}

func NewClient() *FeecenterClient {
	return &FeecenterClient{Client: core.DefaultClient}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("feecenter", action, options...)
}

func (client *FeecenterClient) DoAction(action string, options ...string) ([]byte, error) {
	if region, ok := core.HasRegion(options...); !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return client.Client.DoAction("feecenter", action, options...)
}
