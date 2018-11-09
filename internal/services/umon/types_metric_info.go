package umon

/*
MetricInfo - GetMetricInfo-监控项信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type MetricInfo struct {

	// 指标名称
	MetricName string

	// 单位
	Unit string

	// 仅限内部使用
	ConversionFactor int

	// 数据类型
	Type string

	// 指标组名称，相同指标组前端可合并一张图显示
	MetricGroup string

	// 是否支持告警
	SupportAlarm string

	// 告警设置范围
	AlarmRange interface{} // Backend Bug: backend use dynamic type, object or string

	// 仅限内部使用
	Frequency int

	// 比较参数，可选GE,LE
	CompareOption []string
}
