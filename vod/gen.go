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
	region := "Region=gz"
	videoType := "videoType=mkv"
	coverType := "coverType=jpg"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			`SET vodSessionKey=tcpicli -f "{{.VodSessionKey}}" vod ApplyUpload ` + region + " " + videoType + " " + coverType,
			"CommitUpload",
		},
		FuncMap: map[string][]string{
			"ApplyUpload": []string{"266/9756",
				region,
				videoType,
				coverType,
			},
			"CommitUpload": []string{"266/9757",
				region,
				"vodSessionKey=$vodSessionKey",
			},
		},
		PkgName: "vod",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
