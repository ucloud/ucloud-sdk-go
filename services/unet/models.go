// Code is generated by ucloud-model, DO NOT EDIT IT.



package unet











/*
UnetEIPAddrSet - 
*/
type UnetEIPAddrSet struct {
	
	// 
	IP string 
	
	// 
	OperatorName string 
	
}



/*
UnetAllocateEIPSet - 
*/
type UnetAllocateEIPSet struct {
	
	// 
	EIPAddr []UnetEIPAddrSet 
	
	// 
	EIPId string 
	
}



/*
EIPAddrSet - 
*/
type EIPAddrSet struct {
	
	// 
	IP string 
	
	// 
	OperatorName string 
	
}



/*
UnetBandwidthPackageSet - 
*/
type UnetBandwidthPackageSet struct {
	
	// 
	Bandwidth int 
	
	// 
	BandwidthPackageId string 
	
	// 
	CreateTime int 
	
	// 
	DisableTime int 
	
	// 
	EIPAddr []EIPAddrSet 
	
	// 
	EIPId string 
	
	// 
	EnableTime int 
	
}



/*
UnetBandwidthUsageEIPSet - 
*/
type UnetBandwidthUsageEIPSet struct {
	
	// 
	CurBandwidth float64 
	
	// 
	EIPId string 
	
}



/*
UnetEIPResourceSet - 
*/
type UnetEIPResourceSet struct {
	
	// 
	EIPId string 
	
	// 
	ResourceId string 
	
	// 
	ResourceName string 
	
	// 
	ResourceType string 
	
}



/*
ShareBandwidthSet - 
*/
type ShareBandwidthSet struct {
	
	// 
	ShareBandwidth int 
	
	// 
	ShareBandwidthId string 
	
	// 
	ShareBandwidthName string 
	
}



/*
UnetEIPSet - 
*/
type UnetEIPSet struct {
	
	// 
	Bandwidth int 
	
	// 
	BandwidthType int 
	
	// 
	ChargeType string 
	
	// 
	CreateTime int 
	
	// 
	EIPAddr []UnetEIPAddrSet 
	
	// 
	EIPId string 
	
	// 
	Expire bool 
	
	// 
	ExpireTime int 
	
	// 
	Name string 
	
	// 
	PayMode string 
	
	// 
	Remark string 
	
	// 
	Resource UnetEIPResourceSet 
	
	// 
	ShareBandwidthSet ShareBandwidthSet 
	
	// 
	Status string 
	
	// 
	Tag string 
	
	// 
	Weight int 
	
}



/*
FirewallRuleSet - 
*/
type FirewallRuleSet struct {
	
	// 
	DstPort string 
	
	// 
	Priority string 
	
	// 
	ProtocolType string 
	
	// 
	Remark string 
	
	// 
	RuleAction string 
	
	// 
	SrcIP string 
	
}



/*
FirewallDataSet - 
*/
type FirewallDataSet struct {
	
	// 
	CreateTime int 
	
	// 
	FWId string 
	
	// 
	GroupId string 
	
	// 
	Name string 
	
	// 
	Remark string 
	
	// 
	ResourceCount int 
	
	// 
	Rule []FirewallRuleSet 
	
	// 
	Tag string 
	
	// 
	Type string 
	
}



/*
ResourceSet - 
*/
type ResourceSet struct {
	
	// 
	Name string 
	
	// 
	PrivateIP string 
	
	// 
	Remark string 
	
	// 
	ResourceID string 
	
	// 
	ResourceType string 
	
	// 
	Status int 
	
	// 
	Tag string 
	
	// 
	Zone int 
	
}



/*
EIPSetData - 
*/
type EIPSetData struct {
	
	// 
	Bandwidth int 
	
	// 
	EIPAddr []EIPAddrSet 
	
	// 
	EIPId string 
	
}



/*
UnetShareBandwidthSet - 
*/
type UnetShareBandwidthSet struct {
	
	// 
	BandwidthGuarantee int 
	
	// 
	ChargeType string 
	
	// 
	CreateTime int 
	
	// 
	EIPSet []EIPSetData 
	
	// 
	ExpireTime int 
	
	// 
	Name string 
	
	// 
	PostPayStartTime int 
	
	// 
	ShareBandwidth int 
	
	// 
	ShareBandwidthId string 
	
}



/*
EIPPayModeSet - 
*/
type EIPPayModeSet struct {
	
	// 
	EIPId string 
	
	// 
	EIPPayMode string 
	
}



/*
EIPPriceDetailSet - 
*/
type EIPPriceDetailSet struct {
	
	// 
	ChargeType string 
	
	// 
	Price float64 
	
	// 
	PurchaseValue int 
	
}


