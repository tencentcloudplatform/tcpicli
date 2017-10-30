package ccs

import (
	"encoding/json"
)

type RespDescribeClusterServiceInfo struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Service struct {
			AccessType string `json:"accessType"`
			Containers []struct {
				Arguments     interface{} `json:"arguments"`
				Command       string      `json:"command"`
				ContainerName string      `json:"containerName"`
				CPU           int         `json:"cpu"`
				Envs          []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"envs"`
				Image        string      `json:"image"`
				LiveProbe    interface{} `json:"liveProbe"`
				Memory       int         `json:"memory"`
				Privileged   bool        `json:"privileged"`
				ReadyProbe   interface{} `json:"readyProbe"`
				UnHubID      string      `json:"unHubId"`
				VolumeMounts interface{} `json:"volumeMounts"`
				WorkingDir   string      `json:"workingDir"`
			} `json:"containers"`
			CreatedAt       string `json:"createdAt"`
			CurrentReplicas int    `json:"currentReplicas"`
			DesiredReplicas int    `json:"desiredReplicas"`
			ExternalIP      string `json:"externalIp"`
			Labels          struct {
				QcloudApp string `json:"qcloud-app"`
			} `json:"labels"`
			LbID         string `json:"lbId"`
			LbStatus     string `json:"lbStatus"`
			Namespace    string `json:"namespace"`
			PortMappings []struct {
				ContainerPort int    `json:"containerPort"`
				LbPort        int    `json:"lbPort"`
				NodePort      int    `json:"nodePort"`
				Protocol      string `json:"protocol"`
			} `json:"portMappings"`
			ReasonMap struct {
				NAMING_FAILED int `json:"容器运行中"`
			} `json:"reasonMap"`
			RegionID int `json:"regionId"`
			Selector struct {
				QcloudApp string `json:"qcloud-app"`
			} `json:"selector"`
			ServiceDesc  string `json:"serviceDesc"`
			ServiceIP    string `json:"serviceIp"`
			ServiceName  string `json:"serviceName"`
			SrcReasonMap struct {
				ContainerRunning int `json:"container running"`
			} `json:"srcReasonMap"`
			Status    string `json:"status"`
			Strategy  string `json:"strategy"`
			SubnetID  string `json:"subnetId"`
			SysLabels struct {
				QcloudApp string `json:"qcloud-app"`
			} `json:"sysLabels"`
			UserLabels struct {
			} `json:"userLabels"`
			Volumes []interface{} `json:"volumes"`
		} `json:"service"`
	} `json:"data"`
	Message string `json:"message"`
}

func DescribeClusterServiceInfo(region ...string) (*RespDescribeClusterServiceInfo, error) {
	b, err := DoAction("DescribeClusterServiceInfo", region...)
	if err != nil {
		return nil, err
	}
	m := new(RespDescribeClusterServiceInfo)
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
