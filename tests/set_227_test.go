package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet227(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	testSet227CreateULB00(&ctx)
	testSet227CreateVServer01(&ctx)
	testSet227DescribeVServer02(&ctx)
	testSet227CreateVServer03(&ctx)
	testSet227DescribeVServer04(&ctx)
	testSet227CreateVServer05(&ctx)
	testSet227DescribeVServer06(&ctx)
	testSet227DeleteULB07(&ctx)
}

func testSet227CreateULB00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewCreateULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBName", "测试"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "OuterMode", "Yes"))
	ctx.NoError(utest.SetReqValue(req, "InnerMode", "No"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateULB(req)
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

	ctx.Vars["ULBId_outer"] = ctx.Must(utest.GetValue(resp, "ULBId"))
}

func testSet227CreateVServer01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewCreateVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-http"))
	ctx.NoError(utest.SetReqValue(req, "ListenType", "RequestProxy"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "HTTP"))
	ctx.NoError(utest.SetReqValue(req, "FrontendPort", "80"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Roundrobin"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceType", "ServerInsert"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4107", "str_eq"),
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

func testSet227DescribeVServer02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeVServer(req)
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

func testSet227CreateVServer03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := ulbClient.NewCreateVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-https"))
	ctx.NoError(utest.SetReqValue(req, "ListenType", "RequestProxy"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "HTTPS"))
	ctx.NoError(utest.SetReqValue(req, "FrontendPort", "443"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Source"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceType", "UserDefined"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceInfo", "huangchao"))
	ctx.NoError(utest.SetReqValue(req, "ClientTimeout", "60"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4107", "str_eq"),
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

func testSet227DescribeVServer04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeVServer(req)
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

func testSet227CreateVServer05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := ulbClient.NewCreateVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-https"))
	ctx.NoError(utest.SetReqValue(req, "ListenType", "RequestProxy"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "TCP"))
	ctx.NoError(utest.SetReqValue(req, "FrontendPort", "1024"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Source"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceType", "None"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4107", "str_eq"),
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

func testSet227DescribeVServer06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeVServer(req)
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

func testSet227DeleteULB07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId_outer")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
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
