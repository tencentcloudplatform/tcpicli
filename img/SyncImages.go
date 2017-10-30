package img

import (
	"encoding/json"
)

type SyncImagesResp struct {
	Response struct {
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9417
func SyncImages(options ...string) (*SyncImagesResp, error) {
	resp, err := DoAction("SyncImages", options...)
	if err != nil {
		return nil, err
	}
	var s SyncImagesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
