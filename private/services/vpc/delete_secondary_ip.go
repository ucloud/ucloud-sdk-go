//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC DeleteSecondaryIp

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSecondaryIpRequest is request schema for DeleteSecondaryIp action
type DeleteSecondaryIpRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// ip
	Ip *string `required:"true"`

	// mac
	Mac *string `required:"true"`

	// 子网Id
	SubnetId *string `required:"true"`

	// VPCId
	VPCId *string `required:"false"`

	// 资源Id
	ObjectId *string `required:"false"`
}

// DeleteSecondaryIpResponse is response schema for DeleteSecondaryIp action
type DeleteSecondaryIpResponse struct {
	response.CommonBase
}

// NewDeleteSecondaryIpRequest will create request of DeleteSecondaryIp action.
func (c *VPCClient) NewDeleteSecondaryIpRequest() *DeleteSecondaryIpRequest {
	req := &DeleteSecondaryIpRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSecondaryIp - 删除ip（用于uk8s使用）
func (c *VPCClient) DeleteSecondaryIp(req *DeleteSecondaryIpRequest) (*DeleteSecondaryIpResponse, error) {
	var err error
	var res DeleteSecondaryIpResponse

	err = c.client.InvokeAction("DeleteSecondaryIp", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
