package ubill

/*
ResourceTag - DescribeOrderDetailInfo
*/
type ResourceTag struct {
	//标识名,见下表
	KeyId string

	//标识信息
	Value string
}

/*
KeyId       Description
name        名称(eg. “hello_vm”)
hostname    主机名(eg. “uhost-0zjhzf”)
domain      域名(高防, eg. “www.qq.com”)
private_ip	内网IP
public_ip   外网IP
remark      备注
tag         标签
*/
