package vpc

import ()

type RespDescribeVpcEx struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Data          []RespDescribeVpcExData       `json:"data"`
	Message       string                        `json:"message"`
	TotalCount    int                           `json:"totalCount"`
}
type RespDescribeVpcExData struct {
	CidrBlock       string       `json:"cidrBlock"`
	ClassicLinkNum  int          `json:"classicLinkNum"`
	CreateTime      string       `json:"createTime"`
	IsDefault       bool         `json:"isDefault"`
	IsMulticast     bool         `json:"isMulticast"`
	NatNum          int          `json:"natNum"`
	RouteTableNum   int          `json:"routeTableNum"`
	SubnetNum       int          `json:"subnetNum"`
	UnVpcId         string       `json:"unVpcId"`
	VpcDeviceNum    int          `json:"vpcDeviceNum"`
	VpcId           string       `json:"vpcId"`
	VpcName         string       `json:"vpcName"`
	VpcPeerNum      int          `json:"vpcPeerNum"`
	VpgNum          int          `json:"vpgNum"`
	VpnGwNum        int          `json:"vpnGwNum"`
}

type RespCreateVpc struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
	RouteTableSet []RespCreateVpcRouteTableSet  `json:"routeTableSet"`
	SubnetSet     []interface{}                 `json:"subnetSet"`
	UniqVpcId     string                        `json:"uniqVpcId"`
	VpcCreateTime string                        `json:"vpcCreateTime"`
	VpcId         string                        `json:"vpcId"`
}

type RespCreateVpcRouteTableSet struct {
	RouteTableId    string  `json:"routeTableId"`
	RouteTableName  string  `json:"routeTableName"`
	RouteTableType  int     `json:"routeTableType"`
}

type RespDeleteVpc struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
}

type RespModifyVpcAttribute struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
}

type RespAttachClassicLinkVpc struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
	TaskId        int                           `json:"taskId"`
}

type RespDescribeVpcClassicLink struct {
	Code          int                               `json:"code"`
	CodeDesc      string                            `json:"codeDesc"`
	Data          []RespDescribeVpcClassicLinkData  `json:"data"`
	Message       string                            `json:"message"`
	TotalCount    int                               `json:"totalCount"`
}
type RespDescribeVpcClassicLinkData struct {
	ClassicLinkId  string       `json:"classicLinkId"`
	CreateTime     string       `json:"createTime"`
	InstanceId     string       `json:"instanceId"`
	InstanceName   string       `json:"instanceName"`
	LanIp          string       `json:"lanIp"`
	UnVpcId        string       `json:"unVpcId"`
	VpcId          string       `json:"vpcId"`
}

type RespDetachClassicLinkVpc struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
	TaskId        int                           `json:"taskId"`
}

type RespDescribeVpcLimit struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Data          RespDescribeVpcLimitData      `json:"data"`
	Message       string                        `json:"message"`
}
type RespDescribeVpcLimitData struct {
	Limit         interface{}                   `json:"limit"`
}

type RespCreateSubnet struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
	SubnetSet     []RespCreateSubnetSubnetSet   `json:"subnetSet"`
}
type RespCreateSubnetSubnetSet struct {
	CidrBlock         string                    `json:"cidrBlock"`
	RouteTableId      string                    `json:"routeTableId"`
	SubnetId          string                    `json:"subnetId"`
	SubnetName        string                    `json:"subnetName"`
	UnSubnetId        string                    `json:"unSubnetId"`
	Zone              string                    `json:"zone"`
	ZoneId            int                       `json:"zoneId"`
}

type RespDeleteSubnet struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
}

type RespModifySubnetAttribute struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Message       string                        `json:"message"`
}

type RespDescribeSubnetEx struct {
	Code          int                           `json:"code"`
	CodeDesc      string                        `json:"codeDesc"`
	Data          []RespDescribeSubnetExData    `json:"data"`
	Message       string                        `json:"message"`
}
type RespDescribeSubnetExData struct {
	AvailableIPNum              int                 `json:"availableIPNum"`
	Broadcast                   bool                `json:"broadcast"`
	CidrBlock                   string              `json:"cidrBlock"`
	IsDefault                   bool                `json:"isDefault"`
	NetworkAclId                interface{}         `json:"networkAclId"`
	RouteTableId                string              `json:"routeTableId"`
	RouteTableName              string              `json:"routeTableName"`
	RtbNum                      int                 `json:"rtbNum"`
	SubnetCreateTime            string              `json:"subnetCreateTime"`
	SubnetId                    string              `json:"subnetId"`
	SubnetName                  string              `json:"subnetName"`
	TotalIPNum                  int                 `json:"totalIPNum"`
	UnRouteTableId              string              `json:"unRouteTableId"`
	UnSubnetId                  string              `json:"unSubnetId"`
	UnVpcId                     string              `json:"unVpcId"`
	VpcCidrBlock                string              `json:"vpcCidrBlock"`
	VpcDevices                  int                 `json:"vpcDevices"`
	VpcId                       string              `json:"vpcId"`
	VpcName                     string              `json:"vpcName"`
	Zone                        string              `json:"zone"`
	ZoneId                      int                 `json:"zoneId"`
}

type RespDescribeSubnet struct {
	AvailableIPNum           int            `json:"availableIPNum"`
	CidrBlock                string         `json:"cidrBlock"`
	Code                     int            `json:"code"`
	CodeDesc                 string         `json:"codeDesc"`
	Message                  string         `json:"message"`
	RouteTableId             string         `json:"routeTableId"`
	SubnetCreateTime         string         `json:"subnetCreateTime"`
	SubnetId                 string         `json:"subnetId"`
	SubnetName               string         `json:"subnetName"`
	TotalIPNum               int            `json:"totalIPNum"`
	ZoneId                   int            `json:"zoneId"`
}

type RespCreateRouteTable struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
	RouteTableId             string                               `json:"routeTableId"`
	RouteTableSet            []RespCreateRouteTableRouteTableSet  `json:"routeTableSet"`
	UnRouteTableId           string                               `json:"unRouteTableId"`
}
type RespCreateRouteTableRouteTableSet struct {
	Description           interface{}  `json:"description"`
	DestinationCidrBlock  string       `json:"destinationCidrBlock"`
	NextHub               string       `json:"nextHub"`
	NextType              int          `json:"nextType"`
}

type RespDeleteRouteTable struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespModifyRouteTableAttribute struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespDescribeRouteTable struct {
	Code          int                              `json:"code"`
	CodeDesc      string                           `json:"codeDesc"`
	Data          []RespDescribeRouteTableData     `json:"data"`
	Message       string                           `json:"message"`
	TotalCount    int                              `json:"totalCount"`
}
type RespDescribeRouteTableData struct {
	RouteSet                   []RespDescribeRouteTableRouteSet  `json:"routeSet"`
	RouteTableCreateTime       string                            `json:"routeTableCreateTime"`
	RouteTableId               string                            `json:"routeTableId"`
	RouteTableName             string                            `json:"routeTableName"`
	RouteTableType             int                               `json:"routeTableType"`
	SubnetNum                  int                               `json:"subnetNum"`
	UnRouteTableId             string                            `json:"unRouteTableId"`
	UnVpcId                    string                            `json:"unVpcId"`
	VpcCidrBlock               string                            `json:"vpcCidrBlock"`
	VpcId                      string                            `json:"vpcId"`
	VpcName                    string                            `json:"vpcName"`
}
type RespDescribeRouteTableRouteSet struct {
	Description           interface{}              `json:"description"`
	DestinationCidrBlock  string                   `json:"destinationCidrBlock"`
	nextHub               string                   `json:"nextHub"`
	nextType              int                      `json:"nextType"`
	unNextHub             interface{}              `json:"unNextHub"`
}

type RespAssociateRouteTable struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespCreateNetworkAcl struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Data                     interface{}                          `json:"data"`
	Message                  string                               `json:"message"`
}

type RespDeleteNetworkAcl struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespModifyNetworkAcl struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespDescribeNetworkAcl struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Data                     []RespDescribeNetworkAclData         `json:"data"`
	Message                  string                               `json:"message"`
	TotalCount               int                                  `json:"totalCount"`
}
type RespDescribeNetworkAclData struct {
	CreateTime          string                        `json:"createTime"`
	NetworkAclEntrySet  interface{}                   `json:"networkAclEntrySet"`
	NetworkAclId        string                        `json:"networkAclId"`
	NetworkAclName      string                        `json:"networkAclName"`
	SubnetNum           int                           `json:"subnetNum"`
	SubnetSet           interface{}                   `json:"subnetSet"`
	UnVpcId             string                        `json:"unVpcId"`
	VpcCidrBlock        string                        `json:"vpcCidrBlock"`
	VpcId               string                        `json:"vpcId"`
	VpcName             string                        `json:"vpcName"`
}

type RespModifyNetworkAclEntry struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespCreateSubnetAclRule struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespDeteleSubnetAclRule struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespDescribeVpcPeeringConnections struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
	Data                     []interface{}                        `json:"data"`
	TotalCount               int                                  `json:"totalCount"`
}

type RespCreateVpcPeeringConnection struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
	PeeringConnectionId      string                               `json:"peeringConnectionId"`
}

type RespDeleteVpcPeeringConnection struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespModifyVpcPeeringConnection struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespAcceptVpcPeeringConnection struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespRejectVpcPeeringConnection struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespCreateVpcPeeringConnectionEx struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
	PeeringConnectionId      string                               `json:"peeringConnectionId"`
}

type RespDeleteVpcPeeringConnectionEx struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespModifyVpcPeeringConnectionEx struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespAcceptVpcPeeringConnectionEx struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespRejectVpcPeeringConnectionEx struct {
	Code                     int                                  `json:"code"`
	CodeDesc                 string                               `json:"codeDesc"`
	Message                  string                               `json:"message"`
}

type RespCreateNatGateway struct {
	BillID   string `json:"billId"`
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

type RespDeleteNatGateway struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	TaskID   int    `json:"taskId"`
}

type RespDescribeNatGateway struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     []struct {
		AppID            string   `json:"appId"`
		Bandwidth        int      `json:"bandwidth"`
		CreateTime       string   `json:"createTime"`
		EipCount         int      `json:"eipCount"`
		EipSet           []string `json:"eipSet"`
		MaxConcurrent    int      `json:"maxConcurrent"`
		NatID            string   `json:"natId"`
		NatName          string   `json:"natName"`
		ProductionStatus int      `json:"productionStatus"`
		State            int      `json:"state"`
		UnVpcID          string   `json:"unVpcId"`
		VpcID            int      `json:"vpcId"`
		VpcName          string   `json:"vpcName"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int    `json:"totalCount"`
}

type RespEipBindNatGateway struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	TaskID   int    `json:"taskId"`
}

type RespEipUnBindNatGateway struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
	TaskID   int    `json:"taskId"`
}

type RespModifyNatGateway struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

type RespQueryNatGatewayProductionStatus struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		ErrorCode int `json:"errorCode"`
		Status    int `json:"status"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespUpgradeNatGateway struct {
	BillID   string `json:"billId"`
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

type RespCreateNetworkInterface struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespDescribeNetworkInterfaces struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Data []struct {
			CreateTime     string        `json:"createTime"`
			EniDescription string        `json:"eniDescription"`
			EniName        string        `json:"eniName"`
			GroupSet       []interface{} `json:"groupSet"`
			InstanceSet    struct {
				AttachTime string `json:"attachTime"`
				InstanceID string `json:"instanceId"`
			} `json:"instanceSet"`
			MacAddress            string `json:"macAddress"`
			NetworkInterfaceID    string `json:"networkInterfaceId"`
			Primary               bool   `json:"primary"`
			PrivateIPAddressesSet []struct {
				Description      string `json:"description"`
				EipID            string `json:"eipId"`
				Primary          bool   `json:"primary"`
				PrivateIPAddress string `json:"privateIpAddress"`
				WanIP            string `json:"wanIp"`
			} `json:"privateIpAddressesSet"`
			SubnetID string `json:"subnetId"`
			VpcID    string `json:"vpcId"`
			VpcName  string `json:"vpcName"`
			ZoneID   int    `json:"zoneId"`
		} `json:"data"`
		TotalNum int `json:"totalNum"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespDeleteNetworkInterface struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespModifyNetworkInterface struct {
	Code     int           `json:"code"`
	CodeDesc string        `json:"codeDesc"`
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
}

type RespAssignPrivateIpAddresses struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespUnassignPrivateIpAddresses struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespAttachNetworkInterface struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespDetachNetworkInterface struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespMigrateNetworkInterface struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

type RespMigratePrivateIpAddress struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}