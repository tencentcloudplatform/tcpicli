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
