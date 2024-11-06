// Code is generated by ucloud-model, DO NOT EDIT IT.

package uk8s

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UK8S API Schema

// AddUK8SExistingUHostRequest is request schema for AddUK8SExistingUHost action
type AddUK8SExistingUHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// UK8S集群ID。 可从UK8S控制台获取。
	ClusterId *string `required:"true"`

	// 用于标示添加完节点后是否将节点临时禁用. 传入 "true" 表示禁用,传入其它或不传表示不禁用
	DisableSchedule *bool `required:"false"`

	// 镜像 Id，不填时后台程序会自动选用一个可用的镜像 Id，支持用户自定义镜像，自定义镜像必须基于基础镜像制作。
	ImageId *string `required:"false"`

	// 用户自定义Shell脚本。与UserData的区别在于InitScript在节点初始化完毕后才执行，UserData则是云主机初始化时执行。
	InitScript *string `required:"false"`

	// Node节点标签。key=value形式,多组用”,“隔开，最多5组。 如env=pro,type=game
	Labels *string `required:"false"`

	// 默认110，生产环境建议小于等于110。
	MaxPods *int `required:"false"`

	// Node节点密码。请遵照[[api:uhost-api:specification|字段规范]]设定密码。密码需使用base64进行编码，如下：# echo -n Password1 | base64
	Password *string `required:"true"`

	// 该云主机所属子网Id。
	SubnetId *string `required:"false"`

	// 云主机Id，为了保证节点正常运行，该主机配置不得低于2C4G。
	UHostId *string `required:"true"`

	// 用户自定义数据。当镜像支持Cloud-init Feature时可填写此字段。注意：1、总数据量大小不超过 16K；2、使用base64编码。
	UserData *string `required:"false"`
}

// AddUK8SExistingUHostResponse is response schema for AddUK8SExistingUHost action
type AddUK8SExistingUHostResponse struct {
	response.CommonBase

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string
}

// NewAddUK8SExistingUHostRequest will create request of AddUK8SExistingUHost action.
func (c *UK8SClient) NewAddUK8SExistingUHostRequest() *AddUK8SExistingUHostRequest {
	req := &AddUK8SExistingUHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: AddUK8SExistingUHost

将预先创建好的云主机加入到UK8S集群，需要注意的是，该云主机依然会执行重装系统的操作。
*/
func (c *UK8SClient) AddUK8SExistingUHost(req *AddUK8SExistingUHostRequest) (*AddUK8SExistingUHostResponse, error) {
	var err error
	var res AddUK8SExistingUHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddUK8SExistingUHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// AddUK8SNodeGroupRequest is request schema for AddUK8SNodeGroup action
type AddUK8SNodeGroupRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"false"`

	// 系统盘大小，单位GB。默认40。范围：[40, 500]。注意SSD本地盘无法调整。
	BootDiskSize *int `required:"false"`

	// 磁盘类型
	BootDiskType *string `required:"false"`

	// GPU卡核心数。仅GPU机型支持此字段（可选范围与MachineType+GpuType相关）
	CPU *int `required:"false"`

	// 计费模式
	ChargeType *string `required:"false"`

	// 集群ID
	ClusterId *string `required:"true"`

	// 数据磁盘大小
	DataDiskSize *int `required:"false"`

	// 磁盘类型
	DataDiskType *string `required:"false"`

	// GPU卡核心数
	GPU *int `required:"false"`

	// GPU类型
	GpuType *string `required:"false"`

	// 镜像ID
	ImageId *string `required:"false"`

	// 云主机机型。枚举值["N", "C", "G", "O", "OS"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 内存大小。单位：MB
	Mem *int `required:"false"`

	// 最低cpu平台，枚举值["Intel/Auto", "Intel/IvyBridge", "Intel/Haswell", "Intel/Broadwell", "Intel/Skylake", "Intel/Cascadelake"；"Intel/CascadelakeR"; “Amd/Epyc2”,"Amd/Auto"],默认值是"Intel/Auto"
	MinimalCpuPlatform *string `required:"false"`

	// 节点池名字
	NodeGroupName *string `required:"true"`

	// 子网 ID。默认为集群创建时填写的子网ID，也可以填写集群同VPC内的子网ID。
	SubnetId *string `required:"false"`

	// 业务组
	Tag *string `required:"false"`
}

// AddUK8SNodeGroupResponse is response schema for AddUK8SNodeGroup action
type AddUK8SNodeGroupResponse struct {
	response.CommonBase

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string

	// 节点池ID
	NodeGroupId string
}

// NewAddUK8SNodeGroupRequest will create request of AddUK8SNodeGroup action.
func (c *UK8SClient) NewAddUK8SNodeGroupRequest() *AddUK8SNodeGroupRequest {
	req := &AddUK8SNodeGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: AddUK8SNodeGroup

添加UK8S节点池
*/
func (c *UK8SClient) AddUK8SNodeGroup(req *AddUK8SNodeGroupRequest) (*AddUK8SNodeGroupResponse, error) {
	var err error
	var res AddUK8SNodeGroupResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddUK8SNodeGroup", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// AddUK8SPHostNodeRequest is request schema for AddUK8SPHostNode action
type AddUK8SPHostNodeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 计费模式。枚举值为： \\ > Year，按年付费； \\ > Month，按月付费；\\ 默认为月付
	ChargeType *string `required:"true"`

	// UK8S集群ID。 可从UK8S控制台获取。
	ClusterId *string `required:"true"`

	// 最大创建Node节点数量，取值范围是[1,10]。
	Count *int `required:"true"`

	// 用于标示添加完节点后是否将节点临时禁用. 传入 "true" 表示禁用,传入其它或不传表示不禁用
	DisableSchedule *bool `required:"false"`

	// 镜像 Id，不填时后台程序会自动选用一个可用的镜像 Id，支持用户自定义镜像，自定义镜像必须基于基础镜像制作。
	ImageId *string `required:"false"`

	// 用户自定义Shell脚本。与UserData的区别在于InitScript在节点初始化完毕后才执行。
	InitScript *string `required:"false"`

	// Node节点标签。key=value形式,多组用”,“隔开，最多5组。 如env=pro,type=game
	Labels *string `required:"false"`

	// 默认110，生产环境建议小于等于110。
	MaxPods *int `required:"false"`

	// 网络环境，可选千兆：1G ，万兆：10G， 默认1G。
	NIC *string `required:"false"`

	// Node节点密码。请遵照[[api:uhost-api:specification|字段规范]]设定密码。密码需使用base64进行编码，如下：# echo -n Password1 | base64
	Password *string `required:"true"`

	// 购买时长。默认: 1。月付时，此参数传0，代表了购买至月末。
	Quantity *int `required:"false"`

	// Raid配置，默认Raid10 支持:Raid0、Raid1、Raid5、Raid10，NoRaid
	Raid *string `required:"false"`

	// 子网 ID。默认为集群创建时填写的子网ID，也可以填写集群同VPC内的子网ID。
	SubnetId *string `required:"false"`

	// 物理机类型，默认为：db-2(基础型-SAS-V3)
	Type *string `required:"false"`
}

// AddUK8SPHostNodeResponse is response schema for AddUK8SPHostNode action
type AddUK8SPHostNodeResponse struct {
	response.CommonBase

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string
}

// NewAddUK8SPHostNodeRequest will create request of AddUK8SPHostNode action.
func (c *UK8SClient) NewAddUK8SPHostNodeRequest() *AddUK8SPHostNodeRequest {
	req := &AddUK8SPHostNodeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: AddUK8SPHostNode

为UK8S集群添加一台或多台物理云主机类型的节点。
*/
func (c *UK8SClient) AddUK8SPHostNode(req *AddUK8SPHostNodeRequest) (*AddUK8SPHostNodeResponse, error) {
	var err error
	var res AddUK8SPHostNodeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddUK8SPHostNode", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// AddUK8SUHostNodeRequest is request schema for AddUK8SUHostNode action
type AddUK8SUHostNodeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 系统盘大小，单位GB。默认40。范围：[40, 500]。注意SSD本地盘无法调整。
	BootDiskSize *int `required:"false"`

	// 磁盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	BootDiskType *string `required:"false"`

	// 虚拟CPU核数。可选参数：2-64（具体机型与CPU的对应关系参照控制台）。默认值: 4。
	CPU *int `required:"true"`

	// 计费模式。枚举值为： \\ > Year，按年付费； \\ > Month，按月付费；\\ > Dynamic，按小时预付费 \\ > Postpay，按小时后付费（支持关机不收费，目前仅部分可用区支持，请联系您的客户经理） \\ 默认为月付
	ChargeType *string `required:"true"`

	// UK8S集群ID。 可从UK8S控制台获取。
	ClusterId *string `required:"true"`

	// 创建Node节点数量，取值范围是[1,50]。
	Count *int `required:"true"`

	// 数据磁盘大小，单位GB。默认0。范围 ：[20, 1000]
	DataDiskSize *int `required:"false"`

	// 磁盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	DataDiskType *string `required:"false"`

	// 用于标示添加完节点后是否将节点临时禁用. 传入 "true" 表示禁用,传入其它或不传表示不禁用
	DisableSchedule *bool `required:"false"`

	// GPU卡核心数。仅GPU机型支持此字段（可选范围与MachineType+GpuType相关）
	GPU *int `required:"false"`

	// GPU类型，枚举值["K80", "P40", "V100",]，MachineType为G时必填
	GpuType *string `required:"false"`

	// 镜像 Id，不填时后台程序会自动选用一个可用的镜像 Id，支持用户自定义镜像，自定义镜像必须基于基础镜像制作。
	ImageId *string `required:"false"`

	// 用户自定义Shell脚本。与UserData的区别在于InitScript在节点初始化完毕后才执行，UserData则是云主机初始化时执行。
	InitScript *string `required:"false"`

	// 硬件隔离组id。可通过DescribeIsolationGroup获取。
	IsolationGroup *string `required:"false"`

	// Node节点标签。key=value形式,多组用”,“隔开，最多5组。 如env=pro,type=game
	Labels *string `required:"false"`

	// 云主机机型。枚举值["N", "C", "G", "O", "OS"]。参考[[api:uhost-api:uhost_type|云主机机型说明]]。
	MachineType *string `required:"false"`

	// 默认110，生产环境建议小于等于110。
	MaxPods *int `required:"false"`

	// 内存大小。单位：MB。范围 ：[4096, 262144]，取值为1024的倍数（可选范围参考控制台）。默认值：8192
	Mem *int `required:"true"`

	// 最低cpu平台，枚举值["Intel/Auto", "Intel/IvyBridge", "Intel/Haswell", "Intel/Broadwell", "Intel/Skylake", "Intel/Cascadelake"；"Intel/CascadelakeR"; “Amd/Epyc2”,"Amd/Auto"],默认值是"Intel/Auto"
	MinimalCpuPlatform *string `required:"false"`

	// 【该字段已废弃，请谨慎使用】
	MinmalCpuPlatform *string `required:"false" deprecated:"true"`

	// 节点池id
	NodeGroupId *string `required:"false"`

	// Node节点密码。请遵照[[api:uhost-api:specification|字段规范]]设定密码。密码需使用base64进行编码，如下：# echo -n Password1 | base64
	Password *string `required:"true"`

	// 购买时长。默认: 1。按小时购买(Dynamic)时无需此参数。 月付时，此参数传0，代表了购买至月末。
	Quantity *int `required:"false"`

	// 子网 ID。默认为集群创建时填写的子网ID，也可以填写集群同VPC内的子网ID。
	SubnetId *string `required:"false"`

	// 业务组
	Tag *string `required:"false"`

	// Node节点污点，形式为key=value:effect，多组taints用”,“隔开,最多支持五组。
	Taints *string `required:"false"`

	// 用户自定义数据。当镜像支持Cloud-init Feature时可填写此字段。注意：1、总数据量大小不超过 16K；2、使用base64编码。
	UserData *string `required:"false"`
}

// AddUK8SUHostNodeResponse is response schema for AddUK8SUHostNode action
type AddUK8SUHostNodeResponse struct {
	response.CommonBase

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string

	// Node实例Id集合
	NodeIds []string
}

// NewAddUK8SUHostNodeRequest will create request of AddUK8SUHostNode action.
func (c *UK8SClient) NewAddUK8SUHostNodeRequest() *AddUK8SUHostNodeRequest {
	req := &AddUK8SUHostNodeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: AddUK8SUHostNode

为UK8S集群添加一台Node节点，机型类型为云主机
*/
func (c *UK8SClient) AddUK8SUHostNode(req *AddUK8SUHostNodeRequest) (*AddUK8SUHostNodeResponse, error) {
	var err error
	var res AddUK8SUHostNodeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AddUK8SUHostNode", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

/*
CreateUK8SClusterV2ParamKubeProxy is request schema for complex param
*/
type CreateUK8SClusterV2ParamKubeProxy struct {

	// 集群kube-proxy模式。支持iptables和ipvs，默认为iptables。
	Mode *string `required:"false"`
}

/*
CreateUK8SClusterV2ParamMaster is request schema for complex param
*/
type CreateUK8SClusterV2ParamMaster struct {

	// Master节点所属可用区，需要设置 Master.0.Zone、 Master.1.Zone、Master.2.Zone 三个 Master 节点的可用区。 三个节点可部署在不同可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	Zone *string `required:"true"`
}

/*
CreateUK8SClusterV2ParamNodes is request schema for complex param
*/
type CreateUK8SClusterV2ParamNodes struct {

	// Node节点的系统盘大小，单位GB，默认为40。范围：[40, 500]。注意SSD本地盘无法调整。
	BootDiskSIze *int `required:"false"`

	// 一组Node节点的系统盘类型，请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	BootDiskType *string `required:"false"`

	// 一组Node节点的虚拟CPU核数。单位：核，范围：[2, 64]，可选范围参考控制台。
	CPU *int `required:"true"`

	// 一组Node节点的数量，范围：[1,10]。
	Count *int `required:"true"`

	// 数据磁盘大小，单位GB。默认0。范围 ：[20, 1000]
	DataDiskSize *int `required:"false"`

	// 一组Node节点的数据盘类型，请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	DataDiskType *string `required:"false"`

	// 一组Node节点的GPU卡核心数，仅GPU机型支持此字段。
	GPU *int `required:"false"`

	// 一组Node节点的GPU类型，枚举值["K80", "P40", "V100"]，最新值参考Console。
	GpuType *string `required:"false"`

	// 一组Node节点的隔离组Id，归属于同一隔离组的虚拟机节点将落在不同的物理机上，单个隔离组最多只能容纳8个节点。参见DescribeIsolationGroup。
	IsolationGroup *string `required:"false"`

	// Node节点标签，形式为key=value，多组Labels用”,“隔开,最多支持五组。
	Labels *string `required:"false"`

	// 一组Nodes节点云主机机型，如["N", "C", "O", "OS"]，具体请参照云主机机型。
	MachineType *string `required:"true"`

	// Node节点上可运行最大节点数，默认为110。
	MaxPods *int `required:"false"`

	// 一组Node节点的内存大小。单位：MB,范围 ：[4096, 262144]，取值为1024的倍数，可选范围参考控制台。
	Mem *int `required:"true"`

	// Node节点的最低cpu平台，不选则随机。枚举值["Intel/Auto", "Intel/IvyBridge", "Intel/Haswell", "Intel/Broadwell", "Intel/Skylake", "Intel/Cascadelake"。
	MinimalCpuPlatform *string `required:"false"`

	// 【该字段已废弃，请谨慎使用】
	MinmalCpuPlatform *string `required:"false" deprecated:"true"`

	// Node节点污点，形式为key=value:effect，多组taints用”,“隔开,最多支持五组。
	Taints *string `required:"false"`

	// 一组Nodes节点所属可用区，可创建多组Nodes节点，如一组是CPU Nodes节点，另一组是GPU Nodes节点。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	Zone *string `required:"true"`
}

// CreateUK8SClusterV2Request is request schema for CreateUK8SClusterV2 action
type CreateUK8SClusterV2Request struct {
	request.CommonBase

	// [公共参数] 项目ID。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// 集群所有节点的付费模式。枚举值为： Year，按年付费； Month，按月付费； Dynamic，按小时付费（需开启权限），默认按月。
	ChargeType *string `required:"false"`

	// 创建集群的时候定义clusterdomain
	ClusterDomain *string `required:"false"`

	// 集群名称
	ClusterName *string `required:"true"`

	// 是否允许外网访问apiserver，开启：Yes 不开启：No。默认为No。
	ExternalApiServer *string `required:"false"`

	// Master节点和Node节点的镜像 ID，不填则随机选择可用的基础镜像。支持用户自定义镜像。
	ImageId *string `required:"false"`

	// 用户自定义脚本，与UserData不同，自定义脚本将在集群安装完毕后执行。注意：1、总数据量大小不超多16K；2、使用base64编码。
	InitScript *string `required:"false"`

	// k8s集群的版本，版本信息请参考UK8S集群创建页，不指定的话默认为当前支持的最高版本。
	K8sVersion *string `required:"false"`

	//
	KubeProxy *CreateUK8SClusterV2ParamKubeProxy `required:"false"`

	//
	Master []CreateUK8SClusterV2ParamMaster `required:"false"`

	// Master节点系统盘大小，单位GB，默认为40。范围：[40, 500]。注意SSD本地盘无法调整。
	MasterBootDiskSize *int `required:"false"`

	// Master节点系统盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	MasterBootDiskType *string `required:"false"`

	// Master节点的虚拟CPU核数。可选参数：2-64（具体机型与CPU的对应关系参照控制台）。
	MasterCPU *int `required:"true"`

	// Master节点的数据盘大小，单位GB，默认为0。范围 ：[20, 1000]
	MasterDataDiskSize *int `required:"false"`

	// Master节点数据盘类型。请参考[[api:uhost-api:disk_type|磁盘类型]]。默认为SSD云盘
	MasterDataDiskType *string `required:"false"`

	// 【无效，已删除】当前将自动为Master节点创建隔离组，确保Master节点归属于不同物理机。
	MasterIsolationGroup *string `required:"false"`

	// Master节点的云主机机型（V2.0），如["N", "C", "O", "OS"]，具体请参照云主机机型。
	MasterMachineType *string `required:"true"`

	// Master节点的内存大小。单位：MB。范围 ：[4096, 262144]，取值为1024的倍数（可选范围参考控制台）。
	MasterMem *int `required:"true"`

	// Master节点的最低cpu平台，不选则随机。枚举值["Intel/Auto", "Intel/IvyBridge", "Intel/Haswell", "Intel/Broadwell", "Intel/Skylake", "Intel/Cascadelake"。
	MasterMinimalCpuPlatform *string `required:"false"`

	// 【该字段已废弃，请谨慎使用】
	MasterMinmalCpuPlatform *string `required:"false" deprecated:"true"`

	//
	Nodes []CreateUK8SClusterV2ParamNodes `required:"false"`

	// 集群节点密码，包括Master和Node。密码需包含最少一个大写字母，请使用base64进行编码，举例如下：# echo -n Password1 | base64
	Password *string `required:"true"`

	// 购买时长。默认为1。按小时购买(Dynamic)时无需此参数。 月付时，此参数传0，代表了购买至月末。
	Quantity *int `required:"false"`

	// Service 网段，用于分配ClusterIP，如172.17.0.0/16。该网段不能与集群所属VPC网段重叠。
	ServiceCIDR *string `required:"true"`

	// 集群Node及Pod所属子网
	SubnetId *string `required:"true"`

	// 业务组
	Tag *string `required:"false"`

	// 用户自定义数据。注意：1、总数据量大小不超多16K；2、使用base64编码。
	UserData *string `required:"false"`

	// 集群Node及Pod所属VPC
	VPCId *string `required:"true"`
}

// CreateUK8SClusterV2Response is response schema for CreateUK8SClusterV2 action
type CreateUK8SClusterV2Response struct {
	response.CommonBase

	// 集群ID
	ClusterId string
}

// NewCreateUK8SClusterV2Request will create request of CreateUK8SClusterV2 action.
func (c *UK8SClient) NewCreateUK8SClusterV2Request() *CreateUK8SClusterV2Request {
	req := &CreateUK8SClusterV2Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateUK8SClusterV2

创建UK8S集群
*/
func (c *UK8SClient) CreateUK8SClusterV2(req *CreateUK8SClusterV2Request) (*CreateUK8SClusterV2Response, error) {
	var err error
	var res CreateUK8SClusterV2Response

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUK8SClusterV2", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DelUK8SClusterRequest is request schema for DelUK8SCluster action
type DelUK8SClusterRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 集群id
	ClusterId *string `required:"true"`

	// 是否删除节点挂载的数据盘。枚举值[true:删除，false: 不删除]，默认不删除
	ReleaseUDisk *bool `required:"false"`
}

// DelUK8SClusterResponse is response schema for DelUK8SCluster action
type DelUK8SClusterResponse struct {
	response.CommonBase
}

// NewDelUK8SClusterRequest will create request of DelUK8SCluster action.
func (c *UK8SClient) NewDelUK8SClusterRequest() *DelUK8SClusterRequest {
	req := &DelUK8SClusterRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DelUK8SCluster

删除UK8S集群
*/
func (c *UK8SClient) DelUK8SCluster(req *DelUK8SClusterRequest) (*DelUK8SClusterResponse, error) {
	var err error
	var res DelUK8SClusterResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DelUK8SCluster", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DelUK8SClusterNodeV2Request is request schema for DelUK8SClusterNodeV2 action
type DelUK8SClusterNodeV2Request struct {
	request.CommonBase

	// [公共参数] 项目ID项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// UK8S集群ID。 可从UK8S控制台获取。
	ClusterId *string `required:"true"`

	// Node在UK8S处的唯一标示，如uk8s-reewqe5-sdasadsda。**非云主机或物理云主机资源Id**
	NodeId *string `required:"true"`

	// 删除节点时是否释放数据盘。 枚举值[true:释放，false: 不释放]，默认为true。
	ReleaseDataUDisk *bool `required:"false"`
}

// DelUK8SClusterNodeV2Response is response schema for DelUK8SClusterNodeV2 action
type DelUK8SClusterNodeV2Response struct {
	response.CommonBase

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string
}

// NewDelUK8SClusterNodeV2Request will create request of DelUK8SClusterNodeV2 action.
func (c *UK8SClient) NewDelUK8SClusterNodeV2Request() *DelUK8SClusterNodeV2Request {
	req := &DelUK8SClusterNodeV2Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DelUK8SClusterNodeV2

删除集群中的Node节点，删除前务必先将其中的Pod驱逐。
*/
func (c *UK8SClient) DelUK8SClusterNodeV2(req *DelUK8SClusterNodeV2Request) (*DelUK8SClusterNodeV2Response, error) {
	var err error
	var res DelUK8SClusterNodeV2Response

	reqCopier := *req

	err = c.Client.InvokeAction("DelUK8SClusterNodeV2", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUK8SClusterRequest is request schema for DescribeUK8SCluster action
type DescribeUK8SClusterRequest struct {
	request.CommonBase

	// [公共参数] 项目id
	// ProjectId *string `required:"false"`

	// [公共参数] 所属区域
	// Region *string `required:"true"`

	// k8s集群ID
	ClusterId *string `required:"true"`
}

// DescribeUK8SClusterResponse is response schema for DescribeUK8SCluster action
type DescribeUK8SClusterResponse struct {
	response.CommonBase

	// 集群apiserver地址
	ApiServer string

	// 集群CA根证书
	CACert string

	// 自定义或者默认的clusterdomain
	ClusterDomain string

	// 集群ID
	ClusterId string

	// 资源名字
	ClusterName string

	// 创建时间
	CreateTime int

	// 集群etcd服务证书
	EtcdCert string

	// 集群etcd服务密钥
	EtcdKey string

	// 集群外部apiserver地址
	ExternalApiServer string

	// kube-proxy配置
	KubeProxy KubeProxy

	// Master 节点数量
	MasterCount int

	// Master节点配置信息，具体参考UhostInfo。托管版不返回该信息
	MasterList []UhostInfo

	// Master配置预警：Normal正常；Warning 需要升级；Error    需要紧急升级；
	MasterResourceStatus string

	// Node节点数量
	NodeCount int

	// Node节点配置信息,具体参考UhostInfo
	NodeList []UhostInfo

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

	// K8S版本
	Version string
}

// NewDescribeUK8SClusterRequest will create request of DescribeUK8SCluster action.
func (c *UK8SClient) NewDescribeUK8SClusterRequest() *DescribeUK8SClusterRequest {
	req := &DescribeUK8SClusterRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUK8SCluster

获取集群信息
*/
func (c *UK8SClient) DescribeUK8SCluster(req *DescribeUK8SClusterRequest) (*DescribeUK8SClusterResponse, error) {
	var err error
	var res DescribeUK8SClusterResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUK8SCluster", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUK8SImageRequest is request schema for DescribeUK8SImage action
type DescribeUK8SImageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

}

// DescribeUK8SImageResponse is response schema for DescribeUK8SImage action
type DescribeUK8SImageResponse struct {
	response.CommonBase

	// 虚拟机可用镜像集合, 详见ImageInfo 数组
	ImageSet []ImageInfo

	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息。
	Message string

	// 物理机可用镜像集合, 详见ImageInfo 数组
	PHostImageSet []ImageInfo
}

// NewDescribeUK8SImageRequest will create request of DescribeUK8SImage action.
func (c *UK8SClient) NewDescribeUK8SImageRequest() *DescribeUK8SImageRequest {
	req := &DescribeUK8SImageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUK8SImage

获取UK8S支持的Node节点操作系统，可基于该操作系统制定自定义镜像
*/
func (c *UK8SClient) DescribeUK8SImage(req *DescribeUK8SImageRequest) (*DescribeUK8SImageResponse, error) {
	var err error
	var res DescribeUK8SImageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUK8SImage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeUK8SNodeRequest is request schema for DescribeUK8SNode action
type DescribeUK8SNodeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// UK8S 集群 Id
	ClusterId *string `required:"true"`

	// K8S 节点IP或者节点ID
	Name *string `required:"true"`
}

// DescribeUK8SNodeResponse is response schema for DescribeUK8SNode action
type DescribeUK8SNodeResponse struct {
	response.CommonBase

	// 操作名称
	Action string

	// 已分配到当前节点的 Pod 数量
	AllocatedPodCount int

	// 字符串数组，每一项是类似 "node.alpha.kubernetes.io/ttl=0" 的注解
	Annotations []string

	// 节点 CPU 总量
	CPUCapacity string

	// 节点上已分配 Pod 的 CPU 限制值
	CPULimits string

	// 节点上已分配 Pod 的 CPU 限制值占 CPU 总量的比例
	CPULimitsFraction string

	// 节点上已分配 Pod 的 CPU 请求量
	CPURequests string

	// 节点上已分配 Pod 的 CPU 请求量占 CPU 总量的比例
	CPURequestsFraction string

	// 节点状态数组
	Conditions []K8SNodeCondition

	// 节点上镜像名称数组
	ContainerImages []string

	// 容器运行时版本，如："docker://18.9.9"
	ContainerRuntimeVersion string

	// 时间戳，单位是 秒
	CreationTimestamp int

	// 主机名
	Hostname string

	// 内部 IP 地址
	InternalIP string

	// 内核版本，如："4.19.0-6.el7.ucloud.x86_64"
	KernelVersion string

	// kubeproxy 版本
	KubeProxyVersion string

	// kubelet 版本
	KubeletVersion string

	// 字符串数组，每一项是类似 "kubernetes.io/arch=amd64" 的标签
	Labels []string

	// 节点内存总量
	MemoryCapacity string

	// 节点上已分配 Pod 的内存限制量
	MemoryLimits string

	// 节点上已分配 Pod 的内存限制量占内存总量的比例，如返回值为 "18"，则意味着限制量占总量的 18%
	MemoryLimitsFraction string

	// 节点上已分配 Pod 的内存请求量
	MemoryRequests string

	// 节点上已分配 Pod 的内存请求量占内存总量的比例，如返回值为 "4.5"，则意味着请求量占总量的 4.5%
	MemoryRequestsFraction string

	// 操作出错时的提示信息
	Message string

	// 节点名称
	Name string

	// 操作系统类型，如："CentOS Linux 7 (Core)"
	OSImage string

	// 节点允许的可分配 Pod 最大数量
	PodCapacity int

	// 字符串，如："UCloud://cn-sh2-02//uk8s-vsc0vgob-n-mpzxc"
	ProviderID string

	// 返回码
	RetCode int

	// 字符串数组，每一项是类似 "node-role.kubernetes.io/master:NoSchedule" 的污点
	Taints []string

	// 是否禁止调度
	Unschedulable bool
}

// NewDescribeUK8SNodeRequest will create request of DescribeUK8SNode action.
func (c *UK8SClient) NewDescribeUK8SNodeRequest() *DescribeUK8SNodeRequest {
	req := &DescribeUK8SNodeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeUK8SNode

用于获取 UK8S 节点详情
*/
func (c *UK8SClient) DescribeUK8SNode(req *DescribeUK8SNodeRequest) (*DescribeUK8SNodeResponse, error) {
	var err error
	var res DescribeUK8SNodeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeUK8SNode", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListUK8SClusterNodeV2Request is request schema for ListUK8SClusterNodeV2 action
type ListUK8SClusterNodeV2Request struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// UK8S集群ID
	ClusterId *string `required:"true"`
}

// ListUK8SClusterNodeV2Response is response schema for ListUK8SClusterNodeV2 action
type ListUK8SClusterNodeV2Response struct {
	response.CommonBase

	// 节点详细信息，见NodeInfoV2。
	NodeSet []NodeInfoV2

	// 满足条件的节点数量，包括Master。
	TotalCount int
}

// NewListUK8SClusterNodeV2Request will create request of ListUK8SClusterNodeV2 action.
func (c *UK8SClient) NewListUK8SClusterNodeV2Request() *ListUK8SClusterNodeV2Request {
	req := &ListUK8SClusterNodeV2Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUK8SClusterNodeV2

获取UK8S集群节点信息
*/
func (c *UK8SClient) ListUK8SClusterNodeV2(req *ListUK8SClusterNodeV2Request) (*ListUK8SClusterNodeV2Response, error) {
	var err error
	var res ListUK8SClusterNodeV2Response

	reqCopier := *req

	err = c.Client.InvokeAction("ListUK8SClusterNodeV2", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListUK8SClusterV2Request is request schema for ListUK8SClusterV2 action
type ListUK8SClusterV2Request struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// UK8S集群ID
	ClusterId *string `required:"false"`

	// 返回数据长度，默认为20。
	Limit *int `required:"false"`

	// 列表起始位置偏移量，默认为0。
	Offset *int `required:"false"`
}

// ListUK8SClusterV2Response is response schema for ListUK8SClusterV2 action
type ListUK8SClusterV2Response struct {
	response.CommonBase

	// 满足条件的集群数量
	ClusterCount int

	// 集群信息，具体参考ClusterSet
	ClusterSet []ClusterSet
}

// NewListUK8SClusterV2Request will create request of ListUK8SClusterV2 action.
func (c *UK8SClient) NewListUK8SClusterV2Request() *ListUK8SClusterV2Request {
	req := &ListUK8SClusterV2Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUK8SClusterV2

获取UK8S集群列表信息
*/
func (c *UK8SClient) ListUK8SClusterV2(req *ListUK8SClusterV2Request) (*ListUK8SClusterV2Response, error) {
	var err error
	var res ListUK8SClusterV2Response

	reqCopier := *req

	err = c.Client.InvokeAction("ListUK8SClusterV2", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListUK8SNodeGroupRequest is request schema for ListUK8SNodeGroup action
type ListUK8SNodeGroupRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// 集群ID
	ClusterId *string `required:"true"`
}

// ListUK8SNodeGroupResponse is response schema for ListUK8SNodeGroup action
type ListUK8SNodeGroupResponse struct {
	response.CommonBase

	// 节点池列表
	NodeGroupList []NodeGroupSet
}

// NewListUK8SNodeGroupRequest will create request of ListUK8SNodeGroup action.
func (c *UK8SClient) NewListUK8SNodeGroupRequest() *ListUK8SNodeGroupRequest {
	req := &ListUK8SNodeGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUK8SNodeGroup

列出UK8S节点池
*/
func (c *UK8SClient) ListUK8SNodeGroup(req *ListUK8SNodeGroupRequest) (*ListUK8SNodeGroupResponse, error) {
	var err error
	var res ListUK8SNodeGroupResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListUK8SNodeGroup", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RemoveUK8SNodeGroupRequest is request schema for RemoveUK8SNodeGroup action
type RemoveUK8SNodeGroupRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// 集群id
	ClusterId *string `required:"true"`

	// 节点池Id
	NodeGroupId *string `required:"true"`
}

// RemoveUK8SNodeGroupResponse is response schema for RemoveUK8SNodeGroup action
type RemoveUK8SNodeGroupResponse struct {
	response.CommonBase
}

// NewRemoveUK8SNodeGroupRequest will create request of RemoveUK8SNodeGroup action.
func (c *UK8SClient) NewRemoveUK8SNodeGroupRequest() *RemoveUK8SNodeGroupRequest {
	req := &RemoveUK8SNodeGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RemoveUK8SNodeGroup

删除UK8S节点池
*/
func (c *UK8SClient) RemoveUK8SNodeGroup(req *RemoveUK8SNodeGroupRequest) (*RemoveUK8SNodeGroupResponse, error) {
	var err error
	var res RemoveUK8SNodeGroupResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RemoveUK8SNodeGroup", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
