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
		{
			Name:        "ListInstance",
			Usage:       "List ckafka instances",
			Action:      CkafkaListInstance,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10093",
		},
		{
			Name:        "GetInstanceAttributes",
			Usage:       "Details specific information about specified ckafka instance",
			Action:      CkafkaGetInstanceAttributes,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10094",
		},
		{
			Name:        "CreateTopic",
			Usage:       "Creates a new kafka topic",
			Action:      CkafkaCreateTopic,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10096",
		},
		{
			Name:        "GetTopicAttributes",
			Usage:       "Details specific information about specified ckafka topic",
			Action:      CkafkaGetTopicAttributes,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10102",
		},
		{
			Name:        "DeleteTopic",
			Usage:       "Deletes specified kafka topic",
			Action:      CkafkaDeleteTopic,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10094",
		},
		{
			Name:        "AddPartition",
			Usage:       "Adds partition to specified kafka topic",
			Action:      CkafkaDeleteTopic,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10100",
		},
		{
			Name:        "SetTopicAttributes",
			Usage:       "Sets attributes of specified topic",
			Action:      CkafkaSetTopicAttributes,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10098",
		},
		{
			Name:        "ListTopic",
			Usage:       "Lists all topics on specified instance",
			Action:      CkafkaListTopic,
			Description: "Referrer: https://cloud.tencent.com/document/api/597/10101",
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
func CkafkaListInstance(c *cli.Context) error {
	resp, err := ckafka.ListInstance(c.Args()...)
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
func CkafkaGetInstanceAttributes(c *cli.Context) error {
	resp, err := ckafka.GetInstanceAttributes(c.Args()...)
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
func CkafkaCreateTopic(c *cli.Context) error {
	resp, err := ckafka.CreateTopic(c.Args()...)
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
func CkafkaGetTopicAttributes(c *cli.Context) error {
	resp, err := ckafka.GetTopicAttributes(c.Args()...)
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
func CkafkaDeleteTopic(c *cli.Context) error {
	resp, err := ckafka.DeleteTopic(c.Args()...)
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
func CkafkaListTopic(c *cli.Context) error {
	resp, err := ckafka.ListTopic(c.Args()...)
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
func CkafkaSetTopicAttributes(c *cli.Context) error {
	resp, err := ckafka.SetTopicAttributes(c.Args()...)
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
func CkafkaAddPartition(c *cli.Context) error {
	resp, err := ckafka.AddPartition(c.Args()...)
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
