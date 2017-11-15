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
	domain := fmt.Sprintf("%d.cdn.building9s.io", time.Now().Unix())
	origin := "tencent.com"
	startDate := "startDate=" + time.Now().Add(-48*time.Hour).Format("20060102")
	endDate := "endDate=" + time.Now().Add(-24*time.Hour).Format("20060102")
	startTime := "startTime=" + time.Now().Add(-48*time.Hour).Format("2006-01-02 00:00:00")
	endTime := "endTime=" + time.Now().Add(-24*time.Hour).Format("2006-01-02 00:00:00")
	date := "date=" + time.Now().Add(-24*time.Hour).Format("2006-01-02")
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			`SET projectId=tcpicli -f '{{range .Data}}{{if eq .ProjectName "I01 kiyor"}}{{.ProjectID}}{{end}}{{end}}' account DescribeProject`,
			"AddCdnHost",
			"DO sleep 30",
			"SET hostId=tcpicli -f '{{with (index .Data.Hosts 0)}}{{.HostID}}{{end}}' cdn GetHostInfoByHost " + domain,
			"UpdateCdnConfig",
			"DescribeCdnHosts",
			"GetHostInfoByHost",
			"GetHostInfoById",
			"RefreshCdnUrl",
			"RefreshCdnDir",
			"SET taskId=tcpicli -f {{.Data.TaskID}} cdn RefreshCdnUrl http://" + domain + "/file",
			"DO sleep 30",
			"GetCdnRefreshLog",
			"DescribeCdnHostInfo",
			"DescribeCdnHostDetailedInfo",
			"GetCdnHostsDetailStatistics",
			"GetCdnStatusCode",
			"GetCdnHostsHyStat",
			"GetCdnStatTop",
			"GetCdnProvIspDetailStat",
			"GetCdnLogList",
			"DO sleep 300",
			"OfflineHost",
			"DO sleep 300",
			"OnlineHost",
			"DO sleep 300",
			"OfflineHost",
			"DO sleep 300",
			"DeleteCdnHost",
		},
		SliceOptions: map[string]string{
			"GetHostInfoByHost": "hosts",
			"RefreshCdnUrl":     "urls",
			"RefreshCdnDir":     "dirs",
			"GetHostInfoById":   "ids",
		},
		FuncMap: map[string][]string{
			"AddCdnHost": []string{"228/1406", "host=" + domain, "projectId=0", "hostType=cname", "origin=" + origin},
			"UpdateCdnConfig": []string{"228/3933",
				// 				"host=" + domain,
				"hostId=$hostId",
				`advancedCache=[[4,"",0],[0,"all",1000],[1,".jpg;.js",2000],[2,"/www/html",3000]]`,
				`advancedMaxage=[[0,"all",60]]`,
				`accessIp={"type":1,"list":["1.3.3.4","2.3.4.5"]}`,
				`refer=[1,["qq.baidu.com", "*.baidu.com"],1]`,
				`backupOrigin=qq.com`,
				// 			`statusCache=[[403,"403"]]`,
			},
			"UpdateCdnProject": []string{"228/3935", "projectId=$projectId", "hostId=$hostId"},
			"DescribeCdnHosts": []string{"228/3937"},
			// 		"GetHostInfoByHost": []string{"228/3938", "hosts.0=t1.cdn.building9s.io"},
			"GetHostInfoByHost": []string{"228/3938", "hosts.0=" + domain},
			"GetHostInfoById":   []string{"228/3938", "ids.0=$hostId"},
			"OfflineHost":       []string{"228/1403", "host=" + domain},
			"OnlineHost":        []string{"228/1402", "host=" + domain},
			"DeleteCdnHost":     []string{"228/1396", "host=" + domain},

			"RefreshCdnUrl":    []string{"228/3946", "urls.0=http://" + domain + "/1.jpg"},
			"RefreshCdnDir":    []string{"228/3947", "dirs.0=http://" + domain + "/abc/"},
			"GetCdnRefreshLog": []string{"228/3948", "taskId=$taskId"},

			"DescribeCdnHostInfo":         []string{"228/3941", startDate, endDate, "statType=flux", "projects.0=$projectId"},
			"DescribeCdnHostDetailedInfo": []string{"228/3942", startDate, endDate, "statType=flux", "projects.0=$projectId"},
			"GetCdnHostsDetailStatistics": []string{"228/9703", startTime, endTime, "statType=flux", "hosts.0=" + domain},
			"GetCdnStatusCode":            []string{"228/3943", startDate, endDate, "projects.0=$projectId"},
			"GetCdnHostsHyStat":           []string{"228/7352", date, "statType=bandwidth"},
			"GetCdnStatTop":               []string{"228/3944", startDate, endDate, "statType=flux", "projects.0=$projectId"},
			"GetCdnProvIspDetailStat":     []string{"228/7356", date, "hosts.0=" + domain},
			"GetCdnLogList":               []string{"228/8087", "host=" + domain},
		},
		FixStruct: map[string][]string{
			"DescribeCdnHosts":            []string{"cos_key", "interface{}", "pid_config", "interface{}", "final_pids", "interface{}"},
			"GetCdnHostsHyStat":           []string{"data", "interface{}"},
			"GetCdnStatusCode":            []string{"data", "interface{}"},
			"GetCdnHostsDetailStatistics": []string{"data", "interface{}"},
		},
		PkgName: "cdn",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
