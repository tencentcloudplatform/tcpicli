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
	fileName := "fileName=test0"
	fileId := "fileId=4564972818475166253"
	className := "className=tcpicliclass"
	classNameMod := className + "new"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// -- hold on upload stuff because it has COS API dependency --
			// "ApplyUpload",
			// "CommitUpload",

			"DescribeVodPlayInfo",
			"GetVideoInfo",

			// "CreateClass",
			// "DescribeAllClass",
			// "DescribeClass",
			// `SET classId=tcpicli -f "{{range .Data}}{{.Info.ID}}{{end}}" vod DescribeAllClass ` + region + " " + className,
			// "ModifyClass",
			// "DeleteClass", // doesn't work for some reason. classId=$classId is giving error 1000; invalid param... is API error, not tcpicli bug
		},
		FuncMap: map[string][]string{
			"DescribeVodPlayInfo": []string{"266/7825",
				fileName,
			},
			"GetVideoInfo": []string{"266/8586",
				fileId,
			},
			"CreateClass": []string{"266/7812",
				className,
			},
			"ModifyClass": []string{"266/7815",
				classNameMod,
				"classId=$classId",
			},
			"DescribeAllClass": []string{"266/7813"},
			"DescribeClass":    []string{"266/7814"},
			"DeleteClass": []string{"266/7816",
				"classId=$classId",
			},
		},
		PkgName: "vod",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
