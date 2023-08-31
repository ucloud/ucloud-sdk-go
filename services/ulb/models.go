// Code is generated by ucloud-model, DO NOT EDIT IT.

package ulb

/*
BackendSet -
*/
type BackendSet struct {

	//
	BackendId string

	//
	ResourceId string
}

/*
UlbPolicyBackendSet - DescribePolicyGroup
*/
type UlbPolicyBackendSet struct {

	// 后端资源实例的ID
	BackendId string

	// 后端资源实例的服务端口
	Port int

	// 后端资源实例的内网IP
	PrivateIP string
}

/*
UlbPolicySet - DescribePolicyGroup
*/
type UlbPolicySet struct {

	// 内容转发策略组ID所应用的后端资源列表，具体结构见 UlbPolicyBackendSet
	BackendSet []UlbPolicyBackendSet

	// 内容转发匹配字段
	Match string

	// 内容转发策略组ID
	PolicyId string

	// 内容转发匹配字段的类型，当前只支持按域名转发。枚举值为： Domain，按域名转发
	Type string

	// 内容转发策略组ID应用的VServer实例的ID
	VServerId string
}

/*
UlbPolicyGroupSet - DescribePolicyGroup
*/
type UlbPolicyGroupSet struct {

	// 内容转发策略组ID
	GroupId string

	// 内容转发策略组名称
	GroupName string

	// 内容转发策略组详细信息，具体结构见 UlbPolicySet
	PolicySet []UlbPolicySet
}

/*
SSLBindedTargetSet - DescribeSSL
*/
type SSLBindedTargetSet struct {

	// VServer 所属的ULB实例的资源ID
	ULBId string

	// ULB实例的名称
	ULBName string

	// SSL证书绑定到的VServer的资源ID
	VServerId string

	// 对应的VServer的名字
	VServerName string
}

/*
ULBSSLSet - DescribeSSL
*/
type ULBSSLSet struct {

	// SSL证书绑定到的对象
	BindedTargetSet []SSLBindedTargetSet

	// SSL证书的创建时间
	CreateTime int

	// USSL证书平台的域名,只有当SSLSource为1时才出现
	Domains string

	// SSL证书的HASH值
	HashValue string

	// SSL证书的内容
	SSLContent string

	// SSL证书的Id
	SSLId string

	// SSL证书的名字
	SSLName string

	// SSL证书来源，SSL证书来源，0代表证书来自于ULB平台，1代表证书来自于USSL平台
	SSLSource int

	// SSL证书类型，暂时只有 Pem 一种类型
	SSLType string

	// USSL证书平台的编号,只有当SSLSource为1时才出现
	USSLId string
}

/*
BindVServerInfo - 绑定安全策略的VServer信息
*/
type BindVServerInfo struct {

	// VServer端口
	Port int

	// ULB的ID
	ULBId string

	// 绑定的VServerId
	VServerId string

	// 绑定的VServer名称
	VServerName string
}

/*
SecurityPolicy - 安全策略组
*/
type SecurityPolicy struct {

	// 加密套件
	SSLCiphers []string

	// 安全策略ID
	SecurityPolicyId string

	// 安全策略名称
	SecurityPolicyName string

	// 安全策略类型 0：预定义 1：自定义
	SecurityPolicyType int

	// TLS最低版本
	TLSVersion string

	// 关联的监听
	VServerSet []BindVServerInfo
}

/*
TLSAndCiphers -
*/
type TLSAndCiphers struct {

	// 加密套件
	SSLCiphers []string

	// TLS最低版本
	TLSVersion string
}

/*
PolicyBackendSet - 内容转发下rs详细信息
*/
type PolicyBackendSet struct {

	// 所添加的后端资源在ULB中的对象ID，（为ULB系统中使用，与资源自身ID无关
	BackendId string

	// 后端资源的对象ID
	ObjectId string

	// 所添加的后端资源服务端口
	Port int

	// 后端资源的内网IP
	PrivateIP string

	// 后端资源的实例名称
	ResourceName string

	// 所添加的后端资源的类型，枚举值：UHost -> 云主机；UPM -> 物理云主机； UDHost -> 私有专区主机；UDocker -> 容器；UHybrid->混合云主机；CUBE->Cube；UNI -> 虚拟网卡
	ResourceType string

	// 如果资源绑定了弹性网卡，则展示弹性网卡的资源ID
	SubResourceId string

	// 如果资源绑定了弹性网卡，则展示弹性网卡的资源名称
	SubResourceName string

	// "UNI"或者为空
	SubResourceType string
}

/*
BindSecurityPolicy - VServer绑定的安全策略组信息
*/
type BindSecurityPolicy struct {

	// 加密套件
	SSLCiphers []string

	// 安全策略组ID
	SecurityPolicyId string

	// 安全策略组名称
	SecurityPolicyName string

	// 安全策略类型 0：预定义 1：自定义
	SecurityPolicyType int

	// TLS最低版本
	TLSVersion string
}

/*
ULBPolicySet - 内容转发详细列表
*/
type ULBPolicySet struct {

	// 内容转发下rs的详细信息，参考PolicyBackendSet
	BackendSet []PolicyBackendSet

	// 内容转发规则中域名的匹配方式。枚举值：Regular，正则；Wildcard，泛域名
	DomainMatchMode string

	// 内容转发匹配字段;默认内容转发类型下为空。
	Match string

	// 内容转发Id，默认内容转发类型下为空。
	PolicyId string

	// 内容转发优先级，范围[1,9999]，数字越大优先级越高。默认内容转发规则下为0。
	PolicyPriority int

	// 内容类型，枚举值：Custom -> 客户自定义；Default -> 默认内容转发
	PolicyType string

	// 默认内容转发类型下返回当前rs总数
	TotalCount int

	// 内容转发匹配字段的类型，枚举值：Domain -> 域名；Path -> 路径； 默认内容转发类型下为空
	Type string

	// 所属VServerId
	VServerId string
}

/*
ULBBackendSet - DescribeULB
*/
type ULBBackendSet struct {

	// 后端资源实例的Id
	BackendId string

	// 后端提供服务的实例启用与否，枚举值：0 禁用 1 启用
	Enabled int

	// 是否为backup，只有当vserver的Backup属性为1时才会有此字段，说明：0：主rs1：备rs
	IsBackup int

	// 后端提供服务的端口
	Port int

	// 后端提供服务的内网IP
	PrivateIP string

	// 资源实例的资源Id
	ResourceId string

	// 资源实例的资源名称
	ResourceName string

	// 资源实例的类型
	ResourceType string

	// 后端提供服务的实例运行状态，枚举值：0健康检查健康状态 1 健康检查异常
	Status int

	// 资源绑定的虚拟网卡实例的资源Id
	SubResourceId string

	// 资源绑定的虚拟网卡实例的资源名称
	SubResourceName string

	// 资源绑定的虚拟网卡实例的类型
	SubResourceType string

	// 后端提供服务的资源所在的子网的ID
	SubnetId string

	// 后端服务器所在的VPC
	VPCId string

	// 后端RS权重（在加权轮询算法下有效）
	Weight int
}

/*
FirewallSet - ulb防火墙信息
*/
type FirewallSet struct {

	// 防火墙ID
	FirewallId string

	// 防火墙名称
	FirewallName string
}

/*
ULBIPSet - DescribeULB
*/
type ULBIPSet struct {

	// 弹性IP的带宽值（暂未对外开放）
	Bandwidth int

	// 弹性IP的带宽类型，枚举值：1 表示是共享带宽，0 普通带宽类型（暂未对外开放）
	BandwidthType int

	// 弹性IP地址
	EIP string

	// 弹性IP的ID
	EIPId string

	// 弹性IP的运营商信息，枚举值为：  Bgp：BGP IP International：国际IP
	OperatorName string
}

/*
ULBVServerSet - DescribeULB
*/
type ULBVServerSet struct {

	// 后端资源信息列表，具体结构见下方 ULBBackendSet
	BackendSet []ULBBackendSet

	// 空闲连接的回收时间，单位：秒。
	ClientTimeout int

	// 根据MonitorType确认； 当MonitorType为Port时，此字段无意义。当MonitorType为Path时，代表HTTP检查域名
	Domain string

	// 数据压缩开关，0:关闭 1:开启
	EnableCompression int

	// 0:关闭 1:开启，用于开启http2功能；默认值为0
	EnableHTTP2 int

	// 重定向端口，取值范围[0-65535]；默认值为0，代表关闭；仅HTTP协议支持开启重定向功能
	ForwardPort int

	// VServer服务端口
	FrontendPort int

	// 监听器类型，枚举值为: RequestProxy -> 请求代理；PacketsTransmit -> 报文转发
	ListenType string

	// VServer负载均衡的模式，枚举值：Roundrobin -> 轮询;Source -> 源地址；ConsistentHash -> 一致性哈希；SourcePort -> 源地址（计算端口）；ConsistentHashPort -> 一致性哈希（计算端口）。
	Method string

	// 健康检查类型，枚举值：Port -> 端口检查；Path -> 路径检查；Ping -> Ping探测， Customize -> UDP检查请求代理型默认值为Port，其中TCP协议仅支持Port，其他协议支持Port和Path; 报文转发型TCP协议仅支持Port，UDP协议支持Ping、Port和Customize
	MonitorType string

	// 根据MonitorType确认； 当MonitorType为Port时，此字段无意义。当MonitorType为Path时，代表HTTP检查路径
	Path string

	// 根据PersistenceType确定： None或ServerInsert，此字段为空； UserDefined，此字段展示用户自定义会话string。
	PersistenceInfo string

	// VServer会话保持方式。枚举值为： None -> 关闭会话保持； ServerInsert -> 自动生成； UserDefined -> 用户自定义。
	PersistenceType string

	// 内容转发信息列表，具体结构见下方 ULBPolicySet
	PolicySet []ULBPolicySet

	// VServer实例的协议。 枚举值为：HTTP，TCP，UDP，HTTPS。
	Protocol string

	// 根据MonitorType确认； 当MonitorType为Customize时，此字段有意义，代表UDP检查发出的请求报文
	RequestMsg string

	// 根据MonitorType确认； 当MonitorType为Customize时，此字段有意义，代表UDP检查请求应收到的响应报文
	ResponseMsg string

	// VServer绑定的SSL证书信息，具体结构见下方 ULBSSLSet。
	SSLSet []ULBSSLSet

	// VServer绑定的安全策略,具体结构见BindSecurityPolicy
	SecurityPolicy BindSecurityPolicy

	// VServer的运行状态。枚举值： 0 -> rs全部运行正常;1 -> rs全部运行异常；2 -> rs部分运行异常。
	Status int

	// 负载均衡实例的Id
	ULBId string

	// VServer实例的Id
	VServerId string

	// VServer实例的名字
	VServerName string
}

/*
LoggerSet - ulb日志信息
*/
type LoggerSet struct {

	// ulb日志上传的bucket
	BucketName string

	// 上传到bucket使用的token的tokenid
	TokenID string

	// bucket的token名称
	TokenName string
}

/*
ULBSet - DescribeULB
*/
type ULBSet struct {

	// 带宽
	Bandwidth int

	// 带宽类型，枚举值为： 0，非共享带宽； 1，共享带宽
	BandwidthType int

	// ULB 所属的业务组ID
	BusinessId string

	// ULB的创建时间，格式为Unix Timestamp
	CreateTime int

	// ULB是否开启日志功能。0，关闭；1，开启
	EnableLog int

	// ULB的到期时间，格式为Unix Timestamp
	ExpireTime int `deprecated:"true"`

	// 防火墙信息，具体结构见下方 FirewallSet
	FirewallSet []FirewallSet

	// ULB的详细信息列表，具体结构见下方 ULBIPSet
	IPSet []ULBIPSet

	// ULB ip类型，枚举值：IPv6 / IPv4 （内部测试，暂未对外开放）
	IPVersion string

	// ULB 监听器类型，枚举值：RequestProxy，请求代理； PacketsTransmit ，报文转发；Comprehensive，兼容型；Pending，未定型
	ListenType string

	// 日志功能相关信息，仅当EnableLog为true时会返回，具体结构见下方 LoggerSet
	LogSet LoggerSet

	// 负载均衡的资源名称
	Name string

	// ULB的内网IP，当ULBType为OuterMode时，该值为空
	PrivateIP string

	// 负载均衡的备注
	Remark string

	// ULB的详细信息列表（废弃）
	Resource []string `deprecated:"true"`

	// ULB后向代理IP，仅当有代理IP时返回否
	SnatIps []string

	// ULB 为 InnerMode 时，ULB 所属的子网ID，默认为空
	SubnetId string

	// 负载均衡的业务组名称
	Tag string

	// 负载均衡的资源ID
	ULBId string

	//
	ULBName string `deprecated:"true"`

	// ULB 的类型
	ULBType string

	// ULB所在的VPC的ID
	VPCId string

	// 负载均衡实例中存在的VServer实例列表，具体结构见下方 ULBVServerSet
	VServerSet []ULBVServerSet
}

/*
ULBSimpleSet - ulb简明信息
*/
type ULBSimpleSet struct {

	// 带宽
	Bandwidth int

	// 带宽类型，枚举值为： 0，非共享带宽； 1，共享带宽
	BandwidthType int

	// ULB 所属的业务组ID
	BusinessId string

	// ULB的创建时间，格式为Unix Timestamp
	CreateTime int

	// ULB是否开启日志功能。0，关闭；1，开启
	EnableLog int

	// 防火墙信息，具体结构见下方 FirewallSet
	FirewallSet []FirewallSet

	// ULB的详细信息列表，具体结构见下方 ULBIPSet
	IPSet []ULBIPSet

	// ULB提供服务的IP类型。枚举值，“IPv4”,"IPv6"。默认为“IPv4”
	IPVersion string

	// ULB 监听器类型，枚举值：RequestProxy，请求代理； PacketsTransmit ，报文转发；Comprehensive，兼容型；Pending，未定型
	ListenType string

	// 日志功能相关信息，仅当EnableLog为true时会返回，具体结构见下方 LoggerSet
	LogSet LoggerSet

	// 负载均衡的资源名称
	Name string

	// ULB的内网IP，当ULBType为OuterMode时，该值为空
	PrivateIP string

	// 负载均衡的备注
	Remark string

	// ULB后向代理IP，仅当有代理IP时返回否
	SnatIps []string

	// ULB 为 InnerMode 时，ULB 所属的子网ID
	SubnetId string

	// 负载均衡的业务组名称
	Tag string

	// 负载均衡的资源ID
	ULBId string

	// ULB 的类型（InnerMode or OuterMode）
	ULBType string

	// ULB所在的VPC的ID
	VPCId string

	// ulb下vserver数量
	VServerCount int

	// WAF功能状态，枚举类型：Unavailable：无法创建WAF；NoWAF：未绑定WAF；Intranet：内网回源Waf；Extranet：外网回源Waf
	WAFMode string
}
