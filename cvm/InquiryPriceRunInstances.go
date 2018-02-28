// 2017-12-18 14:43:17.107583 -0800 PST m=+29.322346854
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type InquiryPriceRunInstancesResp struct {
	Response struct {
		Price struct {
			BandwidthPrice struct {
				ChargeUnit string `json:"ChargeUnit"`
				UnitPrice  int64  `json:"UnitPrice"`
			} `json:"BandwidthPrice"`
			InstancePrice struct {
				ChargeUnit string  `json:"ChargeUnit"`
				UnitPrice  float64 `json:"UnitPrice"`
			} `json:"InstancePrice"`
		} `json:"Price,omitempty"`
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9385
func (c *CvmClient) InquiryPriceRunInstances(options ...string) (*InquiryPriceRunInstancesResp, error) {
	resp, err := c.DoAction("InquiryPriceRunInstances", options...)
	if err != nil {
		return nil, err
	}
	var s InquiryPriceRunInstancesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func InquiryPriceRunInstances(options ...string) (*InquiryPriceRunInstancesResp, error) {
	return DefaultClient.InquiryPriceRunInstances(options...)
}

func (r *InquiryPriceRunInstancesResp) String(args ...interface{}) (string, error) {
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
