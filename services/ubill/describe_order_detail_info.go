package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type DescribeOrderDetailInfoRequest struct {
	request.CommonBase

	//开始时间，UNIX time format
	BeginTime *int `required:"true"`
	//结束时间，时间跨度不超过3个月
	EndTime *int     `required:"true"`
	Regions []string `required:"true"`
	//产品类型，默认全部的产品类型
	ResourceTypes []string `required:"false"`
	//订单类型，默认全部订单类型
	OrderTypes []string `required:"false"`
	//付费类型，默认全部的计费方式
	ChargeTypes []string
	//订单状态，默认选中全部的可选参数
	OrderStates []string `required:"false"`
	//是否开过发票，默认选中全部的可选参数
	Invoiceds []string `required:"false"`
	//返回数据长度，默认25
	Limit *int `required:"false"`
	//数据偏移量，默认0
	Offset *int `required:"false"`
	//资源ID
	ResourceIDs []string `required:"false"`
	//交易单号，该字段存在时，可以不传BeginTime和EndTime
	TradeNos []string `required:"false"`
	//true表示查询全部，默认全部，其他选项按照项目自查询
	QueryAll *string `required:"false"`
}

type DescribeOrderDetailInfoResponse struct {
	response.CommonBase

	//json格式的订单信息
	OrderInfo []OrderInfo `json:"order_info"`
}

func (c *UBillClient) NewDescribeOrderDetailInfoRequest() *DescribeOrderDetailInfoRequest {
	req := &DescribeOrderDetailInfoRequest{}

	//设置client请求配置config
	c.client.SetupRequest(req)

	//设置默认重试
	req.SetRetryable(true)
	return req
}

func (c *UBillClient) DescribeOrderDetailInfo(req *DescribeOrderDetailInfoRequest) (*DescribeOrderDetailInfoResponse, error) {
	var err error
	var res DescribeOrderDetailInfoResponse

	err = c.client.InvokeAction("DescribeOrderDetailInfo", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
