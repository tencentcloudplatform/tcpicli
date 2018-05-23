package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/cvm/v3.0"
	"github.com/urfave/cli"
)

var (
	funcCvm30 []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: Cvm30DoAction,
		},
		{
			Name:        "DescribeImportImageOs",
			Usage:       "DescribeImportImageOs",
			Action:      Cvm30DescribeImportImageOs,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/15718",
		},
		{
			Name:        "DescribeImageQuota",
			Usage:       "DescribeImageQuota",
			Action:      Cvm30DescribeImageQuota,
			Description: "Referer: https://cloud.tencent.com/document/api/213/15719",
		},
		{
			Name:        "DescribeImages",
			Usage:       "DescribeImages",
			Action:      Cvm30DescribeImages,
			Description: "Referer: https://cloud.tencent.com/document/api/213/15715",
		},
	}
)

func Cvm30DoAction(c *cli.Context) error {
	resp, err := cvm.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func Cvm30DescribeImportImageOs(c *cli.Context) error {
	resp, err := cvm.DescribeImportImageOs(c.Args()...)
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

func Cvm30DescribeImageQuota(c *cli.Context) error {
	resp, err := cvm.DescribeImageQuota(c.Args()...)
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

func Cvm30DescribeImages(c *cli.Context) error {
	resp, err := cvm.DescribeImages(c.Args()...)
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
