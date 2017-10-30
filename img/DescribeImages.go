package img

import (
	"encoding/json"
)

type DescribeImagesResp struct {
	Response struct {
		Error    interface{} `json:"Error,omitempty"`
		ImageSet []struct {
			CreatedTime      interface{} `json:"CreatedTime"`
			ImageCreator     interface{} `json:"ImageCreator"`
			ImageDescription string      `json:"ImageDescription"`
			ImageID          string      `json:"ImageId"`
			ImageName        string      `json:"ImageName"`
			ImageSize        int         `json:"ImageSize"`
			ImageSource      string      `json:"ImageSource"`
			ImageState       string      `json:"ImageState"`
			ImageType        string      `json:"ImageType"`
			OsName           string      `json:"OsName"`
		} `json:"ImageSet,omitempty"`
		RequestID  string `json:"RequestId"`
		TotalCount int    `json:"TotalCount"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9418
func DescribeImages(options ...string) (*DescribeImagesResp, error) {
	resp, err := DoAction("DescribeImages", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeImagesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
