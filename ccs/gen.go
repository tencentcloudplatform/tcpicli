// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	version := "Version=2017-03-12"
	region := "Region=gz"
	vpcName := "vpcName=ccsgenapivpc"
	cidrBlock := "cidrBlock=10.0.0.0/16"
	subnetCidr := "subnetSet.0.cidrBlock=10.0.0.0/24"
	subnetZone := "subnetSet.0.zoneId=100002"
	subnetName := "subnetSet.0.subnetName=ccsgenapisubnet"
	clusterName := "clusterName=ccsgencluster"
	bandwidthType := "bandwidthType=PayByTraffic"
	bandwidth := "bandwidth=100"
	instanceName := "InstanceName=ccsgenapicvm"
	// keyId := "echo skey-impvf5cj"
	clusterCIDR := "clusterCIDR=172.16.0.0/16"
	isVpcGateway := "isVpcGateway=0"
	zoneId := "zoneId=100002"
	serviceName := "serviceName=apitest"
	serviceDescription := "serviceDescription=descriptionnowhitespace"
	osName := "osName=centos7.2x86_64"
	nginxImage := "containers.0.image=nginx:latest"
	containerPort := "portMappings.0.containerPort=80"
	lbPort := "portMappings.0.lbPort=80"
	serviceProto := "portMappings.0.protocol=TCP"
	cpu := "cpu=2"
	mem := "mem=4"
	storageSize := "storageSize=50"
	rootSize := "rootSize=50"
	goodsNum := "goodsNum=1"
	nodeDeleteMode := "nodeDeleteMode=Return"

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			`SET imageId=tcpicli img DescribeImages ` + region + ` Version=2017-03-12 | grep -E -B1 -i "centos.*7\.3.*64" | grep -i imageid | awk '{ print $2 }' | tr -d "\"," | head -n1`,
			`DO echo $imageId`,
			`DO tcpicli vpc CreateVpc ` + region + " " + vpcName + " " + cidrBlock,
			`SET vpcId=tcpicli -f "{{range .Data}}{{.UnVpcID}}{{end}}" vpc DescribeVpcEx ` + region + " " + vpcName,
			`DO echo $vpcId`,
			`DO tcpicli vpc CreateSubnet vpcId=$vpcId ` + region + " " + subnetCidr + " " + subnetName + " " + subnetZone,
			`SET subnetId=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			`DO echo $subnetId`,
			`SET placementZone=tcpicli -f "{{range .Data}}{{.UnSubnetID}}{{end}}" vpc DescribeSubnetEx vpcId=$vpcId ` + region,
			`DO echo $placementZone`,
			"CreateCluster",
			`DO sleep 5`,
			`SET clusterId=tcpicli -f "{{range .Data.Clusters}}{{.ClusterID}}{{end}}" ccs DescribeCluster ` + clusterName + " " + region,
			`DO tcpicli cvm RunInstances Placement.Zone=$placementZone VirtualPrivateCloud.VpcId=$vpcId VirtualPrivateCloud.SubnetId=$subnetId InstanceChargeType=POSTPAID_BY_HOUR InstanceType=S2.SMALL1 SystemDisk.DiskType=CLOUD_BASIC SystemDisk.DiskSize=50 ` + region + " " + version + " " + instanceName + " ",
			`DO sleep 20`,
			`SET existingCvmId=tcpicli -f '{{range .Response.InstanceSet}}{{if eq .InstanceName "ccsgenapicvm"}}{{.InstanceID}}{{end}}{{end}}' cvm DescribeInstances ` + region,
			"AddClusterInstancesFromExistedCvm",
			`DO sleep 5`,
			"AddClusterInstances",
			`DO sleep 5`,
			"CreateClusterService",
			`DO sleep 2`,
			"DescribeCluster",
			"DescribeClusterInstances",
			"DescribeClusterService",
			"DescribeClusterServiceInfo",
			"ModifyClusterService",
			`DO sleep 10`,
			"PauseClusterService",
			`DO sleep 10`,
			"ResumeClusterService",
			`DO sleep 10`,
			"ModifyServiceDescription",
			"DescribeServiceEvent",
			"RollBackClusterService",
			`DO sleep 10`,
			"DescribeServiceInstance",
			"ModifyServiceReplicas",
			`DO sleep 10`,
			"DeleteClusterService",
			`DO sleep 10`,
			`SET deleteInstances=tcpicli -f "{{range .Data.Instances}}{{.Name}}{{\" \"}}{{end}}" ccs DescribeServiceInstance Region=$Region clusterId=$clusterID serviceName=$serviceName | awk '{ print $1 }'`,
			"DeleteInstances",
			`DO sleep 10`,
			"DeleteClusterInstances",
			"DeleteCluster",
			`DO tcpicli vpc DeleteSubnet vpcId=$vpcId subnetId=$subnetId ` + region,
			`DO tcpicli vpc DeleteVpc vpcId=$vpcId ` + region,
		},
		FuncMap: map[string][]string{
			"CreateCluster": []string{"457/9444",
				region,
				cpu,
				mem,
				osName,
				bandwidth,
				bandwidthType,
				isVpcGateway,
				storageSize,
				rootSize,
				goodsNum,
				clusterCIDR,
				zoneId,
				clusterName,
				"vpcId=$vpcId",
				"subnetId=$subnetId",
			},
			"AddClusterInstances": []string{"457/9447",
				region,
				isVpcGateway,
				zoneId,
				cpu,
				mem,
				bandwidthType,
				bandwidth,
				storageSize,
				rootSize,
				goodsNum,
				"subnetId=$subnetId",
				"clusterId=$clusterId",
			},
			"AddClusterInstancesFromExistedCvm": []string{"457/9458",
				region,
				"clusterId=$clusterId",
				"instanceIds.0=$existingCvmId",
				"keyId=$keyId",
			},
			"DescribeCluster": []string{"457/9448"},
			"DescribeClusterInstances": []string{"457/9449",
				region,
				"clusterId=$clusterId",
			},
			"DeleteClusterInstances": []string{"457/9446",
				region,
				nodeDeleteMode,
				"clusterId=$clusterId",
				"instanceIds.0=ins-mpomofr6",
			},
			"DeleteCluster": []string{"457/9445"},
			"CreateClusterService": []string{"457/9436",
				region,
				serviceName,
				nginxImage,
				containerPort,
				lbPort,
				serviceProto,
				"clusterId=$clusterId",
				"containers.0.containerName=$serviceName",
				"replicas=$replicasMod",
			},
			"DescribeClusterService": []string{"457/9440",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"DescribeClusterServiceInfo": []string{"457/9441",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"ModifyClusterService": []string{"457/9434",
				region,
				serviceName,
				nginxImage,
				containerPort,
				lbPort,
				serviceProto,
				"replicas=$replicasMod",
				"containers.0.containerName=$serviceName",
				"clusterId=$clusterId",
				"strategy=RollingUpdate",
				"minReadySeconds=10",
			},
			"PauseClusterService": []string{"457/9439",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"ResumeClusterService": []string{"457/9442",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"ModifyServiceDescription": []string{"457/9434",
				region,
				serviceName,
				serviceDescription,
				"clusterId=$clusterId",
			},
			"DescribeServiceEvent": []string{"457/9443",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"RollBackClusterService": []string{"457/9438",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"DescribeServiceInstance": []string{"457/9433",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
			"ModifyServiceReplicas": []string{"457/9431",
				region,
				serviceName,
				"clusterId=$clusterId",
				"scaleTo=$replicasMod",
			},
			"DeleteInstances": []string{"457/9432",
				region,
				"clusterId=$clusterId",
				"instances.0=$deleteInstances",
			},
			"DeleteClusterService": []string{"457/9437",
				region,
				serviceName,
				"clusterId=$clusterId",
			},
		},
		PkgName: "ccs",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
