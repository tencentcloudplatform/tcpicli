package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/dfw"
	"github.com/urfave/cli"
)

var (
	funcDfw []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: DfwDoAction,
		},
		{
			Name:        "CreateSecurityGroup",
			Usage:       "create a security group",
			Action:      DfwCreateSecurityGroup,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1238",
		},
		{
			Name:        "DescribeSecurityGroupEx",
			Usage:       "Describes existing security groups",
			Action:      DfwDescribeSecurityGroupEx,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1232",
		},
		{
			Name:        "DeleteSecurityGroup",
			Usage:       "Deletes existing security groups",
			Action:      DfwDeleteSecurityGroup,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1362",
		},
		{
			Name:        "DescribeSecurityGroupPolicys",
			Usage:       "Describes specified security group policy",
			Action:      DfwDescribeSecurityGroupPolicys,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1364",
		},
		{
			Name:        "ModifySecurityGroupPolicys",
			Usage:       "Modifies specified security group policy",
			Action:      DfwModifySecurityGroupPolicys,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1365",
		},
		{
			Name:        "DescribeInstancesOfSecurityGroup",
			Usage:       "Lists instances bound to specified security group",
			Action:      DfwDescribeInstancesOfSecurityGroup,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1366",
		},
		{
			Name:        "ModifySecurityGroupsOfInstance",
			Usage:       "Modify Security Groups bound to particular instance",
			Action:      DfwModifySecurityGroupsOfInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1367",
		},
	}
)

func DfwDoAction(c *cli.Context) error {
	resp, err := dfw.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func DfwCreateSecurityGroup(c *cli.Context) error {
	resp, err := dfw.CreateSecurityGroup(c.Args()...)
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
func DfwDescribeSecurityGroupEx(c *cli.Context) error {
	resp, err := dfw.DescribeSecurityGroupEx(c.Args()...)
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
func DfwDeleteSecurityGroup(c *cli.Context) error {
	resp, err := dfw.DeleteSecurityGroup(c.Args()...)
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
func DfwDescribeSecurityGroupPolicys(c *cli.Context) error {
	resp, err := dfw.DescribeSecurityGroupPolicys(c.Args()...)
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
func DfwModifySecurityGroupPolicys(c *cli.Context) error {
	resp, err := dfw.ModifySecurityGroupPolicys(c.Args()...)
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
func DfwDescribeInstancesOfSecurityGroup(c *cli.Context) error {
	resp, err := dfw.DescribeInstancesOfSecurityGroup(c.Args()...)
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
func DfwModifySecurityGroupsOfInstance(c *cli.Context) error {
	resp, err := dfw.ModifySecurityGroupsOfInstance(c.Args()...)
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
