// +build ignore

package main

import (
	. "."
	// "fmt"
	"github.com/tencentcloudplatform/tcpicli/autogen"
	// "time"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	region := "echo gz"
	vpcId := "echo vpc-l89napt5"       // need hardcode work
	subnetId := "echo subnet-at7hi4d6" // need hardcode work
	clusterID := "apitest"
	bandwidthType := "echo PayByTraffic"
	bandwidth := "echo 100"
	keyId := "echo skey-impvf5cj" // need hardcode work
	clusterCIDR := "echo 172.16.0.0/16"
	isVpcGateway := "echo 0"
	zoneId := "echo 100002"
	serviceName := "echo apitestnginx"
	osName := "echo centos7.2x86_64"
	nginxImage := "echo nginx:latest"
	nginxPort := "echo 80"
	nginxProto := "echo TCP"
	cpu := "echo 2"
	mem := "echo 4"
	storageSize := "echo 50"
	rootSize := "echo 50"
	goodsNum := "echo 1"
	nodeDeleteMode := "echo Return"
	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			`SET clusterID=tcpicli -f "{{range .Data.Clusters}}{{.ClusterID}}{{end}}" ccs DescribeCluster Region=gz clusterName=` + clusterID,
			`SET serviceName=tcpicli -f ""`,
			"SET Region=" + region,
			"SET vpcId=" + vpcId,
			"SET keyId=" + keyId,
			"SET subnetId=" + subnetId,
			"SET bandwidthType=" + bandwidthType,
			"SET bandwidth=" + bandwidth,
			"SET clusterCIDR=" + clusterCIDR,
			"SET isVpcGateway=" + isVpcGateway,
			"SET zoneId=" + zoneId,
			"SET osName=" + osName,
			"SET serviceName=" + serviceName,
			"SET cpu=" + cpu,
			"SET mem=" + mem,
			"SET storageSize=" + storageSize,
			"SET rootSize=" + rootSize,
			"SET goodsNum=" + goodsNum,
			"SET nginxImage=" + nginxImage,
			"SET nginxPort=" + nginxPort,
			"SET nginxProto=" + nginxProto,
			"SET nodeDeleteMode=" + nodeDeleteMode,
			// "CreateCluster",
			// "AddClusterInstances",
			// "AddClusterInstancesFromExistedCvm",
			// "DescribeCluster",
			// "DescribeClusterInstances",
			// "DeleteClusterInstances", // need hardcode work
			// "DeleteCluster",
			// "CreateClusterService",
			// "DescribeClusterService",
			"DescribeClusterServiceInfo",
			// "DeleteClusterService",
		},
		FuncMap: map[string][]string{
			"CreateCluster": []string{"457/9444",
				"Region=$Region",
				"clusterName=$clusterName",
				"cpu=$cpu",
				"mem=$mem",
				"osName=$osName",
				"bandwidth=$bandwidth",
				"bandwidthType=$bandwidthType",
				"vpcId=$vpcId",
				"subnetId=$subnetId",
				"isVpcGateway=$isVpcGateway",
				"storageSize=$storageSize",
				"rootSize=$rootSize",
				"goodsNum=$goodsNum",
				"clusterCIDR=$clusterCIDR",
				"zoneId=$zoneId",
			},
			"AddClusterInstances": []string{"457/9447",
				"Region=$Region",
				"subnetId=$subnetId",
				"isVpcGateway=$isVpcGateway",
				"clusterId=$clusterID",
				"zoneId=$zoneId",
				"cpu=$cpu",
				"mem=$mem",
				"bandwidthType=$bandwidthType",
				"bandwidth=$bandwidth",
				"storageSize=$storageSize",
				"rootSize=$rootSize",
				"goodsNum=$goodsNum",
			},
			"AddClusterInstancesFromExistedCvm": []string{"457/9458",
				"Region=$Region",
				"clusterId=$clusterID",
				"instanceIds.0=ins-8x4kmxlw", // hardcode
				"keyId=$keyId",
			},
			"DescribeCluster": []string{"457/9448"},
			"DescribeClusterInstances": []string{"457/9449",
				"Region=$Region",
				"clusterId=$clusterID",
			},
			"DeleteClusterInstances": []string{"457/9446",
				"Region=$Region",
				"clusterId=$clusterID",
				"instanceIds.0=ins-mpomofr6", // hardcode
				"nodeDeleteMode=$nodeDeleteMode",
			},
			"DeleteCluster": []string{"457/9445"},
			"CreateClusterService": []string{"457/9436",
				"Region=$Region",
				"clusterId=$clusterID",
				"serviceName=$serviceName",
				"replicas=1",
				"containers.0.image=$nginxImage",
				"containers.0.containerName=$serviceName",
				"portMappings.0.containerPort=$nginxPort",
				"portMappings.0.lbPort=$nginxPort",
				"portMappings.0.protocol=$nginxProto",
			},
			"DescribeClusterService": []string{"457/9440",
				"Region=$Region",
				"clusterId=$clusterID",
				"serviceName=$serviceName",
			},
			"DescribeClusterServiceInfo": []string{"457/9441",
				"clusterId=$clusterID",
			},
			"DeleteClusterService": []string{"457/9437",
				"Region=$Region",
				"clusterId=$clusterID",
				"serviceName=$serviceName",
			},
		},
		PkgName: "ccs",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
