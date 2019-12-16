package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet220(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Image_Id_cloud", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))
	ctx.SetVar("saopaulo_image", "uimage-1bkjka")

	testSet220CreateUHostInstance00(&ctx)
	testSet220CreateULB01(&ctx)
	testSet220CreateVServer02(&ctx)
	testSet220DescribeVServer03(&ctx)
	testSet220AllocateBackend04(&ctx)
	testSet220DescribeVServer05(&ctx)
	testSet220CreatePolicy06(&ctx)
	testSet220UpdatePolicy07(&ctx)
	testSet220DeletePolicy08(&ctx)
	testSet220DeleteULB09(&ctx)
	testSet220PoweroffUHostInstance10(&ctx)
	testSet220TerminateUHostInstance11(&ctx)
}

func testSet220CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("Image_Id_cloud")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", "2012_UClou"))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "Name", "ulb-host"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
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

	ctx.Vars["UHostId_01"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
	ctx.Vars["IP_01"] = ctx.Must(utest.GetValue(resp, "IPs.0"))
}

func testSet220CreateULB01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(180) * time.Second)

	req := ulbClient.NewCreateULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBName", "测试"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

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

	ctx.Vars["ULBId"] = ctx.Must(utest.GetValue(resp, "ULBId"))
}

func testSet220CreateVServer02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewCreateVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-test"))
	ctx.NoError(utest.SetReqValue(req, "ListenType", "RequestProxy"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "HTTP"))
	ctx.NoError(utest.SetReqValue(req, "FrontendPort", "80"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Roundrobin"))
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

func testSet220DescribeVServer03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet", "1", "len_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VServerId"] = ctx.Must(utest.GetValue(resp, "DataSet.0.VServerId"))
}

func testSet220AllocateBackend04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := ulbClient.NewAllocateBackendRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "UHost"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("UHostId_01")))
	ctx.NoError(utest.SetReqValue(req, "Port", "80"))
	ctx.NoError(utest.SetReqValue(req, "Enabled", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.AllocateBackend(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "63016", "str_eq"),
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

func testSet220DescribeVServer05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.BackendSet", "1", "len_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["BackendId"] = ctx.Must(utest.GetValue(resp, "DataSet.0.BackendSet.0.BackendId"))
}

func testSet220CreatePolicy06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := ulbClient.NewCreatePolicyRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "BackendId", ctx.GetVar("BackendId")))
	ctx.NoError(utest.SetReqValue(req, "Match", "www.test.com"))
	ctx.NoError(utest.SetReqValue(req, "Type", "Domain"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreatePolicy(req)
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

	ctx.Vars["PolicyId"] = ctx.Must(utest.GetValue(resp, "PolicyId"))
}

func testSet220UpdatePolicy07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewUpdatePolicyRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "PolicyId", ctx.GetVar("PolicyId")))
	ctx.NoError(utest.SetReqValue(req, "BackendId", ctx.GetVar("BackendId")))
	ctx.NoError(utest.SetReqValue(req, "Type", "Domain"))
	ctx.NoError(utest.SetReqValue(req, "Match", "www.testgai.com"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.UpdatePolicy(req)
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

func testSet220DeletePolicy08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewDeletePolicyRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "PolicyId", ctx.GetVar("PolicyId")))

	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeletePolicy(req)
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

func testSet220DeleteULB09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

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

func testSet220PoweroffUHostInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("UHostId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    3,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet220TerminateUHostInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("UHostId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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
