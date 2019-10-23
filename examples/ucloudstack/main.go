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
	cfg.BaseUrl = "http://console.pre.ucloudstack.com/api"

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUDSTACK_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUDSTACK_PUBLIC_KEY")

	return &cfg, &credential
}

func main() {

	// createVM("my-first-vm")

	// stopVM("vm-uGvOT3TZg")

	// deleteVM("vm-lxK-oqTZg")

	// describeVM()

	// describeMetric("vm-hu74T3oZR")

}

func createVM(name string) {

	// 认证
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
	createReq.BootDiskType = ucloud.String("Normal")
	createReq.DataDiskType = ucloud.String("Normal")
	createReq.VMType = ucloud.String("Normal")
	// 网络
	createReq.VPCID = ucloud.String("vpc-1al_S_tbN")
	createReq.SubnetID = ucloud.String("subnet-1al_S_tbN")
	createReq.WANSGID = ucloud.String("sg-1al_S_tbN")
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

	// 认证
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

	// 认证
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

func describeVM() {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeVMInstanceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.Limit = ucloud.Int(10)

	// send request
	descResp, err := ucloudstackClient.DescribeVMInstance(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if descResp.TotalCount > 0 {
		fmt.Printf("fisrt of VMs: %s\n", descResp.Infos[0].VMID)
	}

}

func describeMetric(vmID string) {

	// 认证
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
