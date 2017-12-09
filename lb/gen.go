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
	vpcName := "vpcName=tcpicligen"                        // VPC to create everything in
	vpcCidrBlock := "cidrBlock=10.0.0.0/16"                // VPC cidr
	subnetCidrBlock := "subnetSet.0.cidrBlock=10.0.0.0/24" // Subnet cidr
	subnetName := "subnetSet.0.subnetName=tcpicligensub"   // Subnet Name
	zoneId := "subnetSet.0.zoneId=800002"                  // ap-beijing-2

	region := "Region=bj"                                       // common
	loadBalancerType := "loadBalancerType=3"                    // InquiryLBPrice / CreateLoadBalancer
	loadBalancerName := "loadBalancerName=tcpicligen"           // CreateLoadBalancer
	loadBalancerNameMod := "loadBalancerName=tcpicligenmod"     // ModifyLoadBalancerAttributes
	listenLoadBalancerPort := "listeners.0.loadBalancerPort=80" // CreateLoadBalancerListeners
	listenInstancePort := "listeners.0.instancePort=80"         // CreateLoadBalancerListeners
	listenProto := "listeners.0.protocol=2"                     // CreateLoadBalancerListeners
	listenName := "listeners.0.listenerName=tcpicligen"         // CreateLoadBalancerListeners
	listenNameMod := "listenerName=tcpicligenmod"               // ModifyLoadBalancerListener

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// -- set up vpc & subnet to gen everything --
			`DO tcpicli vpc CreateVpc ` + region + " " + vpcName + " " + vpcCidrBlock, // Set up gen VPC and capture the vpcID, etc
			`SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName,
			`DO echo $vpcId`,
			`DO tcpicli vpc CreateSubnet vpcId=$vpcId ` + region + " " + subnetCidrBlock + " " + subnetName + " " + zoneId,
			`SET subnetId=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			`DO echo $subnetId`,

			"InquiryLBPrice",
			"CreateLoadBalancer",
			"DescribeLoadBalancers",
			`SET loadBalancerId=tcpicli -f "{{range .LoadBalancerSet}}{{.LoadBalancerID}}{{end}}" lb DescribeLoadBalancers ` + region + " " + loadBalancerName,
			`DO echo $loadBalancerId`,
			"ModifyLoadBalancerAttributes",
			"CreateLoadBalancerListeners",
			`DO sleep 5`,
			"DescribeLoadBalancerListeners",
			`SET listenerId=tcpicli -f "{{range .ListenerSet}}{{.listenerId}}{{end}}" lb DescribeLoadBalancerListeners loadBalancerId=$loadBalancerId ` + region + " " + listenName,
			`DO echo $listenerId`,
			"ModifyLoadBalancerListener",
			`DO sleep 5`,

			// -- Clean --
			"DeleteLoadBalancerListeners",
			`DO sleep 5`,
			"DeleteLoadBalancers",
			`DO sleep 5`,
			`DO tcpicli vpc DeleteSubnet vpcId=$vpcId subnetId=$subnetId ` + region,
			`DO tcpicli vpc DeleteVpc vpcId=$vpcId ` + region,
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
			"DescribeLoadBalancers": []string{"214/1261",
				region,
				loadBalancerName,
			},
			"ModifyLoadBalancerAttributes": []string{"214/1263",
				region,
				loadBalancerNameMod,
				"loadBalancerId=$loadBalancerId",
			},
			"DeleteLoadBalancers": []string{"214/1257",
				region,
				"vpcId=$vpcId",
				"loadBalancerIds.0=$loadBalancerId",
			},
			"CreateLoadBalancerListeners": []string{"214/1255",
				region,
				listenLoadBalancerPort,
				listenInstancePort,
				listenProto,
				listenName,
				"loadBalancerId=$loadBalancerId",
			},
			"DescribeLoadBalancerListeners": []string{"214/1260",
				region,
				"loadBalancerId=$loadBalancerId",
			},
			"ModifyLoadBalancerListener": []string{"214/3601",
				region,
				listenNameMod,
				"loadBalancerId=$loadBalancerId",
				"listenerId=$listenerId",
			},
			"DeleteLoadBalancerListeners": []string{"214/1256",
				region,
				"loadBalancerId=$loadBalancerId",
				"listenerIds.0=$listenerId",
			},
		},
		PkgName: "lb",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
