package img

import (
	"encoding/json"
)

type CreateImageResp struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		RetCode int `json:"retCode"`
		TaskID  int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/9415
func CreateImage(options ...string) (*CreateImageResp, error) {
	resp, err := DoAction("CreateImage", options...)
	if err != nil {
		return nil, err
	}
	var s CreateImageResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
