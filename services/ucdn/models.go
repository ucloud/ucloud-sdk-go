// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucdn

/*
CacheConf - 缓存配置
*/
type CacheConf struct {

	// 是否缓存，true为缓存，flase为不缓存。为flase的情况下，CacheTTL和CacheUnit强制不生效
	CacheBehavior bool

	// 缓存时间
	CacheTTL int

	// 缓存时间的单位。sec（秒），min（分钟），hour（小时），day（天）。上限1年。
	CacheUnit string

	// 缓存规则描述
	Description string

	// 是否优先遵循源站头部缓存策略，false为不优先遵循源站，true为优先遵循源站缓存头部。默认为0
	FollowOriginRule bool

	// 状态码模式，非200，206状态码，多个状态码用竖线(|)分隔，该属性仅仅在状态码缓存配置列表中返回
	HttpCodePattern string

	// 路径模式，支持正则
	PathPattern string
}

/*
AccessConf - 访问控制
*/
type AccessConf struct {

	// 多个ip用逗号隔开
	IpBlacklist string
}

/*
DomainInfo - 域名配置
*/
type DomainInfo struct {

	// 访问控制
	AccessConf AccessConf

	// 查询带宽区域 cn代表国内 abroad代表海外 不填默认为全部区域
	AreaCode string

	// 缓存配置规则列表
	CacheConf []CacheConf

	// 缓存Host，不同的域名可以配置为同一个CacheHost来实现缓存共享，默认为加速域名
	CacheHost string

	// 加速类型http,http|https
	CdnProtocol string

	// 加速域名的业务类型，web代表网站，stream代表视频，download代表下载。
	CdnType string

	// 证书名称
	CertName string

	// cdn域名。创建加速域名生成的cdn域名，用于设置CNAME记录
	Cname string

	// 域名创建的时间。格式：时间戳
	CreateTime int

	// 域名，用户创建加速的域名
	Domain string

	// 域名id，创建域名时生成的id
	DomainId string

	// 国外https状态 enableing-开启中  fail-开启失败 enable-启用 disable-未启用
	HttpsStatusAbroad string

	// 国内https状态 enableing-开启中 fail-开启失败 enable-启用 disable-未启用
	HttpsStatusCn string

	// ReferType为白名单时，NullRefer为false代表不允许NULL refer访问，为true代表允许Null refer访问
	NullRefer bool

	// 回源Http请求头部Host，默认是加速域名
	OriginHost string

	// 源站ip即cdn服务器回源访问的ip地址。支持多个源站ip，多个源站ip，可表述为如：[1.1.1.1,2.2.2.2]
	OriginIp []string

	// 回源端口
	OriginPort int

	// 源站协议http，http|https   默认http
	OriginProtocol string

	// Refer列表，支持正则表达式
	ReferList []string

	// refer配置开关，true打开，false关闭
	ReferStatus bool

	// 0白名单，1黑名单
	ReferType int

	// 创建的加速域名的当前的状态。check代表审核中，checkSuccess代表审核通过，checkFail代表审核失败，enable代表加速中，disable代表停止加速，delete代表删除加速 enableing代表正在开启加速，disableing代表正在停止加速中，deleteing代表删除中
	Status string

	// 业务组，默认为Default
	Tag string

	// 测试url，用于域名创建加速时的测试
	TestUrl string

	// 开始分配Cname时间。格式：时间戳
	ValidTime int
}

/*
UrlProgressInfo - UrlProgressInfo
*/
type UrlProgressInfo struct {

	// 刷新任务创建的时间。格式为Unix Timestamp
	CreateTime int

	// 任务完成时间。格式为Unix Timestamp
	FinishTime int

	// 刷新进度，单位%
	Progress int

	// 刷新任务的当前状态，枚举值：success：成功；wait：排队中；process：处理中；failure：失败； unknow：未知
	Status string

	// 刷新的单条url
	Url string
}

/*
TaskInfo - 预取刷新的任务信息
*/
type TaskInfo struct {

	// 刷新任务创建的时间。格式为Unix Timestamp
	CreateTime int

	// 刷新任务的当前状态，枚举值：success：成功；wait：排队中；process：处理中；failure：失败； unknow：未知
	Status string

	// 提交任务时返回的任务ID
	TaskId string

	// file/dir  刷新任务会返回Type，预取任务没有
	Type string `deprecated:"true"`

	// 任务url的信息列表，参考UrlProgressInfo
	UrlLists []UrlProgressInfo
}

/*
KwaiAuthKv - 快手鉴权键值信息
*/
type KwaiAuthKv struct {

	// iv信息
	Iv string

	// key信息
	Key string
}

/*
KwaiAuthConfig - 鉴权信息
*/
type KwaiAuthConfig struct {

	//
	Keys []KwaiAuthKv

	// 类型  pkey / ksc / typeA
	Type string
}

/*
KwaiDomainAuthConfig - 快手域名鉴权信息
*/
type KwaiDomainAuthConfig struct {

	//
	Config []KwaiAuthConfig

	// 域名
	Domain string
}

/*
CertList - 证书信息
*/
type CertList struct {

	// 证书开始时间
	BeginTime int

	// ca证内容
	CaCert string

	// 证书名
	CertName string

	// 通用名
	CommonName string

	// dns名称
	DnsName string

	// 已配置域名个数
	DomainCount int

	// 已配置的域名列表
	Domains []string

	// 证书获取时间
	EndTime int

	// 证书内容
	UserCert string
}

/*
BandwidthInfo - BandwidthInfo
*/
type BandwidthInfo struct {

	// 返回值返回指定时间区间内CDN的带宽峰值，单位Mbps（如果请求参数Type为0，则Value是五分钟粒度的带宽值，如果Type为1，则Value是1小时的带宽峰值，如果Type为2，则Value是一天内的带宽峰值）
	CdnBandwidth float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
HitRateInfo - HitRateInfo
*/
type HitRateInfo struct {

	// 总流量命中率，单位%
	FlowHitRate float64

	// 请求数命中率，单位%
	RequestHitRate float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
HttpCodeInfo - HttpCodeInfo
*/
type HttpCodeInfo struct {

	// 5xx数量
	HttpFiveXX int

	// 4xx数量
	HttpFourXX int

	// 1xx数量
	HttpOneXX int

	// 3xx数量
	HttpThreeXX int

	// 2xx数量
	HttpTwoXX int

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
HttpCodeV2Detail - HTTP状态码详细信息
*/
type HttpCodeV2Detail struct {

	// http100数量
	Http100 int

	// http101数量
	Http101 int

	// http102数量
	Http102 int

	// http200数量
	Http200 int

	// http201数量
	Http201 int

	// http202数量
	Http202 int

	// http203数量
	Http203 int

	// http204数量
	Http204 int

	// http205数量
	Http205 int

	// http206数量
	Http206 int

	// http207数量
	Http207 int

	// http300数量
	Http300 int

	// http301数量
	Http301 int

	// http302数量
	Http302 int

	// http303数量
	Http303 int

	// http304数量
	Http304 int

	// http305数量
	Http305 int

	// http306数量
	Http306 int

	// http307数量
	Http307 int

	// http400数量
	Http400 int

	// http401数量
	Http401 int

	// http402数量
	Http402 int

	// http403数量
	Http403 int

	// http404数量
	Http404 int

	// http405数量
	Http405 int

	// http406数量
	Http406 int

	// http407数量
	Http407 int

	// http408数量
	Http408 int

	// http409数量
	Http409 int

	// http410数量
	Http410 int

	// http411数量
	Http411 int

	// http412数量
	Http412 int

	// http413数量
	Http413 int

	// http414数量
	Http414 int

	// http415数量
	Http415 int

	// http416数量
	Http416 int

	// http417数量
	Http417 int

	// http418数量
	Http418 int

	// http421数量
	Http421 int

	// http422数量
	Http422 int

	// http423数量
	Http423 int

	// http424数量
	Http424 int

	// http425数量
	Http425 int

	// http426数量
	Http426 int

	// http449数量
	Http449 int

	// http451数量
	Http451 int

	// http500数量
	Http500 int

	// http501数量
	Http501 int

	// http502数量
	Http502 int

	// http503数量
	Http503 int

	// http504数量
	Http504 int

	// http505数量
	Http505 int

	// http506数量
	Http506 int

	// http507数量
	Http507 int

	// http509数量
	Http509 int

	// http510数量
	Http510 int

	// 时间
	Time int

	// 当前分组的总状态码数
	Total int
}

/*
RequestInfo - RequestInfo
*/
type RequestInfo struct {

	// 返回值返回指定时间区间内的cdn收到的请求次数之和
	CdnRequest float64

	// 返回值返回指定时间区间内的cdn回源的请求次数之和
	OriginRequest float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
RefererList - RefererList
*/
type RefererList struct {

	// 次数占比，单位%
	Percent float64

	// 客户端请求的referer
	Referer string

	// 次数
	RequestTimes int
}

/*
RefererStatistics - RefererStatistics
*/
type RefererStatistics struct {

	// 日期
	Date string

	// Referer实例表
	RefererList []RefererList
}

/*
DownloadStatisticInfo - DownloadStatisticInfo
*/
type DownloadStatisticInfo struct {

	// 下载次数
	DownloadTimes int

	// 流量占比，单位%
	Percent float64

	// 流量（单位为G）
	Traffic float64

	// 下载链接的url
	Url string
}

/*
UrlStatistics - UrlStatistics
*/
type UrlStatistics struct {

	// 日期
	Date string

	//
	UrlList []DownloadStatisticInfo
}

/*
BandwidthTrafficInfo - BandwidthTrafficInfo
*/
type BandwidthTrafficInfo struct {

	// 返回值返回指定时间区间内CDN的带宽峰值，单位Mbps（如果请求参数Type为0，则Value是五分钟粒度的带宽值，如果Type为1，则Value是1小时的带宽峰值，如果Type为2，则Value是一天内的带宽峰值）
	CdnBandwidth float64

	// 带宽获取的时间点。格式：时间戳
	Time int

	// 对应时间粒度的流量，单位字节
	Traffic float64
}

/*
CacheKeyInfo - 忽略参数缓存配置
*/
type CacheKeyInfo struct {

	// 是否忽略
	Ignore bool

	// 路径模式，支持正则
	PathPattern string

	// 自定义变量,以$符号开头，多个变量用加号(+)连接，$querystring表示所有变量
	QueryString string
}

/*
ReferConf - refer配置
*/
type ReferConf struct {

	// ReferType为白名单时（删除），NullRefer为0代表不允许NULL refer访问，为1代表允许Null refer访问
	NullRefer int

	// Refer防盗链规则列表，支持正则表达式
	ReferList []string

	// Refer防盗链配置  0白名单，1黑名单
	ReferType int
}

/*
OriginConf - 回源配置
*/
type OriginConf struct {

	// 1如果为false表示BackupOriginIp为空，表示没有备份源站，忽略BackupOriginIp，BackupOriginHost字段2如果为true表示BackupOriginIp.n必须至少有一个备份源站地址
	BackupOriginEnable bool

	// 备份回源Http请求头部Host，默认是加速域名
	BackupOriginHost string

	// 备份源站ip即cdn服务器回源访问的ip地址。多个源站ip，可以这样表述，如：["1.1.1.1","2.2.2.2"]
	BackupOriginIpList []string

	// 主源响应的回源错误码（如：404|500），默认空字符串
	OriginErrorCode string

	// 回主源的回源失败数，默认1
	OriginErrorNum int

	// 跟随301跳转  0=不跟随 1=跟随
	OriginFollow301 int

	// 回源Http请求头部Host，默认是加速域名
	OriginHost string

	// 源站ip即cdn服务器回源访问的ip地址。多个源站ip，可以这样表述，如：["1.1.1.1","2.2.2.2"]
	OriginIpList []string

	// 回源端口
	OriginPort int

	// 源站协议http，http|https   默认http
	OriginProtocol string
}

/*
AdvancedConf - 域名高级配置
*/
type AdvancedConf struct {

	// http转https回源 true是，false否
	Http2Https bool

	// 客户端响应http头列表
	HttpClientHeader []string

	// 源站http头列表
	HttpOriginHeader []string

	// 是否开启quic
	QuicEnable bool

	// 是否开启websocket
	WebSocketEnable bool
}

/*
CacheAllConfig - 缓存相关的配置
*/
type CacheAllConfig struct {

	// 缓存Host，不同的域名可以配置为同一个CacheHost来实现缓存共享，默认为加速域名
	CacheHost string

	// 忽略参数缓存配置列表，参见CacheKeyInfo
	CacheKeyList []CacheKeyInfo

	// 缓存配置列表，参见CacheConf
	CacheList []CacheConf

	// 状态码缓存配置列表，参见CacheConf
	HttpCodeCacheList []CacheConf
}

/*
AccessControlConf - 访问控制配置参数
*/
type AccessControlConf struct {

	// ip黑名单，多个ip，可表示为：IpBlackList.0=1.1.1.1，IpBlackList.1=2.2.2.2
	IpBlackList []string

	// refer配置
	ReferConf ReferConf
}

/*
DomainConfigInfo - 更新域名配置
*/
type DomainConfigInfo struct {

	// 访问控制配置 参考AccessControlConf
	AccessControlConf AccessControlConf

	// 高级配置 参考AdvancedConf
	AdvancedConf AdvancedConf

	// 查询带宽区域 cn代表国内 abroad代表海外 all表示全部区域
	AreaCode string

	// 缓存配置 参考CacheAllConfig
	CacheConf CacheAllConfig

	// 加速域名的业务类型，web代表网站，stream代表视频 ，download 代表下载
	CdnType string

	// 国外证书名称
	CertNameAbroad string

	// 国内证书名称
	CertNameCn string

	// cdn域名。创建加速域名生成的cdn域名，用于设置CNAME记录
	Cname string

	// 域名创建的时间。格式：时间戳
	CreateTime int

	// 域名
	Domain string

	// 域名Id
	DomainId string

	// 国外https状态 enableing-开启中  fail-开启失败 enable-启用 disable-未启用
	HttpsStatusAbroad string

	// 国内https状态 enableing-开启中 fail-开启失败 enable-启用 disable-未启用
	HttpsStatusCn string

	// 源站配置 参考OriginConf
	OriginConf OriginConf

	// 创建的加速域名的当前的状态。check代表审核中；checkSuccess代表审核通过；checkFail代表审核失败；enable代表加速中；disable代表停止加速；delete代表删除加速；enableing代表正在开启加速；disableing代表正在停止加速中；deleteing代表删除中；deploying代表部署中
	Status string

	// 业务组：Default
	Tag string

	// 测试url。用于域名创建加速时的测试
	TestUrl string
}

/*
HitRateInfoV2 - HitRateInfoV2
*/
type HitRateInfoV2 struct {

	// 总流量命中率，单位%
	FlowHitRate float64

	// 请求数命中率，单位%
	RequestHitRate float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
HttpCodeInfoV2 - HttpCodeInfoV2
*/
type HttpCodeInfoV2 struct {

	// 1xx信息，参考HttpCodeV2Detail结构
	Http1XX HttpCodeV2Detail

	// 2xx信息，参考HttpCodeV2Detail结构
	Http2XX HttpCodeV2Detail

	// 3xx信息，参考HttpCodeV2Detail结构
	Http3XX HttpCodeV2Detail

	// 4xx信息，参考HttpCodeV2Detail结构
	Http4XX HttpCodeV2Detail

	// 5xx信息，参考HttpCodeV2Detail结构
	Http5XX HttpCodeV2Detail

	// 6xx信息，参考HttpCodeV2Detail结构
	Http6XX HttpCodeV2Detail

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
DomainBaseInfo - 域名基本信息
*/
type DomainBaseInfo struct {

	// 域名
	Domain string

	// 域名的资源id
	DomainId string
}

/*
LogSetInfo - 日志信息
*/
type LogSetInfo struct {

	// 国外日志url列表
	AbroadLog []string

	// 国内日志url列表
	CnLog []string

	// 日志时间UnixTime
	Time int
}

/*
LogSetList - 日志信息列表
*/
type LogSetList struct {

	// 域名
	Domain string

	// 域名信息列表，参考LogSetInfo
	Logs []LogSetInfo
}

/*
LogInfo - 日志信息
*/
type LogInfo struct {

	// Unix时间戳
	LogTime int

	// 日志url地址，多个URL用分号隔开
	LogUrl string
}

/*
DomanLogList - 域名日志列表
*/
type DomanLogList struct {

	// 域名
	Domain string

	// 日志信息列表
	LogList []LogInfo
}

/*
RequestInfoV2 - RequestInfoV2
*/
type RequestInfoV2 struct {

	// 返回值返回指定时间区间内的cdn收到的请求次数之和
	CdnRequest float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
UcdnDomainTrafficSet - GetUcdnDomainTraffic
*/
type UcdnDomainTrafficSet struct {

	// 流量获取的时间点，格式为Unix Timestamp
	Time int

	// 查询每日流量总值，单位：GB
	Value float64
}

/*
BandwidthInfoDetail - 带宽值信息模型(时间-带宽)
*/
type BandwidthInfoDetail struct {

	// 返回值带宽值数据。
	Bandwidth float64

	// 宽获取的时间点。格式：时间戳
	Time int
}

/*
ProIspBandwidthList - 省份带宽流量实例表
*/
type ProIspBandwidthList struct {

	// 返回值返回指定时间区间内CDN的带宽峰值，单位Mbps
	CdnBandwidth float64

	// 带宽获取的时间点。格式：时间戳
	Time int

	// 对应时间粒度的流量，单位字节
	Traffic float64
}

/*
ProIspBandwidthSet - 按省份的带宽流量实例表
*/
type ProIspBandwidthSet struct {

	// 省份带宽流量实例表
	BandwidthTrafficList []ProIspBandwidthList

	// 省份代码
	Province string
}

/*
ProIspRequestListV2 - 省份请求数实例表
*/
type ProIspRequestListV2 struct {

	// 返回值返回指定时间区间内的请求数
	CdnRequest float64

	// 带宽获取的时间点。格式：时间戳
	Time int
}

/*
ProIspRequestNumSetV2 - 按省份的请求数实例表
*/
type ProIspRequestNumSetV2 struct {

	// 省份代码
	Province string

	// 省份请求数实例表 ProIspRequestListV2
	RequestList []ProIspRequestListV2
}

/*
TrafficSet - GetUcdnTraffic
*/
type TrafficSet struct {

	// 购买流量的区域, cn: 国内; abroad: 国外
	Areacode string

	// Areacode区域内总剩余流量, 单位GB
	TrafficLeft float64

	// Areacode区域内总购买流量, 单位GB
	TrafficTotal float64

	// Areacode区域内总使用流量, 单位GB
	TrafficUsed float64
}

/*
IpLocationInfo - ip详细信息
*/
type IpLocationInfo struct {

	// 地区
	Area string

	// 城市
	City string

	// ip是否存在
	Exist bool

	// 客户端请求的ip
	Ip string

	// 运营商
	Isp string
}
