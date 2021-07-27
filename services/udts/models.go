// Code is generated by ucloud-model, DO NOT EDIT IT.

package udts

/*
CheckResultItem - 预检查结果项
*/
type CheckResultItem struct {

	//
	ErrMessage string

	// 状态
	State string
}

/*
CheckResult - 预检查结果
*/
type CheckResult struct {

	//
	Config CheckResultItem

	//
	Connection CheckResultItem

	//
	Privileges CheckResultItem
}

/*
CheckUDTSTaskResult - 预检查返回的结果
*/
type CheckUDTSTaskResult struct {

	//
	Source CheckResult

	//
	Target CheckResult
}

/*
PolicyData - Policy 数据结构
*/
type PolicyData struct {

	// 列名
	Column string

	// 与上面类型对应的值， 比如“1”， “now()”, "Nash".
	Data string

	// 可选值为 CSVData, Function, Fixed
	Type string
}

/*
SyncData - 增量同步数据
*/
type SyncData struct {

	// GTID
	BinlogGTID string

	// Binlog 文件名， 长度不超过128字符
	BinlogName string

	// Binlog Pos
	BinlogPos int

	// 分配给UDTS task的server ID, 必须在MySQL集群中唯一
	ServerId int
}

/*
TiDBNode - TiDB 结点
*/
type TiDBNode struct {

	// DB 名字， 长度不超过63个字符
	DataBase string

	// 地域，只有当 Host 为 UCloud 用户内网地址的时候需要提供
	DataRegion string

	// 数据库地址，长度不能超过 60个字符
	Host string

	// 数据库密码，长度不起来32个字符
	Password string

	// 数据库端口，端口范围 1-65535
	Port int

	// 子网 ID, 只有当 Host 为 UCloud 用户内网地址并且源目属于不同的地域的时候需要提供。
	SubnetId string

	// 表名， 长度不超过64个字符
	Table string

	// 数据库用户名，长度不能超过 32个字符
	User string

	// VPC 资源ID, 只有当 Host 为 UCloud 用户内网地址的时候需要提供。
	VPCId string
}

/*
RedisNode - Redis节点
*/
type RedisNode struct {

	// 数据库地址，只填写主(master)地址，集群模式下，多个地址用 ; 相连
	Address string

	// 数据库所在的地域。 只有当 Host 为 UCloud 用户内网地址的时候需要提供
	DataRegion string

	// Redis2Redis全量迁移是否使用rump，默认是dump-restore
	IsRump string

	// redis密码
	Password string

	// 子网 ID, 只有当 Host 为 UCloud 用户内网地址并且源目属于不同的地域的时候需要提供
	SubnetId string

	// redis模式
	Type string

	// 数据库所在机器的 VPCId, 只有内网跨域迁移的时候需要提供
	VPCId string
}

/*
CSVNode - CSV 结点
*/
type CSVNode struct {

	// 如果 DupAction 为 ignore或者replace,  并且需要调整列的顺序的时候使用。 以逗号分割的列名字符串。
	Columns string

	// 当加载重复数据的时候所采取的行为，有效值有 ignore - 忽略， replace - 替换， update - 更新。 默认为replace
	DupAction string

	// 数据迁移的时候是否保留原有数据， 默认为 false 不保留
	KeepExistData bool

	// 如果 DupAction 为 update, 并且在插入数据的同时想给一些列赋予特定的值的时候使用。
	SetPolicy []PolicyData

	// 数据路径
	URL string

	// 如果 DupAction 为 update, 并且不想用CSV数据完整替换原有数据的时候使用。
	UpdatePolicy []PolicyData
}

/*
UDWNode - UDW 结点
*/
type UDWNode struct {

	// DB 名字， 长度不超过63个字符
	DataBase string

	// 地域
	DataRegion string

	// 数据库地址，长度不能超过 60个字符
	Host string

	// 数据库密码，长度不起来32个字符
	Password string

	// 数据库端口，端口范围 1-65535
	Port int

	// 子网 ID, 只有当源目属于不同的地域的时候需要提供。
	SubnetId string

	// 数据库用户名，长度不能超过 32个字符
	User string

	// VPC 资源ID, 只有当 Host 为 UCloud 用户内网地址的时候需要提供。
	VPCId string
}

/*
MySQLNode - MySQL 结点
*/
type MySQLNode struct {

	// DB 名字， 长度不超过63个字符
	DataBase string

	// 地域，只有当 Host 为 UCloud 用户内网地址的时候需要提供
	DataRegion string

	// 数据库地址，长度不能超过 60个字符
	Host string

	// 数据库密码，长度不起来32个字符
	Password string

	// 数据库端口，端口范围 1-65535
	Port int

	// 子网 ID, 只有当 Host 为 UCloud 用户内网地址并且源目属于不同的地域的时候需要提供。
	SubnetId string

	// 增量同步数据
	SyncData SyncData

	// 表名， 长度不超过64个字符
	Table string

	// 数据库用户名，长度不能超过 32个字符
	User string

	// VPC 资源ID, 只有当 Host 为 UCloud 用户内网地址的时候需要提供。
	VPCId string
}

/*
UFileNode - UFile 结点
*/
type UFileNode struct {

	// Bucket 名称
	BucketName string

	// 地域
	DataRegion string

	// 用户自定义域名
	DomainName string

	// 前缀，utf-8编码，默认为空字符串
	Prefix string

	// 私钥
	PrivateKey string

	// 公钥
	PublicKey string
}

/*
Source - 源端信息
*/
type Source struct {

	// 设置的最大的速率，单位MB/s，公网/专线(0, 56]，用户网(0, 1024]，不填/超过默认是峰值
	BandwidthLimit int

	// 当 DataType 为csv的时候使用。
	CSVNode CSVNode

	// 源数据类型可以是 mysql, tidb, csv, oracle, udb-mysql。目的数据类型可以是 mysql, tidb, udb-mysql, udw.
	DataType string

	// 是否为专线迁移
	IsDedicatedLine bool

	// 当 DataType 为mysql的时候使用。
	MySQLNode MySQLNode

	// 网络类型
	NWType string

	// 当 DataType 为 redis 的时候使用
	RedisNode RedisNode

	// 当 DataType 为tidb的时候使用。
	TiDBNode TiDBNode

	// 当 DataType 为 udw 的时候使用。
	UDWNode UDWNode

	// 当 DataType 为 ufile 的时候使用。
	UFileNode UFileNode
}

/*
ConfigData - Task 的配置信息
*/
type ConfigData struct {

	// 最大失败重试次数
	MaxRetryCount int

	// 任务名称
	Name string

	// Source
	Source Source

	// Target
	Target Source

	// 任务 ID
	TaskId string

	// 任务类型, full全量, incremental增量，full+incremental全量+增量。
	Type string
}

/*
TaskHistoryItem - 任务历史记录中一条数据对应的 Model
*/
type TaskHistoryItem struct {

	// 任务 ID
	AntID string

	// 任务状态
	AntState string

	// 事件时间，值为 timestamp
	CreateTime int

	// 事件时间，为可读的日期时间
	CreateTimeH string
}

/*
Progress - 进度信息
*/
type Progress struct {

	// 已迁移条目数
	CurCount int

	// 已耗时间（单位秒）
	CurDuration int

	// 完成进度
	Percentage float64

	// 总条目数
	TotalCount int

	// 估算总耗时间（单位秒）
	TotalDuration int
}

/*
StatusData - 动态状态信息
*/
type StatusData struct {

	// 当前失败重试次数
	CurRetryCount int

	// 当Status为Failed时, 显示失败原因
	FailedMessage string

	// 用户设置的最大失败重试次数
	MaxRetryCount int

	// Progress 全量迁移进度信息， 当类型为增量迁移时为空
	Progress Progress

	// 任务状态, 状态有 Created:已创建,Checking:检查中,Dumping:转储中,Loading:加载中,Syncing:同步中,Synced:已同步,Done:完成,Failed:失败,Stopping:停止中,Stopped:停止,RetryPending:重试等待中,Starting:启动中,FailedUnrecoverable:异常,StoppedUnrecoverable:异常,Success:成功,Started:已启动
	Status string

	// Binlog 信息
	Sync SyncData
}

/*
ListDataItem - 返回列表的一个 Task 的信息
*/
type ListDataItem struct {

	// 创建时间
	CreateTime int

	// 当前失败重试次数
	CurRetryCount int

	// 最大失败重试次数
	MaxRetryCount int

	// 任务名称
	Name string

	// 全量迁移进度信息，增量迁移时为空
	Progress Progress

	// 任务状态
	Status string

	// 任务 ID
	TaskId string

	// 任务类型, full全量, incremental增量，full+incremental全量+增量。
	Type string
}
