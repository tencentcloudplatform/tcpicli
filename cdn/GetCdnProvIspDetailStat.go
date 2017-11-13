// 2017-11-02 23:46:44.741749842 +0000 UTC m=+89.275018852
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cdn

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type GetCdnProvIspDetailStatResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Date     string        `json:"date"`
		ProvData []interface{} `json:"prov_data"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/228/7356
func GetCdnProvIspDetailStat(options ...string) (*GetCdnProvIspDetailStatResp, error) {
	resp, err := DoAction("GetCdnProvIspDetailStat", options...)
	if err != nil {
		return nil, err
	}
	var s GetCdnProvIspDetailStatResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *GetCdnProvIspDetailStatResp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
