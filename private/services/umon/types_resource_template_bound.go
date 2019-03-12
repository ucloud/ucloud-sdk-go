package umon

/*
ResourceTemplateBound - 资源告警模板绑定关系

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ResourceTemplateBound struct {

	// 是否绑定了告警模板，Yes：是；No：否。若为No，不展示AlarmTemplateId、AlarmTemplateName、Remark
	HasAlarmTemlate string

	// 资源id
	ResourceId string

	// 告警模板id
	AlarmTemplateId int

	// 告警模板名称
	AlarmTemplateName string

	// 是否默认模板，Yes：是；No：否
	IsDefault string

	// 模板备注
	Remark string
}
