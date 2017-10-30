package ccs

import (
	"encoding/json"
)

type RespDescribeClusterService struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Services []struct {
			AccessType      string `json:"accessType"`
			CreatedAt       string `json:"createdAt"`
			CurrentReplicas int    `json:"currentReplicas"`
			DesiredReplicas int    `json:"desiredReplicas"`
			ExternalIP      string `json:"externalIp"`
			Labels          struct {
				QcloudApp string `json:"qcloud-app"`
			} `json:"labels"`
			LbID      string `json:"lbId"`
			LbStatus  string `json:"lbStatus"`
			Namespace string `json:"namespace"`
			ReasonMap struct {
				NAMING_FAILED int `json:"容器运行中"`
			} `json:"reasonMap"`
			ServiceIP    string `json:"serviceIp"`
			ServiceName  string `json:"serviceName"`
			SrcReasonMap struct {
				ContainerRunning int `json:"container running"`
			} `json:"srcReasonMap"`
			Status    string `json:"status"`
			SysLabels struct {
				QcloudApp string `json:"qcloud-app"`
			} `json:"sysLabels"`
			UserLabels struct {
			} `json:"userLabels"`
		} `json:"services"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
	Message string `json:"message"`
}

func DescribeClusterService(region ...string) (*RespDescribeClusterService, error) {
	b, err := DoAction("DescribeClusterService", region...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeClusterService)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
