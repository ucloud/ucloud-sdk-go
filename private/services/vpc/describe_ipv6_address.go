// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeIPv6AddressRequest is request schema for DescribeIPv6Address action
type DescribeIPv6AddressRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// IPv6地址ID。默认为该项目下所有IPv6地址信息。
	IPv6AddressId []string `required:"false"`

	// 数据分页值，默认20
	Limit *int `required:"false"`

	// 数据偏移量，默认0
	Offset *int `required:"false"`
}

// DescribeIPv6AddressResponse is response schema for DescribeIPv6Address action
type DescribeIPv6AddressResponse struct {
	response.CommonBase

	// IPv6地址详情。
	IPv6AddressSet []IPv6AddressInfo

	// IPv6地址总数。
	TotalCount int
}

// NewDescribeIPv6AddressRequest will create request of DescribeIPv6Address action.
func (c *VPCClient) NewDescribeIPv6AddressRequest() *DescribeIPv6AddressRequest {
	req := &DescribeIPv6AddressRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeIPv6Address - 获取IPv6地址详情
func (c *VPCClient) DescribeIPv6Address(req *DescribeIPv6AddressRequest) (*DescribeIPv6AddressResponse, error) {
	var err error
	var res DescribeIPv6AddressResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeIPv6Address", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
