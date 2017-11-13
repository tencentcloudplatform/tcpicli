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
	vpcName := "echo tcpiclivpc"
	cidrBlock := "echo 192.168.0.0/16"
	subnetCidr0 := "echo 192.168.0.0/24"
	subnetCidr1 := "echo 192.168.1.0/24"
	subnetName0 := "echo tcpclisub0"
	subnetName1 := "echo tcpclisub1"

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
			"CreateVpc",
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
		},
		PkgName: "vpc",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
