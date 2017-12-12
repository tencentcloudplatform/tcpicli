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
		{
			Name:        "ModifyLoadBalancerAttributes",
			Usage:       "Change LB attributes",
			Action:      LbModifyLoadBalancerAttributes,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1263",
		},
		{
			Name:        "DeleteLoadBalancers",
			Usage:       "Deletes LB",
			Action:      LbDeleteLoadBalancers,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1257",
		},
		{
			Name:        "CreateLoadBalancerListeners",
			Usage:       "Creates a listener on a LB",
			Action:      LbCreateLoadBalancerListeners,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1255",
		},
		{
			Name:        "DescribeLoadBalancerListeners",
			Usage:       "Describes a listener on a LB",
			Action:      LbDescribeLoadBalancerListeners,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1260",
		},
		{
			Name:        "ModifyLoadBalancerListener",
			Usage:       "Modifies a listener on a LB",
			Action:      LbModifyLoadBalancerListener,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/3601",
		},
		{
			Name:        "DeleteLoadBalancerListeners",
			Usage:       "Deletes a listener from a LB",
			Action:      LbDeleteLoadBalancerListeners,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1256",
		},
		{
			Name:        "RegisterInstancesWithLoadBalancer",
			Usage:       "Binds a CLB with a backend (CVM)",
			Action:      LbRegisterInstancesWithLoadBalancer,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1256",
		},
		{
			Name:        "DescribeLoadBalancerBackends",
			Usage:       "Details backends on specified CLB",
			Action:      LbDescribeLoadBalancerBackends,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1259",
		},
		{
			Name:        "ModifyLoadBalancerBackends",
			Usage:       "Modifies backend details on specified CLB",
			Action:      LbModifyLoadBalancerBackends,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1264",
		},
		{
			Name:        "DeregisterInstancesFromLoadBalancer",
			Usage:       "Removes backends from specified CLB",
			Action:      LbDeregisterInstancesFromLoadBalancer,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1258",
		},
		{
			Name:        "DescribeLBHealthStatus",
			Usage:       "Describes CLB status",
			Action:      LbDescribeLBHealthStatus,
			Description: "Referrer: https://cloud.tencent.com/document/api/214/1326",
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
func LbModifyLoadBalancerAttributes(c *cli.Context) error {
	resp, err := lb.ModifyLoadBalancerAttributes(c.Args()...)
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
func LbDeleteLoadBalancers(c *cli.Context) error {
	resp, err := lb.DeleteLoadBalancers(c.Args()...)
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
func LbCreateLoadBalancerListeners(c *cli.Context) error {
	resp, err := lb.CreateLoadBalancerListeners(c.Args()...)
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
func LbDescribeLoadBalancerListeners(c *cli.Context) error {
	resp, err := lb.DescribeLoadBalancerListeners(c.Args()...)
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
func LbModifyLoadBalancerListener(c *cli.Context) error {
	resp, err := lb.ModifyLoadBalancerListener(c.Args()...)
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
func LbDeleteLoadBalancerListeners(c *cli.Context) error {
	resp, err := lb.DeleteLoadBalancerListeners(c.Args()...)
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
func LbRegisterInstancesWithLoadBalancer(c *cli.Context) error {
	resp, err := lb.RegisterInstancesWithLoadBalancer(c.Args()...)
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
func LbDescribeLoadBalancerBackends(c *cli.Context) error {
	resp, err := lb.DescribeLoadBalancerBackends(c.Args()...)
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
func LbModifyLoadBalancerBackends(c *cli.Context) error {
	resp, err := lb.ModifyLoadBalancerBackends(c.Args()...)
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
func LbDeregisterInstancesFromLoadBalancer(c *cli.Context) error {
	resp, err := lb.DeregisterInstancesFromLoadBalancer(c.Args()...)
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
func LbDescribeLBHealthStatus(c *cli.Context) error {
	resp, err := lb.DescribeLBHealthStatus(c.Args()...)
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
