package ckafka

import (
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CkafkaClient struct {
	core.Client
}

var DefaultClient = CkafkaClient{Client: *core.NewClient()}

func NewClient() *CkafkaClient {
	return &CkafkaClient{Client: *core.NewClient()}
}

func DoAction(action string, options ...string) ([]byte, error) {
	return DefaultClient.Client.DoAction("ckafka", action, options...)
}

func (client *CkafkaClient) DoAction(action string, options ...string) ([]byte, error) {
	region, ok := core.HasRegion(options...)
	if !ok {
		region = core.DefaultRegion()
		options = append(options, "Region="+region)
	}
	return client.Client.DoAction("ckafka", action, options...)
}
