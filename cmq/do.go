package cmq

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CmqClient struct {
	core.Client
}

var DefaultClient = CmqClient{Client: *core.NewClient()}

func NewClient() *CmqClient {
	return &CmqClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("cmq", action, options...)
}

func (client *CmqClient) DoAction(action string, options ...string) ([]byte, error) {
	if region, ok := core.HasRegion(options...); !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return client.Client.DoAction("cmq", action, options...)
}
