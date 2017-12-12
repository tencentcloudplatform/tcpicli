package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/cvm"
	"github.com/urfave/cli"
)

var (
	funcCvm []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: CvmDoAction,
		},
		{
			Name:        "RunInstances",
			Usage:       "Starts new instances",
			Action:      CvmRunInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9384",
		},
		{
			Name:        "DescribeInstances",
			Usage:       "Describes available instances",
			Action:      CvmDescribeInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9388",
		},
		{
			Name:        "DescribeInstancesStatus",
			Usage:       "Displays status of instances",
			Action:      CvmDescribeInstancesStatus,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9389",
		},
		{
			Name:        "StopInstances",
			Usage:       "Stops a running instance",
			Action:      CvmStopInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9383",
		},
		{
			Name:        "StartInstances",
			Usage:       "Starts a stopped instance",
			Action:      CvmStartInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9386",
		},
		{
			Name:        "RebootInstances",
			Usage:       "Reboots a running instance",
			Action:      CvmRebootInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9396",
		},
		{
			Name:        "ResetInstance",
			Usage:       "Resets a running instance",
			Action:      CvmResetInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9396",
		},
		{
			Name:        "InquiryPriceRunInstances",
			Usage:       "Checks the price for running instance",
			Action:      CvmInquiryPriceRunInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9385",
		},
		{
			Name:        "InquiryPriceResetInstance",
			Usage:       "Resets a running instance to original configuration",
			Action:      CvmInquiryPriceResetInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9490",
		},
		{
			Name:        "ResizeInstanceDisks",
			Usage:       "Resizes cloud data disk attached to instance",
			Action:      CvmResizeInstanceDisks,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9387",
		},
		{
			Name:        "InquiryPriceResizeInstanceDisks",
			Usage:       "Checks the cost of resizing instance disks",
			Action:      CvmInquiryPriceResizeInstanceDisks,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9487",
		},
		{
			Name:        "RenewInstances",
			Usage:       "Renews a prepaid instance",
			Action:      CvmRenewInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9392",
		},
		{
			Name:        "InquiryPriceRenewInstances",
			Usage:       "Checks the price of renewing a prepaid instance",
			Action:      CvmInquiryPriceRenewInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9491",
		},
		{
			Name:        "ResetInstancesType",
			Usage:       "Resizes an instance",
			Action:      CvmResetInstancesType,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9394",
		},
		{
			Name:        "InquiryPriceResetInstancesType",
			Usage:       "Checks the price of resizing an instance",
			Action:      CvmInquiryPriceResetInstancesType,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9489",
		},
		{
			Name:        "ModifyInstancesAttribute",
			Usage:       "Changes the name tag of an instance",
			Action:      CvmModifyInstancesAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9381",
		},
		{
			Name:        "ResetInstancesInternetMaxBandwidth",
			Usage:       "Changes the max bandwidth of an instance",
			Action:      CvmResetInstancesInternetMaxBandwidth,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9393",
		},
		{
			Name:        "UpdateInstanceVpcConfig",
			Usage:       "Changes the VPC config of an instance",
			Action:      CvmUpdateInstanceVpcConfig,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9379",
		},
		{
			Name:        "ResetInstancesPassword",
			Usage:       "Changes the instance password",
			Action:      CvmResetInstancesPassword,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9397",
		},
		{
			Name:        "DescribeInstanceInternetBandwidthConfigs",
			Usage:       "Describes prepaid instances bandwidth configuration",
			Action:      CvmDescribeInstanceInternetBandwidthConfigs,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9390",
		},
		{
			Name:        "TerminateInstances",
			Usage:       "Terminates an instance",
			Action:      CvmTerminateInstances,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9395",
		},
		{
			Name:        "DescribeKeyPairs",
			Usage:       "Describes key pairs",
			Action:      CvmDescribeKeyPairs,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9403",
		},
		{
			Name:        "CreateKeyPair",
			Usage:       "Creates new key pairs",
			Action:      CvmCreateKeyPair,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9400",
		},
		{
			Name:        "ModifyKeyPairAttribute",
			Usage:       "Modifies attribute of key pair",
			Action:      CvmModifyKeyPairAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9403",
		},
		{
			Name:        "ImportKeyPair",
			Usage:       "Imports an existing public key",
			Action:      CvmImportKeyPair,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9402",
		},
		{
			Name:        "AssociateInstancesKeyPairs",
			Usage:       "binds a key pair to an instance",
			Action:      CvmAssociateInstancesKeyPairs,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9404",
		},
		{
			Name:        "DisassociateInstancesKeyPairs",
			Usage:       "unbinds a key pair to an instance",
			Action:      CvmDisassociateInstancesKeyPairs,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9405",
		},
	}
)

func CvmDoAction(c *cli.Context) error {
	resp, err := cvm.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func CvmRunInstances(c *cli.Context) error {
	resp, err := cvm.RunInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmDescribeInstances(c *cli.Context) error {
	resp, err := cvm.DescribeInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmDescribeInstancesStatus(c *cli.Context) error {
	resp, err := cvm.DescribeInstancesStatus(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmStopInstances(c *cli.Context) error {
	resp, err := cvm.StopInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmStartInstances(c *cli.Context) error {
	resp, err := cvm.StartInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmRebootInstances(c *cli.Context) error {
	resp, err := cvm.RebootInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmResetInstance(c *cli.Context) error {
	resp, err := cvm.ResetInstance(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmInquiryPriceRunInstances(c *cli.Context) error {
	resp, err := cvm.InquiryPriceRunInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmInquiryPriceResetInstance(c *cli.Context) error {
	resp, err := cvm.InquiryPriceResetInstance(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmResizeInstanceDisks(c *cli.Context) error {
	resp, err := cvm.ResizeInstanceDisks(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmInquiryPriceResizeInstanceDisks(c *cli.Context) error {
	resp, err := cvm.InquiryPriceResizeInstanceDisks(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmRenewInstances(c *cli.Context) error {
	resp, err := cvm.RenewInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmInquiryPriceRenewInstances(c *cli.Context) error {
	resp, err := cvm.InquiryPriceRenewInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmResetInstancesType(c *cli.Context) error {
	resp, err := cvm.ResetInstancesType(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmInquiryPriceResetInstancesType(c *cli.Context) error {
	resp, err := cvm.InquiryPriceResetInstancesType(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmModifyInstancesAttribute(c *cli.Context) error {
	resp, err := cvm.ModifyInstancesAttribute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmResetInstancesInternetMaxBandwidth(c *cli.Context) error {
	resp, err := cvm.ResetInstancesInternetMaxBandwidth(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmUpdateInstanceVpcConfig(c *cli.Context) error {
	resp, err := cvm.UpdateInstanceVpcConfig(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmResetInstancesPassword(c *cli.Context) error {
	resp, err := cvm.ResetInstancesPassword(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmDescribeInstanceInternetBandwidthConfigs(c *cli.Context) error {
	resp, err := cvm.DescribeInstanceInternetBandwidthConfigs(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmTerminateInstances(c *cli.Context) error {
	resp, err := cvm.TerminateInstances(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmDescribeKeyPairs(c *cli.Context) error {
	resp, err := cvm.DescribeKeyPairs(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmCreateKeyPair(c *cli.Context) error {
	resp, err := cvm.CreateKeyPair(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmModifyKeyPairAttribute(c *cli.Context) error {
	resp, err := cvm.ModifyKeyPairAttribute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmImportKeyPair(c *cli.Context) error {
	resp, err := cvm.ImportKeyPair(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmAssociateInstancesKeyPairs(c *cli.Context) error {
	resp, err := cvm.AssociateInstancesKeyPairs(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CvmDisassociateInstancesKeyPairs(c *cli.Context) error {
	resp, err := cvm.DisassociateInstancesKeyPairs(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
