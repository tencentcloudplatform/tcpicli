package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/img"
	"github.com/urfave/cli"
)

var (
	funcImg []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: ImgDoAction,
		},
		{
			Name:        "DescribeImages",
			Usage:       "Describes available images",
			Action:      ImgDescribeImages,
			Description: "Referrer: https://cloud.tencent.com/document/api/213/9418",
		},
	}
)

func ImgDoAction(c *cli.Context) error {
	resp, err := img.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func ImgDescribeImages(c *cli.Context) error {
	resp, err := img.DescribeImages(c.Args()...)
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
