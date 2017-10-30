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
