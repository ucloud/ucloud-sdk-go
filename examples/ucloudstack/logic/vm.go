package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateVM 创建主机
func CreateVM(name string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// create request
	createReq := ucloudstackClient.NewCreateVMInstanceRequest()

	// 地域
	createReq.Region = ucloud.String("cn")
	createReq.Zone = ucloud.String("zone-01")

	// 配置
	createReq.ImageID = ucloud.String("cn-image-centos-74")
	createReq.CPU = ucloud.Int(2)
	createReq.Memory = ucloud.Int(4096)
	createReq.BootDiskSetType = ucloud.String("Normal")
	createReq.DataDiskSetType = ucloud.String("Normal")
	createReq.DataDiskSpace = ucloud.Int(100)
	createReq.VMType = ucloud.String("Normal02")

	// 网络
	createReq.VPCID = ucloud.String("vpc-qPwJrPoNJ")
	createReq.SubnetID = ucloud.String("subnet-qPwJrPoNJ")
	createReq.WANSGID = ucloud.String("sg-qPwJrPoNJ")

	// 认证方式
	createReq.Name = ucloud.String(name)
	createReq.Password = ucloud.String("ucloud.cn132")

	// 计费方式
	createReq.ChargeType = ucloud.String("Month")

	// send request
	newVMInstance, err := ucloudstackClient.CreateVMInstance(createReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if newVMInstance != nil {
		fmt.Printf("resource id of the VM: %+v\n", newVMInstance)
	}

}

// StopVM 关闭主机
func StopVM(vmID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// delete request
	stopReq := ucloudstackClient.NewStopVMInstanceRequest()
	stopReq.Region = ucloud.String("cn")
	stopReq.Zone = ucloud.String("zone-01")
	stopReq.VMID = ucloud.String(vmID)

	// send request
	stopResp, err := ucloudstackClient.StopVMInstance(stopReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if stopResp != nil {
		fmt.Printf("RetCode: %d\n", stopResp.RetCode)
	}

}

// DeleteVM 删除主机
func DeleteVM(vmID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// delete request
	deleteReq := ucloudstackClient.NewDeleteVMInstanceRequest()
	deleteReq.Region = ucloud.String("cn")
	deleteReq.Zone = ucloud.String("zone-01")
	deleteReq.VMID = ucloud.String(vmID)

	// send request
	delResp, err := ucloudstackClient.DeleteVMInstance(deleteReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if delResp != nil {
		fmt.Printf("RetCode: %d\n", delResp.RetCode)
	}

}

// DescribeVM 查询主机信息
func DescribeVM(vmID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeVMInstanceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.Limit = ucloud.Int(10)
	describeReq.VMIDs = []string{vmID}

	// send request
	descResp, err := ucloudstackClient.DescribeVMInstance(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if descResp.TotalCount > 0 {
		fmt.Printf("fisrt of VMs: %+v\n", descResp.Infos)
	}

}

// RestartVMInstance 重启主机
func RestartVMInstance(vmID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// RestartVMInstance request
	restartReq := ucloudstackClient.NewRestartVMInstanceRequest()
	restartReq.Region = ucloud.String("cn")
	restartReq.Zone = ucloud.String("zone-01")
	restartReq.VMID = ucloud.String(vmID)

	// send request
	restartResp, err := ucloudstackClient.RestartVMInstance(restartReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", restartResp)
		return
	}

	fmt.Printf("RestartVMInstance success, vmID: %s", vmID)
}

// ResetVMInstancePassword 重置密码
func ResetVMInstancePassword(vmID, password string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// ResetVMInstancePassword  request
	resetReq := ucloudstackClient.NewResetVMInstancePasswordRequest()
	resetReq.Region = ucloud.String("cn")
	resetReq.Zone = ucloud.String("zone-01")
	resetReq.VMID = ucloud.String(vmID)
	resetReq.Password = ucloud.String(password)

	// send request
	resetResp, err := ucloudstackClient.ResetVMInstancePassword(resetReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resetResp)
		return
	}

	fmt.Printf("ResetVMInstancePassword success, vmID: %s", vmID)
}

// ReinstallVMInstance 重装系统
func ReinstallVMInstance(vmID, imageID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// ReinstallVMInstance  request
	reinstallReq := ucloudstackClient.NewReinstallVMInstanceRequest()
	reinstallReq.Region = ucloud.String("cn")
	reinstallReq.Zone = ucloud.String("zone-01")
	reinstallReq.VMID = ucloud.String(vmID)
	reinstallReq.ImageID = ucloud.String(imageID)

	// send request
	reinstallResp, err := ucloudstackClient.ReinstallVMInstance(reinstallReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", reinstallResp)
		return
	}

	fmt.Printf("ReinstallVMInstance success, vmID: %s", vmID)
}

// ResizeVMConfig 修改系统配置
func ResizeVMConfig(vmID string, cpu, memory int) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// ResizeVMConfig  request
	resizeReq := ucloudstackClient.NewResizeVMConfigRequest()
	resizeReq.Region = ucloud.String("cn")
	resizeReq.Zone = ucloud.String("zone-01")
	resizeReq.VMID = ucloud.String(vmID)
	resizeReq.CPU = ucloud.Int(cpu)
	resizeReq.Memory = ucloud.Int(memory)

	// send request
	resizeResp, err := ucloudstackClient.ResizeVMConfig(resizeReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resizeResp)
		return
	}

	fmt.Printf("ResizeVMConfig success, vmID: %s", vmID)
}

// GetVMInstancePrice 获取虚拟机价格
func GetVMInstancePrice() {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	getVMInstancePriceReq := ucloudstackClient.NewGetVMInstancePriceRequest()
	getVMInstancePriceReq.Region = ucloud.String("cn")
	getVMInstancePriceReq.GPU = ucloud.Int(0)
	getVMInstancePriceReq.CPU = ucloud.Int(1)
	getVMInstancePriceReq.Memory = ucloud.Int(2048)
	getVMInstancePriceReq.DataDiskSpace = ucloud.Int(20)
	getVMInstancePriceReq.DataDiskSetType = ucloud.String("Normal")
	getVMInstancePriceReq.BootDiskSetType = ucloud.String("Normal")
	getVMInstancePriceReq.VMType = ucloud.String("Normal02")
	getVMInstancePriceReq.ChargeType = ucloud.String("Month")
	getVMInstancePriceReq.Quantity = ucloud.Int(1)
	getVMInstancePriceReq.OSType = ucloud.String("Linux")
	getVMInstancePriceReq.ImageID = ucloud.String("cn-image-centos-68")

	getVMInstancePriceResp, err := ucloudstackClient.GetVMInstancePrice(getVMInstancePriceReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, getVMInstancePriceResp.Message)
		return
	}

	fmt.Printf("GetVMInstancePrice success, resp: %+v", getVMInstancePriceResp.Infos)
}

// StartVM 开机
func StartVM(vmID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	startReq := ucloudstackClient.NewStartVMInstanceRequest()
	startReq.Region = ucloud.String("cn")
	startReq.Zone = ucloud.String("zone-01")
	startReq.VMID = ucloud.String(vmID)

	// send request
	startResp, err := ucloudstackClient.StartVMInstance(startReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if startResp != nil {
		fmt.Printf("RetCode: %d\n", startResp.RetCode)
	}
}
