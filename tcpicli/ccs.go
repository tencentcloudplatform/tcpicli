package main

import (
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/ccs"
	"github.com/urfave/cli"
)

var (
	funcCcs []cli.Command = []cli.Command{
		{
			Name:        "do",
			Usage:       "do action",
			Action:      CcsDoAction,
			Description: "do ccs action and output raw json response",
		},
		{
			Name:        "CreateCluster",
			Usage:       "Creates a new CCS cluster",
			Action:      CcsCreateCluster,
			Description: "referer: https://cloud.tencent.com/document/api/457/9444",
		},
		{
			Name:        "AddClusterInstances",
			Usage:       "Adds a new CVM kubernetes node to CCS cluster",
			Action:      CcsAddClusterInstances,
			Description: "referer: https://cloud.tencent.com/document/api/457/9447",
		},
		{
			Name:        "AddClusterInstancesFromExistedCvm",
			Usage:       "Adds an existing CVM as a kubernetes node to a CCS cluster",
			Action:      CcsAddClusterInstancesFromExistedCvm,
			Description: "referer: https://cloud.tencent.com/document/api/457/9450",
		},
		{
			Name:        "DescribeCluster",
			Usage:       "Queries CCS cluster list",
			Action:      CcsDescribeCluster,
			Description: "referer: https://cloud.tencent.com/document/api/457/9448",
		},
		{
			Name:        "DescribeClusterInstances",
			Usage:       "Queries individual CCS cluster and returns node information",
			Action:      CcsDescribeClusterInstances,
			Description: "referer: https://cloud.tencent.com/document/api/457/9449",
		},
		{
			Name:        "DeleteClusterInstances",
			Usage:       "Removes a CVM node from a CCS cluster. If CVM is postpaid, it can be terminated after removal.",
			Action:      CcsDeleteClusterInstances,
			Description: "referer: https://cloud.tencent.com/document/api/457/9446",
		},
	}
)

func CcsDoAction(c *cli.Context) error {
	resp, err := ccs.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func CcsCreateCluster(c *cli.Context) error {
	resp, err := ccs.CreateCluster(c.Args()...)
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
func CcsAddClusterInstances(c *cli.Context) error {
	resp, err := ccs.AddClusterInstances(c.Args()...)
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
func CcsAddClusterInstancesFromExistedCvm(c *cli.Context) error {
	resp, err := ccs.AddClusterInstances(c.Args()...)
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
func CcsDescribeCluster(c *cli.Context) error {
	resp, err := ccs.DescribeCluster(c.Args()...)
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
func CcsDescribeClusterInstances(c *cli.Context) error {
	resp, err := ccs.DescribeClusterInstances(c.Args()...)
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
func CcsDeleteClusterInstances(c *cli.Context) error {
	resp, err := ccs.DeleteClusterInstances(c.Args()...)
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
