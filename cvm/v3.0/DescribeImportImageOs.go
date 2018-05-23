// 2018-05-23 15:01:14.068115068 -0700 PDT m=+1.224780364
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeImportImageOsResp struct {
	Response struct {
		ImportImageOsArchitectureSupported []string `json:"ImportImageOsArchitectureSupported"`
		ImportImageOsListSupported         struct {
			Linux   []string `json:"linux"`
			Windows []string `json:"windows"`
		} `json:"ImportImageOsListSupported"`
		ImportImageOsVersionSupported struct {
			CentOS              []string `json:"CentOS"`
			CoreOS              []string `json:"CoreOS"`
			Debian              []string `json:"Debian"`
			FreeBSD             []string `json:"FreeBSD"`
			OpenSUSE            []string `json:"OpenSUSE"`
			Redhat              []string `json:"Redhat"`
			Suse                []string `json:"SUSE"`
			Ubuntu              []string `json:"Ubuntu"`
			Windows_Server_2003 []string `json:"Windows Server 2003"`
			Windows_Server_2008 []string `json:"Windows Server 2008"`
			Windows_Server_2012 []string `json:"Windows Server 2012"`
			Windows_Server_2016 []string `json:"Windows Server 2016"`
		} `json:"ImportImageOsVersionSupported"`
		RequestID string `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/15718
func (c *CvmClient) DescribeImportImageOs(options ...string) (*DescribeImportImageOsResp, error) {
	resp, err := c.DoAction("DescribeImportImageOs", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeImportImageOsResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DescribeImportImageOs(options ...string) (*DescribeImportImageOsResp, error) {
	return DefaultClient.DescribeImportImageOs(options...)
}

func (r *DescribeImportImageOsResp) String(args ...interface{}) (string, error) {
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