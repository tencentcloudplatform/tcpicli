package main

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/cvm"
	"github.com/urfave/cli"
)

var (
	funcCvm []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action",
			Action: CvmDoAction,
		},
		{
			Name:   "DescribeRegions",
			Usage:  "describes regions",
			Action: CvmDescribeRegions,
		},
		{
			Name:   "DescribeZones",
			Usage:  "describes zones",
			Action: CvmDescribeZones,
		},
		{
			Name:   "RunInstances",
			Usage:  "Run an instance from an image",
			Action: CvmRunInstances,
		},
		{
			Name:   "StopInstances",
			Usage:  "Stop a running instance",
			Action: CvmStopInstances,
		},
		{
			Name:   "StartInstances",
			Usage:  "Start a stopped instance",
			Action: CvmStartInstances,
		},
		{
			Name:   "DescribeInstances",
			Usage:  "Describes instance",
			Action: CvmDescribeInstances,
		},
		{
			Name:   "DescribeInstancesStatus",
			Usage:  "Describes instance",
			Action: CvmDescribeInstancesStatus,
		},
		{
			Name:   "TerminateInstances",
			Usage:  "Terminates a powered-down instance",
			Action: CvmTerminateInstances,
		},
		{
			Name:   "ResetInstances",
			Usage:  "Reset a running instance",
			Action: CvmResetInstances,
		},
		{
			Name:   "RebootInstances",
			Usage:  "Reboot a running instance",
			Action: CvmRebootInstances,
		},
		{
			Name:   "ResizeInstanceDisks",
			Usage:  "Reboot a running instance",
			Action: CvmResizeInstanceDisks,
		},
	}
)

func CvmDoAction(c *cli.Context) error {
	resp, err := cvm.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func CvmDescribeRegions(c *cli.Context) error {
	resp, err := cvm.DescribeRegions(c.Args()...)
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

func CvmDescribeZones(c *cli.Context) error {
	resp, err := cvm.DescribeZones(c.Args()...)
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

func CvmRunInstances(c *cli.Context) error {
	resp, err := cvm.RunInstances(c.Args()...)
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

func CvmStopInstances(c *cli.Context) error {
	resp, err := cvm.StopInstances(c.Args()...)
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

func CvmDescribeInstances(c *cli.Context) error {
	resp, err := cvm.DescribeInstances(c.Args()...)
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

func CvmDescribeInstancesStatus(c *cli.Context) error {
	resp, err := cvm.DescribeInstancesStatus(c.Args()...)
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

func CvmStartInstances(c *cli.Context) error {
	resp, err := cvm.StartInstances(c.Args()...)
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

func CvmTerminateInstances(c *cli.Context) error {
	resp, err := cvm.TerminateInstances(c.Args()...)
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

func CvmResetInstances(c *cli.Context) error {
	resp, err := cvm.ResetInstances(c.Args()...)
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

func CvmRebootInstances(c *cli.Context) error {
	resp, err := cvm.RebootInstances(c.Args()...)
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

func CvmResizeInstanceDisks(c *cli.Context) error {
	resp, err := cvm.ResizeInstanceDisks(c.Args()...)
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
