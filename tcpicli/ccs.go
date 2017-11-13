package main

import (
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
		{
			Name:        "CreateClusterService",
			Usage:       "Creates a new service on a cluster.",
			Action:      CcsCreateClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9436",
		},
		{
			Name:        "DescribeClusterService",
			Usage:       "Describes service.",
			Action:      CcsDescribeClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9440",
		},
		{
			Name:        "DescribeClusterServiceInfo",
			Usage:       "Details particular service.",
			Action:      CcsDescribeClusterServiceInfo,
			Description: "referer: https://cloud.tencent.com/document/api/457/9441",
		},
		{
			Name:        "DescribeServiceEvent",
			Usage:       "Details particular service events in last hour.",
			Action:      CcsDescribeServiceEvent,
			Description: "referer: https://cloud.tencent.com/document/api/457/9443",
		},
		{
			Name:        "RollBackClusterService",
			Usage:       "Roll-back service modification. Only supports one iteration",
			Action:      CcsRollBackClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9438",
		},
		{
			Name:        "DescribeServiceInstance",
			Usage:       "Lists nodes that are used in a particular service",
			Action:      CcsDescribeServiceInstance,
			Description: "referer: https://cloud.tencent.com/document/api/457/9433",
		},
		{
			Name:        "ModifyServiceReplicas",
			Usage:       "Resizes pod replicas in service",
			Action:      CcsModifyServiceReplicas,
			Description: "referer: https://cloud.tencent.com/document/api/457/9431",
		},
		{
			Name:        "DeleteInstances",
			Usage:       "Deletes pods from service",
			Action:      CcsDeleteInstances,
			Description: "referer: https://cloud.tencent.com/document/api/457/9432",
		},
		{
			Name:        "ModifyClusterService",
			Usage:       "Modifies service on a cluster.",
			Action:      CcsModifyClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9434",
		},
		{
			Name:        "PauseClusterService",
			Usage:       "Pauses modification on service.",
			Action:      CcsPauseClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9439",
		},
		{
			Name:        "ResumeClusterService",
			Usage:       "Resumes paused modification.",
			Action:      CcsResumeClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9442",
		},
		{
			Name:        "ModifyServiceDescription", // just add a description to the service. No need to complicate code by changing serviceName or anything like that
			Usage:       "Modifies descriptive data on a service.",
			Action:      CcsModifyServiceDescription,
			Description: "referer: https://cloud.tencent.com/document/api/457/9435",
		},
		{
			Name:        "DeleteClusterService",
			Usage:       "Details particular service.",
			Action:      CcsDeleteClusterService,
			Description: "referer: https://cloud.tencent.com/document/api/457/9437",
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
func CcsCreateClusterService(c *cli.Context) error {
	resp, err := ccs.CreateClusterService(c.Args()...)
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
func CcsDescribeClusterService(c *cli.Context) error {
	resp, err := ccs.DescribeClusterService(c.Args()...)
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
func CcsDescribeClusterServiceInfo(c *cli.Context) error {
	resp, err := ccs.DescribeClusterServiceInfo(c.Args()...)
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
func CcsModifyClusterService(c *cli.Context) error {
	resp, err := ccs.ModifyClusterService(c.Args()...)
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
func CcsPauseClusterService(c *cli.Context) error {
	resp, err := ccs.PauseClusterService(c.Args()...)
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
func CcsResumeClusterService(c *cli.Context) error {
	resp, err := ccs.ResumeClusterService(c.Args()...)
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
func CcsModifyServiceDescription(c *cli.Context) error {
	resp, err := ccs.ModifyServiceDescription(c.Args()...)
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
func CcsDescribeServiceEvent(c *cli.Context) error {
	resp, err := ccs.DescribeServiceEvent(c.Args()...)
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
func CcsRollBackClusterService(c *cli.Context) error {
	resp, err := ccs.RollBackClusterService(c.Args()...)
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
func CcsDescribeServiceInstance(c *cli.Context) error {
	resp, err := ccs.DescribeServiceInstance(c.Args()...)
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
func CcsModifyServiceReplicas(c *cli.Context) error {
	resp, err := ccs.ModifyServiceReplicas(c.Args()...)
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
func CcsDeleteInstances(c *cli.Context) error {
	resp, err := ccs.DeleteInstances(c.Args()...)
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
func CcsDeleteClusterService(c *cli.Context) error {
	resp, err := ccs.DeleteClusterService(c.Args()...)
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
