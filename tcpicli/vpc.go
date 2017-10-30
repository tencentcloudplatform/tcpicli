package main

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/vpc"
	"github.com/urfave/cli"
)

var (
	funcVpc []cli.Command = []cli.Command{
		{
			Name:        "do",
			Usage:       "do <action> <args1=value1> [args2=value2] ...",
			Action:      VpcDoAction,
			Description: "do VPC action and output json response",
		},
		{
			Name:        "CreateVpc",
			Usage:       "Used to create VPCs",
			Action:      VpcCreateVpc,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1309",
		},
		{
			Name:        "DeleteVpc",
			Usage:       "Used to delete VPCs",
			Action:      VpcDeleteVpc,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1307",
		},
		{
			Name:        "DescribeVpcEx",
			Usage:       "Used to query VPCs",
			Action:      VpcDescribeVpcEx,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1372",
		},
		{
			Name:        "ModifyVpcAttribute",
			Usage:       "Used to change VPC attribute. Currently only thing supported is name change???",
			Action:      VpcModifyVpcAttribute,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1310",
		},
		{
			Name:        "AttachClassicLinkVpc",
			Usage:       "Used to connect VPC and basic network devices",
			Action:      VpcAttachClassicLinkVpc,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2098",
		},
		{
			Name:        "DescribeVpcClassicLink",
			Usage:       "Used to describe existing VPC classic links",
			Action:      VpcDescribeVpcClassicLink,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2112",
		},
		{
			Name:        "DetachClassicLinkVpc",
			Usage:       "Used to delete VPC classic links",
			Action:      VpcDetachClassicLinkVpc,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2097",
		},
		{
			Name:        "DescribeVpcLimit",
			Usage:       "Used to describe VPC limits",
			Action:      VpcDescribeVpcLimit,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1844",
		},
		{
			Name:        "CreateSubnet",
			Usage:       "Used to create new subnets",
			Action:      VpcCreateSubnet,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1314",
		},
		{
			Name:        "DeleteSubnet",
			Usage:       "Used to delete existing subnets",
			Action:      VpcDeleteSubnet,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1312",
		},
		{
			Name:        "ModifySubnetAttribute",
			Usage:       "Used to change the name of a subnet",
			Action:      VpcModifySubnetAttribute,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1313",
		},
		{
			Name:        "DescribeSubnetEx",
			Usage:       "Used to query list of subnets against existing VPC",
			Action:      VpcDescribeSubnetEx,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1371",
		},
		{
			Name:        "DescribeSubnet",
			Usage:       "Used to get details about a particular subnet",
			Action:      VpcDescribeSubnet,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1311",
		},
		{
			Name:        "CreateRouteTable",
			Usage:       "Used to create new route tables",
			Action:      VpcCreateRouteTable,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1419",
		},
		{
			Name:        "DeleteRouteTable",
			Usage:       "Used to delete route tables",
			Action:      VpcDeleteRouteTable,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1418",
		},
		{
			Name:        "ModifyRouteTableAttribute",
			Usage:       "Used to modify route tables",
			Action:      VpcModifyRouteTableAttribute,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1417",
		},
		{
			Name:        "DescribeRouteTable",
			Usage:       "Used to query route tables",
			Action:      VpcDescribeRouteTable,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1420",
		},
		{
			Name:        "AssociateRouteTable",
			Usage:       "Used to associate route tables with subnets, and toggle off the association",
			Action:      VpcAssociateRouteTable,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1416",
		},
		{
			Name:        "CreateNetworkAcl",
			Usage:       "Used to create a NACL",
			Action:      VpcCreateNetworkAcl,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1437",
		},
		{
			Name:        "DeleteNetworkAcl",
			Usage:       "Used to delete a NACL",
			Action:      VpcDeleteNetworkAcl,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1439",
		},
		{
			Name:        "ModifyNetworkAcl",
			Usage:       "Used to modify a NACL",
			Action:      VpcModifyNetworkAcl,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1443",
		},
		{
			Name:        "DescribeNetworkAcl",
			Usage:       "Used to describe NACLs",
			Action:      VpcDescribeNetworkAcl,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1441",
		},
		{
			Name:        "ModifyNetworkAclEntry",
			Usage:       "Used to modify NACLs",
			Action:      VpcModifyNetworkAclEntry,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1444",
		},
		{
			Name:        "CreateSubnetAclRule",
			Usage:       "Used to associate NACLs with subnets",
			Action:      VpcCreateSubnetAclRule,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1438",
		},
		{
			Name:        "DeteleSubnetAclRule",
			Usage:       "Used to disassociate NACLs with subnets",
			Action:      VpcDeteleSubnetAclRule,
			Description: "referrer: https://cloud.tencent.com/document/api/215/1442",
		},
		{
			Name:        "DescribeVpcPeeringConnections",
			Usage:       "Used to list VPC peering connections",
			Action:      VpcDescribeVpcPeeringConnections,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2101",
		},
		{
			Name:        "CreateVpcPeeringConnection",
			Usage:       "Used to create new intra-region VPC peering connections (same region)",
			Action:      VpcCreateVpcPeeringConnection,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2107",
		},
		{
			Name:        "DeleteVpcPeeringConnection",
			Usage:       "Used to delete intra-region VPC peering connections (same region)",
			Action:      VpcDeleteVpcPeeringConnection,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2104",
		},
		{
			Name:        "ModifyVpcPeeringConnection",
			Usage:       "Used to modify intra-region VPC peering connections (same region)",
			Action:      VpcModifyVpcPeeringConnection,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2103",
		},
		{
			Name:        "AcceptVpcPeeringConnection",
			Usage:       "Used to accept intra-region VPC peering connections from another account (same region)",
			Action:      VpcAcceptVpcPeeringConnection,
			Description: "referrer: https://cloud.tencent.com/document/api/215/2106",
		},
		{
			Name:        "RejectVpcPeeringConnection",
			Usage:       "Used to reject intra-region VPC peering connections from another account (same region)",
			Action:      VpcRejectVpcPeeringConnection,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/2105",
		},
		{
			Name:        "CreateVpcPeeringConnectionEx",
			Usage:       "Used to create inter-region VPC peering connections from another account (different region)",
			Action:      VpcCreateVpcPeeringConnectionEx,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/4803",
		},
		{
			Name:        "DeleteVpcPeeringConnectionEx",
			Usage:       "Used to delete inter-region VPC peering connections from another account (different region)",
			Action:      VpcDeleteVpcPeeringConnectionEx,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/4804",
		},
		{
			Name:        "ModifyVpcPeeringConnectionEx",
			Usage:       "Used to modify inter-region VPC peering connections from another account (different region)",
			Action:      VpcModifyVpcPeeringConnectionEx,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/4805",
		},
		{
			Name:        "AcceptVpcPeeringConnectionEx",
			Usage:       "Used to accept inter-region VPC peering connections from another account (different region)",
			Action:      VpcAcceptVpcPeeringConnectionEx,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/4806",
		},
		{
			Name:        "RejectVpcPeeringConnectionEx",
			Usage:       "Used to reject inter-region VPC peering connections from another account (different region)",
			Action:      VpcRejectVpcPeeringConnectionEx,
			Description: "referrer: https://intl.cloud.tencent.com/document/api/215/4807",
		},
		{
			Name:        "CreateNatGateway",
			Usage:       "Used to create NAT GW",
			Action:      VpcCreateNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4094",
		},
		{
			Name:        "DeleteNatGateway",
			Usage:       "Used to delete NAT GW",
			Action:      VpcDeleteNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4087",
		},
		{
			Name:        "ModifyNatGateway",
			Usage:       "Used to modify NAT GW",
			Action:      VpcModifyNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4086",
		},
		{
			Name:        "DescribeNatGateway",
			Usage:       "Used to query NAT GW",
			Action:      VpcDescribeNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4088",
		},
		{
			Name:        "QueryNatGatewayProductionStatus",
			Usage:       "Used to query production status of NAT GW",
			Action:      VpcQueryNatGatewayProductionStatus,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4089",
		},
		{
			Name:        "UpgradeNatGateway",
			Usage:       "Used to upgrade NAT GW",
			Action:      VpcUpgradeNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4090",
		},
		{
			Name:        "EipBindNatGateway",
			Usage:       "Used to bind an EIP to a NAT GW",
			Action:      VpcEipBindNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4093",
		},
		{
			Name:        "EipUnBindNatGateway",
			Usage:       "Used to unbind an EIP from a NAT GW",
			Action:      VpcEipUnBindNatGateway,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4092",
		},
		{
			Name:        "CreateNetworkInterface",
			Usage:       "Used to create ENI",
			Action:      VpcCreateNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4811",
		},
		{
			Name:        "DeleteNetworkInterface",
			Usage:       "Used to delete ENI",
			Action:      VpcDeleteNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4813",
		},
		{
			Name:        "ModifyNetworkInterface",
			Usage:       "Used to modify ENI",
			Action:      VpcModifyNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/5372",
		},
		{
			Name:        "AssignPrivateIpAddresses",
			Usage:       "Used to assign private IP address to ENI",
			Action:      VpcAssignPrivateIpAddresses,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4817",
		},
		{
			Name:        "AttachNetworkInterface",
			Usage:       "Used to request private IP to ENI",
			Action:      VpcAttachNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4820",
		},
		{
			Name:        "DescribeNetworkInterfaces",
			Usage:       "Used to query ENI",
			Action:      VpcDescribeNetworkInterfaces,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4814",
		},
		{
			Name:        "DetachNetworkInterface",
			Usage:       "Used to unbind ENIs",
			Action:      VpcDetachNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4822",
		},
		{
			Name:        "MigrateNetworkInterface",
			Usage:       "Used to migrate ENI",
			Action:      VpcMigrateNetworkInterface,
			Description: "referrer: https://cloud.tencent.com/document/api/215/5384",
		},
		{
			Name:        "MigratePrivateIpAddress",
			Usage:       "Used to migrate private IPs",
			Action:      VpcMigratePrivateIpAddress,
			Description: "referrer: https://cloud.tencent.com/document/api/215/5385",
		},
		{
			Name:        "UnassignPrivateIpAddresses",
			Usage:       "Used to return the IP of ENI",
			Action:      VpcUnassignPrivateIpAddresses,
			Description: "referrer: https://cloud.tencent.com/document/api/215/4819",
		},
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
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteVpc(c *cli.Context) error {
	resp, err := vpc.DeleteVpc(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeVpcEx(c *cli.Context) error {
	resp, err := vpc.DescribeVpcEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyVpcAttribute(c *cli.Context) error {
	resp, err := vpc.ModifyVpcAttribute(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAttachClassicLinkVpc(c *cli.Context) error {
	resp, err := vpc.AttachClassicLinkVpc(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeVpcClassicLink(c *cli.Context) error {
	resp, err := vpc.DescribeVpcClassicLink(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDetachClassicLinkVpc(c *cli.Context) error {
	resp, err := vpc.DetachClassicLinkVpc(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeVpcLimit(c *cli.Context) error {
	resp, err := vpc.DescribeVpcLimit(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateSubnet(c *cli.Context) error {
	resp, err := vpc.CreateSubnet(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteSubnet(c *cli.Context) error {
	resp, err := vpc.DeleteSubnet(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifySubnetAttribute(c *cli.Context) error {
	resp, err := vpc.ModifySubnetAttribute(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeSubnetEx(c *cli.Context) error {
	resp, err := vpc.DescribeSubnetEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeSubnet(c *cli.Context) error {
	resp, err := vpc.DescribeSubnet(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateRouteTable(c *cli.Context) error {
	resp, err := vpc.CreateRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteRouteTable(c *cli.Context) error {
	resp, err := vpc.DeleteRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyRouteTableAttribute(c *cli.Context) error {
	resp, err := vpc.ModifyRouteTableAttribute(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeRouteTable(c *cli.Context) error {
	resp, err := vpc.DescribeRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAssociateRouteTable(c *cli.Context) error {
	resp, err := vpc.AssociateRouteTable(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateNetworkAcl(c *cli.Context) error {
	resp, err := vpc.CreateNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteNetworkAcl(c *cli.Context) error {
	resp, err := vpc.DeleteNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyNetworkAcl(c *cli.Context) error {
	resp, err := vpc.ModifyNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeNetworkAcl(c *cli.Context) error {
	resp, err := vpc.DescribeNetworkAcl(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyNetworkAclEntry(c *cli.Context) error {
	resp, err := vpc.ModifyNetworkAclEntry(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateSubnetAclRule(c *cli.Context) error {
	resp, err := vpc.CreateSubnetAclRule(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeteleSubnetAclRule(c *cli.Context) error {
	resp, err := vpc.DeteleSubnetAclRule(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeVpcPeeringConnections(c *cli.Context) error {
	resp, err := vpc.DescribeVpcPeeringConnections(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.CreateVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.DeleteVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.ModifyVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAcceptVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.AcceptVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcRejectVpcPeeringConnection(c *cli.Context) error {
	resp, err := vpc.RejectVpcPeeringConnection(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateVpcPeeringConnectionEx(c *cli.Context) error {
	resp, err := vpc.CreateVpcPeeringConnectionEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteVpcPeeringConnectionEx(c *cli.Context) error {
	resp, err := vpc.DeleteVpcPeeringConnectionEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyVpcPeeringConnectionEx(c *cli.Context) error {
	resp, err := vpc.ModifyVpcPeeringConnectionEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAcceptVpcPeeringConnectionEx(c *cli.Context) error {
	resp, err := vpc.AcceptVpcPeeringConnectionEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcRejectVpcPeeringConnectionEx(c *cli.Context) error {
	resp, err := vpc.RejectVpcPeeringConnectionEx(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateNatGateway(c *cli.Context) error {
	resp, err := vpc.CreateNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteNatGateway(c *cli.Context) error {
	resp, err := vpc.DeleteNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyNatGateway(c *cli.Context) error {
	resp, err := vpc.ModifyNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeNatGateway(c *cli.Context) error {
	resp, err := vpc.DescribeNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcQueryNatGatewayProductionStatus(c *cli.Context) error {
	resp, err := vpc.QueryNatGatewayProductionStatus(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcUpgradeNatGateway(c *cli.Context) error {
	resp, err := vpc.UpgradeNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcEipBindNatGateway(c *cli.Context) error {
	resp, err := vpc.EipBindNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcEipUnBindNatGateway(c *cli.Context) error {
	resp, err := vpc.EipUnBindNatGateway(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcCreateNetworkInterface(c *cli.Context) error {
	resp, err := vpc.CreateNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDeleteNetworkInterface(c *cli.Context) error {
	resp, err := vpc.DeleteNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcModifyNetworkInterface(c *cli.Context) error {
	resp, err := vpc.ModifyNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAssignPrivateIpAddresses(c *cli.Context) error {
	resp, err := vpc.AssignPrivateIpAddresses(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcAttachNetworkInterface(c *cli.Context) error {
	resp, err := vpc.AttachNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDescribeNetworkInterfaces(c *cli.Context) error {
	resp, err := vpc.DescribeNetworkInterfaces(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcDetachNetworkInterface(c *cli.Context) error {
	resp, err := vpc.DetachNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcMigrateNetworkInterface(c *cli.Context) error {
	resp, err := vpc.MigrateNetworkInterface(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcMigratePrivateIpAddress(c *cli.Context) error {
	resp, err := vpc.MigratePrivateIpAddress(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func VpcUnassignPrivateIpAddresses(c *cli.Context) error {
	resp, err := vpc.UnassignPrivateIpAddresses(c.Args()...)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
