package umem

import (
	"github.com/ucloud/ucloud-sdk-go/private/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemRequest is request schema for DescribeUMem action
type DescribeUMemRequest struct {
	request.CommonBase

	// 协议类型: memcache, redis
	Protocol *string `required:"true"`

	// 分页显示的起始偏移, 默认值为0
	Offset *int `required:"false"`

	// 分页显示的条目数, 默认值为20
	Limit *int `required:"false"`

	// 资源ID
	ResourceId *string `required:"false"`

	//
	ResourceType *string `required:"false"`
}

// DescribeUMemResponse is response schema for DescribeUMem action
type DescribeUMemResponse struct {
	response.CommonBase

	// 根据过滤条件得到的总数
	TotalCount int

	// UMem实例列表, 详细参见UMemDataSet
	DataSet []UMemDataSet
}

// NewDescribeUMemRequest will create request of DescribeUMem action.
func (c *UMemClient) NewDescribeUMemRequest() *DescribeUMemRequest {
	req := &DescribeUMemRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMem - 获取UMem列表
func (c *UMemClient) DescribeUMem(req *DescribeUMemRequest) (*DescribeUMemResponse, error) {
	var err error
	var res DescribeUMemResponse

	err = c.Client.InvokeActionWithPatcher("DescribeUMem", req, &res, utils.PortPatcher)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
