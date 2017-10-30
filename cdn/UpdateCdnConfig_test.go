package cdn

import (
	"log"
)

func (t *CdnTests) TestUpdateCdnConfig(domain string) {
	b, err := UpdateCdnConfig("host="+domain, `advancedCache=[[4,"",0],[0,"all",1000],[1,".jpg;.js",2000],[2,"/www/html",3000]]`, `advancedMaxage=[[0,"all",60]]`)
	if err != nil {
		t.Fail()
	}
	log.Println(string(b))
	m, err := GetHostInfoByHost(testdomain)
	if err != nil {
		t.Fail()
	}
	log.Println(m.Data.Hosts[0].HostConfig.AdvancedCache)
}
