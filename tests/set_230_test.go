package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet230(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("sbw_name", "sbw_api_auto")
	ctx.SetVar("sbw_chargetype", "Month")
	ctx.SetVar("sbw_bw", "22")
	ctx.SetVar("resize_bw", "25")
	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("eip_operatorname", "Bgp")
	ctx.SetVar("eip_bw", "2")

	testSet230AllocateShareBandwidth00(&ctx)
	testSet230AllocateEIP01(&ctx)
	testSet230AssociateEIPWithShareBandwidth02(&ctx)
	testSet230DescribeShareBandwidthPrice03(&ctx)
	testSet230DescribeShareBandwidth04(&ctx)
	testSet230ResizeShareBandwidth05(&ctx)
	testSet230DescribeShareBandwidth06(&ctx)
	testSet230DisassociateEIPWithShareBandwidth07(&ctx)
	testSet230ReleaseEIP08(&ctx)
	testSet230ReleaseShareBandwidth09(&ctx)
}

func testSet230AllocateShareBandwidth00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("sbw_name")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("sbw_chargetype")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidth", ctx.GetVar("sbw_bw")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["sbw_id"] = ctx.Must(utest.GetValue(resp, "ShareBandwidthId"))
}

func testSet230AllocateEIP01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "OperatorName", ctx.GetVar("eip_operatorname")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("eip_bw")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "PayMode", "Traffic"))

	ctx.NoError(utest.SetReqValue(req, "Name", "eip_auto"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["eip_id"] = ctx.Must(utest.GetValue(resp, "EIPSet.0.EIPId"))
}

func testSet230AssociateEIPWithShareBandwidth02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := unetClient.NewAssociateEIPWithShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPIds", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthId", ctx.GetVar("sbw_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AssociateEIPWithShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230DescribeShareBandwidthPrice03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iunetClient.NewDescribeShareBandwidthPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("sbw_chargetype")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidth", ctx.GetVar("sbw_bw")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iunetClient.DescribeShareBandwidthPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalPrice", "0", "gt"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230DescribeShareBandwidth04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDescribeShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthIds", ctx.GetVar("sbw_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.ShareBandwidthId", ctx.GetVar("sbw_id"), "str_eq"),
			ctx.NewValidator("DataSet.0.ShareBandwidth", ctx.GetVar("sbw_bw"), "str_eq"),
			ctx.NewValidator("DataSet.0.ChargeType", ctx.GetVar("sbw_chargetype"), "str_eq"),
			ctx.NewValidator("DataSet.0.Name", ctx.GetVar("sbw_name"), "str_eq"),
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("DataSet.0.EIPSet.0.EIPAddr.0.OperatorName", ctx.GetVar("eip_operatorname"), "str_eq"),
			ctx.NewValidator("DataSet.0.EIPSet.0.EIPId", ctx.GetVar("eip_id"), "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230ResizeShareBandwidth05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewResizeShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidth", ctx.GetVar("resize_bw")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthId", ctx.GetVar("sbw_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ResizeShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230DescribeShareBandwidth06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDescribeShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthIds", ctx.GetVar("sbw_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.ShareBandwidth", ctx.GetVar("resize_bw"), "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230DisassociateEIPWithShareBandwidth07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDisassociateEIPWithShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPIds", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthId", ctx.GetVar("sbw_id")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("eip_bw")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DisassociateEIPWithShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230ReleaseEIP08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := unetClient.NewReleaseEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseEIP(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet230ReleaseShareBandwidth09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewReleaseShareBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ShareBandwidthId", ctx.GetVar("sbw_id")))
	ctx.NoError(utest.SetReqValue(req, "EIPBandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseShareBandwidth(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
