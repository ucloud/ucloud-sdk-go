// Code is generated by ucloud-model, DO NOT EDIT IT.

package ubill

/*
ResourceTag -
*/
type ResourceTag struct {

	//
	KeyId string

	//
	Value string
}

/*
OrderDetail -
*/
type OrderDetail struct {

	//
	ProductName string

	//
	Value string
}

/*
OrderInfo -
*/
type OrderInfo struct {

	//
	Amount string

	//
	AmountCoupon string

	//
	AmountFree string

	//
	AmountReal string

	//
	ChargeType string

	//
	Count int

	//
	CouponCode string

	//
	CreateTime int

	//
	EndTime int

	//
	Invoiced string

	//
	OrderDetails []OrderDetail

	//
	OrderNo string

	//
	OrderState string

	//
	OrderType string

	//
	Quantity string

	//
	RegionId string

	//
	ResourceId string

	//
	ResourceTag []ResourceTag

	//
	ResourceType string

	//
	StartTime int

	//
	TradeNo string

	//
	UpdateTime int
}

/*
AccountInfo - 账户信息
*/
type AccountInfo struct {

	// 账户余额
	Amount string

	// 账户可用余额
	AmountAvailable string

	// 信用账户余额
	AmountCredit string

	// 赠送账户余额
	AmountFree string

	// 冻结账户金额
	AmountFreeze string
}

/*
ItemDetail - 产品配置
*/
type ItemDetail struct {

	// 产品小类名称
	ProductName string

	// 产品小类规格
	Value string
}

/*
ResourceExtendInfo - 资源标识
*/
type ResourceExtendInfo struct {

	// 资源标识健
	KeyId string

	// 资源标识值
	Value string
}

/*
BillDetailItem - 账单详情数据
*/
type BillDetailItem struct {

	// 是否为主账号。枚举值：\\ > 0:子账号 \\ > 1:主账号
	Admin int

	// 订单总金额
	Amount string

	// 代金券抵扣
	AmountCoupon string

	// 赠送金额抵扣
	AmountFree string

	// 现金账户支付
	AmountReal string

	// 可用区
	AzGroupCName string

	// 计费方式 (筛选项, 默认全部)。枚举值：\\ > Dynamic:按时 \\ > Month:按月 \\ > Year:按年 \\ > Once:一次性按量 \\ > Used:按量 \\ > Post:后付费
	ChargeType string

	// 创建时间（时间戳）
	CreateTime int

	// 产品配置
	ItemDetails []ItemDetail

	// 订单号
	OrderNo string

	// 订单类型 (筛选项, 默认全部) 。枚举值：\\ > OT_BUY:新购 \\ > OT_RENEW:续费 \\ > OT_UPGRADE:升级 \\ > OT_REFUND:退费 \\ > OT_DOWNGRADE:降级 \\ > OT_SUSPEND:结算 \\ > OT_PAYMENT:删除资源回款 \\ > OT_POSTPAID_PAYMENT:后付费回款 \\ > OT_RECOVER:删除恢复 \\ > OT_POSTPAID_RENEW:过期续费回款
	OrderType string

	// 项目名称
	ProjectName string

	// 资源标识
	ResourceExtendInfo []ResourceExtendInfo

	// 资源ID
	ResourceId string

	// 产品类型。枚举值：\\ > uhost:云主机 \\ > udisk:普通云硬盘 \\ > udb:云数据库 \\ > eip:弹性IP \\ > ufile:对象存储 \\ > fortress_host:堡垒机 \\ > ufs:文件存储 \\ > waf:WEB应用防火墙 \\ > ues:弹性搜索 \\ > udisk_ssd:SSD云硬盘 \\ > rssd:RSSD云硬盘
	ResourceType string

	// 产品类型代码
	ResourceTypeCode int

	// 订单支付状态。枚举值：\\> 0:未支付 \\ > 1:已支付
	ShowHover int

	// 开始时间（时间戳）
	StartTime int

	// 账户昵称
	UserDisplayName string

	// 账户邮箱
	UserEmail string

	// 账户名
	UserName string
}

/*
BillOverviewItem - 账单总览数组内单个结构体数据
*/
type BillOverviewItem struct {

	// 该账户是否为主账号，1 主账号，0 子账号（账单维度按子账号筛选时显示）
	Admin int

	// 订单总金额
	Amount string

	// 代金券抵扣（已入账时显示）
	AmountCoupon string

	// 赠送金额抵扣（已入账时显示）
	AmountFree string

	// 现金账户支付（已入账时显示）
	AmountReal string

	// 账单维度, product 按产品维度聚合,project 按项目维度聚合，user 按子账号维度聚合
	Dimension string

	// 产品分类	（账单维度按产品筛选时显示）
	ProductCategory string

	// 项目名称（账单维度按项目筛选时显示）
	ProjectName string

	// 产品类型	（账单维度按产品筛选时显示）
	ResourceType string

	// 产品类型代码（账单维度按产品筛选时显示）
	ResourceTypeCode int

	// 账户昵称（账单维度按子账号筛选时显示）
	UserDisplayName string

	// 账户邮箱（账单维度按子账号筛选时显示）
	UserEmail string

	// 账户名   （账单维度按子账号筛选时显示）
	UserName string
}
