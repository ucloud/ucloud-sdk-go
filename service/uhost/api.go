package uhost

import (
	"github.com/xiaohui/goucloud/ucloud"
)

// CreateUHostInstance will create instances
type CreateUHostInstanceParams struct {
	ucloud.CommonRequest

	Region    string
	ImageId   string
	LoginMode string
	Password  string
	KeyPair   string
	CPU       int
	Memory    int
	DiskSpace int
	Name      string
	NetworkId string

	SecurityGroupId string
	ChargeType      string
	Quantity        int
	Count           int
	UHostType       string
	NetCapability   string
	Tag             string
	CouponId        string
}

type CreateUHostInstanceResponse struct {
	ucloud.CommonResponse
	HostIds []string
}

func (u *UHost) CreateUHostInstance(params *CreateUHostInstanceParams) (*CreateUHostInstanceResponse, error) {
	response := CreateUHostInstanceResponse{}
	err := u.DoRequest("CreateUHostInstance", params, response)

	return &response, err
}

type DescribeImageParams struct {
	ucloud.CommonRequest

	Region    string
	ImageType string
	OsType    string
	ImageId   string
	Offset    int
	Limit     int
}

type ImageSet struct {
	ImageId   string
	ImageName string
	OsType    string
	OsName    string
	State     string

	ImageDescription string
	CreateTime       string
}

type ImageSetArray []ImageSet

type DescribeImageResponse struct {
	ucloud.CommonResponse

	TotalCount int
	ImageSet   ImageSetArray
}

func (u *UHost) DescribeImage(params *DescribeImageParams) (*DescribeImageResponse, error) {
	response := DescribeImageResponse{}
	err := u.DoRequest("DescribeImage", params, response)

	return &response, err
}

type DescribeUHostInstanceParams struct {
	Region string
	Tag    string
	Offset int
	Limit  int
}

type DiskSet struct {
	Type   string
	DiskId string
	Size   int
}

type DiskSetArray []DiskSet

type IPSet struct {
	Type      string
	IPId      string
	IP        string
	bandwidth int
}

type IPSetArray []IPSet

type UHostSet struct {
	UHostId        string
	UHostType      string
	ImageId        string
	BasicImageId   string
	BasicImageName string
	Tag            string
	Remark         string
	Name           string
	State          string
	CreateTime     int
	ChargeType     string
	ExpireTime     string
	CPU            int
	Memory         int
	DiskSet        DiskSetArray
	IPSet          IPSetArray
	NetCapability  string
}

type UHostSetArray []UHostSet

type DescribeUHostInstanceResponse struct {
	ucloud.CommonResponse

	TotalCount int
	UHostSet   UHostSetArray
}

func (u *UHost) DescribeUHostInstance(params *DescribeUHostInstanceParams) (*DescribeUHostInstanceResponse, error) {
	response := DescribeUHostInstanceResponse{}
	err := u.DoRequest("DescribeUHostInstance", params, response)

	return &response, err
}

type StartUHostInstanceParams struct {
	ucloud.CommonRequest

	Region  string
	UHostId string
}

type StartUHostInstanceResponse struct {
	ucloud.CommonResponse
	UhostId string
}

func (u *UHost) StartUHostInstance(params *StartUHostInstanceParams) (*StartUHostInstanceResponse, error) {
	response := StartUHostInstanceResponse{}
	err := u.DoRequest("StartUHostInstance", params, response)

	return &response, err
}

type StopUHostInstanceParams struct {
	ucloud.CommonRequest

	Region  string
	UHostId string
}

type StopUHostInstanceResponse struct {
	ucloud.CommonResponse

	UhostId string
}

func (u *UHost) StopUHostInstance(params *StopUHostInstanceParams) (*StopUHostInstanceResponse, error) {
	response := StopUHostInstanceResponse{}
	err := u.DoRequest("StopUHostInstance", params, response)

	return &response, err
}

type PoweroffUHostInstanceParams struct {
	ucloud.CommonRequest

	Region  string
	UHostId string
}

type PoweroffUHostInstanceResponse struct {
	ucloud.CommonResponse

	UhostId string
}

func (u *UHost) PoweroffUHostInstance(params *PoweroffUHostInstanceParams) (*PoweroffUHostInstanceResponse, error) {
	response := PoweroffUHostInstanceResponse{}
	err := u.DoRequest("PoweroffUHostInstance", params, response)

	return &response, err
}

type RebootUHostInstanceParams struct {
	ucloud.CommonRequest

	Region  string
	UHostId string
}

type RebootUHostInstanceResponse struct {
	ucloud.CommonResponse

	UhostId string
}

func (u *UHost) RebootUHostInstance(params *RebootUHostInstanceParams) (*RebootUHostInstanceResponse, error) {
	response := RebootUHostInstanceResponse{}
	err := u.DoRequest("RebootUHostInstance", params, response)

	return &response, err
}

type ResetUHostInstancePasswordParams struct {
	ucloud.CommonRequest

	Region   string
	UHostId  string
	Password string
}

type ResetUHostInstancePasswordResponse struct {
	ucloud.CommonResponse

	UhostId string
}

func (u *UHost) ResetUHostInstancePassword(params *ResetUHostInstancePasswordParams) (*ResetUHostInstancePasswordResponse, error) {
	response := ResetUHostInstancePasswordResponse{}
	err := u.DoRequest("ResetUHostInstancePassword", params, response)

	return &response, err
}
