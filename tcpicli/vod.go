package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/vod"
	"github.com/urfave/cli"
)

var (
	funcVod []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action and output raw json response.",
			Action: VodDoAction,
		},
		{
			Name:        "ApplyUpload",
			Usage:       "ApplyUpload",
			Action:      VodApplyUpload,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7788",
		},
	}
)

func VodDoAction(c *cli.Context) error {
	resp, err := vod.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func VodApplyUpload(c *cli.Context) error {
	resp, err := vod.ApplyUpload(c.Args()...)
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
