// +build ignore

package main

import (
	. "."
	// "fmt"
	"github.com/tencentcloudplatform/tcpicli/autogen"
	// "time"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	region := "Region=bj"                      // common
	sgName := "sgName=tcpiclisg"               // CreateSecurityGroup
	sgRemark := "sgRemark=tcpiclimod"          // ModifySecurityGroupAttributes
	iProtocol := "ingress.0.ipProtocol=tcp"    // ModifySecurityGroupPolicys
	iCidr := "ingress.0.cidrIp=192.168.0.0/16" // ModifySecurityGroupPolicys
	iAction := "ingress.0.action=ACCEPT"       // ModifySecurityGroupPolicys
	iPort := "ingress.0.portRange=22"          // ModifySecurityGroupPolicys
	iDesc := "ingress.0.desc=tcpiclisg"        // ModifySecurityGroupPolicys
	eAction := "egress.0.action=ACCEPT"        // ModifySecurityGroupPolicys
	eProtocol := "egress.0.ipProtocol=tcp"     // ModifySecurityGroupPolicys
	eDesc := "egress.0.desc=tcpiclisg"         // ModifySecurityGroupPolicys
	eCidr := "egress.0.cidrIp=192.168.0.0/16"  // ModifySecurityGroupPolicys

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeSecurityGroupEx",
			"CreateSecurityGroup",
			`SET sgId=tcpicli -f "{{range .Data.Detail}}{{.SgID}}{{end}}" dfw DescribeSecurityGroupEx ` + region + " " + sgName,
			`DO echo $sgId`,
			"ModifySecurityGroupAttributes",
			"DescribeSecurityGroupPolicys",
			"ModifySecurityGroupPolicys",
			`SET instanceId=tcpicli -f "{{range .Data}}{{.instanceId}}{{end}}" dfw DescribeInstancesOfSecurityGroup ` + region + " " + sgName,
			`DO echo $instanceId`,
			"DescribeInstancesOfSecurityGroup",
			// "ModifySecurityGroupsOfInstance", // broken. did manually; no way in api to associate sg with instance...
			// "DescribeAssociateSecurityGroups",
			// -- clean up --
			"DeleteSecurityGroup",
		},
		FuncMap: map[string][]string{
			"DescribeSecurityGroupEx": []string{"213/1232",
				region,
			},
			"CreateSecurityGroup": []string{"213/1238",
				region,
				sgName,
			},
			"DeleteSecurityGroup": []string{"213/1362",
				region,
				sgName,
				"sgId=$sgId",
			},
			"ModifySecurityGroupAttributes": []string{"213/1363",
				region,
				sgRemark,
				"sgId=$sgId",
			},
			"DescribeSecurityGroupPolicys": []string{"213/1364",
				region,
				"sgId=$sgId",
			},
			"ModifySecurityGroupPolicys": []string{"213/1365",
				region,
				iProtocol,
				iCidr,
				iAction,
				iPort,
				iDesc,
				eAction,
				eProtocol,
				eDesc,
				eCidr,
				"sgId=$sgId",
			},
			"DescribeInstancesOfSecurityGroup": []string{"213/1366",
				region,
				"sgId=$sgId",
			},
			"ModifySecurityGroupsOfInstance": []string{"213/1367",
				region,
				"sgId=$sgId",
				"instanceSet.0.instanceId=$instanceId",
			},
		},
		PkgName: "dfw",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
