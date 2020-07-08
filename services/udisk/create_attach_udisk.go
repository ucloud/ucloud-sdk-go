package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

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
