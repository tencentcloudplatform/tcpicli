package main

import (
	"errors"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
	"github.com/urfave/cli"
	"log"
	"os"
)

var VER = "1.0.0"
var buildtime string
var formatOut string

func init() {
}

func main() {
	app := cli.NewApp()
	app.Name = "tcpicli"
	app.Usage = "tencent cloud platform cli tool"
	app.EnableBashCompletion = true
	app.Version = VER + "." + buildtime
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "vv", Usage: "verbose"},
		cli.StringFlag{Name: "f", Usage: `filter output
		Example:
		tcpicli -f="The status is: {{.code}}" ccs DescribeCluster
		tcpicli -f="clusterCIDR: {{range .data.clusters}}{{.clusterCIDR}}{{end}}" ccs DescribeCluster
		tcpicli -f="/path/to/file ccs DescribeCluster"
		`},
	}
	app.Before = before
	app.Commands = []cli.Command{
		{
			Name:        "profile",
			Usage:       "profile",
			Subcommands: funcProfile,
		},
		{
			Name:   "do",
			Usage:  "do <service> <action> <args1=value1> [args2=value2] ...",
			Action: do,
			Description: `do ANY action and output json response
Example:
tcpicli do cdn GetHostInfoByHost hosts.0=www.test.com`,
		},
		{
			Name:        "vpc",
			Usage:       "vpc function",
			Subcommands: funcVpc,
		},
		{
			Name:        "cdn",
			Usage:       "cdn function",
			Subcommands: funcCdn,
		},
		{
			Name:        "cvm",
			Usage:       "cvm function",
			Subcommands: funcCvm,
		},
		{
			Name:        "img",
			Usage:       "img function",
			Subcommands: funcImg,
		},
		{
			Name:        "ccs",
			Usage:       "ccs function",
			Subcommands: funcCcs,
		},
		{
			Name:        "cmq",
			Usage:       "cmq function",
			Subcommands: funcCmq,
		},

		{
			Name:        "dfw",
			Usage:       "Cloud FireWall",
			Subcommands: funcDfw,
		},
		{
			Name:        "account",
			Usage:       "account function",
			Subcommands: funcAccount,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func before(c *cli.Context) error {
	formatFlag := c.String("f")
	formatOut = formatFlag

	/*
		if len(formatFlag) > 0 {
			core.Inspect(formatFlag, c.Args()[0], c.Args()[1], c.Args()[2:]...)
			os.Exit(0)
		}
	*/
	if c.Bool("vv") {
		core.Log = log.New(os.Stderr, "[core] ", log.LstdFlags|log.Lshortfile)
	}
	return nil
}

func do(c *cli.Context) error {
	if len(c.Args()) < 2 {
		return errors.New(fmt.Sprintf("%v service action options...", os.Args[0]))
	}
	b, err := core.DoAction(c.Args()[0], c.Args()[1], c.Args()[2:]...)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
