// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

/*
IPv6AddressInfo -

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type IPv6AddressInfo struct {

	// 子网内的IPv6地址的公网带宽值。若DefaultBandwidthType为Bandwidth或Traffic，则带宽值为单个地址的带宽值；若为ShareBandwidth，则带宽值为共享带宽的带宽上限。
	Bandwidth string

	// IPv6地址的公网带宽设置。 ShareBandwidth：共享带宽 Bandwidth：带宽计费 Traffic：流量计费 Private：仅内网访问
	BandwidthType string

	// 付费方式, Year：按年付费; Month, 按月付费; Dynamic：按时。
	ChargeType string

	// IPv6地址
	IPv6Address string

	// IPv6地址Id
	IPv6AddressId string

	// 运营商信息如: 国际: International, BGP: BGP
	OperatorName string

	// IPv6地址关联的资源详情。
	RescourceInfo []IPv6ResourceInfo

	// 若BandwidthType为ShareBandwidth，则为共享带宽ID
	ShareBandwidthId string
}
