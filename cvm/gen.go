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
	// vpcName := "vpcName=tcpiclicvmautogen" // vpc
	// cidrBlock := "cidrBlock=10.0.0.0/16"                        // vpc
	// subnetCidr := "subnetSet.0.cidrBlock=10.0.0.0/24"           // vpc
	// subnetZone := "subnetSet.0.zoneId=800002"                   // vpc
	// subnetName := "subnetSet.0.subnetName=tcpiclicvmautogen"    // vpc
	projectId := "ProjectId=1068783"        // CreateKeyPair
	keyName := "tcpiclikey"                 // CreateKeyPair
	keyNameIn := "KeyName=" + keyName       // CreateKeyPair
	keyNameMod := "tcpiclikeymod"           // ModifyKeyPairAttribute
	keyNameModIn := "KeyName=" + keyNameMod // ModifyKeyPairAttribute
	region := "Region=bj"                   // common
	publicKey := "PublicKey=ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDHoY5NPtrqFyzeZOdSnxn5AA+m3A/vO4REZuVClyKdNbog7CK5RYK6VPcKXc6Dm/RiHkrxp07L7Yw5o80KKSdX2i/trgJ8O4tkmFlpZ0HNGGks/9nbqZDg7h4kLniaYphhhHFckaaM9BbLKnjjjx7NOIBLjLJlQypLpk1DaxM0FFnqvdW/RwMIywMzEAcBoFb2TsQ+Obk0dF2/wEiQt93VppAK/Jua4Xu5q6CFhhWXA2buzw2g1Pd6vPJ+OoIR+QRw4ZfjuH3tDuZXzOr3h11EWJcZQ/apG/fnMoQjng8iTgF+RPcVLZJQ6Jwq1a4JpyG3KL4kRkwFScvLh0cwrCwR"
	instanceChargeType := "InstanceChargeType=POSTPAID_BY_HOUR" // common
	version := "Version=2017-03-12"                             // common
	instanceName := "InstanceName=tcpiclicvmgen"                // RunInstances
	instanceType := "InstanceType=S2.SMALL1"                    // RunInstances
	// vagueInstanceName := "VagueInstanceName=tcpiclicvmgen"                            // DescribeInstances
	diskResize := "DataDisks.0.DiskSize=100"                                          // ResizeInstanceDisks
	diskResizeInquiry := "DataDisks.0.DiskSize=150"                                   // InquiryPriceResizeInstanceDisks
	instanceTypeResize := "InstanceType=S2.SMALL2"                                    // ResizeInstancesType / InquiryPriceResizeInstancesType
	instanceNameMod := "InstanceName=tcpiclicvmgenmod"                                // ModyInstancesAttribute
	internetAccessableMaxBandwidth := "InternetAccessable.InternetMaxBandwidthOut=10" // ResetInstancesInternetMaxBandwidth

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Version: map[string][]string{
			"DescribeInstances": {"InstanceSet"},
		},
		Seq: []string{
			// CreateVpc & CreateSubnet to run autogen cvm stuff on and capture their id and unId
			// `DO tcpicli vpc CreateVpc ` + region + " " + vpcName + " " + cidrBlock,
			// `SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName,
			// `DO echo $vpcId`,
			// `DO tcpicli vpc CreateSubnet vpcId=$vpcId ` + region + " " + subnetCidr + " " + subnetName + " " + subnetZone,
			// `SET subnetId=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			// `DO echo $subnetId`,
			// `SET placementZone=tcpicli -f "{{with index .Data 0}}{{.Zone}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			// `DO echo $placementZone`,
			// `SET imageId=tcpicli img DescribeImages Version=2017-03-12 | grep -E -B1 -i "centos.*7\.3.*64" | grep -i imageid | awk '{ print $2 }' | tr -d "\"," | head -n1`,
			// `DO echo $imageId`,
			// "RunInstances",
			"DescribeInstances",
			// `SET instanceId=tcpicli -f '{{range .Response.InstanceSet}}{{if eq .InstanceName "tcpiclicvmgenmod"}}{{.InstanceID}}{{end}}{{end}}' cvm DescribeInstances ` + region,
			// `DO echo $instanceId`,
			// `SET diskId=tcpicli -f "{{range .Response.InstanceSet}}{{range .DataDisks}}{{.DiskId}}{{end}}{{end}}" cvm DescribeInstances ` + region + " " + vagueInstanceName,
			// `DO echo $diskId`,
			// "DescribeInstancesStatus",
			// "InquiryPriceRunInstances",
			// "StartInstances",
			// "StopInstances",
			// "TerminateInstances",
			// "RebootInstances",
			// "ResetInstance",
			// "InquiryPriceResetInstance",
			// "ResizeInstanceDisks",
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
			// "DescribeKeyPairs",
			// "CreateKeyPair",
			// `SET keyId=tcpicli -f '{{range .Response.KeyPairSet}}{{if eq .KeyName "` + keyName + `"}}{{.KeyID}}{{end}}{{end}}' cvm DescribeKeyPairs ` + version + " " + region,
			// `DO echo $keyId`,
			// "ModifyKeyPairAttribute",
			// "DeleteKeyPairs",
			// "ImportKeyPair",
			// "AssociateInstancesKeyPairs",
			// "DisassociateInstancesKeyPairs",
			// Delete VPC stuff
			// `DO tcpicli vpc DeleteSubnet vpcId=$vpcId subnetId=$subnetId ` + region,
			// `DO tcpicli vpc DeleteVpc vpcId=$vpcId ` + region,
		},
		FuncMap: map[string][]string{
			"RunInstances": []string{"213/9384",
				region,
				version,
				instanceChargeType,
				instanceName,
				instanceType,
				"VirtualPrivateCloud.VpcId=$vpcId",
				"VirtualPrivateCloud.SubnetId=$subnetId",
				"ImageId=$imageId",
				"Placement.Zone=$placementZone",
			},
			"DescribeInstances": []string{"213/9388",
				region,
			},
			"DescribeInstancesStatus": []string{"213/9389",
				version,
			},
			"InquiryPriceRunInstances": []string{"213/9385",
				region,
				version,
				instanceType,
				"Placement.Zone=$placementZone",
				"ImageId=$imageId",
			},
			"StartInstances": []string{"213/9386",
				region,
				version,
				"InstanceIds.1=$instanceId",
			},
			"StopInstances": []string{"213/9383",
				region,
				version,
				"InstanceIds.1=$instanceId",
			},
			"TerminateInstances": []string{"213/9395",
				region,
				version,
				"InstanceIds.0=$instanceId",
			},
			"RebootInstances": []string{"213/9369",
				region,
				version,
				"InstanceIds.1=$instanceId",
			},
			"ResetInstance": []string{"213/9398",
				region,
				version,
				"InstanceId=$instanceId",
			},
			"InquiryPriceResetInstance": []string{"213/9490",
				region,
				version,
				"InstanceId=$instanceId",
			},
			"ResizeInstanceDisks": []string{"213/9387",
				region,
				version,
				diskResize,
				"InstanceId=$instanceId",
			},
			"InquiryPriceResizeInstanceDisks": []string{"213/9487",
				region,
				version,
				diskResizeInquiry,
				"InstanceId=$instanceId",
			},
			"RenewInstances": []string{""},
			"InquiryPriceRenewInstances": []string{"213/9491",
				region,
				version,
				"InstanceIds.1=$instanceId",
			},
			"ResetInstancesType": []string{"213/9394",
				region,
				version,
				instanceTypeResize,
				"InstanceIds.1=$instanceId",
			},
			"InquiryPriceResetInstancesType": []string{"213/9489",
				region,
				version,
				instanceTypeResize,
				"InstanceIds.1=$instanceId",
			},
			"ModifyInstancesRenewFlag": []string{""},
			"ModifyInstancesAttribute": []string{"213/9381",
				region,
				version,
				instanceNameMod,
				"InstanceIds.1=$instanceId",
			},
			"ResetInstancesInternetMaxBandwidth": []string{"213/9393",
				region,
				version,
				internetAccessableMaxBandwidth,
				"InstanceIds.1=$instanceId",
			},
			"ModifyInstancesProject":  []string{""},
			"UpdateInstanceVpcConfig": []string{""},
			"ResetInstancesPassword":  []string{""},
			"DescribeInstanceInternetBandwidthConfigs": []string{"213/9390",
				version,
				region,
				"InstanceId=$instanceId",
			},
			"DescribeKeyPairs": []string{"213/9403",
				region,
				version,
			},
			"CreateKeyPair": []string{"213/9400",
				region,
				version,
				projectId,
				keyNameIn,
			},
			"ModifyKeyPairAttribute": []string{"213/9399",
				region,
				version,
				keyNameModIn,
				"KeyId=$keyId",
			},
			"DeleteKeyPairs": []string{"213/9401",
				region,
				version,
				"KeyIds.1=$keyId",
			},
			"ImportKeyPair": []string{"213/9402",
				version,
				projectId,
				publicKey,
				keyNameIn,
			},
			"AssociateInstancesKeyPairs":    []string{""},
			"DisassociateInstancesKeyPairs": []string{""},
		},
		PkgName: "cvm",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
