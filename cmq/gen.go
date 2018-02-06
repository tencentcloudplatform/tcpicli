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
	queueName := "queueName=tcpicligen"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"CreateQueue",
			"ListQueue",
			"GetQueueAttributes",
			"SetQueueAttributes",
			"SendMessage",
			`DO sleep 3`,
			"ReceiveMessage",
			`SET receiptHandle=tcpicli -f "{{.ReceiptHandle}} cmq ReceiveMessage "` + region + " " + queueName,
			`DO echo $receiptHandle`,
			"DeleteMessage",
			"DeleteQueue",
		},
		FuncMap: map[string][]string{
			"CreateQueue": []string{"406/5832",
				region,
				queueName,
			},
			"ListQueue": []string{"406/5833",
				region,
			},
			"GetQueueAttributes": []string{"406/5834",
				region,
				queueName,
			},
			"SetQueueAttributes": []string{"406/5835",
				region,
				queueName,
				"pollingWaitSeconds=5", // used to arbitrarily change queue attributes for function gen
			},
			"SendMessage": []string{"406/5837",
				region,
				queueName,
				"msgBody=tcpicli gen test",
			},
			"ReceiveMessage": []string{"406/5839",
				region,
				queueName,
			},
			"DeleteMessage": []string{"406/5840",
				region,
				queueName,
				"receiptHandle=$receiptHandle",
			},
			"DeleteQueue": []string{"406/5836",
				region,
				queueName,
			},
		},
		PkgName: "cmq",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
