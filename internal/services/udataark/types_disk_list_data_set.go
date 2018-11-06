package udataark

/*
DiskListDataSet - 备份链记录

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type DiskListDataSet struct {

	// 快照ID
	SnapshotId string

	// 快照名
	SnapshotName string

	// 快照类型，"minor" 小时级别增量备份 ，"major"  天级别增量备份  ，"base" base备份 ，"user" 用户手动备份
	SnapshotType string

	// 快照描述
	Comment string

	// 快照创建时间
	CreateTime int

	// 快照时间
	SnapshotTime int

	// 快照大小
	Size int

	// 是否用户手动快照，1 表示用户手动快照, 0表示非用户手动快照
	User int

	// 快照对应的方舟磁盘版本
	Version int

	// 快照状态，"ready" 表示等待merge ，"finish" 表示已经merge完成
	Status string
}
