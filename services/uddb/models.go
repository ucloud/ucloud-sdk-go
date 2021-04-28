// Code is generated by ucloud-model, DO NOT EDIT IT.

package uddb

/*
SlaveInfo - UDDB只读实例信息
*/
type SlaveInfo struct {

	// 对应数据节点的ID
	DataNodeId string

	// 只读实例ID
	Id string

	// 只读实例状态, 状态列表如下: Init: 初始化中 Fail: 安装失败 Starting: 启动中 Running: 系统正常运行中 Shutdown: 关闭中 Shutoff: 已关闭 Deleted: 已删除 Upgrading: 系统升级中
	State string
}

/*
DataNodeInfo - UDDB存储节点和下挂的只读实例信息
*/
type DataNodeInfo struct {

	// 节点的创建时间
	CreateTime string

	// 数据节点的磁盘大小配置. 单位: GB
	DiskSpace string

	// 数据节点ID
	Id string

	// 最近一次数据迁移任务id
	LastTransTaskId string

	// 数据节点的内存配置, 单位：MB
	Memory string

	// 数据节点的只读实例个数.
	SlaveCount string

	// 只读实例信息列表
	SlaveInfos []SlaveInfo

	// 数据分片状态, 状态列表如下: Init: 初始化中 Fail: 安装失败 Starting: 启动中 Running: 系统正常运行中 Shutdown: 关闭中 Shutoff: 已关闭 Deleted: 已删除 Upgrading: 系统升级中
	State string
}

/*
DataSetUDDB - UDDB信息的DataSet
*/
type DataSetUDDB struct {

	// 管理员帐户名，默认root
	AdminUser string

	// 付费类型，可选值如下: Year: 按年付费 Month: 按月付费 Dynamic: 按需付费(单位: 小时) Trial: 免费试用
	ChargeType string

	// UDDB实例创建时间，采用UTC计时时间戳
	CreateTime string

	// UDDB的数据库版本
	DBTypeId string

	// 数据节点个数
	DataNodeCount string

	// 数据节点的磁盘大小配置. 单位: GB
	DataNodeDiskSpace string

	// UDDB实例的数据节点的信息列表
	DataNodeList []DataNodeInfo

	// 数据节点的内存配置, 单位：MB
	DataNodeMemory string

	// 每个数据节点的只读实例个数.
	DataNodeSlaveCount string

	// UDDB实例过期时间，采用UTC计时时间戳
	ExpiredTime string

	// 存储节点的高可用模式， 分为高可用UDB（HA）和普通UDB（Normal），如果不填， 则默认为HA
	InstanceMode string

	// 存储节点和只读实例的磁盘类型。分为：SSD磁盘（SATA_SSD）或普通磁盘(Normal)。 如果不填，则默认为SATA_SSD
	InstanceType string

	// UDDB实例名称
	Name string

	// UDDB实例访问的端口号
	Port string

	// 参考QPS。 免费版： 15000； 畅享版： 30000 - 100000 （根据节点数而定）； 专享版： 节点数 * 10w qps
	RefQps int

	// 各版本下的节点个数。体验版： 固定为2节点； 畅享版：固定为4节点（后续可通过管理API调整）；专享版：物理机台数
	RouterNodeNum int

	// UDDB路由节点的版本。分为三种： Trival(免费版)： 2中间件节点； QPS：1.5W FellFree(标准版)： 固定为4中间件节点，后续将根据业务请求量自动扩展，最多扩展到12个节点，QPS为3w - 10w； EnjoyAlone(物理机版)：专享物理机，节点数让客户可选
	RouterVersion string

	// UDDB状态, 状态列表如下: Init: 初始化中 InitFail: 初始化失败 Starting: 启动中 Running: 系统正常运行中 Abnormal: 系统运行中, 有异常, 还能提供服务 Error: 系统运行中, 但不能正常提供服务 Shutdown: 关闭中 Shutoff: 已关闭 Deleted: 已删除 UpgradingUDDB: 升降级UDDB配置中 UpgradingDataNode: 升降级UDDB节点配置中 ChangingSlaveCount: 改变只读实例数量中 ScalingOutUDDB: 水平扩展中
	State string

	// UDDB实例ID
	UDDBId string

	// UDDB实例访问的虚IP
	VirtualIP string

	// UDDB实例对应的可用区
	Zone string
}

/*
PriceDetailInfo - UDDB实例计费详情
*/
type PriceDetailInfo struct {

	// 存储节点费用
	DataNodePrice float64

	// 只读实例费用
	DataNodeSlavePrice float64

	// 中间件路由节点费用
	MiddlewarePrice float64
}

/*
PriceInfo - UDDB实例计费详情
*/
type PriceInfo struct {

	// 存储节点费用
	DataNodePrice float64

	// 只读实例费用
	DataNodeSlavePrice float64

	// 中间件路由节点费用
	MiddlewarePrice float64
}
