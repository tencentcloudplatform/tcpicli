package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/feecenter"
	"github.com/urfave/cli"
)

var (
	funcFeecenter []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: FeecenterDoAction,
		},
	}
)

func FeecenterDoAction(c *cli.Context) error {
	resp, err := feecenter.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
