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
		{
			Name:        "CreateQueue",
			Usage:       "Creates a new message queue",
			Action:      CmqCreateQueue,
			Description: "referer https://cloud.tencent.com/document/api/431/5832",
		},
		{
			Name:        "ListQueue",
			Usage:       "Lists existing messages queues in a given region",
			Action:      CmqListQueue,
			Description: "referer https://cloud.tencent.com/document/api/406/5833",
		},
		{
			Name:        "GetQueueAttributes",
			Usage:       "Gets specific attributes about a particular message queue",
			Action:      CmqGetQueueAttributes,
			Description: "referer https://cloud.tencent.com/document/api/406/5834",
		},
		{
			Name:        "SetQueueAttributes",
			Usage:       "Changes a given attribute of a message queue",
			Action:      CmqSetQueueAttributes,
			Description: "referer https://cloud.tencent.com/document/api/406/5835",
		},
		{
			Name:        "SendMessage",
			Usage:       "sends a message to an existing queue",
			Action:      CmqSendMessage,
			Description: "referer https://cloud.tencent.com/document/api/406/5835",
		},
		{
			Name:        "ReceiveMessage",
			Usage:       "Receives a message in specified message queue",
			Action:      CmqReceiveMessage,
			Description: "referer https://cloud.tencent.com/document/api/406/5839",
		},
		{
			Name:        "DeleteQueue",
			Usage:       "Deletes specified message queue",
			Action:      CmqDeleteQueue,
			Description: "referer https://cloud.tencent.com/document/api/406/5836",
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
func CmqCreateQueue(c *cli.Context) error {
	resp, err := cmq.CreateQueue(c.Args()...)
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
func CmqListQueue(c *cli.Context) error {
	resp, err := cmq.ListQueue(c.Args()...)
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
func CmqGetQueueAttributes(c *cli.Context) error {
	resp, err := cmq.GetQueueAttributes(c.Args()...)
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
func CmqSetQueueAttributes(c *cli.Context) error {
	resp, err := cmq.SetQueueAttributes(c.Args()...)
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
func CmqSendMessage(c *cli.Context) error {
	resp, err := cmq.SendMessage(c.Args()...)
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
func CmqReceiveMessage(c *cli.Context) error {
	resp, err := cmq.ReceiveMessage(c.Args()...)
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
func CmqDeleteQueue(c *cli.Context) error {
	resp, err := cmq.DeleteQueue(c.Args()...)
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
