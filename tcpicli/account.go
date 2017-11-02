package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/account"
	"github.com/tencentcloudplatform/tcpicli/core"
	"github.com/urfave/cli"
)

var (
	funcAccount []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: AccountDoAction,
		},
		{
			Name:   "DescribeProject",
			Usage:  "DescribeProject",
			Action: AccountDescribeProject,
		},
	}
)

func AccountDoAction(c *cli.Context) error {
	resp, err := core.DoAction("account", c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func AccountDescribeProject(c *cli.Context) error {
	resp, err := account.DescribeProject(c.Args()...)
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
