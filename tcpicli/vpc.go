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
			Name:        "CreateNetworkAcl",
			Usage:       "Creates a new NACL.",
			Action:      VpcCreateNetworkAcl,
			Description: "Referrer: https://cloud.tencent.com/document/api/215/1437",
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

//func VpcACTION(c *cli.Context) error {
//	resp, err := vpc.ACTION(c.Args()...)
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
