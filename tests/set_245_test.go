package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet245(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")

	testSet245CreateUMemSpace00(&ctx)
	testSet245DescribeUMem01(&ctx)
	testSet245DescribeUMemPrice02(&ctx)
	testSet245DescribeOrderDetailInfo03(&ctx)
	testSet245DescribeUMemUpgradePrice04(&ctx)
	testSet245ResizeUMemSpace05(&ctx)
	testSet245DescribeOrderDetailInfo06(&ctx)
	testSet245GetResourceRenewPrice07(&ctx)
	testSet245CreateRenew08(&ctx)
	testSet245DescribeOrderDetailInfo09(&ctx)
	testSet245DeleteUMemSpace10(&ctx)
}

func testSet245CreateUMemSpace00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewCreateUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "16"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "redis"))
	ctx.NoError(utest.SetReqValue(req, "Type", "double"))
	ctx.NoError(utest.SetReqValue(req, "Name", "Redis_fbs"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.CreateUMemSpace(req)
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

	ctx.Vars["Space_Id"] = ctx.Must(utest.GetValue(resp, "SpaceId"))
}

func testSet245DescribeUMem01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(50) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "redis"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["CreateTime"] = ctx.Must(utest.GetValue(resp, "DataSet.0.CreateTime"))
}

func testSet245DescribeUMemPrice02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeUMemPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "16"))
	ctx.NoError(utest.SetReqValue(req, "Type", "double"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeUMemPrice(req)
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

	ctx.Vars["GetPrice"] = ctx.Must(utest.GetValue(resp, "DataSet.0.Price"))
}

func testSet245DescribeOrderDetailInfo03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_BUY"))
	ctx.NoError(utest.SetReqValue(req, "ChargeTypes", "Month"))
	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("Space_Id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate("+", ctx.GetVar("GetPrice/100"), "0.6", "float")), "le"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("GetPrice/100-0.6"), "float")), "gt"),
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

func testSet245DescribeUMemUpgradePrice04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeUMemUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "18"))
	ctx.NoError(utest.SetReqValue(req, "Type", "double"))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeUMemUpgradePrice(req)
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

	ctx.Vars["UpgradePrice"] = ctx.Must(utest.GetValue(resp, "Price"))
}

func testSet245ResizeUMemSpace05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := umemClient.NewResizeUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "18"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ResizeUMemSpace(req)
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

func testSet245DescribeOrderDetailInfo06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_UPGRADE"))
	ctx.NoError(utest.SetReqValue(req, "ChargeTypes", "Month"))
	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("Space_Id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate("+", ctx.GetVar("UpgradePrice/100"), "0.1", "float")), "lt"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("UpgradePrice/100-0.1"), "float")), "gt"),
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

func testSet245GetResourceRenewPrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iubillClient.NewGetResourceRenewPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iubillClient.GetResourceRenewPrice(req)
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

	ctx.Vars["ReNewPrice"] = ctx.Must(utest.GetValue(resp, "RenewPriceSet.0.Price"))
}

func testSet245CreateRenew08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iubillClient.NewCreateRenewRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iubillClient.CreateRenew(req)
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

func testSet245DescribeOrderDetailInfo09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_RENEW"))
	ctx.NoError(utest.SetReqValue(req, "ChargeTypes", "Month"))
	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("Space_Id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			// ctx.NewValidator("OrderInfos.0.Amount", "ReNewPrice", "str_eq"),
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

func testSet245DeleteUMemSpace10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := umemClient.NewDeleteUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DeleteUMemSpace(req)
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

}
