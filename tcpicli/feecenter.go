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
		{
			Name:        "DescribeBills",
			Usage:       "Describes bills on account",
			Action:      FeecenterDescribeBills,
			Description: "referrer: https://cloud.tencent.com/document/api/378/10770",
		},
		{
			Name:        "DescribeResourceBills",
			Usage:       "Describes resource bills on account",
			Action:      FeecenterDescribeResourceBills,
			Description: "referrer: https://cloud.tencent.com/document/api/378/10772",
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
func FeecenterDescribeBills(c *cli.Context) error {
	resp, err := feecenter.DescribeBills(c.Args()...)
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
func FeecenterDescribeResourceBills(c *cli.Context) error {
	resp, err := feecenter.DescribeResourceBills(c.Args()...)
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
