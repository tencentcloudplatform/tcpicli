// 2017-11-02 23:46:00.113288943 +0000 UTC m=+44.646558033
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cdn

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type GetHostInfoByIdResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Hosts []struct {
			AppID     int64         `json:"app_id"`
			Cache     []interface{} `json:"cache"`
			CacheMode string        `json:"cache_mode"`
			Capping   struct {
				Active    string `json:"active"`
				Bandwidth int64  `json:"bandwidth"`
				Hy        string `json:"hy"`
				Unit      string `json:"unit"`
			} `json:"capping"`
			Cname       string `json:"cname"`
			CosKey      string `json:"cos_key"`
			CreateTime  string `json:"create_time"`
			Deleted     string `json:"deleted"`
			Disabled    int64  `json:"disabled"`
			FinalPids   int64  `json:"final_pids"`
			FurlCache   string `json:"furl_cache"`
			FwdHost     string `json:"fwd_host"`
			FwdHostType string `json:"fwd_host_type"`
			Host        string `json:"host"`
			HostConfig  struct {
				AdvancedCache []struct {
					Rule string `json:"rule"`
					Time int64  `json:"time"`
					Type int64  `json:"type"`
					Unit string `json:"unit"`
				} `json:"advanced_cache"`
				AdvancedMaxage []struct {
					Rule string `json:"rule"`
					Time int64  `json:"time"`
					Type int64  `json:"type"`
					Unit string `json:"unit"`
				} `json:"advanced_maxage"`
				BackupOrigin struct {
					BackupOrigin string `json:"backup_origin"`
				} `json:"backup_origin"`
				Compression struct {
					Switch string   `json:"switch"`
					Type   []string `json:"type"`
				} `json:"compression"`
				IPCc struct {
					List []string `json:"list"`
					Type int64    `json:"type"`
				} `json:"ip_cc"`
				OverseaFeatureSwitch struct {
					OverseaFeatureSwitch string `json:"oversea_feature_switch"`
				} `json:"oversea_feature_switch"`
				SrcMethod struct {
					Backup          string `json:"backup"`
					BackupSrcMethod string `json:"backup_src_method"`
					Master          string `json:"master"`
					SrcMethod       string `json:"src_method"`
				} `json:"src_method"`
			} `json:"host_config"`
			HostID      int64  `json:"host_id"`
			HostType    string `json:"host_type"`
			HTTP2       int64  `json:"http2"`
			HTTPSConfig struct {
				AccessExtension            string `json:"access_extension"`
				ExtraExtension             string `json:"extra_extension"`
				HTTP2                      int64  `json:"http2"`
				HyExtension                string `json:"hy_extension"`
				ProxySslName               string `json:"proxy_ssl_name"`
				ProxySslServerName         string `json:"proxy_ssl_server_name"`
				ProxySslTrustedCertificate string `json:"proxy_ssl_trusted_certificate"`
				ProxySslVerify             string `json:"proxy_ssl_verify"`
				Spdy                       string `json:"spdy"`
				SslStapling                string `json:"ssl_stapling"`
				SslStaplingVerify          string `json:"ssl_stapling_verify"`
				SslVerifyClient            string `json:"ssl_verify_client"`
			} `json:"https_config"`
			ID             int64  `json:"id"`
			Message        string `json:"message"`
			MiddleResource int64  `json:"middle_resource"`
			Origin         string `json:"origin"`
			OwnerUin       int64  `json:"owner_uin"`
			ProjectID      int64  `json:"project_id"`
			Readonly       int64  `json:"readonly"`
			Refer          struct {
				List     []string `json:"list"`
				NullFlag int64    `json:"null_flag"`
				Type     int64    `json:"type"`
			} `json:"refer"`
			Seo           string      `json:"seo"`
			ServiceType   string      `json:"service_type"`
			SslCertID     string      `json:"ssl_cert_id"`
			SslCertName   string      `json:"ssl_cert_name"`
			SslDeployTime interface{} `json:"ssl_deploy_time"`
			SslExpireTime interface{} `json:"ssl_expire_time"`
			SslType       int64       `json:"ssl_type"`
			Status        int64       `json:"status"`
			TegStatus     string      `json:"teg_status"`
			TestURL       string      `json:"test_url"`
			UpdateTime    string      `json:"update_time"`
		} `json:"hosts"`
		Total int64 `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/228/3938
func GetHostInfoById(ids ...string) (*GetHostInfoByIdResp, error) {
	var options []string
	for k, v := range ids {
		options = append(options, fmt.Sprintf("ids.%v=%v", k, v))
	}
	resp, err := DoAction("GetHostInfoById", options...)
	if err != nil {
		return nil, err
	}
	var s GetHostInfoByIdResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *GetHostInfoByIdResp) String(args ...interface{}) (string, error) {
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
