// Code is generated by ucloud-model, DO NOT EDIT IT.

package uhost

/*
KeyPair - 密钥对信息
*/
type KeyPair struct {

	// 密钥对的创建时间，格式为Unix Timestamp。
	CreateTime int

	// 密钥对指纹。md5(ProjectId|KeyPairId|PublicKey)
	KeyPairFingerPrint string

	// 密钥对ID。
	KeyPairId string

	// 密钥对名称。 长度为1~63个英文或中文字符。
	KeyPairName string

	// 密钥对的私钥内容。只有创建接口才会返回。
	PrivateKeyBody string

	// 项目ID。
	ProjectId string
}

/*
FeatureModes - 可以支持的模式类别
*/
type FeatureModes struct {

	// 这个特性必须是列出来的CPU平台及以上的CPU才支持
	MinimalCpuPlatform []string

	// 模式|特性名称
	Name string

	// 为镜像上支持这个特性的标签。例如DescribeImage返回的字段Features包含HotPlug，说明该镜像支持热升级。
	RelatedToImageFeature []string
}

/*
DataDiskInfo - 数据盘信息
*/
type DataDiskInfo struct {

	// 数据盘可支持的服务
	Features []string

	// MaximalSize为磁盘最大值
	MaximalSize int

	// 磁盘最小值，如果没有该字段，最小值取基础镜像Size值即可（linux为20G，windows为40G）。
	MinimalSize int

	// 数据盘类别，包含普通云盘|CLOUD_NORMAL、SSD云盘|CLOUD_SSD和RSSD云盘|CLOUD_RSSD。普通本地盘只包含普通本地盘|LOCAL_NORMAL一种。SSD本地盘只包含SSD本地盘|LOCAL_SSD一种。
	Name string
}

/*
BootDiskInfo - 系统盘信息
*/
type BootDiskInfo struct {

	// 磁盘可支持的服务
	Features []string

	// 系统盘是否允许扩容，如果是本地盘，则不允许扩容，InstantResize为false。
	InstantResize bool

	// MaximalSize为磁盘最大值
	MaximalSize int

	// 系统盘类别，包含普通云盘|CLOUD_NORMAL、SSD云盘|CLOUD_SSD和RSSD云盘|CLOUD_RSSD。普通本地盘只包含普通本地盘|LOCAL_NORMAL一种。SSD本地盘只包含SSD本地盘|LOCAL_SSD一种。
	Name string
}

/*
Collection - CPU和内存可支持的规格
*/
type Collection struct {

	// CPU规格
	Cpu int

	// 内存规格
	Memory []int

	// CPU和内存规格只能在列出来的CPU平台支持
	MinimalCpuPlatform []string
}

/*
GraphicsMemory - GPU的显存指标
*/
type GraphicsMemory struct {

	// 交互展示参数，可忽略
	Rate int

	// 值，单位是GB
	Value int
}

/*
Features - 虚机可支持的特性
*/
type Features struct {

	// 可以提供的模式类别
	Modes []FeatureModes

	// 可支持的特性名称。目前支持的特性网络增强|NetCapability、热升级|Hotplug
	Name string
}

/*
Disks - 磁盘信息
*/
type Disks struct {

	// 系统盘信息
	BootDisk []BootDiskInfo

	// 数据盘信息
	DataDisk []DataDiskInfo

	// 磁盘介质类别信息，磁盘主要分类如下：云盘|cloudDisk、普通本地盘|normalLocalDisk和SSD本地盘|ssdLocalDisk。
	Name string
}

/*
Performance - GPU的性能指标
*/
type Performance struct {

	// 交互展示参数，可忽略
	Rate int

	// 值，单位是TFlops
	Value float64
}

/*
MachineSizes - GPU、CPU和内存信息
*/
type MachineSizes struct {

	// CPU和内存可支持的规格
	Collection []Collection

	// Gpu为GPU可支持的规格即GPU颗数，非GPU机型，Gpu为0
	Gpu int
}

/*
CpuPlatforms - CPU平台信息
*/
type CpuPlatforms struct {

	// 返回AMD的CPU平台信息，例如：AMD: ['Amd/Epyc2']
	Amd []string

	// 返回Arm的CPU平台信息，例如：Ampere: ['Ampere/Altra']
	Ampere []string

	// 返回Intel的CPU平台信息，例如：Intel: ['Intel/CascadeLake','Intel/CascadelakeR','Intel/IceLake']
	Intel []string
}

/*
AvailableInstanceTypes - https://ushare.ucloudadmin.com/pages/viewpage.action?pageId=104662646
*/
type AvailableInstanceTypes struct {

	// 支持的CPU平台，并且按照Intel、AMD和Ampere分类返回
	CpuPlatforms CpuPlatforms

	// 磁盘信息。磁盘主要分类如下：云盘|cloudDisk、普通本地盘|normalLocalDisk和SSD本地盘|ssdLocalDisk。其中云盘主要包含普通云盘|CLOUD_NORMAL、SSD云盘|CLOUD_SSD和RSSD云盘|CLOUD_RSSD。普通本地盘只包含普通本地盘|LOCAL_NORMAL一种。SSD本地盘只包含SSD本地盘|LOCAL_SSD一种。MinimalSize为磁盘最小值，如果没有该字段，最小值取基础镜像Size值即可（linux为20G，windows为40G）。MaximalSize为磁盘最大值。InstantResize表示系统盘是否允许扩容，如果是本地盘，则不允许扩容，InstantResize为false。Features为磁盘可支持的服务：数据方舟|DATAARK，快照服务|SNAPSHOT，加密盘|Encrypted。
	Disks []Disks

	// 虚机可支持的特性。目前支持的特性网络增强|NetCapability、热升级|Hotplug。网络增强分为关闭|Normal、网络增强1.0|Super和网络增强2.0|Ultra。Name为可支持的特性名称，Modes为可以提供的模式类别等，RelatedToImageFeature为镜像上支持这个特性的标签。例如DescribeImage返回的字段Features包含HotPlug，说明该镜像支持热升级。MinimalCpuPlatform表示这个特性必须是列出来的CPU平台及以上的CPU才支持。
	Features []Features

	// GPU的显存指标，value为值，单位是GB。
	GraphicsMemory GraphicsMemory

	// 区分是否是GPU机型：GPU机型|GPU，非GPU机型|Normal。
	MachineClass string

	// GPU、CPU和内存信息。Gpu为GPU可支持的规格，Cpu和Memory分别为CPU和内存可支持的规格。如果非GPU机型，GPU为0。MinimalCpuPlatform代表含义这个CPU和内存规格只能在列出来的CPU平台支持。
	MachineSizes []MachineSizes

	// 机型名称：快杰O型|O 、快杰共享型|OM 、快杰内存型|OMEM 、 快杰PRO型|OPRO、通用N型|N、高主频C型|C和GPU G型|G等
	Name string

	// GPU的性能指标，value为值，单位是TFlops。
	Performance Performance

	// 机型状态：可售|Normal 、 公测|Beta、售罄|Soldout、隐藏|Hidden
	Status string

	// 可用区信息
	Zone string
}

/*
UHostImageSet - DescribeImage
*/
type UHostImageSet struct {

	// 创建时间，格式为Unix时间戳
	CreateTime int

	// 特殊状态标识， 目前包含NetEnhnced（网络增强1.0）, NetEnhanced_Ultra]（网络增强2.0）, HotPlug(热升级), CloudInit, IPv6
	Features []string

	// 行业镜像类型（仅行业镜像将返回这个值）
	FuncType string

	// 镜像描述
	ImageDescription string

	// 镜像ID
	ImageId string

	// 镜像名称
	ImageName string

	// 镜像大小
	ImageSize int

	// 镜像类型 标准镜像：Base， 行业镜像：Business，自定义镜像：Custom
	ImageType string

	// 集成软件名称（仅行业镜像将返回这个值）
	IntegratedSoftware string

	// 介绍链接（仅行业镜像将返回这个值）
	Links string

	// 默认值为空'''。当CentOS 7.3/7.4/7.5等镜像会标记为“Broadwell”
	MinimalCPU string

	// 操作系统名称
	OsName string

	// 操作系统类型：Linux，Windows
	OsType string

	// 镜像状态， 可用：Available，制作中：Making， 不可用：Unavailable
	State string

	// 供应商（仅行业镜像将返回这个值）
	Vendor string

	// 可用区，参见 [可用区列表](../summary/regionlist.html)
	Zone string
}

/*
SpreadInfo - 每个可用区中硬件隔离组信息
*/
type SpreadInfo struct {

	// 当前地域所有可用区中硬件隔离组中云主机的数量，不超过7。
	UHostCount int

	// 可用区信息
	Zone string
}

/*
IsolationGroup - 硬件隔离组信息
*/
type IsolationGroup struct {

	// 硬件隔离组id
	GroupId string

	// 硬件隔离组名称
	GroupName string

	// 备注
	Remark string

	// 每个可用区中的机器数量。参见数据结构SpreadInfo。
	SpreadInfoSet []SpreadInfo
}

/*
UHostDiskSet - DescribeUHostInstance
*/
type UHostDiskSet struct {

	// 备份方案。若开通了数据方舟，则为DATAARK
	BackupType string

	// 磁盘ID
	DiskId string

	// 磁盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。
	DiskType string

	// 磁盘盘符
	Drive string

	// "true": 加密盘 "false"：非加密盘
	Encrypted string

	// 是否是系统盘。枚举值：\\ > True，是系统盘 \\ > False，是数据盘（默认）。Disks数组中有且只能有一块盘是系统盘。
	IsBoot string

	// UDisk名字（仅当磁盘是UDisk时返回）
	Name string

	// 磁盘大小，单位: GB
	Size int

	// 【建议不再使用】磁盘类型。系统盘: Boot，数据盘: Data,网络盘：Udisk
	Type string
}

/*
UHostKeyPair - 主机密钥信息
*/
type UHostKeyPair struct {

	// 密钥对ID
	KeyPairId string

	// 主机密钥对状态，Normal 正常，Deleted 删除
	KeyPairState string
}

/*
UHostIPSet - DescribeUHostInstance
*/
type UHostIPSet struct {

	// IP对应的带宽, 单位: Mb  (内网IP不显示带宽信息)
	Bandwidth int

	// 内网 Private 类型下，表示是否为默认网卡。true: 是默认网卡；其他值：不是。
	Default string

	// IP地址
	IP string

	// 外网IP资源ID 。(内网IP无对应的资源ID)
	IPId string

	// IPv4/IPv6；
	IPMode string

	// 内网 Private 类型下，当前网卡的Mac。
	Mac string

	// 弹性网卡为默认网卡时，返回对应的 ID 值
	NetworkInterfaceId string

	// IP地址对应的子网 ID。（北京一不支持，字段返回为空）
	SubnetId string

	// 国际: Internation，BGP: Bgp，内网: Private
	Type string

	// IP地址对应的VPC ID。（北京一不支持，字段返回为空）
	VPCId string

	// 当前EIP的权重。权重最大的为当前的出口IP。
	Weight int
}

/*
UHostInstanceSet - DescribeUHostInstance
*/
type UHostInstanceSet struct {

	// 是否自动续费，自动续费：“Yes”，不自动续费：“No”
	AutoRenew string

	// 基础镜像ID（指当前自定义镜像的来源镜像）
	BasicImageId string

	// 基础镜像名称（指当前自定义镜像的来源镜像）
	BasicImageName string

	// 系统盘状态 Normal表示初始化完成；Initializing表示在初始化。仍在初始化的系统盘无法制作镜像。
	BootDiskState string

	// 虚拟CPU核数，单位: 个
	CPU int

	// 计费模式，枚举值为： Year，按年付费； Month，按月付费； Dynamic，按需付费（需开启权限）；Preemptive 为抢占式实例；
	ChargeType string

	// true，支持cloutinit方式初始化；false,不支持
	CloudInitFeature bool

	// 云主机CPU平台。参考[[api:uhost-api:uhost_type#主机概念20版本|云主机机型说明]]。
	CpuPlatform string

	// 创建时间，格式为Unix时间戳
	CreateTime int

	//
	DeleteTime int `deprecated:"true"`

	// 磁盘信息见 UHostDiskSet
	DiskSet []UHostDiskSet

	// 到期时间，格式为Unix时间戳
	ExpireTime int

	// GPU个数
	GPU int

	// 【建议不再使用】主机系列：N2，表示系列2；N1，表示系列1
	HostType string

	// true: 开启热升级； false，未开启热升级
	HotplugFeature bool

	// true: 开启 hpc 系列功能；false: 未开启
	HpcFeature bool

	// 详细信息见 UHostIPSet
	IPSet []UHostIPSet

	//
	IPs []string `deprecated:"true"`

	// true:有ipv6特性；false，没有ipv6特性
	IPv6Feature bool

	// 【建议不再使用】主机的系统盘ID。
	ImageId string

	// 隔离组id，不在隔离组则返回""
	IsolationGroup string

	// 密钥信息见 UHostKeyPair
	KeyPair UHostKeyPair

	// 主机的生命周期类型。目前仅支持Normal：普通；
	LifeCycle string

	// 云主机机型（新）。参考[[api:uhost-api:uhost_type#主机概念20版本|云主机机型说明]]。
	MachineType string

	// 内存大小，单位: MB
	Memory int

	// UHost实例名称
	Name string

	// 网络增强。Normal: 无；Super： 网络增强1.0； Ultra: 网络增强2.0
	NetCapability string

	// 【建议不再使用】网络状态。 连接：Connected， 断开：NotConnected
	NetworkState string

	// 创建主机的最初来源镜像的操作系统名称（若直接通过基础镜像创建，此处返回和BasicImageName一致）
	OsName string

	// 操作系统类别。返回"Linux"或者"Windows"
	OsType string

	// RDMA集群id，仅快杰云主机返回该值；其他类型云主机返回""。当云主机的此值与RSSD云盘的RdmaClusterId相同时，RSSD可以挂载到这台云主机。
	RdmaClusterId string

	// 备注
	Remark string

	// 仅抢占式实例返回，LowSpeed为低速模式，PowerOff为关机模式
	RestrictMode string

	// 实例状态，枚举值：\\ >初始化: Initializing; \\ >启动中: Starting; \\> 运行中: Running; \\> 关机中: Stopping; \\ >关机: Stopped \\ >安装失败: Install Fail; \\ >重启中: Rebooting; \\ > 未知(空字符串，获取状态超时或出错)：""
	State string

	// 【建议不再使用】主机磁盘类型。 枚举值为：\\ > LocalDisk，本地磁盘; \\ > UDisk 云盘。\\只要有一块磁盘为本地盘，即返回LocalDisk。
	StorageType string

	// 【建议不再使用】仅北京A的云主机会返回此字段。基础网络模式：Default；子网模式：Private
	SubnetType string

	// 业务组名称
	Tag string

	// 【建议不再使用】数据方舟模式。枚举值：\\ > Yes: 开启方舟； \\ > no，未开启方舟
	TimemachineFeature string

	// 总的数据盘存储空间。
	TotalDiskSpace int

	// UHost实例ID
	UHostId string

	// 【建议不再使用】云主机机型（旧）。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	UHostType string

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone string
}

/*
KeyPairDesc - 密钥对信息，不包含私钥内容。
*/
type KeyPairDesc struct {

	// 密钥对的创建时间，格式为Unix Timestamp。
	CreateTime int

	// 密钥对指纹。md5(ProjectId|KeyPairId|PublicKey)
	KeyPairFingerPrint string

	// 密钥对ID。
	KeyPairId string

	// 密钥对名称。 长度为1~63个英文或中文字符。
	KeyPairName string

	// 项目ID。
	ProjectId string
}

/*
UHostTagSet - DescribeUHostTags
*/
type UHostTagSet struct {

	// 业务组名称
	Tag string

	// 该业务组中包含的主机个数
	TotalCount int

	// 可用区
	Zone string
}

/*
PriceDetail - 价格详细信息
*/
type PriceDetail struct {

	// 快照价格
	Snapshot float64

	// 云盘价格
	UDisk float64

	// 主机价格
	UHost float64

	// 数据卷价格
	Volume float64
}

/*
UHostPriceSet - 主机价格
*/
type UHostPriceSet struct {

	// 计费类型。Year，Month，Dynamic
	ChargeType string

	// 产品列表价。
	ListPrice float64

	// 列表价详细信息（只有询价接口返回）。
	ListPriceDetail PriceDetail

	// 限时优惠的折前原价（即列表价乘以商务折扣后的单价）。
	OriginalPrice float64

	// 原价详细信息（只有询价接口返回）。
	OriginalPriceDetail PriceDetail

	// 价格，单位: 元，保留小数点后两位有效数字
	Price float64

	// 价格详细信息（只有询价接口返回）。
	PriceDetail PriceDetail
}
