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
		{
			Name:        "DescribeVpcEx",
			Usage:       "Queries VPC list.",
			Action:      VpcDescribeVpcEx,
			Description: "Referrer: https://cloud.tencent.com/document/215/1372",
		},
		{
			Name:        "DescribeVpcClassicLink",
			Usage:       "Describes classic link between VPC and CVM",
			Action:      VpcDescribeVpcClassicLink,
			Description: "Referrer: https://cloud.tencent.com/document/215/1844",
		},
		{
			Name:        "DescribeVpcLimit",
			Usage:       "Describes VPC limits.",
			Action:      VpcDescribeVpcLimit,
			Description: "Referrer: https://cloud.tencent.com/document/215/2112",
		},
		{
			Name:        "ModifyVpcAttribute",
			Usage:       "Modifies VPC attribute",
			Action:      VpcModifyVpcAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/215/1310",
		},
		{
			Name:        "AssociateVip",
			Usage:       "Binds EIP to instance",
			Action:      VpcAssociateVip,
			Description: "Referrer: https://cloud.tencent.com/document/215/1361",
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
func VpcDescribeVpcEx(c *cli.Context) error {
	resp, err := vpc.DescribeVpcEx(c.Args()...)
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
func VpcDescribeVpcClassicLink(c *cli.Context) error {
	resp, err := vpc.DescribeVpcClassicLink(c.Args()...)
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
func VpcDescribeVpcLimit(c *cli.Context) error {
	resp, err := vpc.DescribeVpcLimit(c.Args()...)
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
func VpcModifyVpcAttribute(c *cli.Context) error {
	resp, err := vpc.ModifyVpcAttribute(c.Args()...)
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
func VpcAssociateVip(c *cli.Context) error {
	resp, err := vpc.AssociateVip(c.Args()...)
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
