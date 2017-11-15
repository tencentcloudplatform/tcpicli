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
	Region := "echo gz"
	availabilityZone := "echo 100002"
	vpcName := "echo initialtcpiclivpc"
	vpcNameMod := "echo tcpiclivpc"
	vpcId := "echo vpc-ncetpc2r"           // hardcode
	vpcTestInstance := "echo ins-kjhjjaki" // hardcode
	cidrBlock := "echo 10.0.0.0/16"
	subnetCidr0 := "echo 10.0.0.0/24"
	subnetCidr1 := "echo 10.0.1.0/24"
	subnetName0 := "echo tcpclisub0"
	subnetName1 := "echo tcpclisub1"
	vpcType := "echo 1"
	vipId := "echo eip-kw5x9v3i" // hardcode
	lanIp := "echo "             // hardcode need fix
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"SET Region=" + Region,
			"SET vpcName=" + vpcName,
			"SET cidrBlock=" + cidrBlock,
			"SET subnetCidr0=" + subnetCidr0,
			"SET subnetCidr1=" + subnetCidr1,
			"SET subnetName0=" + subnetName0,
			"SET subnetName1=" + subnetName1,
			"SET zoneId=" + availabilityZone,
			"SET vpcType=" + vpcType,
			// "CreateVpc",
			// `SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx Region=$Region vpcName=$vpcName`, // need to fix vpc package
			// "DescribeVpcEx",
			// "DescribeVpcClassicLink",
			// "DescribeVpcLimit",
			"SET vpcNameMod=" + vpcNameMod,
			"SET vpcId=" + vpcId, // hardcode
			// "ModifyVpcAttribute",
			// `SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx Region=$Region vpcName=$vpcNameMod`, // need to fix vpc package
			"SET vipId=" + vipId, // hardcode
			"SET lanIp=" + lanIp, // hardcode
			// "AssociateVip",
			"SET instanceIds.0=" + vpcTestInstance,
			// "AttachClassicLinkVpc",
			"DeleteVpc",
		},
		FuncMap: map[string][]string{
			"CreateVpc": []string{"215/1309",
				"Region=$Region",
				"vpcName=$vpcName",
				"cidrBlock=$cidrBlock",
				"subnetSet.0.cidrBlock=$subnetCidr0",
				"subnetSet.1.cidrBlock=$subnetCidr1",
				"subnetSet.0.subnetName=$subnetName0",
				"subnetSet.1.subnetName=$subnetName1",
				"subnetSet.0.zoneId=$zoneId",
				"subnetSet.1.zoneId=$zoneId",
			},
			"DescribeVpcEx": []string{"215/1372",
				"Region=$Region",
			},
			"DescribeVpcClassicLink": []string{"215/2112",
				"Region=$Region",
			},
			"DescribeVpcLimit": []string{"215/1844",
				"Region=$Region",
				"type.0=$vpcType",
			},
			"ModifyVpcAttribute": []string{"215/1310",
				"Region=$Region",
				"vpcId=$vpcId",
				"vpcName=$vpcNameMod",
			},
			"AssociateVip": []string{"215/1361",
				"Region=$Region",
				"vpcId=$vpcId",
				"vipId=$vipId",
				"lanIp=$lanIp",
			},
			"AttachClassicLinkVpc": []string{"215/2098",
				"Region=$Region",
				"vpcId=$vpcId",
				"instanceIds.0=$vpcTestInstance",
			},
			"DeleteVpc": []string{"215/1307",
				"Region=$Region",
				"vpcId=$vpcId",
			},
		},
		PkgName: "vpc",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
