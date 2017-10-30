package main

import (
	"fmt"
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
