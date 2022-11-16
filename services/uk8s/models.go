// Code is generated by ucloud-model, DO NOT EDIT IT.



package uk8s











/*
DiskSet - 节点磁盘信息
*/
type DiskSet struct {
	
	// 备份方案，枚举类型：BASIC_SNAPSHOT,普通快照；DATAARK,方舟。无快照则不返回该字段。
	BackupType string 
	
	// 磁盘长ID
	DiskId string 
	
	// LOCAL_NOMAL| CLOUD_NORMAL| LOCAL_SSD| CLOUD_SSD|EXCLUSIVE_LOCAL_DISK
	DiskType string 
	
	// 磁盘盘符
	Drive string 
	
	// Yes: 加密 No: 非加密
	Encrypted string 
	
	// 当前主机的IOPS值
	IOPS int 
	
	// True| False
	IsBoot string 
	
	// UDisk名字（仅当磁盘是UDisk时返回）
	Name string 
	
	// 磁盘大小，单位: GB
	Size int 
	
	// 磁盘类型。系统盘: Boot，数据盘: Data,网络盘：Udisk
	Type string 
	
}



/*
IPSet - 节点的IP信息
*/
type IPSet struct {
	
	// IP对应的带宽, 单位: Mb (内网IP不显示带宽信息)
	Bandwidth int 
	
	// 是否默认的弹性网卡的信息。true: 是默认弹性网卡；其他值：不是。
	Default string 
	
	// IP地址
	IP string 
	
	// IP资源ID (内网IP无对应的资源ID)
	IPId string 
	
	// 国际: Internation，BGP: Bgp，内网: Private
	Type string 
	
}



/*
UhostInfo - 机器信息
*/
type UhostInfo struct {
	
	// Cpu数量
	CPU int 
	
	// 创建时间
	CreateTime int 
	
	// 节点磁盘信息
	DiskSet []DiskSet 
	
	// 到期时间
	ExpireTime int 
	
	// 节点IP信息
	IPSet []IPSet 
	
	// 内存
	Memory int 
	
	// 主机名称
	Name string 
	
	// 主机ID
	NodeId string 
	
	// 节点类型：uhost表示云主机;uphost表示物理云主机
	NodeType string 
	
	// 镜像信息
	OsName string 
	
	// 主机状态
	State string 
	
	// 所在机房
	Zone string 
	
}



/*
KubeProxy - KubeProxy信息
*/
type KubeProxy struct {
	
	// KubeProxy模式，枚举值为[ipvs,iptables]
	Mode string 
	
}



/*
ImageInfo - UK8S 可用镜像信息
*/
type ImageInfo struct {
	
	// 镜像 Id
	ImageId string 
	
	// 镜像名称
	ImageName string 
	
	// 该镜像是否支持GPU机型，枚举值[true:不支持，false:支持]。
	NotSupportGPU bool 
	
	// 可用区 Id
	ZoneId int 
	
}



/*
K8SNodeCondition - Kubernetes Node Condition
*/
type K8SNodeCondition struct {
	
	// 最后一次上报状态的时间
	LastProbeTime string 
	
	// 最后一次状态转变时间
	LastTransitionTime string 
	
	// 状态变化的描述信息
	Message string 
	
	// 状态变化的原因
	Reason string 
	
	// 状态，False、True
	Status string 
	
	// Condition 类型，如 MemoryPressure、DiskPressure、PIDPressure、Ready
	Type string 
	
}



/*
UHostIPSet - 云主机IP信息
*/
type UHostIPSet struct {
	
	// IP对应的带宽, 单位: Mb (内网IP不显示带宽信息)
	Bandwidth int 
	
	
	// IP地址
	IP string 
	
	// IP资源ID (内网IP无对应的资源ID)
	IPId string 
	
	// Mac地址
	Mac string 
	
	// IP地址对应的子网 ID
	SubnetId string 
	
	// 国际: Internation，BGP: Bgp，内网: Private
	Type string 
	
	// IP地址对应的VPC ID
	VPCId string 
	
}



/*
NodeInfoV2 - UK8S 节点信息
*/
type NodeInfoV2 struct {
	
	// 节点所属伸缩组ID，非伸缩组创建出来的节点，伸缩组ID为Default。
	AsgId string 
	
	// Node节点CPU核数，单位: 个。
	CPU int 
	
	// 节点创建时间
	CreateTime int 
	
	// 节点计费到期时间
	ExpireTime int 
	
	// 节点的GPU颗数。
	GPU int 
	
	// 节点IP信息，详细信息见 UHostIPSet。
	IPSet []UHostIPSet 
	
	// 资源ID，如uhost-xxxx，或uphost-xxxxx。
	InstanceId string 
	
	// 资源名称，初始值等于NodeId，用户可在UHost或UPHost处修改。
	InstanceName string 
	
	// Node节点的资源类型，枚举值为UHost或UPHost。
	InstanceType string 
	
	// kubeproxy信息，详细信息见KubeProxy。
	KubeProxy KubeProxy 
	
	// 机型类别，分别对应Uhost的MachineType或PHost的PHostType。
	MachineType string 
	
	// 内存大小，单位: MB。
	Memory int 
	
	// NodeId，Node在UK8S处的唯一标示，如uk8s-reewqe5-sdasadsda
	NodeId string 
	
	// 加节点时判断是否没有资源，如果返回NORESOURCE则代表没有资源了
	NodeLogInfo string 
	
	// node角色，枚举值为master、node
	NodeRole string 
	
	// Node的状态：枚举值：初始化："Initializing"；启动中："Starting"；运行："Running"；停止中："Stopping"；停止："Stopped"；待删除："ToBeDeleted"；删除中："Deleting"；异常："Error"；安装失败："Install Fail"；
	NodeStatus string 
	
	// Node节点的镜像名称。
	OsName string 
	
	// Node节点的操作系统类别，如Linux或Windows。
	OsType string 
	
	// 是否允许Pod调度到该节点，枚举值为true或false。
	Unschedulable bool 
	
	// Node所在可用区
	Zone string 
	
}



/*
ClusterSet - 集群信息
*/
type ClusterSet struct {
	
	// 集群apiserver地址
	ApiServer string 
	
	// 集群ID
	ClusterId string 
	
	// 创建集群时判断如果为NORESOURCE则为没资源，否则为空
	ClusterLogInfo string 
	
	// 资源名字
	ClusterName string 
	
	// 创建时间
	CreateTime int 
	
	// 集群外部apiserver地址
	ExternalApiServer string 
	
	// 集群版本
	K8sVersion string 
	
	// Master 节点数量
	MasterCount int 
	
	// Node节点数量
	NodeCount int 
	
	// Pod网段
	PodCIDR string 
	
	// 服务网段
	ServiceCIDR string 
	
	// 集群状态，枚举值：初始化："INITIALIZING"；启动中："STARTING"；创建失败："CREATEFAILED"；正常运行："RUNNING"；添加节点："ADDNODE"；删除节点："DELNODE"；删除中："DELETING"；删除失败："DELETEFAILED"；错误："ERROR"；升级插件："UPDATE_PLUGIN"；更新插件信息："UPDATE_PLUGIN_INFO"；异常："ABNORMAL"；升级集群中："UPGRADING"；容器运行时切换："CONVERTING"
	Status string 
	
	// 所属子网
	SubnetId string 
	
	// 所属VPC
	VPCId string 
	
}



/*
NodeGroupSet - 节点池配置
*/
type NodeGroupSet struct {
	
	// 系统盘类型
	BootDiskType string 
	
	// 虚拟CPU核数
	CPU int 
	
	// 付费方式
	ChargeType string 
	
	// 数据盘大小
	DataDiskSize int 
	
	// 数据盘类型
	DataDiskType string 
	
	// GPU卡核心数
	GPU int 
	
	// GPU类型
	GpuType string 
	
	// 镜像ID
	ImageId string 
	
	// 机型
	MachineType string 
	
	// 内存大小
	Mem int 
	
	// cpu平台
	MinimalCpuPlatform string 
	
	// 节点池ID
	NodeGroupId string 
	
	// 节点池名字
	NodeGroupName string 
	
	// 节点id列表
	NodeList []string 
	
	// 业务组
	Tag string 
	
}


