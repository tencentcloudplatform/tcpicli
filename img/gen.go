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
	version := "Version=2017-03-12"                        // common
	region := "Region=bj"                                  // common
	vagueInstanceName := "VagueInstanceName=tcpiclicvmgen" // DescribeInstances -> CreateImage
	createImageName := "ImageName=tcpicligenimage"         // CreateImage

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// "DescribeImages",
			`SET instanceId=tcpicli -f "{{range .Response.InstanceSet}}{{.InstanceID}}{{end}}" cvm DescribeInstances ` + region + " " + vagueInstanceName,
			`DO echo $instanceId`,
			// "CreateImage",
			`SET imageId=tcpicli img DescribeImages Filters.1.Name=image-name Filters.1.Values.1=tcpicligenimage ` + region + " " + version + ` | grep -i imageid | awk '{ print $2 }' | tr -d "\","`,
			`DO echo $imageId`,
			"DeleteImages",
			// "ModifyImageAttribute",
			// "SyncImages",
			// "ModifyImageSharePermission",
			// "DescribeImageSharePermission",
		},
		FuncMap: map[string][]string{
			"DescribeImages": []string{"213/9418", version},
			"CreateImage": []string{"213/9415",
				region,
				version,
				createImageName,
				"InstanceId=$instanceId",
			},
			"DeleteImages": []string{"213/9418",
				region,
				version,
				"ImageIds.0=$imageId",
			},
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
