// +build ignore

package main

import (
	. "."
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/autogen"
	"time"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	queryName := fmt.Sprintf("query%d", time.Now().Unix())
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"CreateQueue",
			"ListQueue",
			"DeleteQueue",
		},
		FuncMap: map[string][]string{
			"CreateQueue": []string{"431/5832", "queueName=" + queryName},
			"ListQueue":   []string{"406/5833"},
			"DeleteQueue": []string{"406/5836", "queueName=" + queryName},
		},
		PkgName: "cmq",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
