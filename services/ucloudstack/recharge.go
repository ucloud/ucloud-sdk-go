// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RechargeRequest is request schema for Recharge action
type RechargeRequest struct {
	request.CommonBase

	// 充值金额。最少100,最大500000
	Amount *int `required:"true"`

	// 充值来源。INPOUR_FROM_ALIPAY：支付宝，INPOUR_FROM_OFFLINE：银行转账，INPOUR_FROM_SINPAY：新浪支付，INPOUR_FROM_WECHAT_PAY：微信转账。
	FromType *string `required:"true"`

	// 充值单号。充值方式为“账户余额”时为必要参数。
	SerialNo *string `required:"true"`

	// 租户的账户ID。
	UserID *int `required:"true"`
}

// RechargeResponse is response schema for Recharge action
type RechargeResponse struct {
	response.CommonBase

	// 返回信息描述。
	Message string
}

// NewRechargeRequest will create request of Recharge action.
func (c *UCloudStackClient) NewRechargeRequest() *RechargeRequest {
	req := &RechargeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// Recharge - UCloudStack管理员给租户充值
func (c *UCloudStackClient) Recharge(req *RechargeRequest) (*RechargeResponse, error) {
	var err error
	var res RechargeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("Recharge", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
