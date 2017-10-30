package cdn

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
	"github.com/wsxiaoys/terminal/color"
	"log"
	"os"
	"testing"
	"time"
)

type CdnTests struct{ *testing.T }

var testdomain = fmt.Sprintf("%d.dnsv1.com", time.Now().Unix())

func TestMain(m *testing.M) {
	core.Log = log.New(os.Stderr, color.Sprintf("@{y}[core]@{|} "), log.LstdFlags|log.Lshortfile)
	log.SetPrefix(color.Sprintf("@{c}[test]@{|} "))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	os.Exit(m.Run())
}

func sleep(s string) {
	t, _ := time.ParseDuration(s)
	scale := 100
	log.Println("sleep", t)
	for i := 1; i <= scale; i++ {
		a := "["
		for j := 1; j <= i; j++ {
			a += "x"
		}
		for j := 1; j <= scale-i; j++ {
			a += " "
		}
		b := ""
		switch i % 4 {
		case 0:
			b = `-`
		case 1:
			b = `\`
		case 2:
			b = `|`
		case 3:
			b = `/`
		}
		a += fmt.Sprintf("] %s %d%%", b, i)
		fmt.Printf("\r%s", a)
		time.Sleep(t / time.Duration(scale))
	}
	fmt.Println()
}

func TestCdn(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		test := CdnTests{t}
		test.TestAddCdnHost(testdomain)
		test.TestGetHostInfoByHost(testdomain)
	})
	sleep("1m")
	t.Run("update test", func(t *testing.T) {
		test := CdnTests{t}
		test.TestUpdateCdnConfig(testdomain)
	})
	sleep("2s")
	t.Run("offline", func(t *testing.T) {
		test := CdnTests{t}
		test.TestOfflineCdnHost(testdomain)
	})
	sleep("5m")
	t.Run("delete", func(t *testing.T) {
		test := CdnTests{t}
		test.TestDeleteCdnHost(testdomain)
	})
}

func (t *CdnTests) TestAddCdnHost(domain string) {
	// 	b, err := AddCdnHost("host="+testdomain, "projectId=0", "hostType=cname", "origin="+testdomain, "fwdHost="+testdomain, `cache=[[4,"",0],[0,"all",1000],[1,".jpg;.js",2000],[2,"/www/html",3000]]`)
	b, err := AddCdnHost("host="+domain, "projectId=0", "hostType=cname", "origin=origin-"+domain, "fwdHost=origin-"+domain, `advancedCache=[[4,"",0],[0,"all",1000],[1,".jpg;.js",2000],[2,"/www/html",3000]]`)
	if err != nil {
		t.Fail()
	}
	log.Println(string(b))
}

func (t *CdnTests) TestOfflineCdnHost(domain string) {
	b, err := OfflineHost(domain)
	if err != nil {
		t.Fail()
	}
	log.Println(string(b))
}

func (t *CdnTests) TestDeleteCdnHost(domain string) {
	b, err := DeleteCdnHost(domain)
	if err != nil {
		t.Fail()
	}
	log.Println(string(b))
}
