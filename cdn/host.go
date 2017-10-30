package cdn

import ()

type Host struct {
	Id             int64         `json:"id"`
	AppId          int64         `json:"app_id"`
	FinalPids      int           `json:"final_pids"`
	HostId         int           `json:"host_id"`
	ProjectId      int64         `json:"project_id"`
	Cache          []interface{} `json:"cache"`
	CacheMode      string        `json:"cache_mode"`
	Cname          string        `json:"cname"`
	CosKey         interface{}   `json:"cos_key"`
	CreateTime     string        `json:"create_time"`
	Deleted        string        `json:"deleted"`
	Disabled       int           `json:"disabled"`
	FurlCache      string        `json:"furl_cache"`
	FwdHost        string        `json:"fwd_host"`
	FwdHostType    string        `json:"fwd_host_type"`
	Host           string        `json:"host"`
	HostConfig     *HostConfig   `json:"host_config"`
	HostType       string        `json:"host_type"`
	Http2          int           `json:"http2"`
	HttpsConfig    interface{}   `json:"https_config"`
	Message        string        `json:"message"`
	MiddleResource int           `json:"middle_resource"`
	Origin         string        `json:"origin"`
	OwnerUin       int64         `json:"owner_uin"`
	PidConfig      interface{}   `json:"pid_config"`
	Readonly       int           `json:"readonly"`
	Refer          interface{}   `json:"refer"`
	Seo            string        `json:"seo"`
	ServiceType    string        `json:"service_type"`
	SslCertName    string        `json:"ssl_cert_name"`
	SslDeployTime  interface{}   `json:"ssl_deploy_time"`
	SslExpireTime  interface{}   `json:"ssl_expire_time"`
	SslType        int           `json:"ssl_type"`
	Status         int           `json:"status"`
	TegStatus      string        `json:"teg_status"`
	TestUrl        string        `json:"test_url"`
	UpdateTime     string        `json:"update_time"`
}

type HostConfig struct {
	AdvancedCache        []interface{} `json:"advanced_cache"`
	AdvancedMaxage       []interface{} `json:"advanced_maxage"`
	BackupOrigin         interface{}   `json:"backup_origin"`
	ChunkFdSwitch        interface{}   `json:"chunkFd_switch"`
	ClusterSwitch        interface{}   `json:"cluster_switch"`
	Compression          interface{}   `json:"compression"`
	DedicateLine         int           `json:"dedicate_line" console:"International Direct Connect"`
	DetailReqHeader      []interface{} `json:"detail_req_header"`
	DetailRspHeader      []interface{} `json:"detail_rsp_header"`
	Follow302Switch      interface{}   `json:"follow302_switch"`
	IpCc                 interface{}   `json:"ip_cc"`
	IpFreqLimit          interface{}   `json:"ip_freq_limit"`
	OverseaFeatureSwitch interface{}   `json:"oversea_feature_switch"`
	SrcMethod            interface{}   `json:"src_method"`
	StatusCache          []interface{} `json:"status_cache"`
	VideoSwitch          interface{}   `json:"video_switch"`
}

type RespGetHostInfo struct {
	Code     int              `json:"code"`
	CodeDesc string           `json:"codeDesc"`
	Data     *DataGetHostInfo `json:"data"`
	Message  string           `json:"message"`
}
type DataGetHostInfo struct {
	Hosts []Host `json:"hosts"`
	Total int    `json:"total"`
}
