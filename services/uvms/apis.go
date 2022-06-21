// Code is generated by ucloud-model, DO NOT EDIT IT.

package uvms

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UVMS API Schema

// SendUVMSMessageRequest is request schema for SendUVMSMessage action
type SendUVMSMessageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// 被叫号码，采用 E.164 标准，格式为+[国家代码][用户号码]。例如：+8613512345678， 其中前面有一个+号 ，86为国家码，13512345678为手机号
	CalledNumber *string `required:"true"`

	// 调度规则，0-默认（归属地优先），1-随机。当不指定外显号码（主叫号码为空）时生效。如不填写，默认为归属地优先。
	DispatchRule *int `required:"false"`

	// 主叫号码，号码随机时不填。专属号码时传入已购买的号码，仅支持一个号码，在控制台查看已购买的号码。
	FromNumber *string `required:"false"`

	// 号码组类型，1-随机组，2-专属组。如不填写则根据主叫号码判断，若主叫号码为空，则为随机组，若不为空，则为专属组。
	GroupType *int `required:"false"`

	// 模板 ID，在控制台审核通过的模板 ID。
	TemplateId *string `required:"true"`

	// 模板可变参数，以数组的方式填写，举例，TemplateParams.0，TemplateParams.1，... 若模板中无可变参数，则该项可不填写；若模板中有可变参数，则该项为必填项，参数个数需与变量个数保持一致，否则无法发送；
	TemplateParams []string `required:"false"`

	// 自定义的业务标识ID，字符串（ 长度不能超过32 位），不支持 单引号、表情包符号等特殊字符
	UserId *string `required:"false"`
}

// SendUVMSMessageResponse is response schema for SendUVMSMessage action
type SendUVMSMessageResponse struct {
	response.CommonBase

	// 状态码的描述
	Message string

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 ReqUuid。
	ReqUuid string

	// 本次提交发送语音的唯一ID，可根据该值查询本次发送详情
	SessionNo string

	// 本次提交的自定义业务标识ID，仅当发送时传入有效的UserId，才返回该字段。
	UserId string
}

// NewSendUVMSMessageRequest will create request of SendUVMSMessage action.
func (c *UVMSClient) NewSendUVMSMessageRequest() *SendUVMSMessageRequest {
	req := &SendUVMSMessageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: SendUVMSMessage

向指定号码拨打电话
*/
func (c *UVMSClient) SendUVMSMessage(req *SendUVMSMessageRequest) (*SendUVMSMessageResponse, error) {
	var err error
	var res SendUVMSMessageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SendUVMSMessage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
