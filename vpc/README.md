## Virtual Private Cloud API Calls

Documentation: https://cloud.tencent.com/document/api/215/908

The following are examples of using tcpicli to issue API calls to VPC:

#### Query existing VPCs by vpcId (DescribeVpcEx)

Documentation: https://cloud.tencent.com/document/api/215/1372

`tcpicli -vv vpc DescribeVpcEx Region=gz vpcId=vpc-12xvscaz`

**Input:**
```
{
  "Action": "DescribeVpcEx",
  "Region": "gz",
  "vpcId": "vpc-12xvscaz"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": [
    {
      "cidrBlock": "192.168.0.0/16",
      "classicLinkNum": -1,
      "createTime": "2017-09-19 09:49:44",
      "isDefault": false,
      "isMulticast": false,
      "natNum": 0,
      "routeTableNum": 1,
      "subnetNum": 1,
      "unVpcId": "vpc-12xvscaz",
      "vpcDeviceNum": 0,
      "vpcId": "gz_vpc_222312",
      "vpcName": "james",
      "vpcPeerNum": 0,
      "vpgNum": 0,
      "vpnGwNum": 0
    }
  ],
  "message": "",
  "totalCount": 1
}
```

#### Create a new VPC (CreateVpc)

Documentation: https://cloud.tencent.com/document/api/215/1309

`tcpicli -vv vpc CreateVpc Region=gz vpcName=apiTest cidrBlock="10.0.0.0/16"`

**Input:**
```
{
  "Action": "CreateVpc",
  "Region": "gz",
  "cidrBlock": "10.0.0.0/16",
  "vpcName": "apiTest"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": "",
  "routeTableSet": [
    {
      "routeTableId": "gz_rtb_43796",
      "routeTableName": "default",
      "routeTableType": 1
    }
  ],
  "subnetSet": [],
  "uniqVpcId": "vpc-c4hlrvpr",
  "vpcCreateTime": "2017-10-01 16:06:37",
  "vpcId": "gz_vpc_226373"
}
```

#### Delete VPC (DeleteVpc)

Documentation: https://cloud.tencent.com/document/api/215/1307

`tcpicli -vv vpc DeleteVpc Region=gz vpcId=vpc-c4hlrvpr`

**Input:**
```
{
  "Action": "DeleteVpc",
  "Region": "gz",
  "vpcId": "vpc-c4hlrvpr"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Rename VPC (ModifyVpcAttribute)

Documentation: https://cloud.tencent.com/document/api/215/1310

`tcpicli -vv vpc do ModifyVpcAttribute Region=gz vpcId=vpc-9tejit0j vpcName=grobbledongs`

**Input:**
```
{
  "Action": "ModifyVpcAttribute",
  "Region": "gz",
  "vpcId": "vpc-9tejit0j",
  "vpcName": "grobbledongs"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Create Classic Link (AttachClassicLinkVpc)

Documentation: https://cloud.tencent.com/document/api/215/2098

`tcpicli -vv vpc AttachClassicLinkVpc Region=gz vpcId=vpc-e7xezjpn instanceIds.0=ins-kn52i9r6`

**Input:**
```
{
  "Action": "AttachClassicLinkVpc",
  "Region": "gz",
  "instanceIds.0": "ins-kn52i9r6",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": "",
  "taskId": 18176709
}
```

#### Query Existing Classic Links

Documentation: https://cloud.tencent.com/document/api/215/2112

`tcpicli -vv vpc DescribeVpcClassicLink Region=gz`

**Input:**
```
{
  "Action": "DescribeVpcClassicLink",
  "Region": "gz"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": [
    {
      "classicLinkId": "vcx-rw340z73",
      "createTime": "2017-10-02 10:32:24",
      "instanceId": "ins-kn52i9r6",
      "instanceName": "Unnamed",
      "lanIp": "10.104.167.105",
      "unVpcId": "vpc-e7xezjpn",
      "vpcId": "gz_vpc_225886"
    },
    {
      "classicLinkId": "vcx-qfmncw1x",
      "createTime": "2017-09-15 06:03:41",
      "instanceId": "ins-o5hamq0y",
      "instanceName": "kiyor",
      "lanIp": "10.104.98.1",
      "unVpcId": "vpc-jlqxir0t",
      "vpcId": "gz_vpc_214230"
    }
  ],
  "message": "",
  "totalCount": 2
}
```

#### Query Existing Classic Links (DetachClassicLinkVpc)

Documentation: https://cloud.tencent.com/document/api/215/2097

`tcpicli -vv vpc do DetachClassicLinkVpc Region=gz vpcId=vpc-e7xezjpn instanceIds.0=ins-kn52i9r6`

**Input:**
```
{
  "Action": "DetachClassicLinkVpc",
  "Region": "gz",
  "instanceIds.0": "ins-kn52i9r6",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": "",
  "taskId": 18180569
}
```

#### Query VPC limits (DescribeVpcLimit)

Documentation: https://cloud.tencent.com/document/api/215/1844

`tcpicli -vv vpc DescribeVpcLimit Region=gz type.0=1 type.1=2`

**Input:**
```
{
  "Action": "DescribeVpcLimit",
  "Region": "gz",
  "type.0": "1",
  "type.1": "2"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": {
    "limit": {
      "1": 10,
      "2": 10
    }
  },
  "message": ""
}
```

#### Create Subnet (CreateSubnet)

Documentation: https://cloud.tencent.com/document/api/215/1314

`tcpicli -vv vpc CreateSubnet Region=gz vpcId=vpc-e7xezjpn subnetSet.0.subnetName="testApiSubApiCall" subnetSet.0.cidrBlock="10.0.1.0/24" subnetSet.0.zoneId=100002`

**Input:**
```
{
  "Action": "CreateSubnet",
  "Region": "gz",
  "subnetSet.0.cidrBlock": "10.0.1.0/24",
  "subnetSet.0.subnetName": "testApiSubApiCall",
  "subnetSet.0.zoneId": "100002",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": "",
  "subnetSet": [
    {
      "cidrBlock": "10.0.1.0/24",
      "routeTableId": "gz_rtb_43926",
      "subnetId": "gz_subnet_82103",
      "subnetName": "testApiSubApiCall",
      "unSubnetId": "subnet-1pia0da0",
      "zone": "ap-guangzhou-2",
      "zoneId": 100002
    }
  ]
}
```

#### Delete Subnet (DeleteSubnet)

Documentation: https://cloud.tencent.com/document/api/215/1312

`tcpicli -vv vpc DeleteSubnet Region=gz vpcId=vpc-e7xezjpn subnetId=subnet-1pia0da0`

**Input:**
```
{
  "Action": "DeleteSubnet",
  "Region": "gz",
  "subnetId": "subnet-1pia0da0",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Modify Subnet (ModifySubnetAttribute)

Documentation: https://cloud.tencent.com/document/api/215/1313

`tcpicli -vv vpc ModifySubnetAttribute Region=gz vpcId=vpc-e7xezjpn subnetId=subnet-1pia0da0 subnetName=testApiNameChange`

**Input:**
```
{
  "Action": "ModifySubnetAttribute",
  "Region": "gz",
  "subnetId": "subnet-1pia0da0",
  "subnetName": "testApiNameChange",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Query Subnets on VPC (DescribeSubnetEx)

Documentation: https://cloud.tencent.com/document/api/215/1371

`tcpicli -vv vpc DescribeSubnetEx Region=gz vpcId=vpc-e7xezjpn`

**Input:**
```
{
  "Action": "DescribeSubnetEx",
  "Region": "gz",
  "vpcId": "vpc-e7xezjpn"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": [
    {
      "availableIPNum": 252,
      "broadcast": false,
      "cidrBlock": "10.0.0.0/24",
      "isDefault": false,
      "networkAclId": null,
      "routeTableId": "gz_rtb_43926",
      "routeTableName": "default",
      "rtbNum": 0,
      "subnetCreateTime": "2017-10-02 10:18:02",
      "subnetId": "gz_subnet_81968",
      "subnetName": "apiTestSub",
      "totalIPNum": 253,
      "unRouteTableId": "rtb-n590r9fs",
      "unSubnetId": "subnet-358ndkca",
      "unVpcId": "vpc-e7xezjpn",
      "vpcCidrBlock": "10.0.0.0/16",
      "vpcDevices": 1,
      "vpcId": "gz_vpc_225886",
      "vpcName": "apiTest",
      "zone": "ap-guangzhou-2",
      "zoneId": 100002
    }
  ],
  "message": "",
  "totalCount": 1
}
```

#### Create a new Route Table (CreateRouteTable)

Documentation: https://cloud.tencent.com/document/api/215/1419

`tcpicli -vv vpc CreateRouteTable Region=sh vpcId=vpc-70fpkwg8 routeTableName="apiTest" routeSet.0.destinationCidrBlock="192.168.10.0/24" routeSet.0.nextHub=pcx-deafiue0 routeSet.0.nextType=4`

**Input:**
```
{
  "Action": "CreateRouteTable",
  "Region": "sh",
  "routeSet.0.destinationCidrBlock": "192.168.10.0/24",
  "routeSet.0.nextHub": "pcx-deafiue0",
  "routeSet.0.nextType": "4",
  "routeTableName": "apiTest",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": "",
  "routeTableId": "sh_rtb_23969",
  "routeTableSet": [
    {
      "description": null,
      "destinationCidrBlock": "192.168.10.0/24",
      "nextHub": "pcx-deafiue0",
      "nextType": 4
    },
    {
      "description": null,
      "destinationCidrBlock": "Local",
      "nextHub": "Local",
      "nextType": 2
    }
  ],
  "unRouteTableId": "rtb-n66int2n"
}
```

#### Delete an existing Route Table (DeleteRouteTable)

Documentation: https://cloud.tencent.com/document/api/215/1418

`tcpicli -vv vpc DeleteRouteTable Region=sh vpcId=vpc-70fpkwg8 routeTableId=sh_rtb_23971`

**Input:**
```
{
  "Action": "DeleteRouteTable",
  "Region": "sh",
  "routeTableId": "sh_rtb_23971",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Modify an existing Route Table (ModifyRouteTable)

Documentation: https://cloud.tencent.com/document/api/215/1417

`tcpicli -vv vpc ModifyRouteTableAttribute Region=sh vpcId=vpc-70fpkwg8 routeTableId=rtb-0vlchz8j routeSet.0.destinationCidrBlock="192.168.30.0/24" routeSet.0.nextHub=pcx-deafiue0 routeSet.0.nextType=4`

**Input:**
```
{
  "Action": "ModifyRouteTableAttribute",
  "Region": "sh",
  "routeSet.0.destinationCidrBlock": "192.168.30.0/24",
  "routeSet.0.nextHub": "pcx-deafiue0",
  "routeSet.0.nextType": "4",
  "routeTableId": "rtb-0vlchz8j",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Query Route Tables (DescribeRouteTable)

Documentation: https://cloud.tencent.com/document/api/215/1420

`tcpicli -vv vpc DescribeRouteTable Region=sh`

**Input:**
```
{
  "Action": "DescribeRouteTable",
  "Region": "sh"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": [
    {
      "routeSet": [
        {
          "description": null,
          "destinationCidrBlock": "Local",
          "nextHub": "Local",
          "nextType": 2,
          "unNextHub": null
        }
      ],
      "routeTableCreateTime": "2017-10-04 00:17:24",
      "routeTableId": "sh_rtb_23952",
      "routeTableName": "default",
      "routeTableType": 1,
      "subnetNum": 1,
      "unRouteTableId": "rtb-l27ud66h",
      "unVpcId": "vpc-70fpkwg8",
      "vpcCidrBlock": "10.0.0.0/16",
      "vpcId": "sh_vpc_182690",
      "vpcName": "jamesApiTest"
    },
      "routeTableCreateTime": "2017-10-04 00:17:24",
      "routeTableId": "sh_rtb_23952",
      "routeTableName": "default",
      "routeTableType": 1,
      "subnetNum": 1,
      "unRouteTableId": "rtb-l27ud66h",
      "unVpcId": "vpc-70fpkwg8",
      "vpcCidrBlock": "10.0.0.0/16",
      "vpcId": "sh_vpc_182690",
      "vpcName": "jamesApiTest"
    }
  ],
  "message": "",
  "totalCount": 3
}
```

#### Associate Route Tables (AssociateRouteTable)

Documentation: https://cloud.tencent.com/document/api/215/1416

`tcpicli -vv vpc AssociateRouteTable Region=sh vpcId=vpc-70fpkwg8 subnetId=subnet-ff6vw5kh routeTableId=rtb-emo58vwt`

**Input:**
```
{
  "Action": "AssociateRouteTable",
  "Region": "sh",
  "routeTableId": "rtb-emo58vwt",
  "subnetId": "subnet-ff6vw5kh",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Create NACL (CreateNetworkAcl)

Documentation: https://cloud.tencent.com/document/api/215/1437

`tcpicli -vv vpc CreateNetworkAcl Region=sh vpcId=vpc-70fpkwg8 networkAclName="grobbledongs"`

**Input:**
```
{
  "Action": "CreateNetworkAcl",
  "Region": "sh",
  "networkAclName": "grobbledongs",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": {
    "networkAclId": "acl-f1vqk2dn"
  },
  "message": ""
}
```

#### Delete NACL (DeleteNetworkAcl)

Documentation: https://cloud.tencent.com/document/api/215/1439

`tcpicli -vv vpc CreateNetworkAcl Region=sh vpcId=vpc-70fpkwg8 networkAclName="grobbledongs"`

**Input:**
```
 {
  "Action": "DeleteNetworkAcl",
  "Region": "sh",
  "networkAclId": "acl-f1vqk2dn",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Modify NACL (ModifyNetworkAcl) (Only used to rename acls)

Documentation: https://cloud.tencent.com/document/api/215/1443

`tcpicli -vv vpc ModifyNetworkAcl Region=sh vpcId=vpc-70fpkwg8 networkAclId=acl-3bx6kb3d networkAclName="somethingDifferent"`

**Input:**
```
{
  "Action": "ModifyNetworkAcl",
  "Region": "sh",
  "networkAclId": "acl-3bx6kb3d",
  "networkAclName": "somethingDifferent",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Describe NACL list (DescribeNetworkAcl)

Documentation: https://cloud.tencent.com/document/api/215/1441

`tcpicli -vv vpc DescribeNetworkAcl Region=sh vpcId=vpc-70fpkwg8"`

**Input:**
```
{
  "Action": "DescribeNetworkAcl",
  "Region": "sh",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": [
    {
      "createTime": "2017-10-04 16:30:05",
      "networkAclEntrySet": {
        "egress": [
          {
            "action": 1,
            "cidrIp": "0.0.0.0/0",
            "desc": "",
            "ipProtocol": "all",
            "portRange": "ALL"
          }
        ],
        "ingress": [
          {
            "action": 1,
            "cidrIp": "0.0.0.0/0",
            "desc": "",
            "ipProtocol": "all",
            "portRange": "ALL"
          }
        ]
      },
      "networkAclId": "acl-3bx6kb3d",
      "networkAclName": "somethingDifferent",
      "subnetNum": 0,
      "subnetSet": [],
      "unVpcId": "vpc-70fpkwg8",
      "vpcCidrBlock": "10.0.0.0/16",
      "vpcId": "sh_vpc_182690",
      "vpcName": "jamesApiTest"
    }
  ],
  "message": "",
  "totalCount": 1
}
```

#### Set Network ACL Rules (ModifyNetworkAclEntry)

Documentation: https://cloud.tencent.com/document/api/215/1444

`tcpicli -vv vpc ModifyNetworkAclEntry Region=sh vpcId=vpc-70fpkwg8 networkAclId=acl-3bx6kb3d ruleDirection=1 networkAclEntrySet.0.desc="test" networkAclEntrySet.0.ipProtocol=all networkAclEntrySet.0.cidrIp="0.0.0.0/0" networkAclEntrySet.0.portRange=ALL networkAclEntrySet.0.action=1`

**Input:**
```
{
  "Action": "ModifyNetworkAclEntry",
  "Region": "sh",
  "networkAclEntrySet.0.action": "1",
  "networkAclEntrySet.0.cidrIp": "0.0.0.0/0",
  "networkAclEntrySet.0.desc": "test",
  "networkAclEntrySet.0.ipProtocol": "all",
  "networkAclEntrySet.0.portRange": "ALL",
  "networkAclId": "acl-3bx6kb3d",
  "ruleDirection": "1",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Bind Network ACL to subnet (CreateSubnetAclRule)

Documentation: https://cloud.tencent.com/document/api/215/1438

`tcpicli -vv vpc CreateSubnetAclRule Region=sh vpcId=vpc-70fpkwg8 networkAclId=acl-3bx6kb3d subnetIds.0=subnet-ff6vw5kh`

**Input:**
```
{
  "Action": "CreateSubnetAclRule",
  "Region": "sh",
  "networkAclId": "acl-3bx6kb3d",
  "subnetIds.0": "subnet-ff6vw5kh",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Unbind Network ACL to subnet (DeleteSubnetAclRule)

Documentation: https://cloud.tencent.com/document/api/215/1442

`tcpicli -vv vpc deleteSubnetAclRule Region=sh vpcId=vpc-70fpkwg8 networkAclId=acl-3bx6kb3d subnetIds.0=subnet-ff6vw5kh`

**Input:**
```
{
  "Action": "DeleteSubnetAclRule",
  "Region": "sh",
  "networkAclId": "acl-3bx6kb3d",
  "subnetIds.0": "subnet-ff6vw5kh",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```

#### Unbind Network ACL to subnet (DescribeVpcPeeringConnections)

Documentation: https://cloud.tencent.com/document/api/215/2101

`tcpicli -vv vpc deleteSubnetAclRule Region=sh vpcId=vpc-70fpkwg8 networkAclId=acl-3bx6kb3d subnetIds.0=subnet-ff6vw5kh`

**Input:**
```
{
  "Action": "DeleteSubnetAclRule",
  "Region": "sh",
  "networkAclId": "acl-3bx6kb3d",
  "subnetIds.0": "subnet-ff6vw5kh",
  "vpcId": "vpc-70fpkwg8"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "message": ""
}
```
