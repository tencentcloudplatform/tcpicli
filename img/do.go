package img

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ImgClient struct {
	core.Client
}

var DefaultClient = ImgClient{Client: *core.NewClient()}

func NewClient() *ImgClient {
	return &ImgClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("img", action, options...)
}

func (client *ImgClient) DoAction(action string, options ...string) ([]byte, error) {
	return client.Client.DoAction("img", action, options...)
}
