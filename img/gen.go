// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	version := "Version=2017-03-12" // common

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeImages",
			// "CreateImage",
			// "DeleteImages",
			// "ModifyImageAttribute",
			// "SyncImages",
			// "ModifyImageSharePermission",
			// "DescribeImageSharePermission",
		},
		FuncMap: map[string][]string{
			"DescribeImages":               []string{"213/9418", version},
			"CreateImage":                  []string{""},
			"DeleteImages":                 []string{""},
			"ModifyImageAttribute":         []string{""},
			"SyncImages":                   []string{""},
			"ModifyImageSharePermission":   []string{""},
			"DescribeImageSharePermission": []string{""},
		},
		PkgName: "img",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
