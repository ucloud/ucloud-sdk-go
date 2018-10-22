package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1202(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-02")

	ctx.SetVar("Protocol", "memcache")
	ctx.SetVar("ResourceType", "distributed")
	ctx.SetVar("Name", "distributed_memcache")

	testSet1202CreateUMemSpace00(&ctx)
	testSet1202DescribeUMem01(&ctx)
	testSet1202ResizeUMemSpace02(&ctx)
	testSet1202DescribeUMem03(&ctx)
	testSet1202DeleteUMem04(&ctx)
	testSet1202DescribeUMemSpace05(&ctx)
}

func testSet1202CreateUMemSpace00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewCreateUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "8"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))
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

func testSet1202DescribeUMem01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

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

func testSet1202ResizeUMemSpace02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewResizeUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "14"))

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

func testSet1202DescribeUMem03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

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

func testSet1202DeleteUMem04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumemClient.NewDeleteUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("Space_Id")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", ctx.GetVar("ResourceType")))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DeleteUMem(req)
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

func testSet1202DescribeUMemSpace05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeUMemSpaceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))
	ctx.NoError(utest.SetReqValue(req, "SpaceId", ctx.GetVar("Space_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeUMemSpace(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUMemSpaceResponse", "str_eq"),
			ctx.NewValidator("DataSet", ctx.GetVar("Space_Id"), "object_not_contains"),
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
