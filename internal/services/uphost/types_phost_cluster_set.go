package uphost

/*
PHostClusterSet - 物理云主机集群库存信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type PHostClusterSet struct {

	// 集群名。枚举值：千兆网络集群：1G；万兆网络集群：10G
	Name string

	// 库存状态。枚举值：有库存：Available；无库存：SoldOut
	StockStatus string
}
