// 2017-12-12 20:49:31.906852871 +0000 UTC m=+89.902902750
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cdn

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type GetCdnStatTopResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		EndDatetime string `json:"end_datetime"`
		IspData     []struct {
			ID    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"isp_data"`
		Period       int64 `json:"period"`
		ProvinceData []struct {
			ID    int64  `json:"id"`
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"province_data"`
		StartDatetime string `json:"start_datetime"`
		StatType      string `json:"stat_type"`
		URLData       []struct {
			Name  string `json:"name"`
			Value int64  `json:"value"`
		} `json:"url_data"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/228/3944
func (c *CdnClient) GetCdnStatTop(options ...string) (*GetCdnStatTopResp, error) {
	resp, err := c.DoAction("GetCdnStatTop", options...)
	if err != nil {
		return nil, err
	}
	var s GetCdnStatTopResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func GetCdnStatTop(options ...string) (*GetCdnStatTopResp, error) {
	return DefaultClient.GetCdnStatTop(options...)
}

func (r *GetCdnStatTopResp) String(args ...interface{}) (string, error) {
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
