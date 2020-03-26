package ubill

/*
OrderInfo - DescribeOrderDetailInfo
*/
type OrderInfo struct {
	OrderNo      string `json:"OrderNo"`
	OrderType    string
	ChargeType   string
	ResourceId   string
	ResourceTag  []ResourceTag
	OrderState   string
	CreateTime   int
	Amount       string //订单总金额
	AmountReal   string //现金账户金额
	AmountFree   string //赠送账户金额（元）
	AmountCoupon string //优惠金额
	CouponCode   string //如果AmountCoupon不为0，显示代金券号码
	ResourceType string
	UpdateTime   int
	Quantity     string //计费周期数
	Count        int    //资源数量
	Invoiced     string //是否开过发票
	StartTime    int    //开始时间
	EndTime      int    //结束时间
	OrderDetails []OrderDetail
	RegionId     string //region
	TradeNo      string //交易号
}
