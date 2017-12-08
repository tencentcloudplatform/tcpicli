// 2017-12-01 14:00:59.866087 -0800 PST m=+4.402150122
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vod

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeRecordPlayInfoResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	FileSet  []struct {
		Duration string `json:"duration"`
		FileID   string `json:"fileId"`
		FileName string `json:"fileName"`
		ImageURL string `json:"image_url"`
		PlaySet  []struct {
			Definition string `json:"definition"`
			URL        string `json:"url"`
			Vbitrate   int64  `json:"vbitrate"`
			Vheight    int64  `json:"vheight"`
			Vwidth     int64  `json:"vwidth"`
		} `json:"playSet"`
		Status string `json:"status"`
		Vid    string `json:"vid"`
	} `json:"fileSet"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/266/8227
func DescribeRecordPlayInfo(options ...string) (*DescribeRecordPlayInfoResp, error) {
	resp, err := DoAction("DescribeRecordPlayInfo", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeRecordPlayInfoResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeRecordPlayInfoResp) String(args ...interface{}) (string, error) {
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