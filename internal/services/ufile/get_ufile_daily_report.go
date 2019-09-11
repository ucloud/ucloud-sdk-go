// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetUFileDailyReportRequest is request schema for GetUFileDailyReport action
type GetUFileDailyReportRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 空间名称。此字段不为空，返回此Bucket日使用量；否则，返回这个项目的日使用量
	BucketName *string `required:"false"`

	// 查询结束时间;unix时间戳,单位s
	EndTime *int `required:"true"`

	// 查询开始时间;unix时间戳，单位s
	StartTime *int `required:"true"`
}

// GetUFileDailyReportResponse is response schema for GetUFileDailyReport action
type GetUFileDailyReportResponse struct {
	response.CommonBase

	// 消费情况
	DataSet []UFileReportItem
}

// NewGetUFileDailyReportRequest will create request of GetUFileDailyReport action.
func (c *UFileClient) NewGetUFileDailyReportRequest() *GetUFileDailyReportRequest {
	req := &GetUFileDailyReportRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUFileDailyReport - 查看日消费报表
func (c *UFileClient) GetUFileDailyReport(req *GetUFileDailyReportRequest) (*GetUFileDailyReportResponse, error) {
	var err error
	var res GetUFileDailyReportResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUFileDailyReport", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
