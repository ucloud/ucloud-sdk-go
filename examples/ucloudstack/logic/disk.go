package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// GetDiskPrice 获取硬盘价格
func GetDiskPrice() {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	getDiskPriceReq := ucloudstackClient.NewGetDiskPriceRequest()
	getDiskPriceReq.Region = ucloud.String("cn")
	getDiskPriceReq.Zone = ucloud.String("zone-01")
	getDiskPriceReq.ChargeType = ucloud.String("Month")
	getDiskPriceReq.SetType = ucloud.String("Normal")
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

// CreateDisk 创建硬盘
func CreateDisk() {
	cfg, credential := LoadUcloudStackConfig()
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

// AttachDisk 挂载硬盘
func AttachDisk(diskID, resourceID, resourceType string) {
	cfg, credential := LoadUcloudStackConfig()
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

// DetachDisk 卸载硬盘
func DetachDisk(diskID, resourceID string) {
	cfg, credential := LoadUcloudStackConfig()
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

// CloneDisk 克隆硬盘
func CloneDisk(srcID string) {
	cfg, credential := LoadUcloudStackConfig()
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

// DescribeDisk 查询硬盘信息
func DescribeDisk(diskID string) {
	cfg, credential := LoadUcloudStackConfig()
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

// DeleteDisk 删除硬盘
func DeleteDisk(diskID string) {
	cfg, credential := LoadUcloudStackConfig()
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
