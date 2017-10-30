package main

import (
	"encoding/json"
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
			Name:   "DescribeImages",
			Usage:  "Describes available images",
			Action: ImgDescribeImages,
		},
		{
			Name:   "CreateImage",
			Usage:  "Creates images from instances",
			Action: ImgCreateImage,
		},
		{
			Name:   "DeleteImages",
			Usage:  "Creates images from instances",
			Action: ImgDeleteImages,
		},
		{
			Name:   "SyncImages",
			Usage:  "Syncs an Image from one Region to another",
			Action: ImgSyncImages,
		},
		{
			Name:   "DescribeImageSharePermission",
			Usage:  "Describes share permissions",
			Action: ImgDescribeImageSharePermission,
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
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ImgCreateImage(c *cli.Context) error {
	resp, err := img.CreateImage(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ImgDeleteImages(c *cli.Context) error {
	resp, err := img.DeleteImages(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ImgSyncImages(c *cli.Context) error {
	resp, err := img.SyncImages(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ImgDescribeImageSharePermission(c *cli.Context) error {
	resp, err := img.DescribeImageSharePermission(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
