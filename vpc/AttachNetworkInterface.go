package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type AttachNetworkInterfaceResp struct {
	Code     int64       `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	TaskID   int64       `json:"taskId"`
}

// Implement https://cloud.tencent.com/document/api/215/8836
func AttachNetworkInterface(options ...string) (*AttachNetworkInterfaceResp, error) {
	resp, err := DoAction("AttachNetworkInterface", options...)
	if err != nil {
		return nil, err
	}
	var s AttachNetworkInterfaceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *AttachNetworkInterfaceResp) String(args ...interface{}) (string, error) {
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
