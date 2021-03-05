package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateBackendBatchRequest is request schema for UpdateBackendBatch action
type UpdateBackendBatchRequest struct {
	request.CommonBase
	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string    `required:"true"`
	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string    `required:"true"`

	// 负载均衡资源ID
	ULBId *string `required:"true"`
	//
	Attributes []UpdateBackendBatchParamAttributes
}

/*
   Attributes.N is request schema for complex param
*/
type UpdateBackendBatchParamAttributes struct {

	// 后端资源实例的ID(ULB后端ID，非资源自身ID)
	BackendId *string `required:"true"`
	// 后端资源服务端口，取值范围[1-65535]
	Port *int `required:"false"`
	// 所添加的后端RS权重（在加权轮询算法下有效），取值范围[0-100]，默认为1
	Weight *int `required:"false"`
	// 后端实例状态开关
	Enabled *int `required:"false"`
	// 是否为backup 0：主rs 1：备rs 默认为0
	IsBackup *int `required:"false"`
}

// UpdateBackendBatchResponse is response schema for UpdateBackendBatch action
type UpdateBackendBatchResponse struct {
	response.CommonBase

	// 返回ulb对每个rs的修改是否成功
	BackendSet []BackendMsg `required:"false"`
}

// NewUpdateBackendBatchRequest will create request of UpdateBackendBatch action.
func (c *ULBClient) NewUpdateBackendBatchRequest() *UpdateBackendBatchRequest {
	req := &UpdateBackendBatchRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)

	// set encoder for `application/json`
	req.SetEncoder(request.NewJSONEncoder(c.GetConfig(), c.GetCredential()))
	return req
}

/*
   API: UpdateBackendBatch

   批量更新ULB后端资源实例(服务节点)属性
*/
func (c *ULBClient) UpdateBackendBatch(req *UpdateBackendBatchRequest) (*UpdateBackendBatchResponse, error) {
	var err error
	var res UpdateBackendBatchResponse

	reqCopier := *req
	err = c.Client.InvokeAction("UpdateBackendBatch", &reqCopier, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
