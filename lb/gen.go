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
	vpcName := "vpcName=tcpicligen" // VPC to create everything in
	// vpcCidrBlock := "cidrBlock=10.0.0.0/16"                // VPC cidr
	// subnetCidrBlock := "subnetSet.0.cidrBlock=10.0.0.0/24" // Subnet cidr
	subnetName := "subnetSet.0.subnetName=tcpicligensub" // Subnet Name
	zoneId := "subnetSet.0.zoneId=800002"                // ap-beijing-2

	region := "Region=bj"                                   // common
	loadBalancerType := "loadBalancerType=3"                // InquiryLBPrice / CreateLoadBalancer
	loadBalancerName := "loadBalancerName=tcpicligen"       // CreateLoadBalancer
	loadBalancerNameMod := "loadBalancerName=tcpicligenmod" // ModifyLoadBalancerAttributes

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// set up vpc & subnet to gen everything
			// `DO tcpicli vpc CreateVpc ` + region + " " + vpcName + " " + vpcCidrBlock, // Set up gen VPC and capture the vpcID, etc
			`SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName,
			`DO echo $vpcId`,
			// `DO tcpicli vpc CreateSubnet vpcId=$vpcId ` + region + " " + subnetCidrBlock + " " + subnetName + " " + zoneId,
			`SET subnetId=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			`DO echo $subnetId`,

			"InquiryLBPrice",
			"CreateLoadBalancer",
			"DescribeLoadBalancers",
			// `SET `,
			// "ModifyLoadBalancerAttributes",
			// "DeleteLoadBalancers",
			// -- Clean --
			// `DO tcpicli vpc DeleteSubnet vpcId=$vpcId subnetId=$subnetId ` + region,
			// `DO tcpicli vpc DeleteVpc vpcId=$vpcId ` + region,
		},
		FuncMap: map[string][]string{
			"InquiryLBPrice": []string{"214/1328",
				region,
				loadBalancerType,
			},
			"CreateLoadBalancer": []string{"214/1254",
				region,
				loadBalancerType,
				loadBalancerName,
				"vpcId=$vpcId",
				"subnetId=$subnetId",
			},
			"DescribeLoadBalancers": []string{"",
				region,
				loadBalancerName,
			},
			"ModifyLoadBalancerAttributes": []string{"",
				region,
				loadBalancerNameMod,
			},
			"DeleteLoadBalancers": []string{"",
				region,
			},
		},
		PkgName: "lb",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
