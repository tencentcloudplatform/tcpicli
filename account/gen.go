// +build ignore

package main

import (
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

var ()

func main() {
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeProject",
		},
		FuncMap: map[string][]string{
			"DescribeProject": []string{"378/4400"},
		},
		FixStruct: map[string][]string{},
		PkgName:   "account",
	}
}
