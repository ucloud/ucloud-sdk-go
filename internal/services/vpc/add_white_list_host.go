// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// AddWhiteListHostRequest is request schema for AddWhiteListHost action
type AddWhiteListHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NATGateWay Id
	NATGWId *string `required:"true"`

	// 主机资源短ID
	VmIds []string `required:"true"`
}

// AddWhiteListHostResponse is response schema for AddWhiteListHost action
type AddWhiteListHostResponse struct {
	response.CommonBase
}

// NewAddWhiteListHostRequest will create request of AddWhiteListHost action.
func (c *VPCClient) NewAddWhiteListHostRequest() *AddWhiteListHostRequest {
	req := &AddWhiteListHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// AddWhiteListHost - 将指定资源添加到nat网关白名单中，只可添加云主机。
func (c *VPCClient) AddWhiteListHost(req *AddWhiteListHostRequest) (*AddWhiteListHostResponse, error) {
	var err error
	var res AddWhiteListHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddWhiteListHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
