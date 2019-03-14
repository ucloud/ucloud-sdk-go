package umem

/*
UMemDataSet - DescribeUMem

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type UMemSlaveDataSet struct {

	// 资源ID
	GroupId string

	VirtualIP string

	Port int

	// 资源名称
	Name string

	// 创建时间
	CreateTime int

	Role string

	// 到期时间
	ExpireTime int

	// 容量单位GB
	Size int

	// 使用量单位MB
	UsedSize int

	// 实例状态                                  Starting                  // 创建中       Creating                  // 初始化中     CreateFail                // 创建失败     Fail                      // 创建失败     Deleting                  // 删除中       DeleteFail                // 删除失败     Running                   // 运行         Resizing                  // 容量调整中   ResizeFail                // 容量调整失败 Configing                 // 配置中       ConfigFail                // 配置失败Restarting                // 重启中
	State string

	// 计费模式，Year, Month, Dynamic, Trial
	ChargeType string

	// 业务组名称
	Tag string

	// 空间类型:single(无热备),double(热备)
	ResourceType string

	// 节点的配置ID
	ConfigId string

	// Redis版本信息
	Version string

	// 实例所在可用区，或者master redis所在可用区，参见 [可用区列表](../summary/regionlist.html)
	Zone string
}
