package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet687(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Zone", "cn-sh2-02")
	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("VPC_name_1", "VPC_api_test_1")
	ctx.SetVar("remark", "remark_api_test")
	ctx.SetVar("tag", "tag_api_test")
	ctx.SetVar("Subnet_name_1_1", "subnet_1_1")
	ctx.SetVar("subnet_netmask", "24")

	testSet687GetProjectList00(&ctx)
	testSet687CreateVPC01(&ctx)
	testSet687CreateSubnet02(&ctx)
	testSet687UpdateSubnetAttribute03(&ctx)
	testSet687DescribeSubnet04(&ctx)
	testSet687CreateVPC05(&ctx)
	testSet687CreateSubnet06(&ctx)
	testSet687CreateSubnet07(&ctx)
	testSet687DescribeSubnet08(&ctx)
	testSet687AllocateVIP09(&ctx)
	testSet687DescribeVIP10(&ctx)
	testSet687DescribeSubnetResource11(&ctx)
	testSet687ReleaseVIP12(&ctx)
	testSet687DeleteSubnet13(&ctx)
	testSet687DeleteSubnet14(&ctx)
	testSet687DeleteSubnet15(&ctx)
	testSet687AddVPCNetwork16(&ctx)
	testSet687DescribeVPC17(&ctx)
	testSet687CreateVPCIntercom18(&ctx)
	testSet687DescribeVPCIntercom19(&ctx)
	testSet687DeleteVPCIntercom20(&ctx)
	testSet687DeleteVPC21(&ctx)
	testSet687DeleteVPC22(&ctx)
}

func testSet687GetProjectList00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uaccountClient.NewGetProjectListRequest()

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uaccountClient.GetProjectList(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetProjectListResponse", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["project_list"] = ctx.Must(utest.GetValue(resp, "ProjectSet"))
}

func testSet687CreateVPC01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("VPC_name_1")))
	ctx.NoError(utest.SetReqValue(req, "Network", "172.16.16.0/20"))
	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("tag")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("remark")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VPCId_1"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet687CreateSubnet02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewCreateSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))
	ctx.NoError(utest.SetReqValue(req, "Subnet", "172.16.17.0"))
	ctx.NoError(utest.SetReqValue(req, "Netmask", ctx.GetVar("subnet_netmask")))
	ctx.NoError(utest.SetReqValue(req, "SubnetName", ctx.GetVar("Subnet_name_1_1")))
	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("tag")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("remark")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["SubnetId_1_1"] = ctx.Must(utest.GetValue(resp, "SubnetId"))
}

func testSet687UpdateSubnetAttribute03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewUpdateSubnetAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	ctx.NoError(utest.SetReqValue(req, "Tag", "qa"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.UpdateSubnetAttribute(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "UpdateSubnetAttributeResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DescribeSubnet04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDescribeSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "1"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DescribeSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687CreateVPC05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Name", "vpc_2"))
	ctx.NoError(utest.SetReqValue(req, "Network", "192.168.16.0/20"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VPCId_2"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet687CreateSubnet06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewCreateSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_2")))
	ctx.NoError(utest.SetReqValue(req, "Subnet", "192.168.17.0"))
	ctx.NoError(utest.SetReqValue(req, "Netmask", ctx.GetVar("subnet_netmask")))
	ctx.NoError(utest.SetReqValue(req, "SubnetName", "Subnet_2_1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["SubnetId_2_1"] = ctx.Must(utest.GetValue(resp, "SubnetId"))
}

func testSet687CreateSubnet07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_2")))
	ctx.NoError(utest.SetReqValue(req, "Subnet", "192.168.18.0"))
	ctx.NoError(utest.SetReqValue(req, "Netmask", ctx.GetVar("subnet_netmask")))
	ctx.NoError(utest.SetReqValue(req, "SubnetName", "Subnet_2_2"))

	ctx.NoError(utest.SetReqValue(req, "Remark", "remark_2_2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["SubnetId_2_2"] = ctx.Must(utest.GetValue(resp, "SubnetId"))
}

func testSet687DescribeSubnet08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDescribeSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DescribeSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.VPCId", ctx.GetVar("VPCId_1"), "str_eq"),
			ctx.NewValidator("DataSet.0.VPCName", ctx.GetVar("VPC_name_1"), "str_eq"),
			ctx.NewValidator("DataSet.0.SubnetId", ctx.GetVar("SubnetId_1_1"), "str_eq"),
			ctx.NewValidator("DataSet.0.SubnetName", ctx.GetVar("Subnet_name_1_1"), "str_eq"),
			ctx.NewValidator("DataSet.0.Tag", "qa", "str_eq"),
			ctx.NewValidator("DataSet.0.Remark", ctx.GetVar("remark"), "str_eq"),
			ctx.NewValidator("DataSet.0.SubnetType", "2", "str_eq"),
			ctx.NewValidator("DataSet.0.Netmask", "24", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687AllocateVIP09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateVIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))
	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	ctx.NoError(utest.SetReqValue(req, "Name", "vip_api_auto"))

	ctx.NoError(utest.SetReqValue(req, "Remark", "vip_tag1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateVIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VIPId_1"] = ctx.Must(utest.GetValue(resp, "VIPSet.0.VIPId"))
}

func testSet687DescribeVIP10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := unetClient.NewDescribeVIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))
	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeVIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("VIPSet.0.VPCId", ctx.GetVar("VPCId_1"), "str_eq"),
			ctx.NewValidator("VIPSet.0.VIPId", ctx.GetVar("VIPId_1"), "str_eq"),
			ctx.NewValidator("VIPSet.0.SubnetId", ctx.GetVar("SubnetId_1_1"), "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VIP_ip_1"] = ctx.Must(utest.GetValue(resp, "DataSet.0"))
}

func testSet687DescribeSubnetResource11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewDescribeSubnetResourceRequest()

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "20"))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DescribeSubnetResource(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("DataSet.0.ResourceId", ctx.GetVar("VIPId_1"), "str_eq"),
			ctx.NewValidator("DataSet.0.IP", ctx.GetVar("VIP_ip_1"), "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687ReleaseVIP12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewReleaseVIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "VIPId", ctx.GetVar("VIPId_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseVIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteSubnet13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := vpcClient.NewDeleteSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_1_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteSubnet14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := vpcClient.NewDeleteSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_2_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteSubnet15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := vpcClient.NewDeleteSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("SubnetId_2_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687AddVPCNetwork16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewAddVPCNetworkRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))
	ctx.NoError(utest.SetReqValue(req, "Network", "10.100.96.0/20"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.AddVPCNetwork(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DescribeVPC17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDescribeVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VPCIds", ctx.GetVar("VPCId_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DescribeVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687CreateVPCIntercom18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCIntercomRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))
	ctx.NoError(utest.SetReqValue(req, "DstRegion", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "DstVPCId", ctx.GetVar("VPCId_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPCIntercom(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DescribeVPCIntercom19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDescribeVPCIntercomRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DescribeVPCIntercom(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.VPCId", ctx.GetVar("VPCId_2"), "str_eq"),
			ctx.NewValidator("DataSet.0.DstRegion", ctx.GetVar("Region"), "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteVPCIntercom20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDeleteVPCIntercomRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))

	ctx.NoError(utest.SetReqValue(req, "DstRegion", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "DstVPCId", ctx.GetVar("VPCId_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPCIntercom(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteVPC21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet687DeleteVPC22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("VPCId_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
