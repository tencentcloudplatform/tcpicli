package ccs

import (
	"encoding/json"
)

type DescribeClusterResp struct {
	Code     int    `json:"code"`
	Codedesc string `json:"codedesc"`
	Data     struct {
		Clusters []struct {
			ClusterCIDR             string `json:"clusterCIDR"`
			ClusterExternalEndpoint string `json:"clusterExternalEndpoint"`
			ClusterID               string `json:"clusterId"`
			ClusterName             string `json:"clusterName"`
			CreatedAt               string `json:"createdAt"`
			Description             string `json:"description"`
			K8SVersion              string `json:"k8sVersion"`
			NodeNum                 int    `json:"nodeNum"`
			NodeStatus              string `json:"nodeStatus"`
			OpenHTTPS               int    `json:"openHttps"`
			Os                      string `json:"os"`
			ProjectID               int    `json:"projectId"`
			Region                  string `json:"region"`
			RegionID                int    `json:"regionId"`
			Status                  string `json:"status"`
			TotalCPU                int    `json:"totalCpu"`
			TotalMem                int    `json:"totalMem"`
			UnVpcID                 string `json:"unVpcId"`
			UpdatedAt               string `json:"updatedAt"`
			VpcID                   int    `json:"vpcId"`
		} `json:"clusters"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
	Message string `json:"message"`
}

func DescribeCluster(options ...string) (*DescribeClusterResp, error) {
	resp, err := DoAction("DescribeCluster", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeClusterResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
