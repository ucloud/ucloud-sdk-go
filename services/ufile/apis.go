// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UFile API Schema

// CreateBucketRequest is request schema for CreateBucket action
type CreateBucketRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 待创建Bucket的名称，具有全局唯一性
	BucketName *string `required:"true"`

	// Bucket访问类型，public或private; 默认为private
	Type *string `required:"false"`
}

// CreateBucketResponse is response schema for CreateBucket action
type CreateBucketResponse struct {
	response.CommonBase

	// 已创建Bucket的ID
	BucketId string

	// 已创建Bucket的名称
	BucketName string
}

// NewCreateBucketRequest will create request of CreateBucket action.
func (c *UFileClient) NewCreateBucketRequest() *CreateBucketRequest {
	req := &CreateBucketRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateBucket

创建Bucket
*/
func (c *UFileClient) CreateBucket(req *CreateBucketRequest) (*CreateBucketResponse, error) {
	var err error
	var res CreateBucketResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateBucket", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUFileLifeCycleRequest is request schema for CreateUFileLifeCycle action
type CreateUFileLifeCycleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后，自动变为归档存储类型；参数范围：[7,36500]，0代表不启用
	ArchivalDays *int `required:"false"`

	// 存储空间名称
	BucketName *string `required:"true"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后过期，自动删除；参数范围：[7,36500]，0代表不启用
	Days *int `required:"false"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后，自动变为低频存储类型；参数范围：[7,36500]，0代表不启用
	IADays *int `required:"false"`

	// 生命周期名称
	LifeCycleName *string `required:"true"`

	// 生命周期所适用的前缀；*为整个存储空间文件；一条规则只支持一个文件前缀；
	Prefix *string `required:"true"`

	// Enabled -- 启用，Disabled -- 不启用
	Status *string `required:"true"`
}

// CreateUFileLifeCycleResponse is response schema for CreateUFileLifeCycle action
type CreateUFileLifeCycleResponse struct {
	response.CommonBase

	// 生命周期Id
	LifeCycleId string
}

// NewCreateUFileLifeCycleRequest will create request of CreateUFileLifeCycle action.
func (c *UFileClient) NewCreateUFileLifeCycleRequest() *CreateUFileLifeCycleRequest {
	req := &CreateUFileLifeCycleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUFileLifeCycle

创建生命周期管理
*/
func (c *UFileClient) CreateUFileLifeCycle(req *CreateUFileLifeCycleRequest) (*CreateUFileLifeCycleResponse, error) {
	var err error
	var res CreateUFileLifeCycleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUFileLifeCycle", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUFileTokenRequest is request schema for CreateUFileToken action
type CreateUFileTokenRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 令牌允许操作的bucket，默认*表示全部
	AllowedBuckets []string `required:"false"`

	// 令牌允许执行的操作，[ TOKEN_ALLOW_NONE , TOKEN_ALLOW_READ , TOKEN_ALLOW_WRITE , TOKEN_ALLOW_DELETE , TOKEN_ALLOW_LIST, TOKEN_ALLOW_IOP , TOKEN_ALLOW_DP  ，TOKEN_DENY_UPDATE]。默认TOKEN_ALLOW_NONE
	AllowedOps []string `required:"false"`

	// 令牌允许操作的key前缀，默认*表示全部
	AllowedPrefixes []string `required:"false"`

	// 令牌黑名单，支持ipv4，ipv6格式。
	BlackIPList []string `required:"false"`

	// Unix 时间戳，精确到秒，为令牌过期时间点。默认过期时间为一天（即当前Unix时间戳+86400）；注意：过期时间不能超过 4102416000
	ExpireTime *int `required:"false"`

	// 令牌名称
	TokenName *string `required:"true"`

	// 令牌白名单，支持ipv4，ipv6格式。
	WhiteIPList []string `required:"false"`
}

// CreateUFileTokenResponse is response schema for CreateUFileToken action
type CreateUFileTokenResponse struct {
	response.CommonBase

	// 令牌唯一ID
	TokenId string

	// 创建令牌的详细信息
	UFileTokenSet UFileTokenSet
}

// NewCreateUFileTokenRequest will create request of CreateUFileToken action.
func (c *UFileClient) NewCreateUFileTokenRequest() *CreateUFileTokenRequest {
	req := &CreateUFileTokenRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUFileToken

创建US3令牌
*/
func (c *UFileClient) CreateUFileToken(req *CreateUFileTokenRequest) (*CreateUFileTokenResponse, error) {
	var err error
	var res CreateUFileTokenResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUFileToken", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteBucketRequest is request schema for DeleteBucket action
type DeleteBucketRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 待删除Bucket的名称
	BucketName *string `required:"true"`
}

// DeleteBucketResponse is response schema for DeleteBucket action
type DeleteBucketResponse struct {
	response.CommonBase

	// Bucket的ID
	BucketId string

	// Bucket的名称
	BucketName string
}

// NewDeleteBucketRequest will create request of DeleteBucket action.
func (c *UFileClient) NewDeleteBucketRequest() *DeleteBucketRequest {
	req := &DeleteBucketRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteBucket

删除Bucket
*/
func (c *UFileClient) DeleteBucket(req *DeleteBucketRequest) (*DeleteBucketResponse, error) {
	var err error
	var res DeleteBucketResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteBucket", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUFileLifeCycleRequest is request schema for DeleteUFileLifeCycle action
type DeleteUFileLifeCycleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 存储空间名称
	BucketName *string `required:"true"`

	// 生命周期Id
	LifeCycleId *string `required:"true"`
}

// DeleteUFileLifeCycleResponse is response schema for DeleteUFileLifeCycle action
type DeleteUFileLifeCycleResponse struct {
	response.CommonBase
}

// NewDeleteUFileLifeCycleRequest will create request of DeleteUFileLifeCycle action.
func (c *UFileClient) NewDeleteUFileLifeCycleRequest() *DeleteUFileLifeCycleRequest {
	req := &DeleteUFileLifeCycleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUFileLifeCycle

删除生命周期管理
*/
func (c *UFileClient) DeleteUFileLifeCycle(req *DeleteUFileLifeCycleRequest) (*DeleteUFileLifeCycleResponse, error) {
	var err error
	var res DeleteUFileLifeCycleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUFileLifeCycle", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUFileTokenRequest is request schema for DeleteUFileToken action
type DeleteUFileTokenRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 令牌ID
	TokenId *string `required:"true"`
}

// DeleteUFileTokenResponse is response schema for DeleteUFileToken action
type DeleteUFileTokenResponse struct {
	response.CommonBase
}

// NewDeleteUFileTokenRequest will create request of DeleteUFileToken action.
func (c *UFileClient) NewDeleteUFileTokenRequest() *DeleteUFileTokenRequest {
	req := &DeleteUFileTokenRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUFileToken

删除令牌
*/
func (c *UFileClient) DeleteUFileToken(req *DeleteUFileTokenRequest) (*DeleteUFileTokenResponse, error) {
	var err error
	var res DeleteUFileTokenResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUFileToken", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeBucketRequest is request schema for DescribeBucket action
type DescribeBucketRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 如果提供此参数，则获取相应地域下所有空间的空间名称(只返回空间名称信息)
	// Region *string `required:"false"`

	// 待获取Bucket的名称，若不提供，则获取所有Bucket
	BucketName *string `required:"false"`

	// 获取所有Bucket列表的限制数目，默认为20
	Limit *int `required:"false"`

	// 获取所有Bucket列表的偏移数目，默认为0
	Offset *int `required:"false"`
}

// DescribeBucketResponse is response schema for DescribeBucket action
type DescribeBucketResponse struct {
	response.CommonBase

	// Bucket的描述信息 参数见 UFileBucketSet
	DataSet []UFileBucketSet
}

// NewDescribeBucketRequest will create request of DescribeBucket action.
func (c *UFileClient) NewDescribeBucketRequest() *DescribeBucketRequest {
	req := &DescribeBucketRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeBucket

获取Bucket的描述信息
*/
func (c *UFileClient) DescribeBucket(req *DescribeBucketRequest) (*DescribeBucketResponse, error) {
	var err error
	var res DescribeBucketResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeBucket", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUFileLifeCycleRequest is request schema for DescribeUFileLifeCycle action
type DescribeUFileLifeCycleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 存储空间名称
	BucketName *string `required:"true"`

	// 生命周期Id；不传递此参数拉取存储空间下面的所有生命周期信息
	LifeCycleId *string `required:"false"`
}

// DescribeUFileLifeCycleResponse is response schema for DescribeUFileLifeCycle action
type DescribeUFileLifeCycleResponse struct {
	response.CommonBase

	// 生命周期信息
	DateSet []LifeCycleItem
}

// NewDescribeUFileLifeCycleRequest will create request of DescribeUFileLifeCycle action.
func (c *UFileClient) NewDescribeUFileLifeCycleRequest() *DescribeUFileLifeCycleRequest {
	req := &DescribeUFileLifeCycleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUFileLifeCycle

获取生命周期信息
*/
func (c *UFileClient) DescribeUFileLifeCycle(req *DescribeUFileLifeCycleRequest) (*DescribeUFileLifeCycleResponse, error) {
	var err error
	var res DescribeUFileLifeCycleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUFileLifeCycle", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUFileTokenRequest is request schema for DescribeUFileToken action
type DescribeUFileTokenRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// 0表示显示部分token信息；不传递和其他情况表示显示全部token信息
	Display *int `required:"false"`

	// 令牌ID，只返回指定ID信息，否则拉取所有令牌
	TokenId *string `required:"false"`

	// 令牌名称，只返回指定令牌名称信息，否则拉取所有令牌
	TokenName *string `required:"false"`
}

// DescribeUFileTokenResponse is response schema for DescribeUFileToken action
type DescribeUFileTokenResponse struct {
	response.CommonBase

	// 【该字段已废弃，请谨慎使用】
	Action string `deprecated:"true"`

	// 令牌描述信息
	DataSet []UFileTokenSet

	// 【该字段已废弃，请谨慎使用】
	RetCode int `deprecated:"true"`
}

// NewDescribeUFileTokenRequest will create request of DescribeUFileToken action.
func (c *UFileClient) NewDescribeUFileTokenRequest() *DescribeUFileTokenRequest {
	req := &DescribeUFileTokenRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUFileToken

获取令牌信息
*/
func (c *UFileClient) DescribeUFileToken(req *DescribeUFileTokenRequest) (*DescribeUFileTokenResponse, error) {
	var err error
	var res DescribeUFileTokenResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUFileToken", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUFileDailyReportRequest is request schema for GetUFileDailyReport action
type GetUFileDailyReportRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
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

/*
API: GetUFileDailyReport

查看日消费报表
*/
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

// GetUFileQuotaRequest is request schema for GetUFileQuota action
type GetUFileQuotaRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 配额类型，取值为storage-volume, download-traffic或request-count
	QuotaType *string `required:"true"`
}

// GetUFileQuotaResponse is response schema for GetUFileQuota action
type GetUFileQuotaResponse struct {
	response.CommonBase

	// 剩余的配额数值
	LeftQuota float64
}

// NewGetUFileQuotaRequest will create request of GetUFileQuota action.
func (c *UFileClient) NewGetUFileQuotaRequest() *GetUFileQuotaRequest {
	req := &GetUFileQuotaRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUFileQuota

查看配额状态
*/
func (c *UFileClient) GetUFileQuota(req *GetUFileQuotaRequest) (*GetUFileQuotaResponse, error) {
	var err error
	var res GetUFileQuotaResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUFileQuota", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUFileQuotaInfoRequest is request schema for GetUFileQuotaInfo action
type GetUFileQuotaInfoRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 配额类型，取值为storage-volume, download-traffic或request-count
	QuotaType []string `required:"true"`
}

// GetUFileQuotaInfoResponse is response schema for GetUFileQuotaInfo action
type GetUFileQuotaInfoResponse struct {
	response.CommonBase

	// 配额信息数据集
	DataSet []UFileQuotaDataSetItem
}

// NewGetUFileQuotaInfoRequest will create request of GetUFileQuotaInfo action.
func (c *UFileClient) NewGetUFileQuotaInfoRequest() *GetUFileQuotaInfoRequest {
	req := &GetUFileQuotaInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUFileQuotaInfo

获取配额信息
*/
func (c *UFileClient) GetUFileQuotaInfo(req *GetUFileQuotaInfoRequest) (*GetUFileQuotaInfoResponse, error) {
	var err error
	var res GetUFileQuotaInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUFileQuotaInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUFileQuotaPriceRequest is request schema for GetUFileQuotaPrice action
type GetUFileQuotaPriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 下载流量，单位: GB，范围: [0, 60 000]，步长：1GB
	DownloadTraffic *int `required:"false"`

	// 请求次数，单位：万次，范围：[0, 1 000 000]，步长：1万次
	RequestCount *int `required:"false"`

	// 存储容量，单位: GB*天，范围: [0, 30 000 000]，步长：100GB*天
	StorageVolume *int `required:"false"`
}

// GetUFileQuotaPriceResponse is response schema for GetUFileQuotaPrice action
type GetUFileQuotaPriceResponse struct {
	response.CommonBase

	// 待支付价格，单位：分
	Price float64
}

// NewGetUFileQuotaPriceRequest will create request of GetUFileQuotaPrice action.
func (c *UFileClient) NewGetUFileQuotaPriceRequest() *GetUFileQuotaPriceRequest {
	req := &GetUFileQuotaPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUFileQuotaPrice

根据US3的购买配额，查询需要支付的价格。
*/
func (c *UFileClient) GetUFileQuotaPrice(req *GetUFileQuotaPriceRequest) (*GetUFileQuotaPriceResponse, error) {
	var err error
	var res GetUFileQuotaPriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUFileQuotaPrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUFileReportRequest is request schema for GetUFileReport action
type GetUFileReportRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 查询结束时间
	EndTime *int `required:"true"`

	// 查询开始时间
	StartTime *int `required:"true"`
}

// GetUFileReportResponse is response schema for GetUFileReport action
type GetUFileReportResponse struct {
	response.CommonBase

	// 报表内容 参数见 UFileReportSet
	DataSet []UFileReportSet
}

// NewGetUFileReportRequest will create request of GetUFileReport action.
func (c *UFileClient) NewGetUFileReportRequest() *GetUFileReportRequest {
	req := &GetUFileReportRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUFileReport

查看配额使用报表
*/
func (c *UFileClient) GetUFileReport(req *GetUFileReportRequest) (*GetUFileReportResponse, error) {
	var err error
	var res GetUFileReportResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUFileReport", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// SetUFileRefererRequest is request schema for SetUFileReferer action
type SetUFileRefererRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 存储空间名称
	BucketName *string `required:"true"`

	// RefererType为白名单时，RefererAllowNull为false代表不允许空referer访问，为true代表允许空referer访问;此参数默认为 true;
	RefererAllowNull *bool `required:"false"`

	// 开启关闭referer防盗链;关闭防盗链会清空防盗链参数设置，开启防盗链必须指定 RefererType、Referers；开启：on， 关闭：off;
	RefererStatus *string `required:"true"`

	// 防盗链Referer类型，支持两种类型，黑名单和白名单; 1黑名单，2白名单；RefererStatus为"on"时此参数必填；
	RefererType *int `required:"false"`

	// 防盗链Referer规则，支持正则表达式（不支持符号';')
	Referers []string `required:"false"`
}

// SetUFileRefererResponse is response schema for SetUFileReferer action
type SetUFileRefererResponse struct {
	response.CommonBase
}

// NewSetUFileRefererRequest will create request of SetUFileReferer action.
func (c *UFileClient) NewSetUFileRefererRequest() *SetUFileRefererRequest {
	req := &SetUFileRefererRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SetUFileReferer

设置对象存储防盗链
*/
func (c *UFileClient) SetUFileReferer(req *SetUFileRefererRequest) (*SetUFileRefererResponse, error) {
	var err error
	var res SetUFileRefererResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SetUFileReferer", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateBucketRequest is request schema for UpdateBucket action
type UpdateBucketRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 待修改Bucket的名称
	BucketName *string `required:"true"`

	// Bucket访问类型;public或private
	Type *string `required:"true"`
}

// UpdateBucketResponse is response schema for UpdateBucket action
type UpdateBucketResponse struct {
	response.CommonBase

	// Bucket的ID
	BucketId string

	// Bucket的名称
	BucketName string
}

// NewUpdateBucketRequest will create request of UpdateBucket action.
func (c *UFileClient) NewUpdateBucketRequest() *UpdateBucketRequest {
	req := &UpdateBucketRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateBucket

更改Bucket的属性
*/
func (c *UFileClient) UpdateBucket(req *UpdateBucketRequest) (*UpdateBucketResponse, error) {
	var err error
	var res UpdateBucketResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateBucket", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUFileLifeCycleRequest is request schema for UpdateUFileLifeCycle action
type UpdateUFileLifeCycleRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"false"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后过期，自动转换为归档存储类型；范围： [7,36500]，0代表不启用
	ArchivalDays *int `required:"false"`

	// 存储空间名称
	BucketName *string `required:"true"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后过期,自动删除；范围： [7,36500]
	Days *int `required:"false"`

	// 指定一个过期天数N，文件会在其最近更新时间点的N天后过期，自动转换为低频存储类型；范围： [7,36500]，0代表不启用
	IADays *int `required:"false"`

	// 生命周期Id
	LifeCycleId *string `required:"true"`

	// 生命周期名称
	LifeCycleName *string `required:"true"`

	// 生命周期所适用的前缀；*为整个存储空间文件；一条规则只支持一个文件前缀；
	Prefix *string `required:"true"`

	// Enabled -- 启用，Disabled -- 不启用
	Status *string `required:"true"`
}

// UpdateUFileLifeCycleResponse is response schema for UpdateUFileLifeCycle action
type UpdateUFileLifeCycleResponse struct {
	response.CommonBase
}

// NewUpdateUFileLifeCycleRequest will create request of UpdateUFileLifeCycle action.
func (c *UFileClient) NewUpdateUFileLifeCycleRequest() *UpdateUFileLifeCycleRequest {
	req := &UpdateUFileLifeCycleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUFileLifeCycle

更新生命周期管理
*/
func (c *UFileClient) UpdateUFileLifeCycle(req *UpdateUFileLifeCycleRequest) (*UpdateUFileLifeCycleResponse, error) {
	var err error
	var res UpdateUFileLifeCycleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUFileLifeCycle", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUFileTokenRequest is request schema for UpdateUFileToken action
type UpdateUFileTokenRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// 令牌允许操作的bucket
	AllowedBuckets []string `required:"false"`

	// 令牌允许执行的操作，[ TOKEN_ALLOW_NONE , TOKEN_ALLOW_READ , TOKEN_ALLOW_WRITE , TOKEN_ALLOW_DELETE , TOKEN_ALLOW_LIST, TOKEN_ALLOW_IOP , TOKEN_ALLOW_DP ]
	AllowedOps []string `required:"false"`

	// 令牌允许操作的key前缀
	AllowedPrefixes []string `required:"false"`

	// 令牌的超时时间点（时间戳）;注意：过期时间不能超过 4102416000
	ExpireTime *int `required:"false"`

	// 令牌ID
	TokenId *string `required:"true"`

	// 令牌名称
	TokenName *string `required:"false"`
}

// UpdateUFileTokenResponse is response schema for UpdateUFileToken action
type UpdateUFileTokenResponse struct {
	response.CommonBase
}

// NewUpdateUFileTokenRequest will create request of UpdateUFileToken action.
func (c *UFileClient) NewUpdateUFileTokenRequest() *UpdateUFileTokenRequest {
	req := &UpdateUFileTokenRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUFileToken

更新令牌的操作权限，可操作key的前缀，可操作bucket和令牌超时时间点
*/
func (c *UFileClient) UpdateUFileToken(req *UpdateUFileTokenRequest) (*UpdateUFileTokenResponse, error) {
	var err error
	var res UpdateUFileTokenResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUFileToken", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUdsRuleRequest is request schema for UpdateUdsRule action
type UpdateUdsRuleRequest struct {
	request.CommonBase

	// 联系的用户组ID
	ContactGroupId *string `required:"false"`

	// 目标Bucket名字，全局唯一
	DstBucket *string `required:"true"`

	// 解压后的目标目录
	DstDirectory *string `required:"true"`

	// 目标bucket的token之一的tokenId
	DstTokenId *string `required:"true"`

	// 通知的事件数组
	Events []string `required:"false"`

	// 是否以压缩文件的前缀为最后一层目录
	KeepUS3Name *bool `required:"true"`

	// 通知的类型数组
	NotificationTypes []string `required:"false"`

	// 操作的ops数组,"Ops.0":"unzip"
	Ops []string `required:"false"`

	// 触发解压缩的前缀
	Prefixes *string `required:"true"`

	// 规则的唯一Id
	RuleId *string `required:"true"`

	// 规则名称
	RuleName *string `required:"true"`

	// 源Bucket名字，全局唯一
	SrcBucket *string `required:"true"`

	// 源bucket的token之一的tokenId
	SrcTokenId *string `required:"true"`
}

// UpdateUdsRuleResponse is response schema for UpdateUdsRule action
type UpdateUdsRuleResponse struct {
	response.CommonBase

	// 该请求的消息成功或者失败的描述
	Mesage string

	// 返回规则的规则ID
	RuleId string
}

// NewUpdateUdsRuleRequest will create request of UpdateUdsRule action.
func (c *UFileClient) NewUpdateUdsRuleRequest() *UpdateUdsRuleRequest {
	req := &UpdateUdsRuleRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUdsRule

针对对象存储的文件，进行自动触发解压。
*/
func (c *UFileClient) UpdateUdsRule(req *UpdateUdsRuleRequest) (*UpdateUdsRuleResponse, error) {
	var err error
	var res UpdateUdsRuleResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUdsRule", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
