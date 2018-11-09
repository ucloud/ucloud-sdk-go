package udataark

/*
DiskUDataArkStatus - 磁盘方舟状态

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type DiskUDataArkStatus struct {

	// 磁盘Id
	VdiskId string

	// 磁盘名
	VdiskName string

	// 磁盘描述
	VdiskDesc string

	// 是否系统盘
	SysVdisk int

	// 最近一次恢复时间
	RecoverTick int

	// 当前时间
	CurrentTick int

	// "normal" 正常状态"in_init" 初始化中
	UtmStatus string

	// "inprogress" 恢复进行中"succ" 恢复成功 "error" 恢复失败 "cancel" 恢复取消
	RecoverStatus string

	// 备份大小
	UtmSize int
}
