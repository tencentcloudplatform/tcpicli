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
		},
		FuncMap: map[string][]string{
			"DescribeImportImageOs": []string{"213/15718"},
		},
		PkgName: "cvm",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
