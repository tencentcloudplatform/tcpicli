package vpc

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type VpcClient struct {
	*core.Client
}

var DefaultClient = VpcClient{Client: core.DefaultClient}

func NewClient() *VpcClient {
	return &VpcClient{Client: core.DefaultClient}
}

func DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}

	return DefaultClient.Client.DoAction("vpc", action, options...)
}

func (client *VpcClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("vpc", action, options...)
}
