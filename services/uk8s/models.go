// Code is generated by ucloud-model, DO NOT EDIT IT.

package uk8s

/*
ImageInfo - UK8S 可用镜像信息
*/
type ImageInfo struct {

	// 镜像 Id
	ImageId string

	// 镜像名称
	ImageName string

	// 该镜像是否支持GPU机型，枚举值[true:不支持，false:支持]。
	NotSupportGPU bool

	// 可用区 Id
	ZoneId int
}

/*
KubeProxy - KubeProxy信息
*/
type KubeProxy struct {

	// KubeProxy模式，枚举值为[ipvs,iptables]
	Mode string
}

/*
UHostIPSet - 云主机IP信息
*/
type UHostIPSet struct {

	// IP对应的带宽, 单位: Mb (内网IP不显示带宽信息)
	Bandwidth int

	// IP地址
	IP string

	// IP资源ID (内网IP无对应的资源ID)
	IPId string

	// Mac地址
	Mac string

	// IP地址对应的子网 ID
	SubnetId string

	// 国际: Internation，BGP: Bgp，内网: Private
	Type string

	// IP地址对应的VPC ID
	VPCId string
}

/*
NodeInfoV2 - UK8S 节点信息
*/
type NodeInfoV2 struct {

	// 节点所属伸缩组ID，非伸缩组创建出来的节点，伸缩组ID为Default。
	AsgId string

	// Node节点CPU核数，单位: 个。
	CPU int

	// 节点创建时间
	CreateTime int

	// 节点计费到期时间
	ExpireTime int

	// 节点的GPU颗数。
	GPU int

	// 节点IP信息，详细信息见 UHostIPSet。
	IPSet []UHostIPSet

	// 资源ID，如uhost-xxxx，或uphost-xxxxx。
	InstanceId string

	// 资源名称，初始值等于NodeId，用户可在UHost或UPHost处修改。
	InstanceName string

	// Node节点的资源类型，枚举值为UHost或UPHost。
	InstanceType string

	// kubeproxy信息，详细信息见KubeProxy。
	KubeProxy KubeProxy

	// 机型类别，分别对应Uhost的MachineType或PHost的PHostType。
	MachineType string

	// 内存大小，单位: MB。
	Memory int

	// NodeId，Node在UK8S处的唯一标示，如uk8s-reewqe5-sdasadsda
	NodeId string

	// node角色，枚举值为master、node
	NodeRole string

	// Node的状态
	NodeStatus string

	// Node节点的镜像名称。
	OsName string

	// Node节点的操作系统类别，如Linux或Windows。
	OsType string

	// 是否允许Pod调度到该节点，枚举值为true或false。
	Unschedulable bool

	// Node所在可用区
	Zone string
}

/*
ClusterSet - 集群信息
*/
type ClusterSet struct {

	// 集群apiserver地址
	ApiServer string

	// 集群ID
	ClusterId string

	// 资源名字
	ClusterName string

	// 创建时间
	CreateTime int

	// 集群外部apiserver地址
	ExternalApiServer string

	// 集群版本
	K8sVersion string

	// Master 节点数量
	MasterCount int

	// Node节点数量
	NodeCount int

	// Pod网段
	PodCIDR string

	// 服务网段
	ServiceCIDR string

	// 状态
	Status string

	// 所属子网
	SubnetId string

	// 所属VPC
	VPCId string
}
