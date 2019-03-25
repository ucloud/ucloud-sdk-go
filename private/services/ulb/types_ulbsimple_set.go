package ulb

/*
ULBSimpleSet - ulb简明信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ULBSimpleSet struct {

	// 负载均衡的资源ID
	ULBId string

	// 负载均衡的资源名称（资源系统中），缺省值“ULB”
	Name string

	// 负载均衡的业务组名称，缺省值“Default”
	Tag string

	// 负载均衡的备注，缺省值“”
	Remark string

	// ULB的创建时间，格式为Unix Timestamp
	CreateTime int

	// ULB所在的VPC的ID
	VPCId string

	// ULB 为 InnerMode 时，ULB 所属的子网ID，默认为空
	SubnetId string

	// ULB 所属的业务组ID
	BusinessId string

	// ULB的内网IP，当ULBType为OuterMode时，该值为空
	PrivateIP string

	// 带宽类型，枚举值为： 0，非共享带宽； 1，共享带宽
	BandwidthType int

	// 带宽
	Bandwidth int

	// ULB的详细信息列表，具体结构见下方 ULBIPSet
	IPSet []ULBIPSet

	// ulb下vserver数量
	VServerCount int

	// ULB 的类型（InnerMode or OuterMode）
	ULBType string
}
