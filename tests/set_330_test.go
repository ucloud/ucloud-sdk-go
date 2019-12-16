package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet330(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")

	testSet330CreateUMemcacheGroup00(&ctx)
	testSet330DescribeUMem01(&ctx)
	testSet330DescribeUMemcachePrice02(&ctx)
	testSet330DescribeOrderDetailInfo03(&ctx)
	testSet330DescribeUMemcacheUpgradePrice04(&ctx)
	testSet330ResizeUMemcacheGroup05(&ctx)
	testSet330DescribeOrderDetailInfo06(&ctx)
	testSet330GetResourceRenewPrice07(&ctx)
	testSet330CreateRenew08(&ctx)
	testSet330DescribeOrderDetailInfo09(&ctx)
	testSet330DeleteUMem10(&ctx)
}

func testSet330CreateUMemcacheGroup00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewCreateUMemcacheGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "umem_charge"))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.CreateUMemcacheGroup(req)
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

func testSet330DescribeUMem01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "memcache"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
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

	ctx.Vars["CreateTime"] = ctx.Must(utest.GetValue(resp, "DataSet.0.CreateTime"))
}

func testSet330DescribeUMemcachePrice02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeUMemcachePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeUMemcachePrice(req)
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

func testSet330DescribeOrderDetailInfo03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_BUY"))
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
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("GetPrice/100"))), "float_eq"),
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

func testSet330DescribeUMemcacheUpgradePrice04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeUMemcacheUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2"))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeUMemcacheUpgradePrice(req)
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

func testSet330ResizeUMemcacheGroup05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iumemClient.NewResizeUMemcacheGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.ResizeUMemcacheGroup(req)
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

func testSet330DescribeOrderDetailInfo06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

	ctx.NoError(utest.SetReqValue(req, "OrderTypes", "OT_BUY"))
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
			// ctx.NewValidator("OrderInfos.0.Amount", ctx.Must(utest.Calculate(ctx.GetVar("GetPrice/100"))), "float_eq"),
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

func testSet330GetResourceRenewPrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet330CreateRenew08(ctx *utest.TestContext) {
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

func testSet330DescribeOrderDetailInfo09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := ubillClient.NewDescribeOrderDetailInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "BeginTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))
	ctx.NoError(utest.SetReqValue(req, "EndTime", ctx.Must(utest.Calculate(ctx.Must(utest.GetTimestamp("10"))))))

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
			ctx.NewValidator("RetCode", "0", "str_eq"),
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

func testSet330DeleteUMem10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := iumemClient.NewDeleteUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "single"))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DeleteUMem(req)
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
