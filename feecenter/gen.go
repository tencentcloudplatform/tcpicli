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
	payerUin := "payerUin=100000655192" // china account uin

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api",
		Seq: []string{
			`DescribeBills`,
			`DescribeResourceBills`,
		},
		FuncMap: map[string][]string{
			"DescribeBills": []string{"378/10770",
				"startTime=2017-11-01",
				"endTime=2017-11-02",
				payerUin,
			},
			"DescribeResourceBills": []string{"378/10772",
				"startMonth=2017-11",
				"endMonth=2017-12",
			},
		},
		PkgName: "feecenter",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
