package unet

/*
UnetSecurityGroupSet - DescribeSecurityGroup

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type UnetSecurityGroupSet struct {

	// 防火墙ID，可选，为空时不显示Rule
	GroupId string

	// 防火墙组的名称
	GroupName string

	// 防火墙组创建时间，格式为Unix Timestamp
	CreateTime int

	// 防火墙组类型，枚举值为： 0：用户自定义防火墙； 1：默认Web防火墙； 2：默认非Web防火墙
	Type int

	// 防火墙组中的规则表,参见 SecurityGroupRuleSet
	Rule []SecurityGroupRuleSet
}
