package umem

import (
	"github.com/ucloud/ucloud-sdk-go/private/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemSpaceRequest is request schema for DescribeUMemSpace action
type DescribeUMemSpaceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 返回数据长度, 默认为20
	Limit *int `required:"false"`

	// 内存空间ID (无ID，则获取所有)
	SpaceId *string `required:"false"`
}

// DescribeUMemSpaceResponse is response schema for DescribeUMemSpace action
type DescribeUMemSpaceResponse struct {
	response.CommonBase

	// JSON 格式的UMem内存空间实例列表, 详细参见 UMemSpaceSet
	DataSet []UMemSpaceSet

	// 根据过滤条件得到的总数
	TotalCount int
}

// NewDescribeUMemSpaceRequest will create request of DescribeUMemSpace action.
func (c *UMemClient) NewDescribeUMemSpaceRequest() *DescribeUMemSpaceRequest {
	req := &DescribeUMemSpaceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemSpace - 获取UMem内存空间列表
func (c *UMemClient) DescribeUMemSpace(req *DescribeUMemSpaceRequest) (*DescribeUMemSpaceResponse, error) {
	var err error
	var res DescribeUMemSpaceResponse

	err = c.client.InvokeActionWithPatcher("DescribeUMemSpace", req, &res, utils.PortPatcher)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
