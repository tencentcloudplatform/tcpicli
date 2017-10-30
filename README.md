# tcpicli
Tencent Cloud Platform International Command Line Interface

## profile
Select profile to use API keys. Place configuration in ~/.tcpicli/config

Example:

```
[chinatest]
secretid  = ZBcr8RCC0d5LPJ6j9NTo4aF
secretkey = e7OVRFuOiy2E
```

Usage:

`tcpicli -vv profile switch chinatest`

Then issue your API call against the account that's named in the [brakets].

Help: 

`tcpicli`:

```
NAME:
   tcpicli - tencent cloud platform cli tool

USAGE:
   tcpicli [global options] command [command options] [arguments...]

VERSION:
   1.0.0.

COMMANDS:
     profile  profile
     do       do <service> <action> <args1=value1> [args2=value2] ...
     vpc      vpc function
     cdn      cdn function
     cvm      cvm function
     img      img function
     ccs      ccs function
     cmq      cmq function
     dfw      Cloud FireWall
     account  account function
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --vv           verbose
   -f value       filter output
                    Example:
                    tcpicli -f="The status is: {{{.code}}" ccs DescribeCluster
                    tcpicli -f="clusterCIDR: {{range .data.clusters}}{{.clusterCIDR}}{{end}}" ccs DescribeCluster
                    tcpicli -f="/path/to/file ccs DescribeCluster"
   --help, -h     show help
   --version, -v  print the version
```
