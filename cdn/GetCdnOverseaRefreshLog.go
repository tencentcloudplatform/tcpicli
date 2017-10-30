package cdn

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
	"io/ioutil"
	"time"
)

func GetCdnOverseaRefreshLog(options ...string) ([]byte, error) {
	client := core.NewClient(requesturl, true)
	method := "POST"

	params := make(map[string]interface{})
	params["Action"] = "GetCdnOverseaRefreshLog"
	core.AssignParams(params, options...)
	if _, ok := params["date"]; !ok {
		params["date"] = time.Now().Format("20060102")
	}

	resp, err := client.SendRequest(method, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var j interface{}
	json.Unmarshal(b, &j)
	b, _ = json.MarshalIndent(j, "", "  ")
	return b, nil
}
