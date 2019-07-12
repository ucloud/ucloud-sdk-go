//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api PathX DeleteUGAInstance

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteUGAInstanceRequest is request schema for DeleteUGAInstance action
type DeleteUGAInstanceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID,如org-xxxx。请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 全球加速实例ID
	UGAId *string `required:"true"`
}

// DeleteUGAInstanceResponse is response schema for DeleteUGAInstance action
type DeleteUGAInstanceResponse struct {
	response.CommonBase
}

// NewDeleteUGAInstanceRequest will create request of DeleteUGAInstance action.
func (c *PathXClient) NewDeleteUGAInstanceRequest() *DeleteUGAInstanceRequest {
	req := &DeleteUGAInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteUGAInstance - 删除全球加速服务加速配置
func (c *PathXClient) DeleteUGAInstance(req *DeleteUGAInstanceRequest) (*DeleteUGAInstanceResponse, error) {
	var err error
	var res DeleteUGAInstanceResponse

	err = c.Client.InvokeAction("DeleteUGAInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
