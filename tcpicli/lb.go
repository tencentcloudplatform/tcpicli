package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/lb"
	"github.com/urfave/cli"
)

var (
	funcLb []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: LbDoAction,
		},
		{
			Name:        "InquiryLBPrice",
			Usage:       "Describes price of LB using given configuration",
			Action:      LbInquiryLBPrice,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1328",
		},
		{
			Name:        "CreateLoadBalancer",
			Usage:       "Creates a load balancer of specified configuration",
			Action:      LbCreateLoadBalancer,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1254",
		},
		{
			Name:        "DescribeLoadBalancers",
			Usage:       "Describes existing load balancers",
			Action:      LbDescribeLoadBalancers,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1261",
		},
	}
)

func LbDoAction(c *cli.Context) error {
	resp, err := lb.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func LbInquiryLBPrice(c *cli.Context) error {
	resp, err := lb.InquiryLBPrice(c.Args()...)
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
func LbCreateLoadBalancer(c *cli.Context) error {
	resp, err := lb.CreateLoadBalancer(c.Args()...)
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
func LbDescribeLoadBalancers(c *cli.Context) error {
	resp, err := lb.DescribeLoadBalancers(c.Args()...)
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
