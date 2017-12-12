package cdn

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CdnClient struct {
	core.Client
}

var DefaultClient = CdnClient{Client: *core.NewClient()}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("cdn", action, options...)
}

func (client *CdnClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("cdn", action, options...)
}
