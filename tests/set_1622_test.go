package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1622(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("v_peer2", ctx.Must(utest.GetUDPNRegionResource(ctx.GetVar("Region"))))
	ctx.SetVar("create_udpn_bw", "2")
	ctx.SetVar("upgrade_udpn_bw", "3")
	ctx.SetVar("charge_type", "Month")
	ctx.SetVar("project_id", "org-m3kc33")

	testSet1622GetUDPNLineList00(&ctx)
	testSet1622CreateVPC01(&ctx)
	testSet1622CreateVPC02(&ctx)
	testSet1622GetUDPNPrice03(&ctx)
	testSet1622AllocateUDPN04(&ctx)
	testSet1622DescribeUDPN05(&ctx)
	testSet1622CreateVPCIntercom06(&ctx)
	testSet1622GetUDPNUpgradePrice07(&ctx)
	testSet1622ModifyUDPNBandwidth08(&ctx)
	testSet1622DeleteVPCIntercom09(&ctx)
	testSet1622ReleaseUDPN10(&ctx)
	testSet1622DeleteVPC11(&ctx)
	testSet1622DeleteVPC12(&ctx)
}

func testSet1622GetUDPNLineList00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewGetUDPNLineListRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.GetUDPNLineList(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetUDPNLineListResponse", "str_eq"),
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

func testSet1622CreateVPC01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "VPC_TEST"))
	ctx.NoError(utest.SetReqValue(req, "Network", "192.168.0.0/16"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateVPCResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["src_vpc_id"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet1622CreateVPC02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("v_peer2")))

	ctx.NoError(utest.SetReqValue(req, "Name", "VPC_TEST"))
	ctx.NoError(utest.SetReqValue(req, "Network", "192.168.0.0/16"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateVPCResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["dest_vpc_id"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet1622GetUDPNPrice03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewGetUDPNPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Peer1", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Peer2", ctx.GetVar("v_peer2")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.GetUDPNPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetUDPNPriceResponse", "str_eq"),
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

func testSet1622AllocateUDPN04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewAllocateUDPNRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Peer1", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Peer2", ctx.GetVar("v_peer2")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("create_udpn_bw")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("charge_type")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.AllocateUDPN(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "AllocateUDPNResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["UdpnId"] = ctx.Must(utest.GetValue(resp, "UDPNId"))
}

func testSet1622DescribeUDPN05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewDescribeUDPNRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UDPNId", ctx.GetVar("UdpnId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.DescribeUDPN(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDPNResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["peer_1"] = ctx.Must(utest.GetValue(resp, "DataSet.0.Peer1"))
	ctx.Vars["peer_2"] = ctx.Must(utest.GetValue(resp, "DataSet.0.Peer2"))
}

func testSet1622CreateVPCIntercom06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := vpcClient.NewCreateVPCIntercomRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("src_vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "DstVPCId", ctx.GetVar("dest_vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "DstRegion", ctx.Must(utest.GetNotEqual(ctx.GetVar("Region"), ctx.GetVar("peer_1"), ctx.GetVar("peer_2")))))
	ctx.NoError(utest.SetReqValue(req, "DstProjectId", config.ProjectId))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPCIntercom(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateVPCIntercomResponse", "str_eq"),
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

func testSet1622GetUDPNUpgradePrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewGetUDPNUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UDPNId", ctx.GetVar("UdpnId")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("upgrade_udpn_bw")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.GetUDPNUpgradePrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetUDPNUpgradePriceResponse", "str_eq"),
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

func testSet1622ModifyUDPNBandwidth08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udpnClient.NewModifyUDPNBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UDPNId", ctx.GetVar("UdpnId")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("upgrade_udpn_bw")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.ModifyUDPNBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ModifyUDPNBandwidthResponse", "str_eq"),
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

func testSet1622DeleteVPCIntercom09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := vpcClient.NewDeleteVPCIntercomRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("src_vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "DstVPCId", ctx.GetVar("dest_vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "DstRegion", ctx.GetVar("v_peer2")))
	ctx.NoError(utest.SetReqValue(req, "DstProjectId", config.ProjectId))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPCIntercom(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteVPCIntercomResponse", "str_eq"),
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

func testSet1622ReleaseUDPN10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := udpnClient.NewReleaseUDPNRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UDPNId", ctx.GetVar("UdpnId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udpnClient.ReleaseUDPN(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ReleaseUDPNResponse", "str_eq"),
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

func testSet1622DeleteVPC11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("src_vpc_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteVPCResponse", "str_eq"),
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

func testSet1622DeleteVPC12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("v_peer2")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("dest_vpc_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteVPCResponse", "str_eq"),
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
