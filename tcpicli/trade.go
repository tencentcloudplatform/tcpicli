package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/trade"
	"github.com/urfave/cli"
)

var (
	funcTrade []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: TradeDoAction,
		},
	}
)

func TradeDoAction(c *cli.Context) error {
	resp, err := trade.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
