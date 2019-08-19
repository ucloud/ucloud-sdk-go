// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

/*
RemoteVPNGatewayDataSet - DescribeRemoteVPNGateway返回参数

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type RemoteVPNGatewayDataSet struct {

	// 活跃的隧道id
	ActiveTunnels []string

	// 创建时间
	CreateTime int

	// 备注
	Remark string

	// 客户网关IP地址
	RemoteVPNGatewayAddr string

	// 客户网关ID
	RemoteVPNGatewayId string

	// 客户网关名称
	RemoteVPNGatewayName string

	// 用户组
	Tag string

	// 活跃的隧道数量
	TunnelCount int
}
