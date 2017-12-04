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
	region := "Region=bj" // common input param

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeInstances",
			// "DescribeInstancesStatus",
			// "RunInstances",
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
		},
		FuncMap: map[string][]string{
			"RunProcedure": []string{"266/11030",
				region,
				inputType,
				fileStartTimeOffset,
				fileEndTimeOffset,
				procedure,
				"file.id=$fileId",
			},
		},
		PkgName: "cvm",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
