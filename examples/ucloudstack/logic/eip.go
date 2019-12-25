package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateEIP 创建外网IP
func CreateEIP() {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createEIPReq := ucloudstackClient.NewAllocateEIPRequest()
	createEIPReq.Region = ucloud.String("cn")
	createEIPReq.Zone = ucloud.String("zone-01")
	createEIPReq.OperatorName = ucloud.String("Bgp")
	createEIPReq.Bandwidth = ucloud.Int(572)
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

// BindEIP 绑定外网IP
func BindEIP(eipID, resourceID, resourceType string) {
	cfg, credential := LoadUcloudStackConfig()
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

// UnBindEIP 解绑EIP
func UnBindEIP(eipID, resourceID, resourceType string) {
	cfg, credential := LoadUcloudStackConfig()
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

// ReleaseEIP 删除外网IP
func ReleaseEIP(eipID string) {
	cfg, credential := LoadUcloudStackConfig()
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

// DescribeEIP 查询外网IP
func DescribeEIP(eipID string) {
	cfg, credential := LoadUcloudStackConfig()
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

// GetEIPPrice 获取外网IP价格
func GetEIPPrice() {
	cfg, credential := LoadUcloudStackConfig()
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

// ModifyEIPBandwidth 调整外网IP带宽
func ModifyEIPBandwidth(eipID string, bandwidth int) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	modifyEIPReq := ucloudstackClient.NewModifyEIPBandwidthRequest()
	modifyEIPReq.Region = ucloud.String("cn")
	modifyEIPReq.Zone = ucloud.String("zone-01")
	modifyEIPReq.EIPID = ucloud.String(eipID)
	modifyEIPReq.Bandwidth = ucloud.Int(bandwidth)

	// send request
	resp, err := ucloudstackClient.ModifyEIPBandwidth(modifyEIPReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("ModifyEIPBandwidth Success")
}
