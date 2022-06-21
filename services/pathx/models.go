// Code is generated by ucloud-model, DO NOT EDIT IT.

package pathx

/*
GlobalSSHArea -
*/
type GlobalSSHArea struct {

	//
	Area string

	//
	AreaCode string

	//
	RegionSet []string
}

/*
GlobalSSHInfo - GlobalSSH实例信息
*/
type GlobalSSHInfo struct {

	// GlobalSSH分配的加速域名。
	AcceleratingDomain string

	// 被SSH访问的IP所在地区
	Area string

	// globalssh Ultimate带宽包大小
	BandwidthPackage int

	// 支付周期，如Month,Year,Dynamic等
	ChargeType string

	// 资源创建时间戳
	CreateTime int

	// 是否过期
	Expire bool

	// 资源过期时间戳
	ExpireTime int

	// InstanceType为Basic版本时，需要展示具体分配的转发机房
	ForwardRegion string

	// InstanceType等于Free时，由系统自动分配，不等于源站Port值。InstanceType不等于Free时，与源站Port值相同。
	GlobalSSHPort int

	// 实例ID，资源唯一标识
	InstanceId string

	// 枚举值：["Enterprise","Basic","Free","Welfare"], 分别代表企业版，基础版本，免费版本，较早的公测免费版
	InstanceType string

	// 源站服务器监听的SSH端口，windows系统为RDP端口
	Port int

	// 备注信息
	Remark string

	// 被SSH访问的源站 IPv4地址。
	TargetIP string
}

/*
LineDetail - 子线路
*/
type LineDetail struct {

	// 线路源
	LineFrom string

	// 线路源中文名称
	LineFromName string

	// 线路计费Id
	LineId string

	// 线路目的
	LineTo string

	// 线路目的中文名称
	LineToName string
}

/*
UGAALine - UGAA加速线路
*/
type UGAALine struct {

	// 子线路信息
	LineDetail []LineDetail

	// 线路源
	LineFrom string

	// 线路源中文名称
	LineFromName string

	// 线路计费Id
	LineId string

	// 线路目的
	LineTo string

	// 线路目的中文名称
	LineToName string

	// 线路可售最大带宽
	MaxBandwidth int
}

/*
SSLBindedTargetSet - Describle SSL Bind UAG Info
*/
type SSLBindedTargetSet struct {

	// SSL证书绑定到的实例ID
	ResourceId string

	// SSL证书绑定到的实例名称
	ResourceName string
}

/*
PathXSSLSet - Describle PathX SSL
*/
type PathXSSLSet struct {

	// SSL证书的创建时间 时间戳
	CreateTime int

	// 证书过期时间 时间戳
	ExpireTime int

	// SSL证书绑定的对象
	SSLBindedTargetSet []SSLBindedTargetSet

	// SSL证书内容
	SSLContent string

	// SSL证书的Id
	SSLId string

	// SSL证书（用户证书、私钥、ca证书合并）内容md5值
	SSLMD5 string

	// SSL证书的名字
	SSLName string

	// 证书来源，0：用户上传 1: 免费颁发
	SourceType int

	// 证书域名
	SubjectName string
}

/*
ForwardArea - 全地域加速源站区域
*/
type ForwardArea struct {

	// 源站区域中文
	Area string

	// 源站区域代码
	AreaCode string

	// 大陆代码
	ContinentCode string

	// 国家代码
	CountryCode string

	// 国旗 emoji
	FlagEmoji string

	// 国旗unicode
	FlagUnicode string
}

/*
SrcAreaInfo - 接入地域信息
*/
type SrcAreaInfo struct {

	// AreaCode对应城市名
	Area string

	// AreaCode ,城市机场代码
	AreaCode string

	// 国旗Emoji
	FlagEmoji string

	// 国旗Unicode
	FlagUnicode string
}

/*
OutPublicIpInfo - 线路回源IP信息
*/
type OutPublicIpInfo struct {

	// 线路回源节点机房代号
	Area string

	// 线路回源节点EIP
	IP string
}

/*
ForwardTask - 全球统一接入转发端口任务信息
*/
type ForwardTask struct {

	// 加速端口
	Port int

	// 转发协议，枚举值["TCP"，"UDP"，"HTTPHTTP"，"HTTPSHTTP"，"HTTPSHTTPS"，"WSWS"，"WSSWS"，"WSSWSS"]。TCP和UDP代表四层转发，其余为七层转发。
	Protocol string

	// 源站服务器监听的端口号
	RSPort int
}

/*
AccelerationAreaInfos - 加速大区信息
*/
type AccelerationAreaInfos struct {

	// 加速区code
	AccelerationArea string

	// 加速节点信息
	AccelerationNodes []SrcAreaInfo
}

/*
ForwardInfo - 全球统一接入加速实例配置信息
*/
type ForwardInfo struct {

	// 加速大区代码
	AccelerationArea string

	// 加速节点列表
	AccelerationAreaInfos []AccelerationAreaInfos

	// 加速大区名称
	AccelerationAreaName string

	// 购买的带宽值
	Bandwidth int

	// 加速域名
	CName string

	// 计费方式
	ChargeType string

	// 资源创建时间
	CreateTime int

	// 源站域名
	Domain string

	// 回源出口IP地址
	EgressIpList []OutPublicIpInfo

	// 资源过期时间
	ExpireTime int

	// 源站IP列表，多个值由半角英文逗号相隔
	IPList []string

	// 加速配置ID
	InstanceId string

	// 加速实例名称
	Name string

	// 源站中文名
	OriginArea string

	// 源站AreaCode
	OriginAreaCode string

	// 端口列表
	PortSets []ForwardTask

	// 备注
	Remark string
}

/*
NodeDelays - 全地域加速各个区域加速延迟情况
*/
type NodeDelays struct {

	// 加速区域
	Area string

	// 加速区域Code
	AreaCode string

	// 国家代码
	CountryCode string

	// 国旗Emoji
	FlagEmoji string

	// 国旗Code
	FlagUnicode string

	// 加速延迟
	Latency float64

	// 公网延迟
	LatencyInternet float64

	// 加速提升比例
	LatencyOptimization float64

	// 加速后丢包率
	Loss float64

	// 原始丢包率
	LossInternet float64

	// 丢包下降比例
	LossOptimization float64
}

/*
AccelerationInfo - 加速提升信息
*/
type AccelerationInfo struct {

	// 加速大区代码
	AccelerationArea string

	// 加速大区名称
	AccelerationName string

	// 加速提升情况
	NodeInfo []NodeDelays
}

/*
UGAL7Forwarder - UGA实例 7层转发器信息
*/
type UGAL7Forwarder struct {

	// 接入端口
	Port int

	// 转发协议，枚举值["TCP"，"UDP"，"HTTPHTTP"，"HTTPSHTTP"，"HTTPSHTTPS"]。TCP和UDP代表四层转发，其余为七层转发
	Protocol string

	// RSPort，源站监听端口
	RSPort int

	// 证书ID
	SSLId string

	// 证书名称
	SSLName string
}

/*
UGAL4Forwarder - UGA实例 4层转发器信息
*/
type UGAL4Forwarder struct {

	// 接入端口
	Port int

	// 转发协议，枚举值["TCP"，"UDP"，"HTTPHTTP"，"HTTPSHTTP"，"HTTPSHTTPS"]。TCP和UDP代表四层转发，其余为七层转发
	Protocol string

	// RSPort，源站监听端口
	RSPort int
}

/*
UPathSet - uga关联的upath信息
*/
type UPathSet struct {

	// 带宽 Mbps, 1~800Mbps
	Bandwidth int

	// 线路起点英文代号，加速区域
	LineFrom string

	// 线路起点中文名字，加速区域
	LineFromName string

	// 线路ID
	LineId string

	// 线路对端英文代号，源站区域
	LineTo string

	// 线路对端中文名字，源站区域
	LineToName string

	// UPath 实例ID
	UPathId string

	// UPath名字
	UPathName string
}

/*
UGAATask - 用户在UGAA实例下配置的多端口任务
*/
type UGAATask struct {

	// 接入端口
	Port int

	// 转发协议，枚举值["TCP"，"UDP"，"HTTPHTTP"，"HTTPSHTTP"，"HTTPSHTTPS"]。TCP和UDP代表四层转发，其余为七层转发
	Protocol string
}

/*
UGAAInfo - 全球加速实例信息
*/
type UGAAInfo struct {

	// 加速域名，请在加速区域配置您的业务域名的CName记录值为加速域名
	CName string

	// 源站域名
	Domain string

	// 源站IP列表，多个值由半角英文逗号相隔
	IPList []string

	// UGA 4层转发器配置，记录接入或回源端口，接入或回源协议信息
	L4ForwarderSet []UGAL4Forwarder

	// UGA 7层转发器配置，记录接入或回源端口，接入或回源协议信息 如绑定证书会返回证书ID
	L7ForwarderSet []UGAL7Forwarder

	// 源站所在区域，加速实例在绑定线路后会自动设置该值。console页面上通过该值过滤加速实例可以绑定的upath实例。注意：缺少该值会导致在console上无法修改线路
	Location string

	// 线路出口IP地址
	OutPublicIpList []OutPublicIpInfo

	// 端口配置信息（不再维护，建议使用ForwarderSet）
	TaskSet []UGAATask

	// 加速配置实例ID
	UGAId string

	// 加速配置名称
	UGAName string

	// 绑定的加速线路
	UPathSet []UPathSet
}

/*
PathXUGAInfo - 加速实例配置信息
*/
type PathXUGAInfo struct {

	// 源站域名
	Domain string

	// 源站IP列表，多个值由半角英文逗号相隔
	IPList []string

	// 加速配置ID
	UGAId string
}

/*
UPathInfo - 加速线路信息
*/
type UPathInfo struct {

	// 带宽，单位Mbps
	Bandwidth int

	// 计费模式，默认为Month 按月收费,可选范围['Month','Year','Dynamic']
	ChargeType string

	// UPath创建的时间，10位时间戳
	CreateTime int

	// UPath的过期时间，10位时间戳
	ExpireTime int

	// 线路入口名称
	LineFromName string

	// 选择的线路
	LineId string

	// 线路出口名称
	LineToName string

	// UPath实例名字
	Name string

	// 线路出口IP数组
	OutPublicIpList []OutPublicIpInfo

	// 是否为后付费实例
	PostPaid bool

	// 与该UPath绑定的UGA列表
	UGAList []PathXUGAInfo

	// UPath加速线路实例ID
	UPathId string
}

/*
AlarmRuler - 告警详情
*/
type AlarmRuler struct {

	// 告警探测周期，单位秒
	AlarmFrequency int

	// 收敛策略，可选范围 ['Exponential','Continuous','Once']，分别对应指数递增、连续告警、单次告警
	AlarmStrategy string

	// 告警模板策略ID
	AlarmTemplateRuleId int

	// 比较策略，可选 ['GE','LE']  分别代表不小于和不大于
	Compare string

	// 联系组ID
	ContactGroupId int

	// 告警指标名称, 所有n的个数必须一致。目前仅允许以下四项：UpathNetworkOut:出向带宽，UpathNetworkIn:入向带宽，UpathNetworkOutUsage:出向带宽使用率，UpathNetworkInUsage:入向带宽使用率
	MetricName string

	// 资源类型
	ResourceType string

	// 告警阈值，带宽使用率的阈值范围是[50,100]的正整数，带宽告警阈值为1000000的倍数, 如大于2Mbps则告警 阈值应该传 2000000
	Threshold int

	// 告警触发周期（次数）
	TriggerCount int
}

/*
MatricPoint - 某一时刻的监控数据
*/
type MatricPoint struct {

	// 时间戳
	Timestamp int

	// 监控点数值
	Value int
}

/*
MetricPeriod - 一段时间内的监控数据
*/
type MetricPeriod struct {

	// 入向带宽
	NetworkIn []MatricPoint

	// 入向带宽使用率
	NetworkInUsage []MatricPoint

	// 出向带宽
	NetworkOut []MatricPoint

	// 出向带宽使用率
	NetworkOutUsage []MatricPoint
}

/*
UGA3Metric - 一段时间内的监控数据
*/
type UGA3Metric struct {

	// 当前连接数
	ConnectCount []MatricPoint

	// 子线路当前连接数
	ConnectCountSubline []MatricPoint

	// 线路平均延迟
	Delay []MatricPoint

	// 延迟提升
	DelayPromote []MatricPoint

	// 子线路延迟提升
	DelayPromoteSubline []MatricPoint

	// 子线路延迟
	DelaySubline []MatricPoint

	// 入向带宽
	NetworkIn []MatricPoint

	// 子线路入口带宽
	NetworkInSubline []MatricPoint

	// 入向带宽使用率
	NetworkInUsage []MatricPoint

	// 出向带宽
	NetworkOut []MatricPoint

	// 子线路出口带宽
	NetworkOutSubline []MatricPoint

	// 出向带宽使用率
	NetworkOutUsage []MatricPoint
}

/*
UGA3Price -
*/
type UGA3Price struct {

	// 加速大区代码
	AccelerationArea string

	// 加速大区名称
	AccelerationAreaName string

	// 加速配置带宽价格
	AccelerationBandwidthPrice float64

	// 转发配置价格
	AccelerationForwarderPrice float64
}
