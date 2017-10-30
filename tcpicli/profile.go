package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var (
	funcProfile []cli.Command = []cli.Command{
		{
			Name:   "list",
			Usage:  "list",
			Action: ProfileList,
		},
		{
			Name:   "switch",
			Usage:  "switch",
			Action: ProfileSwitch,
		},
		{
			Name:   "create",
			Usage:  "create [name]",
			Action: ProfileCreate,
		},
	}
)

func ProfileList(c *cli.Context) error {
	for _, v := range core.ListProfiles() {
		mark := "-"
		if v.Name() == core.CurrentProfile() {
			mark = "*"
		}
		fmt.Println(mark, v.Name())
	}
	return nil
}

func ProfileSwitch(c *cli.Context) error {
	err := core.SwitchProfile(c.Args().First())
	if err != nil {
		return err
	}
	return ProfileList(nil)
}

func ProfileCreate(c *cli.Context) error {
	name := c.Args().First()
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	}
	var conf core.Configure
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Secrect ID: ")
	conf.SecretId, _ = reader.ReadString('\n')
	conf.SecretId = strings.TrimRight(conf.SecretId, "\n")
	fmt.Print("Secrect Key: ")
	conf.SecretKey, _ = reader.ReadString('\n')
	conf.SecretKey = strings.TrimRight(conf.SecretKey, "\n")
	fmt.Print("Default Region: ")
	conf.DefaultRegion, _ = reader.ReadString('\n')
	conf.DefaultRegion = strings.TrimRight(conf.DefaultRegion, "\n")
	fmt.Print("Use Internal Call (Enter 'true' if you do call in tencent CVM, only work some service like CMQ): ")
	internal, _ := reader.ReadString('\n')
	if internal == "true" {
		conf.Internal = true
	}
	return core.CreateProfile(name, &conf)
}
