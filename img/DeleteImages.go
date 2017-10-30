package img

import (
	"encoding/json"
)

type DeleteImagesResp struct {
	Code     int         `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Detail   interface{} `json:"detail"`
	Message  string      `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/9415
func DeleteImages(options ...string) (*DeleteImagesResp, error) {
	resp, err := DoAction("DeleteImages", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteImagesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
