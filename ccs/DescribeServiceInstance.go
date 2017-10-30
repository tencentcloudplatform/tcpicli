package ccs

import (
	"encoding/json"
)

type DescribeServiceInstanceResp struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Instances []struct {
			Containers []struct {
				ContainerID string `json:"containerId"`
				Image       string `json:"image"`
				Name        string `json:"name"`
				Reason      string `json:"reason"`
				Status      string `json:"status"`
			} `json:"containers"`
			CreatedAt    string `json:"createdAt"`
			IP           string `json:"ip"`
			Name         string `json:"name"`
			NodeIP       string `json:"nodeIp"`
			NodeName     string `json:"nodeName"`
			ReadyCount   int    `json:"readyCount"`
			Reason       string `json:"reason"`
			RestartCount int    `json:"restartCount"`
			SrcReason    string `json:"srcReason"`
			Status       string `json:"status"`
		} `json:"instances"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9433
func DescribeServiceInstance(options ...string) (*DescribeServiceInstanceResp, error) {
	resp, err := DoAction("DescribeServiceInstance", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeServiceInstanceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
