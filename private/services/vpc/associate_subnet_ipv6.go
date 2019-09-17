// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// AssociateSubnetIPv6Request is request schema for AssociateSubnetIPv6 action
type AssociateSubnetIPv6Request struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 子网关联的IPv6网段，固定为/64掩码。该参数不填写，则默认从VPC关联的IPv6网段中分配。
	IPv6Network *string `required:"false"`

	// 需要关联IPv6网段的子网
	SubnetworkId *string `required:"true"`
}

// AssociateSubnetIPv6Response is response schema for AssociateSubnetIPv6 action
type AssociateSubnetIPv6Response struct {
	response.CommonBase

	// 子网关联的IPv6网段
	IPv6Network string
}

// NewAssociateSubnetIPv6Request will create request of AssociateSubnetIPv6 action.
func (c *VPCClient) NewAssociateSubnetIPv6Request() *AssociateSubnetIPv6Request {
	req := &AssociateSubnetIPv6Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// AssociateSubnetIPv6 - 子网关联IPv6网段
func (c *VPCClient) AssociateSubnetIPv6(req *AssociateSubnetIPv6Request) (*AssociateSubnetIPv6Response, error) {
	var err error
	var res AssociateSubnetIPv6Response

	reqCopier := *req

	err = c.Client.InvokeAction("AssociateSubnetIPv6", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
