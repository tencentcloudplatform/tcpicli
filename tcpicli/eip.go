package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/eip"
	"github.com/urfave/cli"
)

var (
	funcEip []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: EipDoAction,
		},
		{
			Name:        "DescribeEip",
			Usage:       "Describes allocated EIPs",
			Action:      EipDescribeEip,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1379",
		},
		{
			Name:        "DescribeEipQuota",
			Usage:       "Describes EIP quota in specified region",
			Action:      EipDescribeEipQuota,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1378",
		},
		{
			Name:        "CreateEip",
			Usage:       "Creates EIP",
			Action:      EipCreateEip,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1381",
		},
		{
			Name:        "DeleteEip",
			Usage:       "Delete specified EIP",
			Action:      EipDeleteEip,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1380",
		},
		{
			Name:        "ModifyEipAttributes",
			Usage:       "Changed the name of EIP",
			Action:      EipModifyEipAttributes,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1375",
		},
		{
			Name:        "EipBindInstance",
			Usage:       "Bind existing EIP to existing instance",
			Action:      EipEipBindInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1377",
		},
		{
			Name:        "EipUnBindInstance",
			Usage:       "UnBind EIP from instance",
			Action:      EipEipUnBindInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/1376",
		},
	}
)

func EipDoAction(c *cli.Context) error {
	resp, err := eip.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func EipDescribeEip(c *cli.Context) error {
	resp, err := eip.DescribeEip(c.Args()...)
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
func EipDescribeEipQuota(c *cli.Context) error {
	resp, err := eip.DescribeEipQuota(c.Args()...)
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
func EipCreateEip(c *cli.Context) error {
	resp, err := eip.CreateEip(c.Args()...)
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
func EipDeleteEip(c *cli.Context) error {
	resp, err := eip.DeleteEip(c.Args()...)
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
func EipModifyEipAttributes(c *cli.Context) error {
	resp, err := eip.DescribeEipQuota(c.Args()...)
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
func EipEipBindInstance(c *cli.Context) error {
	resp, err := eip.DescribeEipQuota(c.Args()...)
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
func EipEipUnBindInstance(c *cli.Context) error {
	resp, err := eip.DescribeEipQuota(c.Args()...)
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
