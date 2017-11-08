package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/ccs"
	"github.com/urfave/cli"
)

var (
	funcCcs []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: CcsDoAction,
		},
		{
			Name:        "DescribeCluster",
			Usage:       "Query clusters",
			Action:      CcsDescribeCluster,
			Description: "referer: https://www.qcloud.com/document/api/457/9448",
		},
		// {
		// 	Name:        "DescribeClusterInstances",
		// 	Usage:       "Query cluster CVMs",
		// 	Action:      CcsDescribeCluster,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9449",
		// },
		{
			Name:        "CreateCluster",
			Usage:       "Create new cluster",
			Action:      CcsCreateCluster,
			Description: "referer: https://www.qcloud.com/document/api/457/9444",
		},
		// {
		// 	Name:        "CreateClusterService",
		// 	Usage:       "Create new cluster service",
		// 	Action:      CcsCreateClusterService,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9436",
		// },
		// {
		// 	Name:        "DeleteCluster",
		// 	Usage:       "Delete a cluster",
		// 	Action:      CcsDeleteCluster,
		// 	Description: "referer: https://cloud.tencent.com/document/api/457/9445",
		// },
		// {
		// 	Name:        "AddClusterInstances",
		// 	Usage:       "Add node to existing cluster",
		// 	Action:      CcsAddClusterInstances,
		// 	Description: "referer: https://cloud.tencent.com/document/api/457/9447",
		// },
		// {
		// 	Name:        "DescribeClusterService",
		// 	Usage:       "Describes services",
		// 	Action:      CcsDescribeClusterService,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9444",
		// },
		// {
		// 	Name:        "DescribeClusterServiceInfo",
		// 	Usage:       "Describes service details",
		// 	Action:      CcsDescribeClusterServiceInfo,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9441",
		// },
		// {
		// 	Name:        "ModifyClusterService",
		// 	Usage:       "Modifies existing service specifications",
		// 	Action:      CcsModifyClusterService,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9434",
		// },
		// {
		// 	Name:        "DeleteClusterService",
		// 	Usage:       "Deletes existing service",
		// 	Action:      CcsDeleteClusterService,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9437",
		// },
		// {
		// 	Name:        "DescribeServiceInstance",
		// 	Usage:       "Describes pods running in service",
		// 	Action:      CcsDescribeServiceInstance,
		// 	Description: "referer: https://www.qcloud.com/document/api/457/9433",
		// },
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

// func CcsDescribeClusterInstances(c *cli.Context) error {
// 	resp, err := ccs.DescribeClusterInstances(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

func CcsCreateCluster(c *cli.Context) error {
	resp, err := ccs.CreateCluster(c.Args()...)
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

// func CcsCreateClusterService(c *cli.Context) error {
// 	resp, err := ccs.CreateClusterService(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsDeleteCluster(c *cli.Context) error {
// 	resp, err := ccs.DeleteCluster(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsAddClusterInstances(c *cli.Context) error {
// 	resp, err := ccs.AddClusterInstances(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsDescribeClusterService(c *cli.Context) error {
// 	resp, err := ccs.DescribeClusterService(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsDescribeClusterServiceInfo(c *cli.Context) error {
// 	resp, err := ccs.DescribeClusterServiceInfo(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsModifyClusterService(c *cli.Context) error {
// 	resp, err := ccs.ModifyClusterService(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsDeleteClusterService(c *cli.Context) error {
// 	resp, err := ccs.DeleteClusterService(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }

// func CcsDescribeServiceInstance(c *cli.Context) error {
// 	resp, err := ccs.DescribeServiceInstance(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		return err
// 	}
//
// 	fmt.Println(string(b))
// 	return nil
// }
