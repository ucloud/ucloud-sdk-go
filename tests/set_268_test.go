package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet268(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("eip_bw", "2")
	ctx.SetVar("bwp1_bw", "3")
	ctx.SetVar("eip_operatorname", "Bgp")
	ctx.SetVar("bwp2_bw", "4")

	testSet268AllocateEIP00(&ctx)
	testSet268CreateULB01(&ctx)
	testSet268BindEIP02(&ctx)
	testSet268CreateBandwidthPackage03(&ctx)
	testSet268DescribeBandwidthPackage04(&ctx)
	testSet268DeleteBandwidthPackage05(&ctx)
	testSet268CreateBandwidthPackage06(&ctx)
	testSet268DescribeBandwidthPackage07(&ctx)
	testSet268DeleteBandwidthPackage08(&ctx)
	testSet268UnBindEIP09(&ctx)
	testSet268DeleteULB10(&ctx)
	testSet268ReleaseEIP11(&ctx)
}

func testSet268AllocateEIP00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "OperatorName", ctx.GetVar("eip_operatorname")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("eip_bw")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "PayMode", "Bandwidth"))

	ctx.NoError(utest.SetReqValue(req, "Name", "eip_auto"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateEIP(req)
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

	ctx.Vars["eip_id"] = ctx.Must(utest.GetValue(resp, "EIPSet.0.EIPId"))
}

func testSet268CreateULB01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewCreateULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBName", "ulb_for_bwp"))

	ctx.NoError(utest.SetReqValue(req, "OuterMode", "Yes"))

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

	ctx.Vars["ulb_id"] = ctx.Must(utest.GetValue(resp, "ULBId"))
}

func testSet268BindEIP02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := unetClient.NewBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "ulb"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("ulb_id")))

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

func testSet268CreateBandwidthPackage03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewCreateBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("bwp1_bw")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "TimeRange", "2"))
	ctx.NoError(utest.SetReqValue(req, "EnableTime", ctx.Must(utest.Calculate("+", "30", ctx.Must(utest.GetTimestamp("10"))))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.CreateBandwidthPackage(req)
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

	ctx.Vars["bwp_id1"] = ctx.Must(utest.GetValue(resp, "BandwidthPackageId"))
}

func testSet268DescribeBandwidthPackage04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewDescribeBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeBandwidthPackage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSets.0.EIPId", ctx.GetVar("eip_id"), "str_eq"),
			ctx.NewValidator("DataSets.0.Bandwidth", ctx.GetVar("bwp1_bw"), "str_eq"),
			ctx.NewValidator("DataSets.0.BandwidthPackageId", ctx.GetVar("bwp_id1"), "str_eq"),
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

func testSet268DeleteBandwidthPackage05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewDeleteBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "BandwidthPackageId", ctx.GetVar("bwp_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DeleteBandwidthPackage(req)
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

func testSet268CreateBandwidthPackage06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewCreateBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Bandwidth", ctx.GetVar("bwp2_bw")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "TimeRange", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.CreateBandwidthPackage(req)
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

	ctx.Vars["bwp_id2"] = ctx.Must(utest.GetValue(resp, "BandwidthPackageId"))
}

func testSet268DescribeBandwidthPackage07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewDescribeBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeBandwidthPackage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSets.0.EIPId", ctx.GetVar("eip_id"), "str_eq"),
			ctx.NewValidator("DataSets.0.Bandwidth", ctx.GetVar("bwp2_bw"), "str_eq"),
			ctx.NewValidator("DataSets.0.BandwidthPackageId", ctx.GetVar("bwp_id2"), "str_eq"),
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

func testSet268DeleteBandwidthPackage08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewDeleteBandwidthPackageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "BandwidthPackageId", ctx.GetVar("bwp_id2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DeleteBandwidthPackage(req)
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

func testSet268UnBindEIP09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewUnBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "ulb"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("ulb_id")))

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

func testSet268DeleteULB10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ulb_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
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

func testSet268ReleaseEIP11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := unetClient.NewReleaseEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseEIP(req)
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
