package vpc

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type VpcClient struct {
	core.Client
}

var DefaultClient = VpcClient{Client: *core.NewClient()}

func NewClient() *VpcClient {
	return &VpcClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("vpc", action, options...)
}

func (client *VpcClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("vpc", action, options...)
}
