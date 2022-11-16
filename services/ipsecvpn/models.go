// Code is generated by ucloud-model, DO NOT EDIT IT.



package ipsecvpn











/*
RemoteVPNGatewayDataSet - DescribeRemoteVPNGateway返回参数
*/
type RemoteVPNGatewayDataSet struct {
	
	// 
	ActiveTunnels []string `deprecated:"true"`
	
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



/*
VPNGatewayDataSet - DescribeVPNGateway返回参数
*/
type VPNGatewayDataSet struct {
	
	// 是否自动续费
	AutoRenew string 
	
	// 付费类型
	ChargeType string 
	
	// 创建时间
	CreateTime int 
	
	// 绑定EIP的IP地址
	EIP string 
	
	// EIPID
	EIPId string 
	
	// EIP类型
	EIPType string 
	
	// 到期时间
	ExpireTime int 
	
	// 网关类型
	Grade string 
	
	// 网关备注
	Remark string 
	
	// 网关业务组
	Tag string 
	
	// 所属VPCId
	VPCId string 
	
	// 所属VPC名字
	VPCName string 
	
	// 网关Id
	VPNGatewayId string 
	
	// 网关名字
	VPNGatewayName string 
	
}



/*
IKEData - IKE信息
*/
type IKEData struct {
	
	// IKE认证算法
	IKEAuthenticationAlgorithm string 
	
	// IKEDH组
	IKEDhGroup string 
	
	// IKE加密算法
	IKEEncryptionAlgorithm string 
	
	// IKEv1协商模式
	IKEExchangeMode string 
	
	// IKE本地ID标识
	IKELocalId string 
	
	// IKE预共享秘钥
	IKEPreSharedKey string 
	
	// IKE对端ID标识
	IKERemoteId string 
	
	// IKE秘钥生存时间
	IKESALifetime string 
	
	// IKE版本
	IKEVersion string 
	
}



/*
IPSecData - IPSec参数
*/
type IPSecData struct {
	
	// IPSec通道中使用的认证算法
	IPSecAuthenticationAlgorithm string 
	
	// IPSec通道中使用的加密算法
	IPSecEncryptionAlgorithm string 
	
	// 指定VPN连接的本地子网，用逗号分隔
	IPSecLocalSubnetIds []string 
	
	// 是否开启PFS功能,Disable表示关闭，数字表示DH组
	IPSecPFSDhGroup string 
	
	// 使用的安全协议，ESP或AH
	IPSecProtocol string 
	
	// 指定VPN连接的客户网段，用逗号分隔
	IPSecRemoteSubnets []string 
	
	// IPSec中SA的生存时间
	IPSecSALifetime string 
	
	// IPSec中SA的生存时间（以字节计）
	IPSecSALifetimeBytes string 
	
}



/*
VPNTunnelDataSet - DescribeVPNTunnel信息
*/
type VPNTunnelDataSet struct {
	
	// 创建时间
	CreateTime int 
	
	// IKE参数
	IKEData IKEData 
	
	// IPSec参数
	IPSecData IPSecData 
	
	// 备注
	Remark string 
	
	// 对端网关Id
	RemoteVPNGatewayId string 
	
	// 对端网关名字
	RemoteVPNGatewayName string 
	
	// 用户组
	Tag string 
	
	// 所属VPCId
	VPCId string 
	
	// 所属VOC名字
	VPCName string 
	
	// 所属VPN网关id
	VPNGatewayId string 
	
	// VPN网关名字
	VPNGatewayName string 
	
	// 隧道id
	VPNTunnelId string 
	
	// 隧道名称
	VPNTunnelName string 
	
}



/*
VPNGatewayPriceSet - VPN网关的价格信息
*/
type VPNGatewayPriceSet struct {
	
	// VPN网关付费方式
	ChargeType string 
	
	// VPN网关价格, 单位"元"
	Price float64 
	
	// 资源有效期, 以Unix Timestamp表示
	PurchaseValue int 
	
}


