package ufs

/*
MountPointInfo - 挂载点信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type MountPointInfo struct {

	// 挂载点名称
	MountPointName string

	// Vpc ID
	VpcId string

	// Subnet ID
	SubnetId string

	// 挂载点IP
	MountPointIp string

	// 文件系统创建时间（unix时间戳）
	CreateTime int

	// Subnet ID + 网段的形式，方便前端展示
	SubnetDescription string
}
