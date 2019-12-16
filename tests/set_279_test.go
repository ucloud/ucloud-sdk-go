package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet279(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Image_Id", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet279CreateUHostInstance00(&ctx)
	testSet279CreateULB01(&ctx)
	testSet279AllocateEIP02(&ctx)
	testSet279DescribeEIP03(&ctx)
	testSet279UpdateEIPAttribute04(&ctx)
	testSet279GetEIPPrice05(&ctx)
	testSet279BindEIP06(&ctx)
	testSet279ModifyEIPBandwidth07(&ctx)
	testSet279SetEIPPayMode08(&ctx)
	testSet279ModifyEIPWeight09(&ctx)
	testSet279UnBindEIP10(&ctx)
	testSet279BindEIP11(&ctx)
	testSet279UnBindEIP12(&ctx)
	testSet279ReleaseEIP13(&ctx)
	testSet279DescribeEIP14(&ctx)
	testSet279DeleteULB15(&ctx)
	testSet279PoweroffUHostInstance16(&ctx)
	testSet279TerminateUHostInstance17(&ctx)
}

func testSet279CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Name", "eip-s1-bgp"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("Image_Id")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", "2012_UClou"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostIds", "0", "len_ge"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["UHostId"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet279CreateULB01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := ulbClient.NewCreateULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBName", "ulb-eip"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateULB(req)
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

	ctx.Vars["ULBId"] = ctx.Must(utest.GetValue(resp, "ULBId"))
}

func testSet279AllocateEIP02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewAllocateEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "OperatorName", "Bgp"))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Dynamic"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "PayMode", "Bandwidth"))

	ctx.NoError(utest.SetReqValue(req, "Name", "eip-bgp-01"))
	ctx.NoError(utest.SetReqValue(req, "Remark", "test"))

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

	ctx.Vars["EIPId_01"] = ctx.Must(utest.GetValue(resp, "EIPSet.0.EIPId"))
}

func testSet279DescribeEIP03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewDescribeEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPIds", ctx.GetVar("EIPId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeEIP(req)
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

func testSet279UpdateEIPAttribute04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewUpdateEIPAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "Name", "eip-auto-gai"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "huangchao"))
	ctx.NoError(utest.SetReqValue(req, "Remark", "test-gai"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.UpdateEIPAttribute(req)
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

func testSet279GetEIPPrice05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewGetEIPPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "OperatorName", "Bgp"))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GetEIPPrice(req)
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

func testSet279BindEIP06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "ulb"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.BindEIP(req)
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

func testSet279ModifyEIPBandwidth07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewModifyEIPBandwidthRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
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

func testSet279SetEIPPayMode08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewSetEIPPayModeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "PayMode", "Traffic"))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.SetEIPPayMode(req)
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

func testSet279ModifyEIPWeight09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewModifyEIPWeightRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "Weight", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ModifyEIPWeight(req)
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

func testSet279UnBindEIP10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewUnBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "ulb"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.UnBindEIP(req)
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

func testSet279BindEIP11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "uhost"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("UHostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.BindEIP(req)
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

func testSet279UnBindEIP12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewUnBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "uhost"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("UHostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.UnBindEIP(req)
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

func testSet279ReleaseEIP13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewReleaseEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("EIPId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "8039", "str_eq"),
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

func testSet279DescribeEIP14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDescribeEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPIds", ctx.GetVar("EIPId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "0", "str_eq"),
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

func testSet279DeleteULB15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
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

func testSet279PoweroffUHostInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("UHostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
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

func testSet279TerminateUHostInstance17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("UHostId")))

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
