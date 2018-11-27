package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet242(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")

	testSet242DescribeURedisPrice00(&ctx)
	testSet242CreateURedisGroup01(&ctx)
	testSet242DescribeUMem02(&ctx)
	testSet242DescribeOrderDetailInfo03(&ctx)
	testSet242DescribeURedisUpgradePrice04(&ctx)
	testSet242ResizeURedisGroup05(&ctx)
	testSet242DescribeOrderDetailInfo06(&ctx)
	testSet242GetResourceRenewPrice07(&ctx)
	testSet242CreateRenew08(&ctx)
	testSet242DescribeOrderDetailInfo09(&ctx)
	testSet242DeleteURedisGroup10(&ctx)
}

func testSet242DescribeURedisPrice00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeURedisPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeURedisPrice(req)
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

func testSet242CreateURedisGroup01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewCreateURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "Redis_zb"))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))

	ctx.NoError(utest.SetReqValue(req, "HighAvailability", "disable"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	// ctx.NoError(utest.SetReqValue(req, "Protocol", "redis"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.CreateURedisGroup(req)
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

	ctx.Vars["group_id"] = ctx.Must(utest.GetValue(resp, "GroupId"))
}

func testSet242DescribeUMem02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))

	ctx.NoError(utest.SetReqValue(req, "ResourceType", "single"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "redis"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["CreateTime"] = ctx.Must(utest.GetValue(resp, "DataSet.0.CreateTime"))
}

func testSet242DescribeOrderDetailInfo03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_BUY"))

	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("group_id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("GetPrice/100"))), "float_eq"),
		},
		MaxRetries:    5,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet242DescribeURedisUpgradePrice04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeURedisUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2"))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeURedisUpgradePrice(req)
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

func testSet242ResizeURedisGroup05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := umemClient.NewResizeURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ResizeURedisGroup(req)
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

func testSet242DescribeOrderDetailInfo06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_UPGRADE"))
	ctx.NoError(utest.SetReqValue(req, "ChargeTypes", "Month"))
	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("group_id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate("+", ctx.GetVar("UpgradePrice/100"), "0.1", "float")), "lt"),
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("UpgradePrice/100-0.1"), "float")), "gt"),
		},
		MaxRetries:    5,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet242GetResourceRenewPrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iubillClient.NewGetResourceRenewPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("group_id")))
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

func testSet242CreateRenew08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iubillClient.NewCreateRenewRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
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

func testSet242DescribeOrderDetailInfo09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.GetTimestamp("10"))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_RENEW"))
	ctx.NoError(utest.SetReqValue(req, "ChargeTypes", "Month"))
	ctx.NoError(utest.SetReqValue(req, "OrderStates", "OS_FINISHED"))

	ctx.NoError(utest.SetReqValue(req, "ResourceIds", ctx.GetVar("group_id")))

	ctx.NoError(utest.SetReqValue(req, "QueryAll", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ubillClient.DescribeOrderDetailInfo(req)
		},
		Validators: []utest.TestValidator{
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.GetVar("ReNewPrice"), "str_eq"),
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

func testSet242DeleteURedisGroup10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := umemClient.NewDeleteURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DeleteURedisGroup(req)
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
