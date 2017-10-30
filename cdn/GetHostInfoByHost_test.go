/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : GetHostInfoByHost_test.go

* Purpose :

* Creation Date : 09-06-2017

* Last Modified : Wed Sep  6 15:46:26 2017

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package cdn

import (
	"log"
)

func (t *CdnTests) TestGetHostInfoByHost(domain string) {
	b, err := GetHostInfoByHost(domain)
	if err != nil {
		t.Fail()
	}
	log.Println(b)
}
