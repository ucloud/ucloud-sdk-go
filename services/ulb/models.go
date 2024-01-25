// Code is generated by ucloud-model, DO NOT EDIT IT.

package ulb

/*
TargetSet - 服务节点信息
*/
type TargetSet struct {

	// 服务节点是否启用。 默认值：true
	Enabled bool

	// 服务节点的标识ID。
	Id string

	// 服务节点是否为备节点。 默认值：false
	IsBackup bool

	// 服务节点的端口。限定取值：[1-65535]； 默认值：80
	Port int

	// 服务节点的IP。在IP类型时，必传
	ResourceIP string

	// 服务节点的资源ID。在非IP类型时，必传
	ResourceId string

	// 服务节点的类型。限定枚举值：UHost -> 云主机，UNI -> 虚拟网卡，UPM -> 物理云主机，IP ->  IP类型； 默认值："UHost"； 非IP类型，如果该资源有多个IP，将只能添加主IP； 非IP类型，展示时，会显示相关资源信息，IP类型只展示IP信息。 在相关资源被删除时，非IP类型会把相关资源从lb中剔除，IP类型不保证这个逻辑
	ResourceType string

	// 服务节点的健康检查状态。限定枚举值：Healthy -> 健康，Unhealthy -> 不健康
	State string

	// 服务节点的子网资源ID。在IP类型时，必传
	SubnetId string

	// 服务节点的VPC资源ID。在IP类型时，必传
	VPCId string

	// 服务节点的权重。限定取值：[1-100]； 仅在加权轮询算法时有效； 默认值：1
	Weight int
}

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
ForwardTargetSet - 转发的后端服务节点
*/
type ForwardTargetSet struct {

	// 服务节点的标识ID
	Id string

	// 权重。仅监听器负载均衡算法是加权轮询是有效；取值范围[1-100]，默认值为1
	Weight int
}

/*
ForwardConfigSet - 转发服务节点相关配置
*/
type ForwardConfigSet struct {

	// 转发的后端服务节点。限定在监听器的服务节点池里；数组长度可以为0。具体结构详见 ForwardTargetSet
	Targets []ForwardTargetSet
}

/*
PathConfigSet - 路径相关配置
*/
type PathConfigSet struct {

	// 取值。暂时只支持数组长度为1； 取值需符合相关匹配方式的条件
	Values []string
}

/*
HostConfigSet - 域名相关配置
*/
type HostConfigSet struct {

	// 匹配方式。限定枚举值：Regular-正则，Wildcard-泛域名； 默认值：Regular
	MatchMode string

	// 取值。暂时只支持数组长度为1； 取值需符合相关匹配方式的条件
	Values []string
}

/*
RuleAction - 转发动作
*/
type RuleAction struct {

	// 转发服务节点相关配置。 具体结构详见 ForwardConfigSet
	ForwardConfig ForwardConfigSet

	// 动作类型。限定枚举值：Forward
	Type string
}

/*
RuleCondition - 转发规则匹配条件
*/
type RuleCondition struct {

	// 域名相关配置。Type为Host时必填。具体结构详见 HostConfigSet
	HostConfig HostConfigSet

	// 路径相关配置。Type为Path时必填。具体结构详见 PathConfigSet
	PathConfig PathConfigSet

	// 匹配条件类型。限定枚举值：Host，Path
	Type string
}

/*
HealthCheckConfigSet - 健康检查相关配置
*/
type HealthCheckConfigSet struct {

	// （应用型专用）HTTP检查域名。 当Type为HTTP时，此字段有意义，代表HTTP检查域名
	Domain string

	// 是否开启健康检查功能。暂时不支持关闭。 默认值为：true
	Enabled bool

	// （应用型专用）HTTP检查路径。当Type为HTTP时，此字段有意义，代表HTTP检查路径
	Path string

	// 健康检查方式。应用型限定取值： Port -> 端口检查；HTTP -> HTTP检查； 默认值：Port
	Type string
}

/*
Certificate - （应用型专用）服务器证书信息
*/
type Certificate struct {

	// 是否为默认证书
	IsDefault bool

	// 证书ID
	SSLId string
}

/*
Rule - （应用型专用）转发规则信息
*/
type Rule struct {

	// 是否为默认转发规则
	IsDefault bool

	// 当转发的服务节点为空时，规则是否忽略
	Pass bool

	// 转发动作。具体规则详见RuleAction
	RuleActions []RuleAction

	// 转发规则匹配条件。具体结构详见 RuleCondition
	RuleConditions []RuleCondition

	// 转发规则的ID
	RuleId string
}

/*
StickinessConfigSet - 会话保持相关配置
*/
type StickinessConfigSet struct {

	// （应用型专用）自定义Cookie。当StickinessType取值"UserDefined"时有效
	CookieName string

	// 是否开启会话保持功能。应用型负载均衡实例基于Cookie实现
	Enabled bool

	// （应用型专用）Cookie处理方式。限定枚举值： ServerInsert -> 自动生成KEY；UserDefined -> 用户自定义KEY
	Type string
}

/*
Target - 服务节点信息
*/
type Target struct {

	// 服务节点是否启用
	Enabled bool

	// 服务节点的标识ID。为ALB/NLB中使用，与资源自身ID无关，可用于UpdateTargetsAttribute/RemoveTargets
	Id string

	// 服务节点是否为备节点
	IsBackup bool

	// 服务节点的端口
	Port int

	// 服务节点的IP
	ResourceIP string

	// 服务节点的资源ID
	ResourceId string

	// 服务节点的资源名称
	ResourceName string

	// 服务节点的类型。限定枚举值：UHost -> 云主机，UNI -> 虚拟网卡，UPM -> 物理云主机，IP ->  IP类型； 默认值："UHost"； 非IP类型，如果该资源有多个IP，将只能添加主IP； 非IP类型，展示时，会显示相关资源信息，IP类型只展示IP信息。 在相关资源被删除时，非IP类型会把相关资源从lb中剔除，IP类型不保证这个逻辑
	ResourceType string

	// 服务节点的健康检查状态。限定枚举值：Healthy -> 健康，Unhealthy -> 不健康
	State string

	// 服务节点的子网资源ID
	SubnetId string

	// 服务节点的VPC资源ID
	VPCId string

	// 服务节点的权重。仅在加权轮询算法时有效
	Weight int
}

/*
Listener - 负载均衡监听器信息
*/
type Listener struct {

	// （应用型专用）服务器默认证书ID。仅HTTPS监听支持。具体接口详见 Certificate
	Certificates []Certificate

	// （应用型专用）是否开启数据压缩功能。目前只支持使用gzip对特定文件类型进行压缩
	CompressionEnabled bool

	// （应用型专用）是否开启HTTP/2特性。仅HTTPS监听支持开启
	HTTP2Enabled bool

	// 健康检查相关配置。具体结构详见 HealthCheckConfigSet
	HealthCheckConfig HealthCheckConfigSet

	// 连接空闲超时时间。单位：秒
	IdleTimeout int

	// 监听器的ID
	ListenerId string

	// 监听器的监听端口
	ListenerPort int

	// 监听协议。应用型限定取值： HTTP、HTTPS
	ListenerProtocol string

	// 监听器的名称
	Name string

	// （应用型专用）是否开启HTTP重定向到HTTPS。仅HTTP监听支持开启
	RedirectEnabled bool

	// （应用型专用）重定向端口
	RedirectPort int

	// 监听器的备注信息
	Remark string

	// （应用型专用）转发规则信息
	Rules []Rule

	// 负载均衡算法。应用型限定取值：Roundrobin -> 轮询;Source -> 源地址； WeightRoundrobin -> 加权轮询; Leastconn -> 最小连接数；Backup ->主备模式
	Scheduler string

	// （应用型专用）安全策略组ID。仅HTTPS监听支持绑定；Default -> 原生策略
	SecurityPolicyId string

	// listener健康状态。限定枚举值：Healthy -> 健康，Unhealthy -> 不健康，PartialHealth -> 部分健康，None -> 无节点状态
	State string

	// 会话保持相关配置。具体结构详见 StickinessConfigSet
	StickinessConfig StickinessConfigSet

	// 添加的服务节点信息。具体结构详见 Target
	Targets []Target
}

/*
AccessLogConfigSet - （应用型专用）访问日志相关配置
*/
type AccessLogConfigSet struct {

	// （应用型专用）是否开启访问日志记录功能
	Enabled bool

	// （应用型专用）用于存储访问日志的bucket
	US3BucketName string

	// （应用型专用）上传访问日志到bucket所需的token
	US3TokenId string
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
IPInfo - 绑定的IP信息
*/
type IPInfo struct {

	// 网络模式。 限定枚举值：Internet -> 互联网，Intranet -> 内联网
	AddressType string

	// 带宽值。单位M
	Bandwidth int

	// 带宽类型。限定枚举值：1 -> 共享带宽，0 -> 普通带宽类型
	BandwidthType int

	// IP地址
	IP string

	// IP协议版本
	IPVersion string

	// 唯一标识ID
	Id string

	// 外网IP的运营商信息。枚举值为：Telecom -> 电信，Unicom -> 联通，International -> 国际IP，Bgp -> BGP，Duplet -> 双线（电信+联通双线路），BGPPro -> 精品BGP，China-mobile -> 中国移动，Anycast -> AnycastEIP
	OperatorName string
}

/*
LoadBalancer - 负载均衡实例信息
*/
type LoadBalancer struct {

	// （应用型专用）访问日志相关配置
	AccessLogConfig AccessLogConfigSet

	// 是否开启自动续费
	AutoRenewEnabled bool

	// 付费模式
	ChargeType string

	// 负载均衡实例创建时间。格式为Unix Timestamp
	CreateTime int

	// 防火墙信息
	Firewall FirewallSet

	// 绑定的IP信息。具体结构详见 IPInfo
	IPInfos []IPInfo

	// 负载均衡实例支持的IP协议版本
	IPVersion string

	// 监听器信息。当ShowDetail为false时，为空
	Listeners []Listener

	// 负载均衡实例的ID
	LoadBalancerId string

	// 负载均衡实例的名称
	Name string

	// 有效期（计费）。格式为Unix Timestamp
	PurchaseValue int

	// 负载均衡实例的备注信息
	Remark string

	// 应用型实例的代理IP或网络型FULLNAT模式下snat所用的IP
	SnatIPs []string

	// lb状态：Normal-正常；Arrears-欠费停服
	Status string

	// 负载均衡实例所属的子网资源ID。负载均衡实例的内网VIP和SNAT场景的源IP限定在该子网内；指定子网不影响添加后端服务节点时的范围，依旧是整个VPC下支持的资源
	SubnetId string

	// 负载均衡实例所属的业务组ID
	Tag string

	// 负载均衡实例的类型。限定枚举值：Application -> 应用型，Network -> 网络型
	Type string

	// 负载均衡实例所属的VPC资源ID
	VPCId string
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

	// 证书过期时间,只有当SSLSource为1时才出现
	NotAfter int

	// 证书颁发时间,只有当SSLSource为1时才出现
	NotBefore int

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
SSLRelation - SSL证书绑定到的对象
*/
type SSLRelation struct {

	// 是否为监听器默认SSL证书
	IsDefault bool

	// 监听器的ID
	ListenerId string

	// 监听器的名称
	ListenerName string

	// 负载均衡实例的ID
	LoadBalancerId string

	// 负载均衡实例的名称
	LoadBalancerName string
}

/*
SSLInfo - SSL返回信息
*/
type SSLInfo struct {

	// SSL证书的创建时间
	CreateTime int

	// USSL证书平台的域名,只有当SSLSource为1时才出现
	Domains string

	// SSL证书的HASH值
	HashValue string

	// 证书过期时间,只有当SSLSource为1时才出现
	NotAfter int

	// 证书颁发时间,只有当SSLSource为1时才出现
	NotBefore int

	// SSL绑定ULB和ALB的关系
	Relations []SSLRelation

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
SecurityPolicyRelation - 关联的监听
*/
type SecurityPolicyRelation struct {

	// 监听器的ID
	ListenerId string

	// 监听器的名称
	ListenerName string

	// 监听端口
	ListenerPort int

	// 负载均衡实例的ID
	LoadBalancerId string
}

/*
SecurityPolicyInfo - 安全策略信息
*/
type SecurityPolicyInfo struct {

	// 关联的监听
	Relations []SecurityPolicyRelation

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

/*
BackendMsg - ulb修rs状态时返回的信息
*/
type BackendMsg struct {

	// rs的资源ID
	BackendId string

	// 修改rs返回的消息
	SubMessage string

	// 修改rs的返回值
	SubRetCode int
}
