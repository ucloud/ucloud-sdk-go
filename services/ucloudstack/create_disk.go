// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateDiskRequest is request schema for CreateDisk action
type CreateDiskRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值：cn,表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType *string `required:"true"`

	// 磁盘大小
	DiskSpace *int `required:"true"`

	// 磁盘名称
	Name *string `required:"true"`

	// 购买时长。默认值1。小时不生效，月范围【1，11】，年范围【1，5】。
	Quantity *int `required:"false"`

	// 磁盘类型。例如：Normal,SSD
	SetType *string `required:"true"`
}

// CreateDiskResponse is response schema for CreateDisk action
type CreateDiskResponse struct {
	response.CommonBase

	// 创建的磁盘ID
	DiskID string

	// 返回信息描述。
	Message string
}

// NewCreateDiskRequest will create request of CreateDisk action.
func (c *UCloudStackClient) NewCreateDiskRequest() *CreateDiskRequest {
	req := &CreateDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateDisk - 创建UCloudStack硬盘
func (c *UCloudStackClient) CreateDisk(req *CreateDiskRequest) (*CreateDiskResponse, error) {
	var err error
	var res CreateDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
