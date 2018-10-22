package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet149(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Image_Id_ucloud", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))
	ctx.SetVar("saopaulo_image", "uimage-1bkjka")

	testSet149CreateUHostInstance00(&ctx)
	testSet149CreateULB01(&ctx)
	testSet149DescribeULB02(&ctx)
	testSet149CreateVServer03(&ctx)
	testSet149DescribeVServer04(&ctx)
	testSet149AllocateBackendBatch05(&ctx)
	testSet149DescribeVServer06(&ctx)
	testSet149UpdateBackendAttribute07(&ctx)
	testSet149ReleaseBackend08(&ctx)
	testSet149DescribeVServer09(&ctx)
	testSet149AllocateBackend10(&ctx)
	testSet149DescribeVServer11(&ctx)
	testSet149GetBackendMetricInfo12(&ctx)
	testSet149GetVServerMetricInfo13(&ctx)
	testSet149UpdateULBAttribute14(&ctx)
	testSet149UpdateVServerAttribute15(&ctx)
	testSet149DescribeVServer16(&ctx)
	testSet149DeleteVServer17(&ctx)
	testSet149DescribeVServer18(&ctx)
	testSet149DeleteULB19(&ctx)
	testSet149DeleteULB20(&ctx)
	testSet149DeleteULB21(&ctx)
	testSet149DescribeULBSimple22(&ctx)
	testSet149PoweroffUHostInstance23(&ctx)
	testSet149TerminateUHostInstance24(&ctx)
}

func testSet149CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("Image_Id_ucloud")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", "Y2VzaGkxMjMrKw=="))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "Name", "ulb-host"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "UHostType", "Normal"))

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

func testSet149CreateULB01(ctx *utest.TestContext) {
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

func testSet149DescribeULB02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := ulbClient.NewDescribeULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "60"))
	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.ULBId", ctx.GetVar("ULBId"), "str_eq"),
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

func testSet149CreateVServer03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet149DescribeVServer04(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.ListenType", "RequestProxy", "str_eq"),
			ctx.NewValidator("DataSet.0.VServerName", "vserver-test", "str_eq"),
			ctx.NewValidator("DataSet.0.Protocol", "HTTP", "str_eq"),
			ctx.NewValidator("DataSet.0.FrontendPort", "80", "str_eq"),
			ctx.NewValidator("DataSet.0.Method", "Roundrobin", "str_eq"),
			ctx.NewValidator("DataSet.0.ClientTimeout", "60", "str_eq"),
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

func testSet149AllocateBackendBatch05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewAllocateBackendBatchRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	ctx.NoError(utest.SetReqValue(req, "Backends", ctx.Must(utest.ConcatWithVertical(ctx.GetVar("UHostId_01"), "UHost", "80", "1", ctx.GetVar("IP_01")))))
	ctx.NoError(utest.SetReqValue(req, "ApiVersion", "3"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.AllocateBackendBatch(req)
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

func testSet149DescribeVServer06(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.ListenType", "RequestProxy", "str_eq"),
			ctx.NewValidator("DataSet.0.VServerName", "vserver-test", "str_eq"),
			ctx.NewValidator("DataSet.0.Protocol", "HTTP", "str_eq"),
			ctx.NewValidator("DataSet.0.FrontendPort", "80", "str_eq"),
			ctx.NewValidator("DataSet.0.Method", "Roundrobin", "str_eq"),
			ctx.NewValidator("DataSet.0.ClientTimeout", "60", "str_eq"),
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

func testSet149UpdateBackendAttribute07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewUpdateBackendAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "BackendId", ctx.GetVar("BackendId")))
	ctx.NoError(utest.SetReqValue(req, "Port", "80"))
	ctx.NoError(utest.SetReqValue(req, "Enabled", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.UpdateBackendAttribute(req)
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

func testSet149ReleaseBackend08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewReleaseBackendRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "BackendId", ctx.GetVar("BackendId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.ReleaseBackend(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149DescribeVServer09(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.BackendSet", "0", "len_eq"),
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

func testSet149AllocateBackend10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet149DescribeVServer11(ctx *utest.TestContext) {
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

}

func testSet149GetBackendMetricInfo12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iulbClient.NewGetBackendMetricInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "BackendIds", ctx.GetVar("BackendId")))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iulbClient.GetBackendMetricInfo(req)
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

func testSet149GetVServerMetricInfo13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iulbClient.NewGetVServerMetricInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iulbClient.GetVServerMetricInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("DataSet.0.VServerId", ctx.GetVar("VServerId"), "str_eq"),
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

func testSet149UpdateULBAttribute14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewUpdateULBAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "Name", "测试-改"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.UpdateULBAttribute(req)
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

func testSet149UpdateVServerAttribute15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewUpdateVServerAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-gai"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "HTTPS"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Source"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceType", "ServerInsert"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.UpdateVServerAttribute(req)
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

func testSet149DescribeVServer16(ctx *utest.TestContext) {
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
			ctx.NewValidator("TotalCount", "1", "str_eq"),
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

func testSet149DeleteVServer17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDeleteVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet149DescribeVServer18(ctx *utest.TestContext) {
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
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149DeleteULB19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149DeleteULB20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149DeleteULB21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149DescribeULBSimple22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iulbClient.NewDescribeULBSimpleRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iulbClient.DescribeULBSimple(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "4103", "str_eq"),
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

func testSet149PoweroffUHostInstance23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("UHostId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet149TerminateUHostInstance24(ctx *utest.TestContext) {
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
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
