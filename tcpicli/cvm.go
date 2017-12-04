package main

import (
	"encoding/json"
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
			Name:        "DescribeInstances",
			Usage:       "Describes available instances",
			Action:      CvmDescribeInstances,
			Describtion: "Referrer: https://cloud.tencent.com/document/api/213/9388",
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
