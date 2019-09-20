package vpc

/*
ResourceInfo - 资源信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ResourceInfo struct {
	// 资源ip
	IP string

	// 资源的IPv6地址
	IPv6Address string

	// 资源名称
	Name string

	// 资源Id
	ResourceId string

	// 资源类型。对应的资源类型：UHOST，云主机；PHOST，物理云主机；ULB，负载均衡；UHADOOP_HOST，hadoop节点；UFORTRESS_HOST，堡垒机；UNATGW，NAT网关；UKAFKA，Kafka消息队列；UMEM，内存存储；DOCKER，容器集群；UDB，数据库；UDW，数据仓库；VIP，内网VIP.
	ResourceType string

	// 资源绑定的虚拟网卡的实例ID
	SubResourceId string

	// 资源绑定的虚拟网卡的实例名称
	SubResourceName string

	// 资源绑定的虚拟网卡的类型
	SubResourceType string
}
