// +build ignore
package main

import (
	. "."
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"ListInstance",
			"CreateTopic",
			"ListTopic",
			"SetInstanceAttributes",
			"GetInstanceAttributes",
			"SetTopicAttributes",
			"GetTopicAttributes",
			"AddPartition",
			"AddTopicIpWhitelist",
			// clean up
			"DeleteTopic",
			"DeleteTopicIpWhitelist",
		},
		FuncMap: map[string][]string{
			"SetInstanceAttributes": []string{"597/10095",
				region,
			},
			"GetInstanceAttributes": []string{"597/10094",
				region,
			},
			"ListInstance": []string{"597/10093",
				region,
			},
			"CreateTopic": []string{"597/10096",
				region,
			},
			"DeleteTopic": []string{"597/10099",
				region,
			},
			"SetTopicAttributes": []string{"597/10098",
				region,
			},
			"GetTopicAttributes": []string{"597/10102",
				region,
			},
			"ListTopic": []string{"597/10101",
				region,
			},
			"AddPartition": []string{"597/10100",
				region,
			},
			"AddTopicIpWhitelist": []string{"597/10103",
				region,
			},
			"DeleteTopicIpWhitelist": []string{"597/10104",
				region,
			},
		},
		PkgName: "ckafka",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
