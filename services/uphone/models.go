// Code is generated by ucloud-model, DO NOT EDIT IT.

package uphone

/*
UPhoneInstance -
*/
type UPhoneInstance struct {

	// 虚拟CPU核数。
	CPU int

	// 云手机异步任务回调
	Callback string

	// 计费模式。枚举值为： > 年 Year，按年付费； > Month，按月付费； > Dynamic，按小时预付费; 默认为月付
	ChargeType string

	// 城市Id，eg: cn-shanghai, cn-jinan
	CityId string

	// 城市名称，eg:上海二、济南市
	CityName string

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 磁盘大小，单位: GB
	DiskSize int

	// 到期时间；格式为Unix时间戳
	ExpireTime int

	// 云手机镜像ID，不超过32个字节。
	ImageId string

	// 云手机IP地址
	Ip string

	// IP所属地域Id，eg: hk，th-bkk
	IpRegion string

	// 内存大小。单位MB
	Memory int

	// 云手机镜像系统，如"Android armv8"
	OsType string

	// 刷新率
	Refresh int

	// 备注
	Remark string

	// 分辨率
	Resolution string

	// 云手机所在的服务器ID，不超过32个字节。
	ServerId string

	// 云手机启动图片URL链接
	SplashScreen string

	// 云手机状态<br />* 启动中: STARTING; <br />* 运行中: RUNNING; <br />* 关机中: STOPPING; <br />* 关机: STOPPED <br />* 重启中: REBOOTING; <br />* 重置中: RESETTING; <br />* 启动失败: START_FAILED; <br />* 关机失败: STOP_FAILED; <br />* 重启失败: REBOOT_FAILED; <br />* 重置失败: RESET_FAILED; <br />* 未知状态：UNDEFINED_STATE或""
	State string

	// 业务组名称
	Tag string

	// 云手机的唯一标识，不超过32个字节。
	UPhoneId string

	// 云手机规格名称
	UPhoneModelName string

	// 云手机的名称，不超过65个字符。
	UPhoneName string
}

/*
AppInstance -
*/
type AppInstance struct {

	// 应用的唯一标识ID
	AppId string

	// 应用名称。
	AppName string

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 应用描述。
	Description string

	// 修改时间，格式为Unix时间戳。
	ModifyTime int
}

/*
AppVersionInstance -
*/
type AppVersionInstance struct {

	// 应用的唯一标识ID
	AppId string

	// 应用版本的唯一标识ID
	AppVersionId string

	// 应用版本名。
	AppVersionName string

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 应用版本描述。
	Description string

	// 修改时间，格式为Unix时间戳。
	ModifyTime int

	// 应用包名。
	PackageName string

	// 应用版本相关的Apk文件存放的公网URL地址。
	URL string
}

/*
CityInstance -
*/
type CityInstance struct {

	// 城市别名。如cn-sh2
	CityAlias string

	// 城市Id，eg: cn-shanghai, cn-jinan
	CityId string

	// 城市名称，eg:上海二、济南市
	CityName string

	// 城市类型。枚举值： <br /> * CENTER，中心城市  <br />* EDGE，边缘计算城市
	CityType string

	// 表示该城市资源是否售罄
	IsSoldOut bool
}

/*
UPhoneDetailInstance -
*/
type UPhoneDetailInstance struct {

	// 应用版本实例，每项参数可见数据模型 [AppVersionInstance](#AppVersionInstance)
	AppVersion AppVersionInstance

	// 虚拟CPU核数。
	CPU int

	// 城市Id，eg: cn-shanghai, cn-jinan
	CityId string

	// 城市名称，eg:上海二、济南市
	CityName string

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 磁盘大小，单位: GB
	DiskSize int

	// 云手机镜像ID，不超过32个字节。
	ImageId string

	// 内存大小。单位MB
	Memory int

	// 云手机镜像系统，如"Android armv8"
	OsType string

	// 刷新率
	Refresh int

	// 备注
	Remark string

	// 分辨率
	Resolution string

	// 云手机所在的服务器ID，不超过32个字节。
	ServerId string

	// 云手机状态<br />* 启动中: STARTING; <br />* 运行中: RUNNING; <br />* 关机中: STOPPING; <br />* 关机: STOPPED <br />* 重启中: REBOOTING; <br />* 重置中: RESETTING; <br />* 启动失败: START_FAILED; <br />* 关机失败: STOP_FAILED; <br />* 重启失败: REBOOT_FAILED; <br />* 重置失败: RESET_FAILED; <br />* 未知状态：UNDEFINED或""
	State string

	// 云手机的唯一标识，不超过32个字节。
	UPhoneId string

	// 云手机规格名称
	UPhoneModelName string

	// 云手机的名称，不超过65个字符。
	UPhoneName string
}

/*
UPhoneImageInstance -
*/
type UPhoneImageInstance struct {

	// 云手机镜像所安装的应用版本列表，具体参数见 [AppVersionInstance](#appversioninstance)
	AppVersions []AppVersionInstance

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 云手机镜像描述信息
	Description string

	// 云手机镜像资源Id
	ImageId string

	// 云手机镜像名称
	ImageName string

	// 云手机镜像状态<br />* 制作中: 制作中; <br />* 可用的: CREATING; <br />* 制作失败: CREATE_FAILED; <br />* 上传中: UPLOADING <br />* 上传失败: UPLOAD_FAILED; <br />* 未知状态：UNDEFINED或""
	ImageState string

	// 云手机镜像类型，枚举值：<br />> 用户自制镜像: CUSTOM;  <br />> 标准镜像: BASE;;
	ImageType string

	// 修改时间，格式为Unix时间戳。
	ModifyTime int

	// 云手机镜像系统
	OsType string

	// 云手机镜像创建时所属云手机资源 Id
	UPhoneId string
}

/*
IpRegion - 独立IP地域信息
*/
type IpRegion struct {

	// 已开通地域。参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	Id string

	// 库存状态。枚举值：有库存：Available；无库存：SoldOut
	StockStatus string
}

/*
Task -
*/
type Task struct {

	// 安装/卸载任务相关的应用版本唯一标识ID
	AppVersionId string

	// 任务处理开始时间，格式为Unix时间戳。
	BeginTime int

	// 任务处理结束时间，格式为Unix时间戳。
	EndTime int

	// Task错误信息
	ErrorMsg string

	// 异步任务执行结果
	ExecuteMsg string

	// 任务状态* 等待中：PENDING* 运行中：RUNNING* 成功：SUCCESS* 失败：FAILED
	State string

	// Task的唯一标识
	TaskId string

	// 云手机的唯一标识，云手机相关任务包含此字段。
	UPhoneId string
}

/*
Job -
*/
type Job struct {

	// 安装/卸载Job相关的应用版本唯一标识
	AppVersionId string

	// Job处理开始时间，格式为Unix时间戳。
	BeginTime int

	// Job创建时间，格式为Unix时间戳。
	CreateTime int

	// Job处理结束时间，格式为Unix时间戳。
	EndTime int

	// Job的唯一标识Id
	JobId string

	// Job类型，仅支持INSTALL_APP、UNINSTALL_APP、RUN_ASYNC_COMMAND。
	JobType string

	// 任务状态* 等待中：PENDING* 运行中：RUNNING* 成功：SUCCESS* 部分成功：PARTIAL_SUCCESS* 失败：FAILED
	State string

	// Task信息，详情见数据结构Task（如果State为PENDING，此字段为空）
	Tasks []Task

	// 此Job相关的云手机唯一标识
	UPhoneIds []string
}

/*
UPhoneModelInstance -
*/
type UPhoneModelInstance struct {

	// 虚拟CPU核数。
	CPU int

	// 型号描述信息
	Description string

	// 磁盘大小，单位: GB
	DiskSize int

	// DPI
	Dpi int

	// 内存大小。单位MB
	Memory int

	// 刷新率
	Refresh int

	// 分辨率
	Resolution string

	// UPhoneModel名称
	UPhoneModelName string
}

/*
ServerDiskSet -
*/
type ServerDiskSet struct {

	// 磁盘类型。请参考磁盘类型。
	DiskType string

	// 是否是系统盘。枚举值：> True，是系统盘> False，是数据盘（默认）。Disks数组中有且只能有一块盘是系统盘。
	IsBoot bool

	// 磁盘大小，单位: GB
	Size int
}

/*
UPhoneSpec -
*/
type UPhoneSpec struct {

	// 手机开数，即该服务器规格能生成对应手机规格的云手机个数
	UPhoneCount int

	// 手机规格名
	UPhoneModelName string
}

/*
ServerModelInstance -
*/
type ServerModelInstance struct {

	// 虚拟CPU核数。可选参数：1-64（具体机型与CPU的对应关系参照控制台）。
	CPU int

	// 磁盘信息见 UPhoneDiskSet
	DiskSet []ServerDiskSet

	// GPU个数
	GPU int

	// GPU类型
	GPUType string

	// 内存大小。单位：MB。
	Memory int

	// ServerModel名称
	ServerModelName string

	// 表示该机型是否上线，用于前端判断是否开放给用户。枚举值：>AVAILABLE，开放>UNAVAILABLE, 不开放
	ServerModelState string

	// 【数组】手机说明，包含该服务器规格所能创建的手机规格名及对应手机开数。每项参数可见数据模型 [UPhoneSpec](#UPhoneSpec)
	UPhoneSpecs []UPhoneSpec
}

/*
IpSet -
*/
type IpSet struct {

	// IP地址
	Ip string

	// ipv4或者ipv6
	IpMode string

	// 共有或私有
	IpType string

	// 运营商
	Isp string
}

/*
ServerInstance -
*/
type ServerInstance struct {

	// 计费模式。枚举值为：<br />  *Year，按年付费；<br />*  Month，按月付费；<br />默认为月付
	ChargeType string

	// 城市Id，eg: cn-shanghai, cn-jinan
	CityId string

	// 城市名称，eg:上海二、济南市
	CityName string

	// 创建时间，格式为Unix时间戳。
	CreateTime int

	// 到期时间，格式为Unix时间戳。
	ExpireTime int

	// 服务器IP信息
	IpSet []IpSet

	// 修改时间，格式为Unix时间戳。
	ModifyTime int

	// 云服务器备注。
	Remark string

	// 云手机服务器的唯一标识。
	ServerId string

	// 云服务器规格。
	ServerModel ServerModelInstance

	// 云手机服务器名称。
	ServerName string

	// 实例状态，枚举值： * 初始化: Initializing; * 云手机创建中: UPhoneCreating;* 运行中: Running; * 删除中: Deleting; * 创建失败: CreateFailed * 未知(空字符串，获取状态超时或出错)：""
	State string

	// 云手机开数。
	UPhoneCount int

	// 云手机规格名称。
	UPhoneModelName string
}

/*
StockInfo - model的可用量信息
*/
type StockInfo struct {

	// ServerModel名称
	ModelName string

	// 资源余量
	StockCount int
}

/*
UPhoneAllowance - 云手机余量结构体
*/
type UPhoneAllowance struct {

	// 可创建云手机个数
	Allowance int

	// 枚举值，云手机型号名称，取值：UPhone X，UPhone Plus，UPhone Pro
	ModelName string
}

/*
UPhonePriceSet - 云手机价格列表
*/
type UPhonePriceSet struct {

	// 计费类型，枚举值：Year，Month，Dynamic
	ChargeType string

	// 产品列表价
	ListPrice float64

	// 限时优惠的折前原价（即列表价乘以商务折扣后的单价）
	OriginalPrice float64

	// 价格，单位: 元，保留小数点后两位有效数字
	Price float64
}

/*
UPhoneServerPriceSet - 价格列表
*/
type UPhoneServerPriceSet struct {

	// 计费类型，枚举值：Year，Month, Dynamic
	ChargeType string

	// 产品列表价
	ListPrice float64

	// 限时优惠的折前原价（即列表价乘以商务折扣后的单价）
	OriginalPrice float64

	// 价格，单位: 元，保留小数点后两位有效数字
	Price float64
}

/*
UPhoneCommandResult -
*/
type UPhoneCommandResult struct {

	// 同步shell命令的执行结果。
	ExecuteMsg string

	// 云手机实例的资源ID。
	UPhoneId string
}

/*
Results -
*/
type Results struct {

	//
	ExecuteMsg string

	//
	UPhoneId string
}
