package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

type AllocateEIPParams struct {
	ucloud.CommonRequest

	OperatorName     string
	Region           string
	Bandwidth        int
	Tag              string
	ChargeType       string
	Quantity         int
	PayMode          string
	ShareBandwidthId string
	CouponId         string
	Name             string
	Remark           string
}

type EIPAddr struct {
	OperatorName string
	IP           string
}

type ShareBandwidthSet struct {
	ShareBandwidth     int
	ShareBandwidthId   string
	ShareBandwidthName string
}

type EIPSet struct {
	EIPId             string
	CurBandwidth      float32
	Weight            int
	BandwidthType     int
	Bandwidth         int
	Status            string
	ChargeType        string
	CreateTime        int
	ExpireTime        int
	Name              string
	Tag               string
	Remark            string
	PayMode           string
	ShareBandwidthSet *ShareBandwidthSet
	EIPAddr           *[]EIPAddr
	Resource          *[]ucloud.Resource
}

type AllocateEIPResponse struct {
	ucloud.CommonResponse

	EIPSet *[]EIPSet
}

func (u *UNet) AllocateEIP(params *AllocateEIPParams) (*AllocateEIPResponse, error) {
	response := &AllocateEIPResponse{}
	err := u.DoRequest("AllocateEIP", params, response)

	return response, err
}

type DescribeEIPParams struct {
	ucloud.CommonRequest

	Region string
	EIPIds []string
	OffSet int
	Limit  int
}

type DescribeEIPResponse struct {
	ucloud.CommonResponse

	TotalCount     int
	TotalBandwidth int
	EIPSet         *[]EIPSet
}

func (u *UNet) DescribeEIP(params *DescribeEIPParams) (*DescribeEIPResponse, error) {
	response := &DescribeEIPResponse{}
	err := u.DoRequest("DescribeEIP", params, response)

	return response, err
}

type UpdateEIPAttributeParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
	Name   string
	Tag    string
	Remark string
}

type UpdateEIPAttributeResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) UpdateEIPAttribute(params *UpdateEIPAttributeParams) (*UpdateEIPAttributeResponse, error) {
	response := &UpdateEIPAttributeResponse{}
	err := u.DoRequest("UpdateEIPAttribute", params, response)

	return response, err
}

type ReleaseEIPParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
}

type ReleaseEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ReleaseEIP(params *ReleaseEIPParams) (*ReleaseEIPResponse, error) {
	response := &ReleaseEIPResponse{}
	err := u.DoRequest("ReleaseEIP", params, response)

	return response, err
}

type BindEIPParams struct {
	ucloud.CommonRequest

	Region       string
	EIPId        string
	ResourceType string
	ResourceId   string
}

type BindEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) BindEIP(params *BindEIPParams) (*BindEIPResponse, error) {
	response := &BindEIPResponse{}
	err := u.DoRequest("BindEIP", params, response)

	return response, err
}

type UnBindEIPParams struct {
	ucloud.CommonRequest

	Region       string
	EIPId        string
	ResourceType string
	ResourceId   string
}

type UnBindEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) UnBindEIP(params *UnBindEIPParams) (*UnBindEIPResponse, error) {
	response := &UnBindEIPResponse{}
	err := u.DoRequest("UnBindEIP", params, response)

	return response, err
}

type ModifyEIPBandwidthParams struct {
	ucloud.CommonRequest

	Region    string
	EIPId     string
	Bandwidth int
}

type ModifyEIPBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ModifyEIPBandwidth(params *ModifyEIPBandwidthParams) (*ModifyEIPBandwidthResponse, error) {
	response := &ModifyEIPBandwidthResponse{}
	err := u.DoRequest("ModifyEIPBandwidth", params, response)

	return response, err
}

type ModifyEIPWeightParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
	Weight int
}

type ModifyEIPWeightResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ModifyEIPWeight(params *ModifyEIPWeightParams) (*ModifyEIPWeightResponse, error) {
	response := &ModifyEIPWeightResponse{}
	err := u.DoRequest("ModifyEIPWeight", params, response)

	return response, err
}

type GetEIPPriceParams struct {
	ucloud.CommonRequest

	OperatorName string
	Bandwidth    int
	ChargeType   string
	PayMode      string
}

type PriceSet struct {
	ChargeType    string
	Price         float32
	PurchaseValue int
}
type GetEIPPriceResponse struct {
	ucloud.CommonResponse

	PriceSet *[]PriceSet
}

func (u *UNet) GetEIPPrice(params *GetEIPPriceParams) (*GetEIPPriceResponse, error) {
	response := &GetEIPPriceResponse{}
	err := u.DoRequest("GetEIPPrice", params, response)

	return response, err
}

type GetEIPUpgradePriceParams struct {
	ucloud.CommonRequest

	Region    string
	EIPId     string
	Bandwidth int
}

type GetEIPUpgradePriceResponse struct {
	ucloud.CommonResponse

	Price float64
}

func (u *UNet) GetEIPUpgradePrice(params *GetEIPUpgradePriceParams) (*GetEIPUpgradePriceResponse, error) {
	response := &GetEIPUpgradePriceResponse{}
	err := u.DoRequest("GetEIPUpgradePrice", params, response)

	return response, err
}

type GetEIPPayModeEIPParams struct {
	ucloud.CommonRequest

	Region string
	EIPIds []string
}

type EIPPayMode struct {
	EIPId      string
	EIPPaymode string
}

type GetEIPPayModeEIPResponse struct {
	ucloud.CommonResponse

	EIPPayMode *[]EIPPayMode
}

func (u *UNet) GetEIPPayModeEIP(params *GetEIPUpgradePriceParams) (*GetEIPPayModeEIPResponse, error) {
	response := &GetEIPPayModeEIPResponse{}
	err := u.DoRequest("GetEIPPayModeEIP", params, response)

	return response, err
}

type SetEIPPayModeParams struct {
	ucloud.CommonRequest

	Region    string
	EIPId     string
	Bandwidth int
	PayMode   string
}

type SetEIPPayModeResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) SetEIPPayMode(params *SetEIPPayModeParams) (*SetEIPPayModeResponse, error) {
	response := &SetEIPPayModeResponse{}
	err := u.DoRequest("SetEIPPayMode", params, response)

	return response, err
}

type AllocateVIPParams struct {
	ucloud.CommonRequest

	Region string
	Zone   string
	Count  int
}

type AllocateVIPResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) AllocateVIP(params *AllocateVIPParams) (*AllocateVIPResponse, error) {
	response := &AllocateVIPResponse{}
	err := u.DoRequest("AllocateVIP", params, response)

	return response, err
}

type DescribeVIPParams struct {
	ucloud.CommonRequest

	Region string
	Zone   string
}

type DescribeVIPResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) DescribeVIP(params *DescribeVIPParams) (*DescribeVIPResponse, error) {
	response := &DescribeVIPResponse{}
	err := u.DoRequest("DescribeVIP", params, response)

	return response, err
}

type ReleaseVIPParams struct {
	ucloud.CommonRequest

	Region string
	Zone   string
	VIP    string
}

type ReleaseVIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ReleaseVIP(params *ReleaseVIPParams) (*ReleaseVIPResponse, error) {
	response := &ReleaseVIPResponse{}
	err := u.DoRequest("ReleaseVIP", params, response)

	return response, err
}

type CreateBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region     string
	EIPId      string
	Bandwidth  int
	EnableTime int
	TimeRange  int
	CouponId   string
}

type CreateBandwidthPackageResponse struct {
	ucloud.CommonResponse

	BandwidthPackageId string
}

func (u *UNet) CreateBandwidthPackage(params *CreateBandwidthPackageParams) (*CreateBandwidthPackageResponse, error) {
	response := &CreateBandwidthPackageResponse{}
	err := u.DoRequest("CreateBandwidthPackage", params, response)

	return response, err
}

type DescribeBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region string
	Limit  int
	OffSet int
}

type BandwidthPackageDataSet struct {
	BandwidthPackageId string
	EnableTime         int
	DisableTime        int
	CreateTime         int
	Bandwidth          int
	EIPId              string
	EIPAddr            *[]EIPAddr
}

type DescribeBandwidthPackageResponse struct {
	ucloud.CommonResponse

	TotalCount int
	DataSets   *[]BandwidthPackageDataSet
}

func (u *UNet) DescribeBandwidthPackage(params *DescribeBandwidthPackageParams) (*DescribeBandwidthPackageResponse, error) {
	response := &DescribeBandwidthPackageResponse{}
	err := u.DoRequest("DescribeBandwidthPackage", params, response)

	return response, err
}

type DeleteBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region             string
	BandwidthPackageId string
}

type DeleteBandwidthPackageResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) DeleteBandwidthPackage(params *DeleteBandwidthPackageParams) (*DeleteBandwidthPackageResponse, error) {
	response := &DeleteBandwidthPackageResponse{}
	err := u.DoRequest("DeleteBandwidthPackage", params, response)

	return response, err
}

type AllocateShareBandwidthParams struct {
	ucloud.CommonRequest

	Region         string
	ShareBandwidth int
	ChargeType     string
	Quantity       int
	Name           string
}

type AllocateShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) AllocateShareBandwidth(params *AllocateShareBandwidthParams) (*AllocateShareBandwidthResponse, error) {
	response := &AllocateShareBandwidthResponse{}
	err := u.DoRequest("AllocateShareBandwidth", params, response)

	return response, err
}

type DescribeShareBandwidthParams struct {
	ucloud.CommonRequest

	Region            string
	ShareBandwidthIds []string
}

type UNetShareBandwidthSet struct {
	Bandwidth        int
	ShareBandwidthId string
	ChargeType       string
	CreateTime       int
	ExpireTime       int
	EIPSet           *[]EIPAddr
}

type DescribeShareBandwidthResponse struct {
	ucloud.CommonResponse

	DataSet *[]UNetShareBandwidthSet
}

func (u *UNet) DescribeShareBandwidth(params *DescribeShareBandwidthParams) (*DescribeShareBandwidthResponse, error) {
	response := &DescribeShareBandwidthResponse{}
	err := u.DoRequest("DescribeShareBandwidth", params, response)

	return response, err
}

type ResizeShareBandwidthParams struct {
	ucloud.CommonRequest

	Region           string
	ShareBandwidth   int
	ShareBandwidthId string
}

type ResizeShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ResizeShareBandwidth(params *ResizeShareBandwidthParams) (*ResizeShareBandwidthResponse, error) {
	response := &ResizeShareBandwidthResponse{}
	err := u.DoRequest("ResizeShareBandwidth", params, response)

	return response, err
}

type ReleaseShareBandwidthParams struct {
	ucloud.CommonRequest

	Region           string
	EIPBandwidth     int
	ShareBandwidthId string
}

type ReleaseShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ReleaseShareBandwidth(params *ReleaseShareBandwidthParams) (*ReleaseShareBandwidthResponse, error) {
	response := &ReleaseShareBandwidthResponse{}
	err := u.DoRequest("ReleaseShareBandwidth", params, response)

	return response, err
}

type AssociateEIPWithShareBandwidthParams struct {
	ucloud.CommonRequest

	Region           string
	EIPIds           []string
	ShareBandwidthId string
}

type AssociateEIPWithShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) AssociateEIPWithShareBandwidth(params *AssociateEIPWithShareBandwidthParams) (*AssociateEIPWithShareBandwidthResponse, error) {
	response := &AssociateEIPWithShareBandwidthResponse{}
	err := u.DoRequest("AssociateEIPWithShareBandwidth", params, response)

	return response, err
}

type DisassociateEIPWithShareBandwidthParams struct {
	ucloud.CommonRequest

	Region           string
	EIPIds           []string
	ShareBandwidthId string
	Bandwidth        int
}

type DisassociateEIPWithShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) DisassociateEIPWithShareBandwidth(params *DisassociateEIPWithShareBandwidthParams) (*DisassociateEIPWithShareBandwidthResponse, error) {
	response := &DisassociateEIPWithShareBandwidthResponse{}
	err := u.DoRequest("DisassociateEIPWithShareBandwidth", params, response)

	return response, err
}

type DescribeBandwidthUsageParams struct {
	ucloud.CommonRequest

	OffSet int
	Limit  int
	EIPIds []string
}

type UnetBandwidthUsageEIPSet struct {
	CurBandwidth float64
	EIPId        string
}
type DescribeBandwidthUsageResponse struct {
	ucloud.CommonResponse

	TotalCount int
	EIPSet     *[]UnetBandwidthUsageEIPSet
}

func (u *UNet) DescribeBandwidthUsage(params *DescribeBandwidthUsageParams) (*DescribeBandwidthUsageResponse, error) {
	response := &DescribeBandwidthUsageResponse{}
	err := u.DoRequest("DescribeBandwidthUsage", params, response)

	return response, err
}

type DescribeFirewallParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	FWId      string
	OffSet    int
	Limit     int
}

type FirewallDataSet struct {
	Name          string
	Tag           string
	Remark        string
	ResourceCount int
	CreateTime    string
	Type          string
	Rule          *[]FirewallRuleSet
}

type FirewallRuleSet struct {
	SrcIP        string
	Priority     int
	ProtocolType string
	DstPort      string
	RuleAction   string
}
type DescribeFirewallResponse struct {
	ucloud.CommonResponse

	DataSet *[]FirewallDataSet
}

func (u *UNet) DescribeFirewall(params *DescribeFirewallParams) (*DescribeFirewallResponse, error) {
	response := &DescribeFirewallResponse{}
	err := u.DoRequest("DescribeFirewall", params, response)

	return response, err
}

type UpdateFirewallParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	FWId      string
	Rule      []string
}

type UpdateFirewallResponse struct {
	ucloud.CommonResponse

	FWId string
}

func (u *UNet) UpdateFirewall(params *UpdateFirewallParams) (*UpdateFirewallResponse, error) {
	response := &UpdateFirewallResponse{}
	err := u.DoRequest("UpdateFirewall", params, response)

	return response, err
}

type DescribeFirewallResourceParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	FWId      string

	Limit  string
	OffSet string
}

type ResourceSet struct {
	Name         string
	PrivateIP    string
	Remark       string
	ResourceID   string
	ResourceType string
	Status       string
	Tag          string
	Zone         string
}

type DescribeFirewallResourceResponse struct {
	ucloud.CommonResponse

	ResourceSet *[]ResourceSet
}

func (u *UNet) DescribeFirewallResource(params *DescribeFirewallResourceParams) (*DescribeFirewallResourceResponse, error) {
	response := &DescribeFirewallResourceResponse{}
	err := u.DoRequest("DescribeFirewallResource", params, response)

	return response, err
}

type DeleteFirewallParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	FWId      string
}

type DeleteFirewallResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) DeleteFirewall(params *DeleteFirewallParams) (*DeleteFirewallResponse, error) {
	response := &DeleteFirewallResponse{}
	err := u.DoRequest("DeleteFirewall", params, response)

	return response, err
}

type GrantFirewallParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	FWId      string

	ResourceType string
	ResourceId   string
}

type GrantFirewallResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) GrantFirewall(params *GrantFirewallParams) (*GrantFirewallResponse, error) {
	response := &GrantFirewallResponse{}
	err := u.DoRequest("GrantFirewall", params, response)

	return response, err
}

type CreateFirewallParams struct {
	ucloud.CommonRequest

	Region    string
	ProjectId string
	Rule      []string
	Name      string
	Tag       string
	Remark    string
}

type CreateFirewallResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) CreateFirewall(params *CreateFirewallParams) (*CreateFirewallResponse, error) {
	response := &CreateFirewallResponse{}
	err := u.DoRequest("CreateFirewall", params, response)

	return response, err
}
