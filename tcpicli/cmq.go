// implement https://cloud.tencent.com/document/api/406/5851

package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/cmq"
	"github.com/urfave/cli"
)

var (
	funcCmq []cli.Command = []cli.Command{
		{
			Name:        "do",
			Usage:       "do <action> <args1=value1> [args2=value2] ...",
			Action:      CmqDoAction,
			Description: "do CMQ action and output json response",
		},
		// {
		// 	Name:        "CreateQueue",
		// 	Usage:       "CreateQueue",
		// 	Action:      CmqCreateQueue,
		// 	Description: "referer https://cloud.tencent.com/document/api/431/5832",
		// },
		// {
		// 	Name:        "ListQueue",
		// 	Usage:       "ListQueue",
		// 	Action:      CmqListQueue,
		// 	Description: "referer https://cloud.tencent.com/document/api/406/5833",
		// },
		// {
		// 	Name:        "DeleteQueue",
		// 	Usage:       "DeleteQueue",
		// 	Action:      CmqDeleteQueue,
		// 	Description: "referer https://cloud.tencent.com/document/api/406/5836",
		// },
	}
)

func CmqDoAction(c *cli.Context) error {
	resp, err := cmq.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

// func CmqCreateQueue(c *cli.Context) error {
// 	resp, err := cmq.CreateQueue(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	r, err := resp.String(formatOut)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(r)
// 	return nil
// }
//
// func CmqListQueue(c *cli.Context) error {
// 	resp, err := cmq.ListQueue(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	r, err := resp.String(formatOut)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(r)
// 	return nil
// }
//
// func CmqDeleteQueue(c *cli.Context) error {
// 	resp, err := cmq.DeleteQueue(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	r, err := resp.String(formatOut)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(r)
// 	return nil
// }
