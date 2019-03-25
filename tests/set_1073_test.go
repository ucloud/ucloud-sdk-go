package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1073(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "dXFhQHVjbG91ZA==")
	ctx.SetVar("Type", "DB-2")
	ctx.SetVar("Name", "test123123")
	ctx.SetVar("Remark", "test123123")
	ctx.SetVar("ChargeType", "Month")
	ctx.SetVar("ProjectId", "org-fthbzm")

	testSet1073DescribePHostImage00(&ctx)
	testSet1073GetPHostPrice01(&ctx)
	testSet1073DescribePHostTags02(&ctx)
	testSet1073CreatePHost03(&ctx)
	testSet1073DescribePHost04(&ctx)
	testSet1073RebootPHost05(&ctx)
	testSet1073DescribePHost06(&ctx)
	testSet1073StopPHost07(&ctx)
	testSet1073DescribePHost08(&ctx)
	testSet1073ReinstallPHost09(&ctx)
	testSet1073DescribePHost10(&ctx)
	testSet1073PoweroffPHost11(&ctx)
	testSet1073DescribePHost12(&ctx)
	testSet1073TerminatePHost13(&ctx)
}

func testSet1073DescribePHostImage00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageType", "Base"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHostImage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostImageResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

	ctx.Vars["ImageID1"] = ctx.Must(utest.GetValue(resp, "ImageSet.0.ImageId"))
	ctx.Vars["ImageID2"] = ctx.Must(utest.GetValue(resp, "ImageSet.2.ImageId"))
}

func testSet1073GetPHostPrice01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewGetPHostPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.GetPHostPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetPHostPriceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073DescribePHostTags02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostTagsRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHostTags(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostTagsResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073CreatePHost03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewCreatePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID1")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("Remark")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.CreatePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreatePHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

	ctx.Vars["PHost"] = ctx.Must(utest.GetValue(resp, "PHostId.0"))
}

func testSet1073DescribePHost04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PHostType", ctx.GetVar("Type"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Name", ctx.GetVar("Name"), "str_eq"),
			ctx.NewValidator("PHostSet.0.PHostId", ctx.GetVar("PHost"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Remark", ctx.GetVar("Remark"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Zone", ctx.GetVar("Zone"), "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073RebootPHost05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(600) * time.Second)

	req := uphostClient.NewRebootPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("ProjectId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.RebootPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "RebootPHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073DescribePHost06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073StopPHost07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iuphostClient.NewStopPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("ProjectId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuphostClient.StopPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopPHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073DescribePHost08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Stopped", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073ReinstallPHost09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewReinstallPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.ReinstallPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ReinstallPHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073DescribePHost10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PHostType", ctx.GetVar("Type"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Name", ctx.GetVar("Name"), "str_eq"),
			ctx.NewValidator("PHostSet.0.PHostId", ctx.GetVar("PHost"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Remark", ctx.GetVar("Remark"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Zone", ctx.GetVar("Zone"), "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073PoweroffPHost11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewPoweroffPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("ProjectId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.PoweroffPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "PoweroffPHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073DescribePHost12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Stopped", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet1073TerminatePHost13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewTerminatePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("PHost")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("ProjectId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.TerminatePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminatePHostResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}
