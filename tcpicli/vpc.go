package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/vpc"
	"github.com/urfave/cli"
)

var (
	funcVpc []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action and output raw json response.",
			Action: VpcDoAction,
		},
		{
			Name:        "CreateVpc",
			Usage:       "Creates a new VPC.",
			Action:      VpcCreateVpc,
			Description: "Referrer: https://cloud.tencent.com/document/215/1309",
		},
		//{
		//	Name:        "action",
		//	Usage:       "",
		//	Action:      Vpcaction,
		//	Description: "Referrer: https://cloud.tencent.com/document/api/uri",
		//},
	}
)

func VpcDoAction(c *cli.Context) error {
	resp, err := vpc.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func VpcCreateVpc(c *cli.Context) error {
	resp, err := vpc.CreateVpc(c.Args()...)
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

//func VpcACTION(c *cli.Context) error {
//	resp, err := vpc.ACTION(c.Args()...)
//	if err != nil {
//		return err
//	}
//	r, err := resp.String(formatOut)
//	if err != nil {
//		return err
//	}
//	fmt.Println(r)
//	return nil
//}
