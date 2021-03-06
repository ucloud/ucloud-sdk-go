//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS DescribeUFSVPCSubnet

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUFSVPCSubnetRequest is request schema for DescribeUFSVPCSubnet action
type DescribeUFSVPCSubnetRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"true"`
}

// DescribeUFSVPCSubnetResponse is response schema for DescribeUFSVPCSubnet action
type DescribeUFSVPCSubnetResponse struct {
	response.CommonBase

	// 用户的vpc信息
	DataSet []VPCInfo
}

// NewDescribeUFSVPCSubnetRequest will create request of DescribeUFSVPCSubnet action.
func (c *UFSClient) NewDescribeUFSVPCSubnetRequest() *DescribeUFSVPCSubnetRequest {
	req := &DescribeUFSVPCSubnetRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUFSVPCSubnet - 展示可供文件系统挂载的VPC和SUBNET
func (c *UFSClient) DescribeUFSVPCSubnet(req *DescribeUFSVPCSubnetRequest) (*DescribeUFSVPCSubnetResponse, error) {
	var err error
	var res DescribeUFSVPCSubnetResponse

	err = c.Client.InvokeAction("DescribeUFSVPCSubnet", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
