package vpc

/*
VPCSubnetInfoSet - DescribeSubnet

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type VPCSubnetInfoSet struct {

	// 创建时间
	CreateTime int

	// 子网网关
	Gateway string

	// 是否有natgw
	HasNATGW bool

	// 子网关联的IPv6网段
	IPv6Network string

	// 子网掩码
	Netmask string

	// 子网关联的IPv6网段所属运营商
	OperatorName string

	// 备注
	Remark string

	// 路由表Id
	RouteTableId string

	// 子网网段
	Subnet string

	// 子网Id
	SubnetId string

	// 子网名称
	SubnetName string

	// 子网类型
	SubnetType int

	// 业务组
	Tag string

	// VPCId
	VPCId string

	// VPC名称
	VPCName string

	// 可用区名称
	Zone string

	// 虚拟路由 id
	VRouterId string

	// 名称
	Name string
}
