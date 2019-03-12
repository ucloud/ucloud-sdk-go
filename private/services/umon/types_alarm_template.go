package umon

/*
AlarmTemplate - 单个告警模板信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type AlarmTemplate struct {

	// 是否为默认模板
	IsDefault string

	// 告警模板id
	AlarmTemplateId int

	// 告警模板名称
	AlarmTemplateName string

	// 备注
	Remark string

	// 资源类型
	ResourceType string

	// 绑定的资源数量
	BoundResourceCount int
}
