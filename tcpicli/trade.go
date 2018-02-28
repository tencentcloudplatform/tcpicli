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
		{
			Name:        "DescribeUserInfo",
			Usage:       "Describes user info",
			Action:      TradeDescribeUserInfo,
			Description: "Referrer: https://cloud.tencent.com/document/api/378/4397",
		},
		{
			Name:        "DescribeDealsByCond",
			Usage:       "Describes user info",
			Action:      TradeDescribeDealsByCond,
			Description: "Referrer: https://cloud.tencent.com/document/api/378/4403",
		},
		{
			Name:        "DescribeAccountBalance",
			Usage:       "Describes account balance",
			Action:      TradeDescribeAccountBalance,
			Description: "Referrer: https://cloud.tencent.com/document/api/378/4397",
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

func TradeDescribeUserInfo(c *cli.Context) error {
	resp, err := trade.DescribeUserInfo(c.Args()...)
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
func TradeDescribeDealsByCond(c *cli.Context) error {
	resp, err := trade.DescribeDealsByCond(c.Args()...)
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
func TradeDescribeAccountBalance(c *cli.Context) error {
	resp, err := trade.DescribeAccountBalance(c.Args()...)
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
