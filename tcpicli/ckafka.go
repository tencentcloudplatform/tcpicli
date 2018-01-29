package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/ckafka"
	"github.com/urfave/cli"
)

var (
	funcCkafka []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: CkafkaDoAction,
		},
	}
)

func CkafkaDoAction(c *cli.Context) error {
	resp, err := ckafka.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
