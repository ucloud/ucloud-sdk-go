package ufs

/*
UFSVolumeInfo - 文件系统信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type UFSVolumeInfo struct {

	// 文件系统名称
	VolumeName string

	// 文件系统ID
	VolumeId string

	// 文件系统是否已添加挂载点
	VolumeStatus string

	// 文件系统已添加挂载点时返回
	MountPoint string

	// 文件系统备注信息
	Remark string

	// 文件系统所属业务组
	Tag string

	// 文件系统创建时间（unix时间戳）
	CreateTime int

	// 文件系统过期时间（unix时间戳）
	ExpiredTime int

	// 文件系统大小，单位GB
	Size int

	// 是否过期
	IsExpired string
}
