// implement https://cloud.tencent.com/document/api/228/1723

package main

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/cdn"
	"github.com/urfave/cli"
	"strconv"
)

var (
	funcCdn []cli.Command = []cli.Command{
		{
			Name:        "do",
			Usage:       "do <action> <args1=value1> [args2=value2] ...",
			Action:      CdnDoAction,
			Description: "do CDN action and output json response",
		},
		{
			Name:        "AddCdnHost",
			Usage:       "help info",
			Action:      CdnAddCdnHost,
			Description: "referer https://cloud.tencent.com/document/product/228/1406",
		},
		{
			Name:        "DescribeCdnHosts",
			Usage:       "DescribeCdnHosts",
			Action:      CdnDescribeCdnHosts,
			Description: "referer https://cloud.tencent.com/document/product/228/3937",
		},
		{
			Name:        "GetHostInfoByHost",
			Usage:       "GetHostInfoByHost",
			Action:      CdnGetHostInfoByHost,
			Description: "referer https://cloud.tencent.com/document/product/228/3938",
		},
		{
			Name:        "GetHostInfoById",
			Usage:       "GetHostInfoById",
			Action:      CdnGetHostInfoById,
			Description: "referer https://cloud.tencent.com/document/product/228/3939",
		},
		{
			Name:        "GetHostInfo",
			Usage:       "GetHostInfo id/host ...",
			Action:      CdnGetHostInfo,
			Description: "implement GetHostInfoByHost or GetHostInfoById by auto detect args",
		},
		{
			Name:        "OnlineHost",
			Usage:       "OnlineHost id/host",
			Action:      CdnOnlineHost,
			Description: "referer https://cloud.tencent.com/document/product/228/1402",
		},
		{
			Name:        "OfflineHost",
			Usage:       "OfflineHost id/host",
			Action:      CdnOfflineHost,
			Description: "referer https://cloud.tencent.com/document/product/228/1403",
		},
		{
			Name:        "DeleteCdnHost",
			Usage:       "DeleteCdnHost id/host",
			Action:      CdnDeleteCdnHost,
			Description: "referer https://cloud.tencent.com/document/product/228/1396",
		},
		{
			Name:        "UpdateCdnConfig",
			Usage:       "UpdateCdnConfig host=xxx yyy=zzz ...",
			Action:      CdnUpdateCdnConfig,
			Description: "referer https://cloud.tencent.com/document/api/228/3933",
		},
		{
			Name:        "UpdateCache",
			Usage:       "UpdateCache",
			Action:      CdnUpdateCache,
			Description: "referer https://cloud.tencent.com/document/api/228/3934",
		},
		{
			Name:        "RefreshCdnUrl",
			Usage:       "RefreshCdnUrl url1 url2",
			Action:      CdnRefreshCdnUrl,
			Description: "referer https://cloud.tencent.com/document/api/231/3946",
		},
		{
			Name:        "RefreshCdnDir",
			Usage:       "RefreshCdnDir dir1 dir2",
			Action:      CdnRefreshCdnDir,
			Description: "referer https://cloud.tencent.com/document/api/228/3947",
		},
		{
			Name:   "RefreshCdnOverseaUrl",
			Usage:  "RefreshCdnOverseaUrl",
			Action: CdnRefreshCdnOverseaUrl,
		},
		{
			Name:   "RefreshCdnOverseaDir",
			Usage:  "RefreshCdnOverseaDir",
			Action: CdnRefreshCdnOverseaDir,
		},
		{
			Name:        "GetCdnRefreshLog",
			Usage:       "GetCdnRefreshLog",
			Action:      CdnGetCdnRefreshLog,
			Description: "referer https://cloud.tencent.com/document/api/228/3948",
		},
		{
			Name:   "CdnOverseaPushser",
			Usage:  "CdnOverseaPushser",
			Action: CdnCdnOverseaPushser,
		},
		{
			Name:   "GetCdnLogList",
			Usage:  "GetCdnLogList",
			Action: CdnGetCdnLogList,
			Description: `referer https://cloud.tencent.com/document/api/228/8087
# download logs with aria2
tcpicli -f '{{range .Data.List}}{{.Link}}
{{end}}' cdn GetCdnLogList host=$host|aria2c -i-`,
		},

		{
			Name:        "DescribeCdnHostDetailedInfo",
			Usage:       "DescribeCdnHostDetailedInfo",
			Action:      CdnDescribeCdnHostDetailedInfo,
			Description: "referer https://cloud.tencent.com/document/api/228/3942",
		},

		{
			Name:        "GetCdnHostsDetailStatistics",
			Usage:       "GetCdnHostsDetailStatistics",
			Action:      CdnGetCdnHostsDetailStatistics,
			Description: "referer https://cloud.tencent.com/document/api/228/9703",
		},

		{
			Name:        "GetCdnStatusCode",
			Usage:       "GetCdnStatusCode",
			Action:      CdnGetCdnStatusCode,
			Description: "referer https://cloud.tencent.com/document/api/228/3943",
		},

		{
			Name:        "GetCdnHostsHyStat",
			Usage:       "GetCdnHostsHyStat",
			Action:      CdnGetCdnHostsHyStat,
			Description: "referer https://cloud.tencent.com/document/api/228/7352",
		},

		{
			Name:        "GetCdnStatTop",
			Usage:       "GetCdnStatTop",
			Action:      CdnGetCdnStatTop,
			Description: "referer https://cloud.tencent.com/document/api/228/3944",
		},

		{
			Name:        "GetCdnProvIspDetailStat",
			Usage:       "GetCdnProvIspDetailStat",
			Action:      CdnGetCdnProvIspDetailStat,
			Description: "referer https://cloud.tencent.com/document/api/228/7356",
		},

		{
			Name:   "GetCdnMiddleSourceList",
			Usage:  "GetCdnMiddleSourceList",
			Action: CdnGetCdnMiddleSourceList,
			Description: `
referer https://cloud.tencent.com/document/api/
# list without json format
tcpicli -f '{{range .Data.MiddleSourceList}}{{.}}
{{end}}' cdn GetCdnMiddleSourceList format=ip_block
`,
		},
	}
)

func CdnDoAction(c *cli.Context) error {
	resp, err := cdn.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func CdnDescribeCdnHosts(c *cli.Context) error {
	resp, err := cdn.DescribeCdnHosts(c.Args()...)
	if err != nil {
		return err
	}
	b, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func CdnAddCdnHost(c *cli.Context) error {
	resp, err := cdn.AddCdnHost(c.Args()...)
	if err != nil {
		return err
	}
	b, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func CdnGetHostInfoByHost(c *cli.Context) error {
	resp, err := cdn.GetHostInfoByHost(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnGetHostInfoById(c *cli.Context) error {
	resp, err := cdn.GetHostInfoById(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func CdnGetHostInfo(c *cli.Context) error {
	if _, err := strconv.Atoi(c.Args().First()); err == nil {
		return CdnGetHostInfoById(c)
	}
	return CdnGetHostInfoByHost(c)
}

func CdnOnlineHost(c *cli.Context) error {
	prefix := "host="
	if _, err := strconv.Atoi(c.Args().First()); err == nil {
		prefix = "hostId="
	}
	resp, err := cdn.OnlineHost(prefix + c.Args().First())
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnOfflineHost(c *cli.Context) error {
	prefix := "host="
	if _, err := strconv.Atoi(c.Args().First()); err == nil {
		prefix = "hostId="
	}
	resp, err := cdn.OfflineHost(prefix + c.Args().First())
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnDeleteCdnHost(c *cli.Context) error {
	prefix := "host="
	if _, err := strconv.Atoi(c.Args().First()); err == nil {
		prefix = "hostId="
	}
	resp, err := cdn.DeleteCdnHost(prefix + c.Args().First())
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnUpdateCdnConfig(c *cli.Context) error {
	resp, err := cdn.UpdateCdnConfig(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnUpdateCache(c *cli.Context) error {
	resp, err := cdn.UpdateCache(c.Args()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func CdnRefreshCdnUrl(c *cli.Context) error {
	resp, err := cdn.RefreshCdnUrl(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnRefreshCdnDir(c *cli.Context) error {
	resp, err := cdn.RefreshCdnDir(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnRefreshCdnOverseaUrl(c *cli.Context) error {
	resp, err := cdn.RefreshCdnOverSeaUrl(c.Args()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func CdnRefreshCdnOverseaDir(c *cli.Context) error {
	resp, err := cdn.RefreshCdnOverSeaDir(c.Args()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func CdnGetCdnRefreshLog(c *cli.Context) error {
	resp, err := cdn.GetCdnRefreshLog(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func CdnCdnOverseaPushser(c *cli.Context) error {
	resp, err := cdn.CdnOverseaPushser(c.Args()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func CdnGetCdnLogList(c *cli.Context) error {
	resp, err := cdn.GetCdnLogList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnDescribeCdnHostDetailedInfo(c *cli.Context) error {
	resp, err := cdn.DescribeCdnHostDetailedInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnHostsDetailStatistics(c *cli.Context) error {
	resp, err := cdn.GetCdnHostsDetailStatistics(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnStatusCode(c *cli.Context) error {
	resp, err := cdn.GetCdnStatusCode(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnHostsHyStat(c *cli.Context) error {
	resp, err := cdn.GetCdnHostsHyStat(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnStatTop(c *cli.Context) error {
	resp, err := cdn.GetCdnStatTop(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnProvIspDetailStat(c *cli.Context) error {
	resp, err := cdn.GetCdnProvIspDetailStat(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func CdnGetCdnMiddleSourceList(c *cli.Context) error {
	resp, err := cdn.GetCdnMiddleSourceList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
