package vpc

import (
	"encoding/json"
)

type CreateVpcResp struct {
	Code          int    `json:"code"`
	CodeDesc      string `json:"codeDesc"`
	Message       string `json:"message"`
	RouteTableSet []struct {
		RouteTableID   string `json:"routeTableId"`
		RouteTableName string `json:"routeTableName"`
		RouteTableType int    `json:"routeTableType"`
	} `json:"routeTableSet"`
	SubnetSet     []interface{} `json:"subnetSet"`
	UniqVpcID     string        `json:"uniqVpcId"`
	VpcCreateTime string        `json:"vpcCreateTime"`
	VpcID         string        `json:"vpcId"`
}

func CreateVpc(options ...string) (*CreateVpcResp, error) {
	resp, err := DoAction("CreateVpc", options...)
	if err != nil {
		return nil, err
	}
	var s CreateVpcResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}
