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
	imgId := ""
	vpcName := "vpcName=tcpiclicvmautogen"                      // vpc
	cidrBlock := "cidrBlock=10.0.0.0/16"                        // vpc
	subnetCidr := "subnetSet.0.cidrBlock=10.0.0.0/24"           // vpc
	subnetZone := "subnetSet.0.zoneId=800001"                   // vpc
	subnetName := "subnetSet.0.subnetName=tcpiclicvmautogen"    // vpc
	region := "Region=bj"                                       // common
	instanceChargeType := "instanceChargeType=POSTPAID_BY_HOUR" // common
	version := "Version=2017-03-12"                             // common

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// CreateVpc & CreateSubnet to run autogen cvm stuff on and capture their id and unId
			`DO tcpicli vpc CreateVpc ` + region + " " + vpcName + " " + cidrBlock,
			`SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName,
			`DO tcpicli vpc CreateSubnet vpcId=$vpcId ` + region + " " + subnetCidr + " " + subnetName + " " + subnetZone,
			`SET subnetId=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			`SET placementZone=tcpicli -f "{{with index .Data 1}}{{.Zone}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			"RunInstances",
			"DescribeInstances",
			// "DescribeInstancesStatus",
			// "InquiryPriceRunInstances",
			// "StartInstances",
			// "StopInstances",
			// "TerminateInstances",
			// "RebootInstances",
			// "ResetInstance",
			// "InquiryPriceResetInstance",
			// "ResizeInstanceDisks",
			// "InquiryPriceResizeInstanceDisks",
			// "RenewInstances",
			// "InquiryPriceRenewInstances",
			// "ResetInstancesType",
			// "InquiryPriceResetInstancesType",
			// "ModifyInstancesRenewFlag",
			// "ModifyInstancesAttribute",
			// "ResetInstancesInternetMaxBandwidth",
			// "ModifyInstancesProject",
			// "UpdateInstanceVpcConfig",
			// "ResetInstancesPassword",
			// "DescribeInstanceInternetBandwidthConfigs",
			// "DescribeImages",
			// "CreateImage",
			// "DeleteImages",
			// "ModifyImageAttribute",
			// "SyncImages",
			// "ModifyImageSharePermission",
			// "DescribeImageSharePermission",
			// "AttachNetworkInterface",
			// "DescribeSecurityGroups",
			// "CreateSecurityGroup",
			// "DeleteSecurityGroup",
			// "ModifySecurityGroupAttribute",
			// "DescribeSecurityGroupPolicy",
			// "ModifySecurityGroupPolicy",
			// "DescribeInstancesOfSecurityGroup",
			// "ModifySecurityGroupsOfInstance",
			// "DescribeAssociateSecurityGroups",
			// "DescribeEip",
			// "DescribeEipQuota",
			// "ModifyEipAttributes",
			// "CreateEip",
			// "DeleteEip",
			// "EipBindInstance",
			// "EipUnBindInstance",
			// "TransformWanIpToEip",
			// "DescribeKeyPairs",
			// "CreateKeyPair",
			// "ModifyKeyPairAttribute",
			// "DeleteKeyPairs",
			// "ImportKeyPair",
			// "AssociateInstancesKeyPairs",
			// "DisassociateInstancesKeyPairs",
			// Delete VPC stuff
			`DO tcpicli vpc DeleteSubnet vpcId=$vpcId subnetId=$subnetId ` + region,
			`DO tcpicli vpc DeleteVpc vpcId=$vpcId ` + region,
		},
		FuncMap: map[string][]string{
			"DescribeInstances": []string{"213/9388",
				region,
				"vpcId=$vpcId",
			},
			"RunInstances": []string{"213/9384",
				region,
				version,
				instanceChargeType,
				"VirtualPrivateCloud.VpcId=$vpcId",
				"VirtualPrivateCloud.SubnetId=$subnetId",
			},
			"DescribeInstancesStatus":                  []string{""},
			"InquiryPriceRunInstances":                 []string{""},
			"StartInstances":                           []string{""},
			"StopInstances":                            []string{""},
			"TerminateInstances":                       []string{""},
			"RebootInstances":                          []string{""},
			"ResetInstance":                            []string{""},
			"InquiryPriceResetInstance":                []string{""},
			"ResizeInstanceDisks":                      []string{""},
			"InquiryPriceResizeInstanceDisks":          []string{""},
			"RenewInstances":                           []string{""},
			"InquiryPriceRenewInstances":               []string{""},
			"ResetInstancesType":                       []string{""},
			"InquiryPriceResetInstancesType":           []string{""},
			"ModifyInstancesRenewFlag":                 []string{""},
			"ModifyInstancesAttribute":                 []string{""},
			"ResetInstancesInternetMaxBandwidth":       []string{""},
			"ModifyInstancesProject":                   []string{""},
			"UpdateInstanceVpcConfig":                  []string{""},
			"ResetInstancesPassword":                   []string{""},
			"DescribeInstanceInternetBandwidthConfigs": []string{""},
			"DescribeImages":                           []string{""},
			"CreateImage":                              []string{""},
			"DeleteImages":                             []string{""},
			"ModifyImageAttribute":                     []string{""},
			"SyncImages":                               []string{""},
			"ModifyImageSharePermission":               []string{""},
			"DescribeImageSharePermission":             []string{""},
			"AttachNetworkInterface":                   []string{""},
			"DescribeSecurityGroups":                   []string{""},
			"CreateSecurityGroup":                      []string{""},
			"DeleteSecurityGroup":                      []string{""},
			"ModifySecurityGroupAttribute":             []string{""},
			"DescribeSecurityGroupPolicy":              []string{""},
			"ModifySecurityGroupPolicy":                []string{""},
			"DescribeInstancesOfSecurityGroup":         []string{""},
			"ModifySecurityGroupsOfInstance":           []string{""},
			"DescribeAssociateSecurityGroups":          []string{""},
			"DescribeEip":                              []string{""},
			"DescribeEipQuota":                         []string{""},
			"ModifyEipAttributes":                      []string{""},
			"CreateEip":                                []string{""},
			"DeleteEip":                                []string{""},
			"EipBindInstance":                          []string{""},
			"EipUnBindInstance":                        []string{""},
			"TransformWanIpToEip":                      []string{""},
			"DescribeKeyPairs":                         []string{""},
			"CreateKeyPair":                            []string{""},
			"ModifyKeyPairAttribute":                   []string{""},
			"DeleteKeyPairs":                           []string{""},
			"ImportKeyPair":                            []string{""},
			"AssociateInstancesKeyPairs":               []string{""},
			"DisassociateInstancesKeyPairs":            []string{""},
		},
		PkgName: "cvm",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
