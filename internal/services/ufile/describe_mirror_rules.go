// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeMirrorRulesRequest is request schema for DescribeMirrorRules action
type DescribeMirrorRulesRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 获取回源规则的bucket的名称
	BucketName *string `required:"true"`
}

// DescribeMirrorRulesResponse is response schema for DescribeMirrorRules action
type DescribeMirrorRulesResponse struct {
	response.CommonBase

	// 回源规则信息
	Mirrors []Mirrors
}

// NewDescribeMirrorRulesRequest will create request of DescribeMirrorRules action.
func (c *UFileClient) NewDescribeMirrorRulesRequest() *DescribeMirrorRulesRequest {
	req := &DescribeMirrorRulesRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeMirrorRules - 获取回源规则描述信息
func (c *UFileClient) DescribeMirrorRules(req *DescribeMirrorRulesRequest) (*DescribeMirrorRulesResponse, error) {
	var err error
	var res DescribeMirrorRulesResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeMirrorRules", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
