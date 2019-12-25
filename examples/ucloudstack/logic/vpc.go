package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateVPC 创建私有网络(VPC)
func CreateVPC() {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createVPCReq := ucloudstackClient.NewCreateVPCRequest()
	createVPCReq.Region = ucloud.String("cn")
	createVPCReq.Zone = ucloud.String("zone-01")
	createVPCReq.Name = ucloud.String("test-sdk")
	createVPCReq.Network = ucloud.String("172.16.0.0/16")
	createVPCReq.Remark = ucloud.String("test-sdk")

	// send request
	createVPCResp, err := ucloudstackClient.CreateVPC(createVPCReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("CreateVPC Success, Infos: %+v", createVPCResp.VPCID)
}

// DescribeVPC 创建私有网络(VPC)
func DescribeVPC(vpcID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeVPCReq := ucloudstackClient.NewDescribeVPCRequest()
	describeVPCReq.Region = ucloud.String("cn")
	describeVPCReq.Zone = ucloud.String("zone-01")
	describeVPCReq.VPCIDs = []string{vpcID}

	// send request
	describeVPCResp, err := ucloudstackClient.DescribeVPC(describeVPCReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("DescribeVPC Success, Infos: %+v", describeVPCResp.Infos)
}

// DeleteVPC 删除私有网络(VPC)
func DeleteVPC(vpcID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeVPCReq := ucloudstackClient.NewDeleteVPCRequest()
	describeVPCReq.Region = ucloud.String("cn")
	describeVPCReq.Zone = ucloud.String("zone-01")
	describeVPCReq.VPCID = ucloud.String(vpcID)

	// send request
	describeVPCResp, err := ucloudstackClient.DeleteVPC(describeVPCReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("DeleteVPC Success, Infos: %+v", describeVPCResp)
}

// CreateSubnet 创建子网
func CreateSubnet(vpcID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createSubnetReq := ucloudstackClient.NewCreateSubnetRequest()
	createSubnetReq.Region = ucloud.String("cn")
	createSubnetReq.Zone = ucloud.String("zone-01")
	createSubnetReq.VPCID = ucloud.String(vpcID)
	createSubnetReq.Name = ucloud.String("test")
	createSubnetReq.Remark = ucloud.String("test")
	createSubnetReq.Network = ucloud.String("172.16.12.0/25")

	// send request
	createSubnetResp, err := ucloudstackClient.CreateSubnet(createSubnetReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("CreateSubnet Success, Infos: %+v", createSubnetResp)
}

// DescribeSubnet 查询子网
func DescribeSubnet(vpcID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeSubnetReq := ucloudstackClient.NewDescribeSubnetRequest()
	describeSubnetReq.Region = ucloud.String("cn")
	describeSubnetReq.Zone = ucloud.String("zone-01")
	describeSubnetReq.VPCID = ucloud.String(vpcID)

	// send request
	describeSubnetResp, err := ucloudstackClient.DescribeSubnet(describeSubnetReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("DescribeSubnet Success, Infos: %+v", describeSubnetResp)
}

// DeleteSubnet 查询子网
func DeleteSubnet(subnetID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeVPCReq := ucloudstackClient.NewDeleteSubnetRequest()
	describeVPCReq.Region = ucloud.String("cn")
	describeVPCReq.Zone = ucloud.String("zone-01")
	describeVPCReq.SubnetID = ucloud.String(subnetID)

	// send request
	describeVPCResp, err := ucloudstackClient.DeleteSubnet(describeVPCReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	fmt.Printf("DeleteSubnet Success, Infos: %+v", describeVPCResp)
}
