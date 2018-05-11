// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	region := "Region=gz"
	topicName := "topicName=apitest"
	partitionNum := "partitionNum=1"       // used with CreateTopic
	replicaNum := "replicaNum=1"           // used with CreateTopic
	enableWhiteList := "enableWhiteList=1" // used with SetTopicAttributes

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"ListInstance",
			`SET instanceId=tcpicli -f '{{range .Data.InstanceList}}{{if eq .InstanceName "test"}}{{.InstanceID}}{{end}}{{end}}' ckafka ListInstance ` + region,
			`DO echo $instanceId`,
			"GetInstanceAttributes",
			"CreateTopic",
			"ListTopic",
			"GetTopicAttributes",
			"SetTopicAttributes",
			"SetInstanceAttributes",
			"AddPartition",
			"AddTopicIpWhitelist",

			// clean up
			"DeleteTopicIpWhitelist",
			"DeleteTopic",
		},
		FuncMap: map[string][]string{
			"ListInstance": []string{"597/10093",
				region,
			},
			"GetInstanceAttributes": []string{"597/10094",
				region,
				"instanceId=$instanceId",
			},
			"CreateTopic": []string{"597/10096",
				region,
				topicName,
				partitionNum,
				replicaNum,
				"instanceId=$instanceId",
			},
			"ListTopic": []string{"597/10101",
				region,
				"instanceId=$instanceId",
			},
			"GetTopicAttributes": []string{"597/10102",
				region,
				topicName,
				"instanceId=$instanceId",
			},
			"SetTopicAttributes": []string{"597/10098",
				region,
				topicName,
				enableWhiteList,
				"instanceId=$instanceId",
			},
			"SetInstanceAttributes": []string{"597/10095",
				region,
				"instanceId=$instanceId",
				"msgRetentionTime=5",
			},
			"DeleteTopic": []string{"597/10099",
				region,
				topicName,
				"instanceId=$instanceId",
			},
			"AddPartition": []string{"597/10100",
				region,
				topicName,
				"partitionNum=2",
				"instanceId=$instanceId",
			},
			"AddTopicIpWhitelist": []string{"597/10103",
				region,
				topicName,
				"instanceId=$instanceId",
				"ipWhiteList.0=172.168.8.8",
			},
			"DeleteTopicIpWhitelist": []string{"597/10104",
				region,
				topicName,
				"instanceId=$instanceId",
				"ipWhiteList.0=172.168.8.8",
			},
		},
		PkgName: "ckafka",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
