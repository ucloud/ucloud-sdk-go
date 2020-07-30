// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UDisk API Schema

// AttachUDiskRequest is request schema for AttachUDisk action
type AttachUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 是否允许多点挂载（Yes: 允许多点挂载， No: 不允许多点挂载， 不填默认Yes ）
	MultiAttach *string `required:"false"`

	// 需要挂载的UDisk实例ID.
	UDiskId *string `required:"true"`

	// UHost实例ID
	UHostId *string `required:"true"`
}

// AttachUDiskResponse is response schema for AttachUDisk action
type AttachUDiskResponse struct {
	response.CommonBase

	// 挂载的设备名称
	DeviceName string

	// 挂载的UDisk实例ID
	UDiskId string

	// 挂载的UHost实例ID
	UHostId string
}

// NewAttachUDiskRequest will create request of AttachUDisk action.
func (c *UDiskClient) NewAttachUDiskRequest() *AttachUDiskRequest {
	req := &AttachUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: AttachUDisk

将一个可用的UDisk挂载到某台主机上，当UDisk挂载成功后，还需要在主机内部进行文件系统操作
*/
func (c *UDiskClient) AttachUDisk(req *AttachUDiskRequest) (*AttachUDiskResponse, error) {
	var err error
	var res AttachUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AttachUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CloneUDiskRequest is request schema for CloneUDisk action
type CloneUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic，Postpay，Trial 默认: Month
	ChargeType *string `required:"false"`

	// Disk注释
	Comment *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 克隆父Disk的Id
	SourceId *string `required:"true"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`
}

// CloneUDiskResponse is response schema for CloneUDisk action
type CloneUDiskResponse struct {
	response.CommonBase

	// 创建UDisk Id
	UDiskId []string
}

// NewCloneUDiskRequest will create request of CloneUDisk action.
func (c *UDiskClient) NewCloneUDiskRequest() *CloneUDiskRequest {
	req := &CloneUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: CloneUDisk

从UDisk创建UDisk克隆
*/
func (c *UDiskClient) CloneUDisk(req *CloneUDiskRequest) (*CloneUDiskResponse, error) {
	var err error
	var res CloneUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CloneUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CloneUDiskSnapshotRequest is request schema for CloneUDiskSnapshot action
type CloneUDiskSnapshotRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic，Postpay 默认: Dynamic
	ChargeType *string `required:"false"`

	// Disk注释
	Comment *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 购买UDisk大小,单位:GB,范围[1~8000]。(UDisk大小设定对本地盘快照有效，对云盘快照无效)
	Size *int `required:"false"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 克隆父Snapshot的Id
	SourceId *string `required:"true"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`
}

// CloneUDiskSnapshotResponse is response schema for CloneUDiskSnapshot action
type CloneUDiskSnapshotResponse struct {
	response.CommonBase

	// 创建UDisk Id
	UDiskId []string
}

// NewCloneUDiskSnapshotRequest will create request of CloneUDiskSnapshot action.
func (c *UDiskClient) NewCloneUDiskSnapshotRequest() *CloneUDiskSnapshotRequest {
	req := &CloneUDiskSnapshotRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: CloneUDiskSnapshot

从快照创建UDisk克隆
*/
func (c *UDiskClient) CloneUDiskSnapshot(req *CloneUDiskSnapshotRequest) (*CloneUDiskSnapshotResponse, error) {
	var err error
	var res CloneUDiskSnapshotResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CloneUDiskSnapshot", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CloneUDiskUDataArkRequest is request schema for CloneUDiskUDataArk action
type CloneUDiskUDataArkRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic，Postpay 默认: Dynamic
	ChargeType *string `required:"false"`

	// Disk注释
	Comment *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 购买UDisk大小,单位:GB,范围[1~8000]。(UDisk大小设定对本地盘备份有效，对云盘备份无效)
	Size *int `required:"false"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 指定从方舟克隆的备份时间点
	SnapshotTime *int `required:"true"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`

	// 需要克隆的源盘id
	UDiskId *string `required:"true"`
}

// CloneUDiskUDataArkResponse is response schema for CloneUDiskUDataArk action
type CloneUDiskUDataArkResponse struct {
	response.CommonBase

	// 创建UDisk Id
	UDiskId []string
}

// NewCloneUDiskUDataArkRequest will create request of CloneUDiskUDataArk action.
func (c *UDiskClient) NewCloneUDiskUDataArkRequest() *CloneUDiskUDataArkRequest {
	req := &CloneUDiskUDataArkRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: CloneUDiskUDataArk

从数据方舟的备份创建UDisk
*/
func (c *UDiskClient) CloneUDiskUDataArk(req *CloneUDiskUDataArkRequest) (*CloneUDiskUDataArkResponse, error) {
	var err error
	var res CloneUDiskUDataArkResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CloneUDiskUDataArk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateAttachUDiskRequest is request schema for CreateAttachUDisk action
type CreateAttachUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic, Postpay, Trial 。 Size小于等于2000时，默认为Dynamic；Size大于2000时，默认为Month。
	ChargeType *string `required:"false"`

	// 加密需要的cmk id，UKmsMode为Yes时，必填
	CmkId *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 是否允许多点挂载（Yes: 允许多点挂载， No: 不允许多点挂载， 不填默认Yes ）
	MultiAttach *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 购买UDisk大小,单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`

	// 当创建云盘类型为RSSDDataDisk时，根据传入的UHostId，创建与虚机在同一PodId下的云盘
	UHostId *string `required:"true"`

	// 是否加密。Yes：加密，No：不加密，默认值（No）
	UKmsMode *string `required:"false"`
}

// CreateAttachUDiskResponse is response schema for CreateAttachUDisk action
type CreateAttachUDiskResponse struct {
	response.CommonBase

	// 挂载设备名称
	DeviceName string

	// 挂载的UDisk实例ID
	UDiskId string

	// 挂载的UHost实例ID
	UHostId string
}

// NewCreateAttachUDiskRequest will create request of CreateAttachUDisk action.
func (c *UDiskClient) NewCreateAttachUDiskRequest() *CreateAttachUDiskRequest {
	req := &CreateAttachUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateAttachUDisk

创建并挂载UDisk磁盘
*/
func (c *UDiskClient) CreateAttachUDisk(req *CreateAttachUDiskRequest) (*CreateAttachUDiskResponse, error) {
	var err error
	var res CreateAttachUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateAttachUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUDiskRequest is request schema for CreateUDisk action
type CreateUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic, Postpay, Trial 。 Size小于等于2000时，默认为Dynamic；Size大于2000时，默认为Month。
	ChargeType *string `required:"false"`

	// 加密需要的cmk id，UKmsMode为Yes时，必填
	CmkId *string `required:"false"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 实例名称
	Name *string `required:"true"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 购买UDisk大小,单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 业务组 默认：Default
	Tag *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`

	// 是否加密。Yes：加密，No：不加密，默认值（No）
	UKmsMode *string `required:"false"`
}

// CreateUDiskResponse is response schema for CreateUDisk action
type CreateUDiskResponse struct {
	response.CommonBase

	// UDisk实例Id
	UDiskId []string
}

// NewCreateUDiskRequest will create request of CreateUDisk action.
func (c *UDiskClient) NewCreateUDiskRequest() *CreateUDiskRequest {
	req := &CreateUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUDisk

创建UDisk磁盘
*/
func (c *UDiskClient) CreateUDisk(req *CreateUDiskRequest) (*CreateUDiskResponse, error) {
	var err error
	var res CreateUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateUDiskSnapshotRequest is request schema for CreateUDiskSnapshot action
type CreateUDiskSnapshotRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic 默认: Dynamic  (已废弃)
	ChargeType *string `required:"false"`

	// 快照描述
	Comment *string `required:"false"`

	// 快照名称
	Name *string `required:"true"`

	// 购买时长 默认: 1  (已废弃)
	Quantity *int `required:"false"`

	// 快照的UDisk的Id
	UDiskId *string `required:"true"`
}

// CreateUDiskSnapshotResponse is response schema for CreateUDiskSnapshot action
type CreateUDiskSnapshotResponse struct {
	response.CommonBase

	// 快照Id
	SnapshotId []string
}

// NewCreateUDiskSnapshotRequest will create request of CreateUDiskSnapshot action.
func (c *UDiskClient) NewCreateUDiskSnapshotRequest() *CreateUDiskSnapshotRequest {
	req := &CreateUDiskSnapshotRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUDiskSnapshot

创建snapshot快照
*/
func (c *UDiskClient) CreateUDiskSnapshot(req *CreateUDiskSnapshotRequest) (*CreateUDiskSnapshotResponse, error) {
	var err error
	var res CreateUDiskSnapshotResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUDiskSnapshot", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUDiskRequest is request schema for DeleteUDisk action
type DeleteUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 要删除的UDisk的Id
	UDiskId *string `required:"true"`
}

// DeleteUDiskResponse is response schema for DeleteUDisk action
type DeleteUDiskResponse struct {
	response.CommonBase
}

// NewDeleteUDiskRequest will create request of DeleteUDisk action.
func (c *UDiskClient) NewDeleteUDiskRequest() *DeleteUDiskRequest {
	req := &DeleteUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUDisk

删除UDisk
*/
func (c *UDiskClient) DeleteUDisk(req *DeleteUDiskRequest) (*DeleteUDiskResponse, error) {
	var err error
	var res DeleteUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteUDiskSnapshotRequest is request schema for DeleteUDiskSnapshot action
type DeleteUDiskSnapshotRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 快照Id(填写后不能填写UDisk Id)
	SnapshotId *string `required:"false"`

	// UDisk Id,删除该盘所创建出来的所有快照(填写后不能填写SnapshotId)
	UDiskId *string `required:"false"`
}

// DeleteUDiskSnapshotResponse is response schema for DeleteUDiskSnapshot action
type DeleteUDiskSnapshotResponse struct {
	response.CommonBase
}

// NewDeleteUDiskSnapshotRequest will create request of DeleteUDiskSnapshot action.
func (c *UDiskClient) NewDeleteUDiskSnapshotRequest() *DeleteUDiskSnapshotRequest {
	req := &DeleteUDiskSnapshotRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUDiskSnapshot

删除Snapshot
*/
func (c *UDiskClient) DeleteUDiskSnapshot(req *DeleteUDiskSnapshotRequest) (*DeleteUDiskSnapshotResponse, error) {
	var err error
	var res DeleteUDiskSnapshotResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteUDiskSnapshot", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeRecycleUDiskRequest is request schema for DescribeRecycleUDisk action
type DescribeRecycleUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 返回数据长度, 默认为20
	Limit *int `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`
}

// DescribeRecycleUDiskResponse is response schema for DescribeRecycleUDisk action
type DescribeRecycleUDiskResponse struct {
	response.CommonBase

	// 回收站磁盘列表
	DataSet []RecycleUDiskSet

	// 磁盘数量
	TotalCount int
}

// NewDescribeRecycleUDiskRequest will create request of DescribeRecycleUDisk action.
func (c *UDiskClient) NewDescribeRecycleUDiskRequest() *DescribeRecycleUDiskRequest {
	req := &DescribeRecycleUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeRecycleUDisk

拉取回收站中云硬盘列表
*/
func (c *UDiskClient) DescribeRecycleUDisk(req *DescribeRecycleUDiskRequest) (*DescribeRecycleUDiskResponse, error) {
	var err error
	var res DescribeRecycleUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeRecycleUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUDiskRequest is request schema for DescribeUDisk action
type DescribeUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// ProtocolVersion字段为1时，需结合IsBoot确定具体磁盘类型:普通数据盘：DiskType:"CLOUD_NORMAL",IsBoot:"False"；普通系统盘：DiskType:"CLOUD_NORMAL",IsBoot:"True"；SSD数据盘：DiskType:"CLOUD_SSD",IsBoot:"False"；SSD系统盘：DiskType:"CLOUD_SSD",IsBoot:"True"；RSSD数据盘：DiskType:"CLOUD_RSSD",IsBoot:"False"；为空拉取所有。ProtocolVersion字段为0或没有该字段时，可设为以下几个值:普通数据盘：DataDisk；普通系统盘：SystemDisk；SSD数据盘：SSDDataDisk；SSD系统盘：SSDSystemDisk；RSSD数据盘：RSSDDataDisk；为空拉取所有。
	DiskType *string `required:"false"`

	// 是否忽略计费信息。Yes：忽略，No：不忽略，默认值（No）。（如不关心账单信息，建议选填“Yes”，可降低请求延时）
	IgnoreUBillInfo *string `required:"false"`

	// ProtocolVersion字段为1且DiskType不为空时，必须设置，设置规则请参照DiskType；ProtocolVersion字段为1且DiskType为空时，该字段无效。ProtocolVersion字段为0或没有该字段时，该字段无效。
	IsBoot *string `required:"false"`

	// 返回数据长度, 默认为20
	Limit *int `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 请求协议版本，建议升级为1，为1时DiskType与UHost磁盘类型定义一致；默认为0
	ProtocolVersion *int `required:"false"`

	// UDisk Id(留空返回全部)
	UDiskId *string `required:"false"`

	// 根据传入的UHostIdForAttachment，筛选出虚机在同一PodId下的云盘
	UHostIdForAttachment *string `required:"false"`
}

// DescribeUDiskResponse is response schema for DescribeUDisk action
type DescribeUDiskResponse struct {
	response.CommonBase

	// JSON 格式的UDisk数据列表, 每项参数可见下面 UDiskDataSet
	DataSet []UDiskDataSet

	// 根据过滤条件得到的总数
	TotalCount int
}

// NewDescribeUDiskRequest will create request of DescribeUDisk action.
func (c *UDiskClient) NewDescribeUDiskRequest() *DescribeUDiskRequest {
	req := &DescribeUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUDisk

获取UDisk实例
*/
func (c *UDiskClient) DescribeUDisk(req *DescribeUDiskRequest) (*DescribeUDiskResponse, error) {
	var err error
	var res DescribeUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUDiskPriceRequest is request schema for DescribeUDiskPrice action
type DescribeUDiskPriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic，Postpay，Trial 默认: Month
	ChargeType *string `required:"false"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），SystemDisk（普通系统盘），SSDSystemDisk（SSD系统盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 是否将快照服务(数据方舟)，云硬盘放入一张订单, 是："Yes",否："No"，默认是"No"
	IsTotalPrice *string `required:"false"`

	// 云主机机型（V2.0），枚举值["N", "C", "G", "O", "OM"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 购买UDisk的时长，默认值为1
	Quantity *int `required:"false"`

	// 购买UDisk大小,单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；普通系统盘：范围[1~8000]；SSD系统盘：范围[1~4000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`
}

// DescribeUDiskPriceResponse is response schema for DescribeUDiskPrice action
type DescribeUDiskPriceResponse struct {
	response.CommonBase

	// 价格参数列表，具体说明见 UDiskPriceDataSet
	DataSet []UDiskPriceDataSet
}

// NewDescribeUDiskPriceRequest will create request of DescribeUDiskPrice action.
func (c *UDiskClient) NewDescribeUDiskPriceRequest() *DescribeUDiskPriceRequest {
	req := &DescribeUDiskPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUDiskPrice

获取UDisk实例价格信息
*/
func (c *UDiskClient) DescribeUDiskPrice(req *DescribeUDiskPriceRequest) (*DescribeUDiskPriceResponse, error) {
	var err error
	var res DescribeUDiskPriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUDiskPrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUDiskSnapshotRequest is request schema for DescribeUDiskSnapshot action
type DescribeUDiskSnapshotRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 返回数据长度, 默认为20
	Limit *int `required:"false"`

	// 数据偏移量, 默认为0
	Offset *int `required:"false"`

	// 快照id，SnapshotId , UDiskId 同时传SnapshotId优先
	SnapshotId *string `required:"false"`

	// UDiskId,返回该盘所做快照.(必须同时传Zone)
	UDiskId *string `required:"false"`
}

// DescribeUDiskSnapshotResponse is response schema for DescribeUDiskSnapshot action
type DescribeUDiskSnapshotResponse struct {
	response.CommonBase

	// JSON 格式的Snapshot列表, 详细参见 UDiskSnapshotSet
	DataSet []UDiskSnapshotSet

	// 根据过滤条件得到的总数
	TotalCount int
}

// NewDescribeUDiskSnapshotRequest will create request of DescribeUDiskSnapshot action.
func (c *UDiskClient) NewDescribeUDiskSnapshotRequest() *DescribeUDiskSnapshotRequest {
	req := &DescribeUDiskSnapshotRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUDiskSnapshot

获取UDisk快照
*/
func (c *UDiskClient) DescribeUDiskSnapshot(req *DescribeUDiskSnapshotRequest) (*DescribeUDiskSnapshotResponse, error) {
	var err error
	var res DescribeUDiskSnapshotResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUDiskSnapshot", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUDiskUpgradePriceRequest is request schema for DescribeUDiskUpgradePrice action
type DescribeUDiskUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），SystemDisk（普通系统盘），SSDSystemDisk（SSD系统盘），RSSDDataDisk（RSSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`

	// 云主机机型（V2.0），枚举值["N", "C", "G", "O", "OM"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 购买UDisk大小,单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；普通系统盘：范围[1~8000]；SSD系统盘：范围[1~4000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// 是否开启快照服务。Yes：开启，No：不开启，默认值：No
	SnapshotService *string `required:"false"`

	// 升级目标UDisk ID
	SourceId *string `required:"true"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`
}

// DescribeUDiskUpgradePriceResponse is response schema for DescribeUDiskUpgradePrice action
type DescribeUDiskUpgradePriceResponse struct {
	response.CommonBase

	// 用户折后价 (对应计费CustomPrice)
	OriginalPrice int

	// 价格
	Price int
}

// NewDescribeUDiskUpgradePriceRequest will create request of DescribeUDiskUpgradePrice action.
func (c *UDiskClient) NewDescribeUDiskUpgradePriceRequest() *DescribeUDiskUpgradePriceRequest {
	req := &DescribeUDiskUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUDiskUpgradePrice

获取UDisk升级价格信息
*/
func (c *UDiskClient) DescribeUDiskUpgradePrice(req *DescribeUDiskUpgradePriceRequest) (*DescribeUDiskUpgradePriceResponse, error) {
	var err error
	var res DescribeUDiskUpgradePriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUDiskUpgradePrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DetachUDiskRequest is request schema for DetachUDisk action
type DetachUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 需要卸载的UDisk实例ID
	UDiskId *string `required:"true"`

	// UHost实例ID
	UHostId *string `required:"true"`
}

// DetachUDiskResponse is response schema for DetachUDisk action
type DetachUDiskResponse struct {
	response.CommonBase

	// 卸载的UDisk实例ID
	UDiskId string

	// 卸载的UHost实例ID
	UHostId string
}

// NewDetachUDiskRequest will create request of DetachUDisk action.
func (c *UDiskClient) NewDetachUDiskRequest() *DetachUDiskRequest {
	req := &DetachUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DetachUDisk

卸载某个已经挂载在指定UHost实例上的UDisk
*/
func (c *UDiskClient) DetachUDisk(req *DetachUDiskRequest) (*DetachUDiskResponse, error) {
	var err error
	var res DetachUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DetachUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RecoverUDiskRequest is request schema for RecoverUDisk action
type RecoverUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// Year , Month, Dynamic 默认: Dynamic
	ChargeType *string `required:"false"`

	// 购买时长 默认: 1
	Quantity *int `required:"false"`

	// 云硬盘资源ID
	UDiskId *string `required:"true"`
}

// RecoverUDiskResponse is response schema for RecoverUDisk action
type RecoverUDiskResponse struct {
	response.CommonBase
}

// NewRecoverUDiskRequest will create request of RecoverUDisk action.
func (c *UDiskClient) NewRecoverUDiskRequest() *RecoverUDiskRequest {
	req := &RecoverUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RecoverUDisk

从回收站中恢复云硬盘
*/
func (c *UDiskClient) RecoverUDisk(req *RecoverUDiskRequest) (*RecoverUDiskResponse, error) {
	var err error
	var res RecoverUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RecoverUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RenameUDiskRequest is request schema for RenameUDisk action
type RenameUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 重命名的UDisk的Id
	UDiskId *string `required:"true"`

	// 重命名UDisk的name
	UDiskName *string `required:"true"`
}

// RenameUDiskResponse is response schema for RenameUDisk action
type RenameUDiskResponse struct {
	response.CommonBase
}

// NewRenameUDiskRequest will create request of RenameUDisk action.
func (c *UDiskClient) NewRenameUDiskRequest() *RenameUDiskRequest {
	req := &RenameUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RenameUDisk

重命名UDisk
*/
func (c *UDiskClient) RenameUDisk(req *RenameUDiskRequest) (*RenameUDiskResponse, error) {
	var err error
	var res RenameUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RenameUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ResizeUDiskRequest is request schema for ResizeUDisk action
type ResizeUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 云主机机型（V2.0），枚举值["N", "C", "G", "O", "OM"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 调整后大小, 单位:GB,普通数据盘：范围[1~8000]；SSD数据盘：范围[1~8000]；RSSD数据盘：范围[1~32000]。
	Size *int `required:"true"`

	// UDisk Id
	UDiskId *string `required:"true"`
}

// ResizeUDiskResponse is response schema for ResizeUDisk action
type ResizeUDiskResponse struct {
	response.CommonBase
}

// NewResizeUDiskRequest will create request of ResizeUDisk action.
func (c *UDiskClient) NewResizeUDiskRequest() *ResizeUDiskRequest {
	req := &ResizeUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ResizeUDisk

调整UDisk容量
*/
func (c *UDiskClient) ResizeUDisk(req *ResizeUDiskRequest) (*ResizeUDiskResponse, error) {
	var err error
	var res ResizeUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ResizeUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RestoreUDiskRequest is request schema for RestoreUDisk action
type RestoreUDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 从指定的快照恢复
	SnapshotId *string `required:"false"`

	// 指定从方舟恢复的备份时间点
	SnapshotTime *int `required:"false"`

	// 需要恢复的盘ID
	UDiskId *string `required:"true"`
}

// RestoreUDiskResponse is response schema for RestoreUDisk action
type RestoreUDiskResponse struct {
	response.CommonBase
}

// NewRestoreUDiskRequest will create request of RestoreUDisk action.
func (c *UDiskClient) NewRestoreUDiskRequest() *RestoreUDiskRequest {
	req := &RestoreUDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: RestoreUDisk

从备份恢复数据至UDisk
*/
func (c *UDiskClient) RestoreUDisk(req *RestoreUDiskRequest) (*RestoreUDiskResponse, error) {
	var err error
	var res RestoreUDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RestoreUDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// SetUDiskUDataArkModeRequest is request schema for SetUDiskUDataArkMode action
type SetUDiskUDataArkModeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 使用的代金券id
	CouponId *string `required:"false"`

	// 【即将废弃，开启快照服务时，免费开启数据方舟】是否开启数据方舟。Yes：开启，No：不开启，默认值：No
	UDataArkMode *string `required:"false"`

	// 需要设置数据方舟的UDisk的Id
	UDiskId *string `required:"true"`
}

// SetUDiskUDataArkModeResponse is response schema for SetUDiskUDataArkMode action
type SetUDiskUDataArkModeResponse struct {
	response.CommonBase
}

// NewSetUDiskUDataArkModeRequest will create request of SetUDiskUDataArkMode action.
func (c *UDiskClient) NewSetUDiskUDataArkModeRequest() *SetUDiskUDataArkModeRequest {
	req := &SetUDiskUDataArkModeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SetUDiskUDataArkMode

设置UDisk数据方舟的状态
*/
func (c *UDiskClient) SetUDiskUDataArkMode(req *SetUDiskUDataArkModeRequest) (*SetUDiskUDataArkModeResponse, error) {
	var err error
	var res SetUDiskUDataArkModeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SetUDiskUDataArkMode", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
