package main

import (
	"encoding/json"
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
			Name:   "CreateSecurityGroup",
			Usage:  "create a security group",
			Action: DfwCreateSecurityGroup,
		},
		{
			Name:   "DescribeSecurityGroupEx",
			Usage:  "Describes existing security groups",
			Action: DfwDescribeSecurityGroupEx,
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
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func DfwDescribeSecurityGroupEx(c *cli.Context) error {
	resp, err := dfw.DescribeSecurityGroupEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
