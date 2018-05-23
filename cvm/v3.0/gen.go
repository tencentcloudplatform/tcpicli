// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
	"log"
	"os"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	client := NewClient()
	client.SetLog(os.Stderr, "[debug] ", log.LstdFlags|log.Lshortfile)
	return client.DoAction(action, query...)
}

func main() {
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			`DescribeImportImageOs`,
			`DescribeImageQuota`,
			`DescribeImages`,
		},
		FuncMap: map[string][]string{
			"DescribeImportImageOs": []string{"213/15718"},
			"DescribeImageQuota":    []string{"213/15719"},
			"DescribeImages":        []string{"213/15715"},
		},
		PkgName: "cvm",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
