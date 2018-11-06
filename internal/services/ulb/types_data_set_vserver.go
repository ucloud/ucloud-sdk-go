package ulb

/*
DataSetVServer - GetVServerMetricInfo的DataSet

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type DataSetVServer struct {

	// 所添加的后端资源ID，（为ULB系统中使用，与资源自身ID无关）
	BackendId string

	// 监听器类型RequestProxy或者PacketsTransmit
	ListenType string

	// 后端节点名称
	Name string

	// 后端节点业务组
	Tag string

	// 后端节点监听端口
	Port int

	// 资源系统记录的后端节点id
	MetricId string

	// VServer 的 ID
	VServerId string
}
