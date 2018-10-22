package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1203(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-04")

	ctx.SetVar("Protocol", "redis")
	ctx.SetVar("Type", "double")
	ctx.SetVar("Name", "distributed_redis")

	testSet1203CreateUMemSpace00(&ctx)
	testSet1203DescribeUMem01(&ctx)
	testSet1203ModifyUMemSpaceName02(&ctx)
	testSet1203DescribeUMem03(&ctx)
	testSet1203GetUMemSpaceState04(&ctx)
	testSet1203ResizeUMemSpace05(&ctx)
	testSet1203DescribeUMem06(&ctx)
	testSet1203DeleteUMemSpace07(&ctx)
	testSet1203DescribeUMem08(&ctx)
}

func testSet1203CreateUMemSpace00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := umemClient.NewCreateUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "16"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))
	ctx.NoError(utest.SetReqValue(req, "Type", ctx.GetVar("Type")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
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

func testSet1203DescribeUMem01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

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

}

func testSet1203ModifyUMemSpaceName02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewModifyUMemSpaceNameRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "Rename_FBS"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ModifyUMemSpaceName(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ModifyUMemSpaceNameResponse", "str_eq"),
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

func testSet1203DescribeUMem03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
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

func testSet1203GetUMemSpaceState04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewGetUMemSpaceStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.GetUMemSpaceState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetUMemSpaceStateResponse", "str_eq"),
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

func testSet1203ResizeUMemSpace05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := umemClient.NewResizeUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "17"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ResizeUMemSpace(req)
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

func testSet1203DescribeUMem06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
		},
		MaxRetries:    40,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1203DeleteUMemSpace07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(50) * time.Second)

	req := umemClient.NewDeleteUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DeleteUMemSpace(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteUMemSpaceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1203DescribeUMem08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUMemResponse", "str_eq"),
			ctx.NewValidator("DataSet", ctx.GetVar("Space_Id"), "object_not_contains"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
