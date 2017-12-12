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
	region := "Region=bj"                                     // common
	eipName := "eipName=tcpiclimod"                           // ModifyEipAttribute
	vagueInstanceName := "VagueInstanceName=tcpiclicvmgenmod" // DescribeInstances / EipBindInstance ->

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeEip",
			"DescribeEipQuota",
			"CreateEip",
			`SET eipId=tcpicli -f "{{range .Data.EipSet}}{{.eipId}}{{end}}" eip DescribeEip ` + region,
			`DO echo $eipId`,
			"ModifyEipAttributes",
			`SET instanceId=tcpicli -f "{{range .Response.InstanceSet}}{{.InstanceID}}{{end}}" cvm DescribeInstances ` + region + " " + vagueInstanceName,
			`DO echo $instanceId`,
			"EipBindInstance",
			`DO sleep 15`,
			"EipUnBindInstance",
			`DO sleep 15`,
			"DeleteEip",
		},
		FuncMap: map[string][]string{
			"DescribeEip": []string{"213/1379",
				region,
			},
			"DescribeEipQuota": []string{"213/1378",
				region,
			},
			"CreateEip": []string{"213/1381",
				region,
			},
			"ModifyEipAttributes": []string{"213/1375",
				region,
				eipName,
				"eipId=$eipId",
			},
			"EipBindInstance": []string{"213/1377",
				region,
				"eipId=$eipId",
				"unInstanceId=$instanceId",
			},
			"EipUnBindInstance": []string{"213/1376",
				region,
				"eipId=$eipId",
			},
			"DeleteEip": []string{"213/1380",
				region,
				"eipIds.0=$eipId",
			},
		},
		PkgName: "eip",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
