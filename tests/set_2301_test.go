package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2301(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("begin_time", ctx.Must(utest.Calculate("-", ctx.Must(utest.GetTimestamp("10")), 86400*7)))

	testSet2301GetEIPPrice00(&ctx)
	testSet2301AllocateEIP01(&ctx)
	testSet2301DescribeOrderDetailInfo02(&ctx)
	testSet2301GetEIPUpgradePrice03(&ctx)
	testSet2301ModifyEIPBandwidth04(&ctx)
	testSet2301DescribeOrderDetailInfo05(&ctx)
	testSet2301GetEIPUpgradePrice06(&ctx)
	testSet2301ModifyEIPBandwidth07(&ctx)
	testSet2301DescribeOrderDetailInfo08(&ctx)
	testSet2301ReleaseEIP09(&ctx)
}

func testSet2301GetEIPPrice00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewGetEIPPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "OperatorName", "Bgp"))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GetEIPPrice(req)
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

	ctx.Vars["buy_price"] = ctx.Must(utest.GetValue(resp, "PriceSet.0.Price"))
}

func testSet2301AllocateEIP01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "OperatorName", "Bgp"))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateEIP(req)
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

	ctx.Vars["eip_id"] = ctx.Must(utest.GetValue(resp, "EIPSet.0.EIPId"))
}

func testSet2301DescribeOrderDetailInfo02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.GetVar("begin_time")))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "ResourceTypes", "EIP"))
	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_BUY"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("OrderInfos.0.Amount", ctx.GetVar("buy_price"), "float_eq"),
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

func testSet2301GetEIPUpgradePrice03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewGetEIPUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "3"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GetEIPUpgradePrice(req)
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

	ctx.Vars["upgrade_price"] = ctx.Must(utest.GetValue(resp, "Price"))
}

func testSet2301ModifyEIPBandwidth04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewModifyEIPBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "3"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ModifyEIPBandwidth(req)
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

func testSet2301DescribeOrderDetailInfo05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.GetVar("begin_time")))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "ResourceTypes", "EIP"))
	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_UPGRADE"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("OrderInfos.0.Amount", ctx.GetVar("upgrade_price"), "float_eq"),
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

func testSet2301GetEIPUpgradePrice06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewGetEIPUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GetEIPUpgradePrice(req)
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

	ctx.Vars["down_price"] = ctx.Must(utest.GetValue(resp, "Price"))
}

func testSet2301ModifyEIPBandwidth07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := unetClient.NewModifyEIPBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ModifyEIPBandwidth(req)
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

func testSet2301DescribeOrderDetailInfo08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.GetVar("begin_time")))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "ResourceTypes", "EIP"))
	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_DOWNGRADE"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("OrderInfos.0.Amount", ctx.GetVar("down_price"), "float_eq"),
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

func testSet2301ReleaseEIP09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewReleaseEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseEIP(req)
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
