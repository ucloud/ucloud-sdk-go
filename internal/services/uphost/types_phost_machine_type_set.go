package uphost

/*
PHostMachineTypeSet - 物理云主机机型列表

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type PHostMachineTypeSet struct {

	// 机型名
	Name string

	// 物理云主机机型别名，全网唯一。
	Type string

	// CPU信息
	CPU PHostCPUSet

	// 内存大小，单位GB
	Memory int

	// 磁盘信息
	Disks []PHostDiskSet

	// 其他组件信息
	Components PHostComponentSet

	// 集群库存信息
	Clusters []PHostClusterSet

	// 是否支持Raid。枚举值：支持：YES；不支持：NO
	RaidSupported string
}
