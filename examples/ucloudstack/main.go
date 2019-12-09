package main

import (
	"fmt"
	"os"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func loadUcloudStackConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.BaseUrl = "http://console.dev.ucloudstack.com/api"

	credential := auth.NewCredential()
	credential.PublicKey = os.Getenv("UCLOUDSTACK_PUBLIC_KEY")
	credential.PrivateKey = os.Getenv("UCLOUDSTACK_PRIVATE_KEY")


	return &cfg, &credential
}

func main() {

	// createVM("my-first-vm")

	// stopVM("vm-QcJZVKbZg")

	// startVM("vm-QcJZVKbZg")

	// deleteVM("vm-QcJZVKbZg")

	// describeVM("vm-QcJZVKbZg")

	// describeMetric("vm-QcJZVKbZg")

	// createUser("sdktest@ucloud.cn", "ucloud.cn")

	// recharge(200000259, "alipay-1234")

	// describeUser(200000259)

	// loginByPassword("sdktest@ucloud.cn", "ucloud.cn")

	// getVMInstancePrice()

	// createEIP()

	// bindEIP("eip-nkFMVFbZR", "vm-QcJZVKbZg", "VM")

	// unBindEIP("eip-nkFMVFbZR", "vm-QcJZVKbZg", "VM")

	// releaseEIP("eip-nkFMVFbZR")

	// describeEIP("eip-nkFMVFbZR")

	// getEIPPrice()

	// getDiskPrice()

	// createDisk()

	// attachDisk("disk-AzKn4FbZg", "vm-QcJZVKbZg", "VM")

	// detachDisk("disk-AzKn4FbZg", "vm-QcJZVKbZg")

	// describeDisk("disk-AzKn4FbZg")

	// cloneDisk("disk-AzKn4FbZg")

	// deleteDisk("disk-AzKn4FbZg")

}

func createVM(name string) {
	cfg, credential := loadUcloudStackConfig()
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
	createReq.VMType = ucloud.String("Normal")

	// 网络
	createReq.VPCID = ucloud.String("vpc-dEBdhMBZO")
	createReq.SubnetID = ucloud.String("subnet-dEBdhMBZO")
	createReq.WANSGID = ucloud.String("sg-dEBdhMBZO")

	// 认证方式
	createReq.Name = ucloud.String(name)
	createReq.Password = ucloud.String("ucloud.cn132")

	// 计费方式
	createReq.ChargeType = ucloud.String("Month")

	// send request
	newVMInstance, err := ucloudstackClient.CreateVMInstance(createReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if newVMInstance != nil {
		fmt.Printf("resource id of the VM: %s\n", newVMInstance.VMID)
	}

}

func stopVM(vmID string) {
	cfg, credential := loadUcloudStackConfig()
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
	}

	if stopResp != nil {
		fmt.Printf("RetCode: %d\n", stopResp.RetCode)
	}

}

func deleteVM(vmID string) {
	cfg, credential := loadUcloudStackConfig()
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
	}

	if delResp != nil {
		fmt.Printf("RetCode: %d\n", delResp.RetCode)
	}

}

func describeVM(vmID string) {
	cfg, credential := loadUcloudStackConfig()
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
	}

	if descResp.TotalCount > 0 {
		fmt.Printf("fisrt of VMs: %+v\n", descResp.Infos)
	}

}

func describeMetric(vmID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewDescribeMetricRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceID = ucloud.String(vmID)
	metricReq.MetricName = []string{"CPUUtilization"}
	metricReq.BeginTime = ucloud.String("1571819416")
	metricReq.EndTime = ucloud.String("1571823016")

	// send request
	metricResp, err := ucloudstackClient.DescribeMetric(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if metricResp.TotalCount > 0 && len(metricResp.Infos[0].Infos) > 0 {
		fmt.Printf("value of %s at %d: %f\n",
			metricResp.Infos[0].MetricName,
			metricResp.Infos[0].Infos[0].Timestamp,
			metricResp.Infos[0].Infos[0].Value)
	}

}

// 管理员创建租户
func createUser(email, password string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	createUserReq := ucloudstackClient.NewCreateUserRequest()
	createUserReq.UserEmail = ucloud.String(email)
	createUserReq.PassWord = ucloud.String(password)

	createUserResp, err := ucloudstackClient.CreateUser(createUserReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, createUserResp.Message)
		return
	}

	fmt.Printf("New User's ID: %d", createUserResp.UserID)
}

func describeUser(userID int) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	describeUserReq := ucloudstackClient.NewDescribeUserRequest()
	describeUserReq.Limit = ucloud.Int(10)
	describeUserReq.Offset = ucloud.Int(0)
	describeUserReq.UserIDs = []int{userID}
	describeUserResp, err := ucloudstackClient.DescribeUser(describeUserReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, describeUserResp.Message)
		return
	}

	fmt.Printf("User's info: %v", describeUserResp.Infos)
}

func recharge(userID int, serialNo string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	chargeReq := ucloudstackClient.NewRechargeRequest()
	chargeReq.Amount = ucloud.Int(100000)
	chargeReq.FromType = ucloud.String("INPOUR_FROM_ALIPAY")
	chargeReq.SerialNo = ucloud.String(serialNo)
	chargeReq.UserID = ucloud.Int(userID)

	chargeResp, err := ucloudstackClient.Recharge(chargeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, chargeResp.Message)
		return
	}

	fmt.Printf("Recharge success")
}

func loginByPassword(email, password string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	loginByPasswordReq := ucloudstackClient.NewLoginByPasswordRequest()
	loginByPasswordReq.UserEmail = ucloud.String(email)
	loginByPasswordReq.Password = ucloud.String(password)

	loginByPasswordResp, err := ucloudstackClient.LoginByPassword(loginByPasswordReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, loginByPasswordResp.Message)
		return
	}

	fmt.Printf("Recharge success")
}

func getVMInstancePrice() {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	getVMInstancePriceReq := ucloudstackClient.NewGetVMInstancePriceRequest()
	getVMInstancePriceReq.Region = ucloud.String("cn")
	getVMInstancePriceReq.GPU = ucloud.Int(0)
	getVMInstancePriceReq.CPU = ucloud.Int(1)
	getVMInstancePriceReq.Memory = ucloud.Int(2048)
	getVMInstancePriceReq.DataDiskSpace = ucloud.Int(20)
	getVMInstancePriceReq.DataDiskSetType = ucloud.String("Normal")
	getVMInstancePriceReq.BootDiskSetType = ucloud.String("Normal")
	getVMInstancePriceReq.VMType = ucloud.String("SSD")
	getVMInstancePriceReq.ChargeType = ucloud.String("Month")
	getVMInstancePriceReq.Quantity = ucloud.Int(1)
	getVMInstancePriceReq.OSType = ucloud.String("Linux")
	getVMInstancePriceReq.ImageID = ucloud.String("cn-image-centos-68")

	getVMInstancePriceResp, err := ucloudstackClient.GetVMInstancePrice(getVMInstancePriceReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, getVMInstancePriceResp.Message)
		return
	}

	fmt.Printf("Recharge success, resp: %+v", getVMInstancePriceResp.Infos)
}

func startVM(vmID string) {
	cfg, credential := loadUcloudStackConfig()
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
	}

	if startResp != nil {
		fmt.Printf("RetCode: %d\n", startResp.RetCode)
	}
}

func createEIP() {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createEIPReq := ucloudstackClient.NewAllocateEIPRequest()
	createEIPReq.Region = ucloud.String("cn")
	createEIPReq.Zone = ucloud.String("zone-01")
	createEIPReq.OperatorName = ucloud.String("Bgp")
	createEIPReq.Bandwidth = ucloud.Int(10)
	createEIPReq.ChargeType = ucloud.String("Month")
	createEIPReq.Quantity = ucloud.Int(1)
	createEIPReq.Name = ucloud.String("测试EIP")

	// send request
	createEIPResp, err := ucloudstackClient.AllocateEIP(createEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if createEIPResp != nil {
		fmt.Printf("Allocate EIP Success, EIPID: %s", createEIPResp.EIPID)
	}
}

func bindEIP(eipID, resourceID, resourceType string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	bindEIPReq := ucloudstackClient.NewBindEIPRequest()
	bindEIPReq.Region = ucloud.String("cn")
	bindEIPReq.Zone = ucloud.String("zone-01")
	bindEIPReq.EIPID = ucloud.String(eipID)
	bindEIPReq.ResourceID = ucloud.String(resourceID)
	bindEIPReq.ResourceType = ucloud.String(resourceType)

	// send request
	_, err := ucloudstackClient.BindEIP(bindEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("Bind EIP Success")
}

func unBindEIP(eipID, resourceID, resourceType string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	unBindEIPReq := ucloudstackClient.NewUnBindEIPRequest()
	unBindEIPReq.Region = ucloud.String("cn")
	unBindEIPReq.Zone = ucloud.String("zone-01")
	unBindEIPReq.EIPID = ucloud.String(eipID)
	unBindEIPReq.ResourceID = ucloud.String(resourceID)
	unBindEIPReq.ResourceType = ucloud.String(resourceType)

	// send request
	_, err := ucloudstackClient.UnBindEIP(unBindEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("UnBind EIP Success")
}

func releaseEIP(eipID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	releaseEIPReq := ucloudstackClient.NewReleaseEIPRequest()
	releaseEIPReq.Region = ucloud.String("cn")
	releaseEIPReq.Zone = ucloud.String("zone-01")
	releaseEIPReq.EIPID = ucloud.String(eipID)

	// send request
	_, err := ucloudstackClient.ReleaseEIP(releaseEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("Release EIP Success")
}

func describeEIP(eipID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeEIPReq := ucloudstackClient.NewDescribeEIPRequest()
	describeEIPReq.Region = ucloud.String("cn")
	describeEIPReq.Zone = ucloud.String("zone-01")
	describeEIPReq.EIPIDs = []string{eipID}

	// send request
	describeEIPResp, err := ucloudstackClient.DescribeEIP(describeEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("Describe EIP Success, Infos: %+v", describeEIPResp.Infos)
}

func getEIPPrice() {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	getEIPPriceReq := ucloudstackClient.NewGetEIPPriceRequest()
	getEIPPriceReq.Region = ucloud.String("cn")
	getEIPPriceReq.Zone = ucloud.String("zone-01")
	getEIPPriceReq.ChargeType = ucloud.String("Month")
	getEIPPriceReq.OpertatorName = ucloud.String("Bgp")
	getEIPPriceReq.Quantity = ucloud.Int(1)
	getEIPPriceReq.Bandwidth = ucloud.Int(10)

	// send request
	getEIPPriceResp, err := ucloudstackClient.GetEIPPrice(getEIPPriceReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("Get EIP price Success, Infos: %+v", getEIPPriceResp.Infos)
}

func getDiskPrice() {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	getDiskPriceReq := ucloudstackClient.NewGetDiskPriceRequest()
	getDiskPriceReq.Region = ucloud.String("cn")
	getDiskPriceReq.Zone = ucloud.String("zone-01")
	getDiskPriceReq.ChargeType = ucloud.String("Month")
	getDiskPriceReq.SetType = ucloud.String("SSD")
	getDiskPriceReq.DiskSpace = ucloud.Int(100)
	getDiskPriceReq.Quantity = ucloud.Int(1)

	// send request
	getDiskPriceResp, err := ucloudstackClient.GetDiskPrice(getDiskPriceReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if getDiskPriceResp.RetCode != 0 {
		fmt.Printf("get disk price fail, err: %s", getDiskPriceResp.Message)
		return
	}

	fmt.Printf("Get disk price Success, Infos: %+v", getDiskPriceResp.Infos)
}

func createDisk() {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createDiskReq := ucloudstackClient.NewCreateDiskRequest()
	createDiskReq.Region = ucloud.String("cn")
	createDiskReq.Zone = ucloud.String("zone-01")
	createDiskReq.ChargeType = ucloud.String("Month")
	createDiskReq.SetType = ucloud.String("Normal")
	createDiskReq.DiskSpace = ucloud.Int(10)
	createDiskReq.Quantity = ucloud.Int(1)
	createDiskReq.Name = ucloud.String("硬盘测试")

	// send request
	createDiskResp, err := ucloudstackClient.CreateDisk(createDiskReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if createDiskResp.RetCode != 0 {
		fmt.Printf("create disk  fail, err: %s", createDiskResp.Message)
		return
	}

	fmt.Printf("create disk Success, Infos: %+v", createDiskResp.DiskID)
}

func attachDisk(diskID, resourceID, resourceType string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	attachDiskReq := ucloudstackClient.NewAttachDiskRequest()
	attachDiskReq.Region = ucloud.String("cn")
	attachDiskReq.Zone = ucloud.String("zone-01")
	attachDiskReq.DiskID = ucloud.String(diskID)
	attachDiskReq.ResourceID = ucloud.String(resourceID)
	attachDiskReq.ResourceType = ucloud.String(resourceType)

	// send request
	attachDiskResp, err := ucloudstackClient.AttachDisk(attachDiskReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if attachDiskResp.RetCode != 0 {
		fmt.Printf("attach disk  fail, err: %s", attachDiskResp.Message)
		return
	}

	fmt.Printf("Attach disk Success")
}

func detachDisk(diskID, resourceID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	detachDiskReq := ucloudstackClient.NewDetachDiskRequest()
	detachDiskReq.Region = ucloud.String("cn")
	detachDiskReq.Zone = ucloud.String("zone-01")
	detachDiskReq.DiskID = ucloud.String(diskID)
	detachDiskReq.ResourceID = ucloud.String(resourceID)

	// send request
	detachDiskResp, err := ucloudstackClient.DetachDisk(detachDiskReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if detachDiskResp.RetCode != 0 {
		fmt.Printf("detach disk  fail, err: %s", detachDiskResp.Message)
		return
	}

	fmt.Printf("Detach disk Success")
}

func cloneDisk(srcID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	cloneDiskReq := ucloudstackClient.NewCloneDiskRequest()
	cloneDiskReq.Region = ucloud.String("cn")
	cloneDiskReq.Zone = ucloud.String("zone-01")
	cloneDiskReq.ChargeType = ucloud.String("Month")
	cloneDiskReq.Quantity = ucloud.Int(1)
	cloneDiskReq.Name = ucloud.String("硬盘克隆测试")
	cloneDiskReq.SrcID = ucloud.String(srcID)

	// send request
	cloneDiskResp, err := ucloudstackClient.CloneDisk(cloneDiskReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if cloneDiskResp.RetCode != 0 {
		fmt.Printf("clone disk  fail, err: %s", cloneDiskResp.Message)
		return
	}

	fmt.Printf("create disk Success, Infos: %+v", cloneDiskResp.DiskID)
}

func describeDisk(diskID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeDiskReq := ucloudstackClient.NewDescribeDiskRequest()
	describeDiskReq.Region = ucloud.String("cn")
	describeDiskReq.Zone = ucloud.String("zone-01")
	describeDiskReq.DiskIDs = []string{diskID}

	// send request
	describeDiskResp, err := ucloudstackClient.DescribeDisk(describeDiskReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}
	if describeDiskResp.RetCode != 0 {
		fmt.Printf("detach disk  fail, err: %s", describeDiskResp.Message)
		return
	}

	fmt.Printf("Describe disk Success, Totalcount: %d Infos: %+v", describeDiskResp.TotalCount, describeDiskResp.Infos)
}

func deleteDisk(diskID string) {
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	deleteDiskReq := ucloudstackClient.NewDeleteDiskRequest()
	deleteDiskReq.Region = ucloud.String("cn")
	deleteDiskReq.Zone = ucloud.String("zone-01")
	deleteDiskReq.DiskID = ucloud.String(diskID)

	// send request
	deleteDiskResp, err := ucloudstackClient.DeleteDisk(deleteDiskReq)
	if err != nil {
		fmt.Printf("Something bad happened: %s\n, err: %s", err, deleteDiskResp.Message)
		return
	}
	if deleteDiskResp.RetCode != 0 {
		fmt.Printf("Delete disk  fail, err: %s", deleteDiskResp.Message)
		return
	}

	fmt.Printf("Delete disk Success")
}
