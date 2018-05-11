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

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			"DescribeUserInfo",
			"DescribeDealsByCond",
			"DescribeAccountBalance",
			// "PayDeals", don't get this one
		},
		FuncMap: map[string][]string{
			"DescribeUserInfo": []string{"378/4397",
				region,
			},
			"DescribeDealsByCond": []string{"378/4403",
				region,
				"startTime=2018-01-01",
				"endTime=2018-01-02",
			},
			"DescribeAccountBalance": []string{"378/4397",
				region,
			},
			"PayDeals": []string{"378/4394",
				region,
				"dealIds.1=18003",
			},
		},
		PkgName: "trade",
		Pkg:     new(Pkg),
	}
	gen.Run()
}
