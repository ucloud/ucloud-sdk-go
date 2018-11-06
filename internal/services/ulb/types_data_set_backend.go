package ulb

/*
DataSetBackend - GetBackendMetricInfo的DataSet

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type DataSetBackend struct {

	// 所添加的后端资源ID，（为ULB系统中使用，与资源自身ID无关）
	BackendId string

	// 资源系统记录的后端节点id
	MetricId string

	// 后端节点名称
	Name string

	// 后端节点业务组
	Tag string

	// 后端节点备注
	Remark string

	// 后端节点监听端口
	Port int

	// 后端节点内网ip
	PrivateIP string
}
