package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/vpc"
	"github.com/urfave/cli"
)

var (
	funcVpc []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action and output raw json response.",
			Action: VpcDoAction,
		},
		{
			Name:        "CreateVpc",
			Usage:       "Creates a new VPC.",
			Action:      VpcCreateVpc,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1309",
		},
		{
			Name:        "DescribeVpcEx",
			Usage:       "Queries VPC list.",
			Action:      VpcDescribeVpcEx,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1372",
		},
		{
			Name:        "DescribeVpcClassicLink",
			Usage:       "Describes classic link between VPC and CVM",
			Action:      VpcDescribeVpcClassicLink,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1844",
		},
		{
			Name:        "DescribeVpcLimit",
			Usage:       "Describes VPC limits.",
			Action:      VpcDescribeVpcLimit,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/2112",
		},
		{
			Name:        "ModifyVpcAttribute",
			Usage:       "Modifies VPC attribute",
			Action:      VpcModifyVpcAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1310",
		},
		// {
		// 	Name:        "AssociateVip",
		// 	Usage:       "Binds EIP to instance",
		// 	Action:      VpcAssociateVip,
		// 	Description: "Referrer: https://cloud.tencent.com/document/api/215/1361",
		// },
		//{
		//	Name:        "AttachClassicLinkVpc",
		//	Usage:       "Binds a CVM to a VPC.",
		//	Action:      VpcAttachClassicLinkVpc,
		//	Description: "Referrer: https://cloud.tencent.com/document/api/215/2098",
		//},
		{
			Name:        "DeleteVpc",
			Usage:       "Deletes a VPC. Can't have any resources in VPC.",
			Action:      VpcDeleteVpc,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1307",
		},
		{
			Name:        "CreateSubnet",
			Usage:       "Creates subnets in existing VPCs",
			Action:      VpcCreateSubnet,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1314",
		},
		{
			Name:        "DescribeSubnetEx",
			Usage:       "Queries subnet list",
			Action:      VpcDescribeSubnetEx,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1371",
		},
		{
			Name:        "ModifySubnetAttribute",
			Usage:       "Queries subnet list",
			Action:      VpcModifySubnetAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1313",
		},
		{
			Name:        "DescribeSubnet",
			Usage:       "Gives detailed infromation about particular subnet",
			Action:      VpcDescribeSubnet,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1311",
		},
		{
			Name:        "DeleteSubnet",
			Usage:       "Deletes subnet.",
			Action:      VpcDeleteSubnet,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1312",
		},
		{
			Name:        "CreateRouteTable",
			Usage:       "Creates a new route table",
			Action:      VpcCreateRouteTable,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1419",
		},
		{
			Name:        "AssociateRouteTable",
			Usage:       "Associates a route table",
			Action:      VpcAssociateRouteTable,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1416",
		},
		{
			Name:        "InquiryVpnPrice",
			Usage:       "Queries the cost of a VPN using period of time and bandwidth",
			Action:      VpcInquiryVpnPrice,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5104",
		},
		{
			Name:        "DescribeVpnGw",
			Usage:       "Returns information on existing VPN GWs",
			Action:      VpcDescribeVpnGw,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5108",
		},
		{
			Name:        "ModifyVpnGw",
			Usage:       "Changes attributes of existing VPN GWs.",
			Action:      VpcModifyVpnGw,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5107",
		},
		{
			Name:        "DescribeUserGwVendor",
			Usage:       "Acquires information on supported peer GW vendors.",
			Action:      VpcDescribeUserGwVendor,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5120",
		},
		{
			Name:        "DescribeUserGw",
			Usage:       "Lists user GWs",
			Action:      VpcDescribeUserGw,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5119",
		},
		{
			Name:        "ModifyUserGw",
			Usage:       "Acquires information on supported peer GW vendors.",
			Action:      VpcModifyUserGw,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5118",
		},
		{
			Name:        "CreateDirectConnectGateway",
			Usage:       "Creates direct connect GW",
			Action:      VpcCreateDirectConnectGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4824",
		},
		{
			Name:        "DescribeDirectConnectGateway",
			Usage:       "Queries list of direct connect GW",
			Action:      VpcDescribeDirectConnectGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4827",
		},
		{
			Name:        "ModifyDirectConnectGateway",
			Usage:       "Changes attribute of direct connect GW",
			Action:      VpcModifyDirectConnectGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4826",
		},
		{
			Name:        "CreateLocalIPTranslationNatRule",
			Usage:       "Creates a NAT rule on a direct connect GW",
			Action:      VpcCreateLocalIPTranslationNatRule,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5185",
		},
		{
			Name:        "DescribeLocalIPTranslationNatRule",
			Usage:       "Describes a NAT rule on a direct connect GW",
			Action:      VpcDescribeLocalIPTranslationNatRule,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5188",
		},
		{
			Name:        "InquiryNatPrice",
			Usage:       "Returns the cost of a NAT GW given n-concurrent connections",
			Action:      VpcInquiryNatPrice,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4091",
		},
		{
			Name:        "CreateNatGateway",
			Usage:       "Creates a NAT gateway",
			Action:      VpcCreateNatGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4094",
		},
		{
			Name:        "DescribeNatGateway",
			Usage:       "Describes a NAT gateway",
			Action:      VpcDescribeNatGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4088",
		},
		{
			Name:        "ModifyNatGateway",
			Usage:       "Modifies a NAT gateway",
			Action:      VpcModifyNatGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4086",
		},
		{
			Name:        "EipBindNatGateway",
			Usage:       "Binds EIP to a  NAT gateway",
			Action:      VpcEipBindNatGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4093",
		},
		// {
		// 	Name:        "EipUnBindNatGateway",
		// 	Usage:       "Unbinds EIP from a NAT gateway",
		// 	Action:      VpcEipUnBindNatGateway,
		// 	Description: "Referrer: https://cloud.tencent.com/document/api/215/4092",
		// },
		{
			Name:        "DeleteNatGateway",
			Usage:       "Deletes a NAT gateway",
			Action:      VpcDeleteNatGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4087",
		},
		{
			Name:        "ModifyLocalIPTranslationNatRule",
			Usage:       "Modifies a NAT rule on a direct connect GW",
			Action:      VpcModifyLocalIPTranslationNatRule,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5187",
		},
		{
			Name:        "DeleteLocalIPTranslationNatRule",
			Usage:       "Deletes a NAT rule on a direct connect GW",
			Action:      VpcDeleteLocalIPTranslationNatRule,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5186",
		},
		{
			Name:        "DeleteDirectConnectGateway",
			Usage:       "Deletes an existing direct connect GW",
			Action:      VpcDeleteDirectConnectGateway,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/4825",
		},
		{
			Name:        "DeleteUserGw",
			Usage:       "Deletes peered GW.",
			Action:      VpcDeleteUserGw,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5117",
		},
		{
			Name:        "CreateNetworkAcl",
			Usage:       "Creates a new NACL.",
			Action:      VpcCreateNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1437",
		},
		{
			Name:        "CreateSubnetAclRule",
			Usage:       "Binds NACL to subnet.",
			Action:      VpcCreateNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1438",
		},
		{
			Name:        "DeteleSubnetAclRule",
			Usage:       "Unbinds NACL from subnet.",
			Action:      VpcDeleteNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1442",
		},
		{
			Name:        "DescribeNetworkAcl",
			Usage:       "Queries NACL list",
			Action:      VpcDescribeNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1441",
		},
		{
			Name:        "ModifyNetworkAcl",
			Usage:       "Changes the name of NACL",
			Action:      VpcModifyNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1443",
		},
		{
			Name:        "ModifyNetworkAclEntry",
			Usage:       "Add rules to NACL",
			Action:      VpcModifyNetworkAclEntry,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1444",
		},
		{
			Name:        "DeleteNetworkAcl",
			Usage:       "Deletes existing NACL.",
			Action:      VpcDeleteNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1439",
		},
		{
			Name:        "CreateVpcPeeringConnection",
			Usage:       "Creates a VPC peering connection",
			Action:      VpcCreateVpcPeeringConnection,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/2107",
		},
		{
			Name:        "DescribeVpcPeeringConnections",
			Usage:       "Lists all existing VPC peering connections",
			Action:      VpcDescribeVpcPeeringConnections,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/2101",
		},
		{
			Name:        "DeleteVpcPeeringConnection",
			Usage:       "Deletes a VPC peering connection.",
			Action:      VpcDeleteVpcPeeringConnection,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/2104",
		},
		{
			Name:        "DescribeRouteTable",
			Usage:       "Describes RTs.",
			Action:      VpcDescribeRouteTable,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1420",
		},
		{
			Name:        "CreateRoute",
			Usage:       "Creates a route in specified RT",
			Action:      VpcCreateRoute,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5688",
		},
		{
			Name:        "DeleteRoute",
			Usage:       "Deletes a route in specified RT",
			Action:      VpcDeleteRoute,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/5689",
		},
		{
			Name:        "DeleteRouteTable",
			Usage:       "Deletes RT.",
			Action:      VpcDeleteRouteTable,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1418",
		},
		{
			Name:        "ModifyRouteTableAttribute",
			Usage:       "Modifies RT attribute.",
			Action:      VpcModifyRouteTableAttribute,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1417",
		},
		//{
		//	Name:        "action",
		//	Usage:       "",
		//	Action:      Vpcaction,
		//	Description: "Referrer: https://cloud.tencent.com/document/api/uri",
		//},
	}
)

func VpcDoAction(c *cli.Context) error {
	resp, err := vpc.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func VpcCreateVpc(c *cli.Context) error {
	resp, err := vpc.CreateVpc(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeVpcEx(c *cli.Context) error {
	resp, err := vpc.DescribeVpcEx(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeVpcClassicLink(c *cli.Context) error {
	resp, err := vpc.DescribeVpcClassicLink(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeVpcLimit(c *cli.Context) error {
	resp, err := vpc.DescribeVpcLimit(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyVpcAttribute(c *cli.Context) error {
	resp, err := vpc.ModifyVpcAttribute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

// func VpcAssociateVip(c *cli.Context) error {
// 	resp, err := vpc.AssociateVip(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	r, err := resp.String(formatOut)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(r)
// 	return nil
// }
//func VpcAttachClassicLinkVpc(c *cli.Context) error {
//	resp, err := vpc.AttachClassicLinkVpc(c.Args()...)
//	if err != nil {
//		return err
//	}
//	r, err := resp.String(formatOut)
//	if err != nil {
//		return err
//	}
//	fmt.Println(r)
//	return nil
//}
func VpcDeleteVpc(c *cli.Context) error {
	resp, err := vpc.DeleteVpc(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateSubnet(c *cli.Context) error {
	resp, err := vpc.CreateSubnet(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeSubnetEx(c *cli.Context) error {
	resp, err := vpc.DescribeSubnetEx(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeSubnet(c *cli.Context) error {
	resp, err := vpc.DescribeSubnet(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteSubnet(c *cli.Context) error {
	resp, err := vpc.DeleteSubnet(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateRouteTable(c *cli.Context) error {
	resp, err := vpc.CreateRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcAssociateRouteTable(c *cli.Context) error {
	resp, err := vpc.AssociateRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcInquiryVpnPrice(c *cli.Context) error {
	resp, err := vpc.InquiryVpnPrice(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeVpnGw(c *cli.Context) error {
	resp, err := vpc.DescribeVpnGw(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyVpnGw(c *cli.Context) error {
	resp, err := vpc.ModifyVpnGw(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeUserGwVendor(c *cli.Context) error {
	resp, err := vpc.DescribeUserGwVendor(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeUserGw(c *cli.Context) error {
	resp, err := vpc.DescribeUserGw(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyUserGw(c *cli.Context) error {
	resp, err := vpc.ModifyUserGw(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateDirectConnectGateway(c *cli.Context) error {
	resp, err := vpc.CreateDirectConnectGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeDirectConnectGateway(c *cli.Context) error {
	resp, err := vpc.DescribeDirectConnectGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyDirectConnectGateway(c *cli.Context) error {
	resp, err := vpc.ModifyDirectConnectGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateLocalIPTranslationNatRule(c *cli.Context) error {
	resp, err := vpc.CreateLocalIPTranslationNatRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeLocalIPTranslationNatRule(c *cli.Context) error {
	resp, err := vpc.DescribeLocalIPTranslationNatRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcInquiryNatPrice(c *cli.Context) error {
	resp, err := vpc.InquiryNatPrice(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateNatGateway(c *cli.Context) error {
	resp, err := vpc.CreateNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeNatGateway(c *cli.Context) error {
	resp, err := vpc.DescribeNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyNatGateway(c *cli.Context) error {
	resp, err := vpc.ModifyNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcEipBindNatGateway(c *cli.Context) error {
	resp, err := vpc.EipBindNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

// func VpcEipUnBindNatGateway(c *cli.Context) error {
// 	resp, err := vpc.EipBindNatGateway(c.Args()...)
// 	if err != nil {
// 		return err
// 	}
// 	r, err := resp.String(formatOut)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(r)
// 	return nil
// }
func VpcDeleteNatGateway(c *cli.Context) error {
	resp, err := vpc.DeleteNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyLocalIPTranslationNatRule(c *cli.Context) error {
	resp, err := vpc.ModifyLocalIPTranslationNatRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteLocalIPTranslationNatRule(c *cli.Context) error {
	resp, err := vpc.DeleteLocalIPTranslationNatRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteDirectConnectGateway(c *cli.Context) error {
	resp, err := vpc.DeleteDirectConnectGateway(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteUserGw(c *cli.Context) error {
	resp, err := vpc.DeleteUserGw(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.CreateVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeVpcPeeringConnections(c *cli.Context) error {
	resp, err := vpc.DescribeVpcPeeringConnections(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.DeleteVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeRouteTable(c *cli.Context) error {
	resp, err := vpc.DescribeRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteRouteTable(c *cli.Context) error {
	resp, err := vpc.DeleteRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateRoute(c *cli.Context) error {
	resp, err := vpc.CreateRoute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteRoute(c *cli.Context) error {
	resp, err := vpc.DeleteRoute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyRouteTableAttribute(c *cli.Context) error {
	resp, err := vpc.ModifyRouteTableAttribute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateNetworkAcl(c *cli.Context) error {
	resp, err := vpc.CreateNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDescribeNetworkAcl(c *cli.Context) error {
	resp, err := vpc.DescribeNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeleteNetworkAcl(c *cli.Context) error {
	resp, err := vpc.DeleteNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyNetworkAcl(c *cli.Context) error {
	resp, err := vpc.ModifyNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifyNetworkAclEntry(c *cli.Context) error {
	resp, err := vpc.ModifyNetworkAclEntry(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcCreateSubnetAclRule(c *cli.Context) error {
	resp, err := vpc.CreateSubnetAclRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcDeteleSubnetAclRule(c *cli.Context) error {
	resp, err := vpc.DeteleSubnetAclRule(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VpcModifySubnetAttribute(c *cli.Context) error {
	resp, err := vpc.ModifySubnetAttribute(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
