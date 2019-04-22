package ufs

/*
UFSVolumeInfo2 - 文件系统信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type UFSVolumeInfo2 struct {

	// 文件系统名称
	VolumeName string

	// 文件系统ID
	VolumeId string

	// 当前文件系统已创建的挂载点数目
	TotalMountPointNum int

	// 文件系统允许创建的最大挂载点数目
	MaxMountPointNum int

	// 文件系统存储类型，枚举值，Basic表示容量型，Advanced表示性能型
	StorageType string

	// 文件系统协议，枚举值，NFSv3表示NFS V3协议，NFSv4表示NFS V4协议
	ProtocolType string

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
