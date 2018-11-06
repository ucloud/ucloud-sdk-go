package uhost

/*
ResourceInfoSet - 资源信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ResourceInfoSet struct {

	// 地域, 参见 [地域和可用区列表](../summary/regionlist.html)
	Region string

	// 可用区，参见 [可用区列表](../summary/regionlist.html)
	Zone string

	// Normal:标准机型, SATA_SSD: SATA SSD机型， BigData: 大数据机型，GPU: GPU型G1(原GPU型)，GPU_G2：GPU型G2
	UHostType string
}
