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
		{
			Name:        "DescribeVodPlayInfo",
			Usage:       "DescribeVodPlayInfo",
			Action:      VodDescribeVodPlayInfo,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7825",
		},
		{
			Name:        "GetVideoInfo",
			Usage:       "GetVideoInfo",
			Action:      VodGetVideoInfo,
			Description: "Referer: https://cloud.tencent.com/document/api/266/8586",
		},
		{
			Name:        "CreateClass",
			Usage:       "CreateClass",
			Action:      VodCreateClass,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7812",
		},
		{
			Name:        "DescribeAllClass",
			Usage:       "DescribeAllClass",
			Action:      VodDescribeAllClass,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7813",
		},
		{
			Name:        "DescribeClass",
			Usage:       "DescribeClass",
			Action:      VodDescribeClass,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7814",
		},
		{
			Name:        "ModifyClass",
			Usage:       "ModifyClass",
			Action:      VodModifyClass,
			Description: "Referer: https://cloud.tencent.com/document/api/266/7815",
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
func VodDescribeVodPlayInfo(c *cli.Context) error {
	resp, err := vod.DescribeVodPlayInfo(c.Args()...)
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
func VodGetVideoInfo(c *cli.Context) error {
	resp, err := vod.GetVideoInfo(c.Args()...)
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
func VodCreateClass(c *cli.Context) error {
	resp, err := vod.CreateClass(c.Args()...)
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
func VodDescribeAllClass(c *cli.Context) error {
	resp, err := vod.DescribeAllClass(c.Args()...)
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
func VodDescribeClass(c *cli.Context) error {
	resp, err := vod.DescribeClass(c.Args()...)
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
func VodModifyClass(c *cli.Context) error {
	resp, err := vod.ModifyClass(c.Args()...)
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
