## Cloud Container Service API Calls

Documentation: https://cloud.tencent.com/document/api/457/9427

The following are examples of using tcpicli to issue API calls to CCS:

#### Query existing clusters per region (DescribeCluster)

Documentation: https://cloud.tencent.com/document/api/457/9448

`tcpicli -vv ccs DescribeCluster Region=gz`

**Input:** 
```
{
  "Action": "DescribeCluster",
  "ClusterId": "cls-0shevcf0",
  "Region": "gz"
}
```

**Output:**
```
{
  "code": 0,
  "codedesc": "Success",
  "data": {
    "clusters": [
      {
        "clusterCIDR": "10.0.0.0/14",
        "clusterExternalEndpoint": "",
        "clusterId": "cls-0shevcf0",
        "clusterName": "jamesTest",
        "createdAt": "2017-09-26 12:53:51",
        "description": "",
        "k8sVersion": "1.4.6",
        "nodeNum": 2,
        "nodeStatus": "AllNormal",
        "openHttps": 1,
        "os": "ubuntu16.04.1 LTSx86_64",
        "projectId": 0,
        "region": "ap-guangzhou",
        "regionId": 1,
        "status": "Running",
        "totalCpu": 4,
        "totalMem": 8,
        "unVpcId": "vpc-12xvscaz",
        "updatedAt": "2017-09-26 12:58:02",
        "vpcId": 222312
      }
    ],
    "totalCount": 3
  },
  "message": ""
}
```

#### Query instances that are running in particular clusters (DescribeClusterInstances)

Documentation: https://cloud.tencent.com/document/api/457/9449

`tcpicli -vv ccs DescribeClusterInstances Region=gz clusterId=cls-0shevcf0`

**Input:**
```
{
  "Action": "DescribeCluster",
  "ClusterId": "cls-0shevcf0",
  "Region": "gz"
}
```

**Output:**
```
{
  "code": 0,
  "codedesc": "Success",
  "data": {
    "clusters": [
      {
        "clusterCIDR": "10.0.0.0/14",
        "clusterExternalEndpoint": "",
        "clusterId": "cls-0shevcf0",
        "clusterName": "jamesTest",
        "createdAt": "2017-09-26 12:53:51",
        "description": "",
        "k8sVersion": "1.4.6",
        "nodeNum": 2,
        "nodeStatus": "AllNormal",
        "openHttps": 1,
        "os": "ubuntu16.04.1 LTSx86_64",
        "projectId": 0,
        "region": "ap-guangzhou",
        "regionId": 1,
        "status": "Running",
        "totalCpu": 4,
        "totalMem": 8,
        "unVpcId": "vpc-12xvscaz",
        "updatedAt": "2017-09-26 12:58:02",
        "vpcId": 222312
      }
    ],
    "totalCount": 1
  },
  "message": ""
}
```

#### Create new cluster (CreateCluster)

CreateCluster API Documentation: https://cloud.tencent.com/document/api/457/9444
DescribeAvailibilityZones API Documentation (Relevant to cluster CVM configuration): https://cloud.tencent.com/document/api/213/1286

`tcpicli -vv ccs CreateCluster Region=gz clusterName=jamesTest cpu=2 mem=4 "osName=ubuntu16.04.1 LTSx86_64" bandwidth=5 bandwidthType=PayByHour vpcId=vpc-12xvscaz zoneId=100002 subnetId=subnet-45a5fp4q isVpcGateway=1 storageSize=50 rootSize=50 goodsNum=2 clusterCIDR=10.0.0.0/14`

**Input:**
```
{
  "Action": "CreateCluster",
  "Region": "gz",
  "bandwidth": "5",
  "bandwidthType": "PayByHour",
  "clusterCIDR": "10.0.0.0/14",
  "clusterName": "jamesTest",
  "cpu": "2",
  "goodsNum": "2",
  "isVpcGateway": "1",
  "mem": "4",
  "osName": "ubuntu16.04.1 LTSx86_64",
  "rootSize": "50",
  "storageSize": "50",
  "subnetId": "subnet-45a5fp4q",
  "vpcId": "vpc-12xvscaz",
  "zoneId": "100002"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": {
    "clusterId": "cls-8rzaq0eu",
    "requestId": 2946
  },
  "message": ""
}
```

#### Delete cluster (DeleteCluster)

`tcpicli -vv ccs DeleteCluster Region=gz clusterId=cls-0shevcf0`

**Input:**
```
{
  "Action": "DeleteCluster",
  "Region": "gz",
  "clusterId": "cls-0shevcf0"
}
```

**Output:**
```
{
  "code": 0,
  "codeDesc": "Success",
  "data": {
    "requestId": 2945
  },
  "message": ""
}
```
