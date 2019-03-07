package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2377(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "YW4mODE1MDI5")
	ctx.SetVar("Type", "DISK-48T-V4")
	ctx.SetVar("ModifyName", "TestName")
	ctx.SetVar("ModifyRemark", "TestRemark")
	ctx.SetVar("Cluster", "10G")
	ctx.SetVar("Raid", "NoRaid")

	testSet2377DescribePHostImage00(&ctx)
	testSet2377DescribePHostMachineType01(&ctx)
	testSet2377DescribePHostResourceInfo02(&ctx)
	testSet2377CreatePHost03(&ctx)
	testSet2377DescribePHost04(&ctx)
	testSet2377GetPHostKVMInfo05(&ctx)
	testSet2377ModifyPHostInfo06(&ctx)
	testSet2377DescribePHost07(&ctx)
	testSet2377StopPHost08(&ctx)
	testSet2377DescribePHost09(&ctx)
	testSet2377ReinstallPHost10(&ctx)
	testSet2377DescribePHost11(&ctx)
	testSet2377PoweroffPHost12(&ctx)
	testSet2377DescribePHost13(&ctx)
	testSet2377StartPHost14(&ctx)
	testSet2377DescribePHost15(&ctx)
	testSet2377TerminatePHost16(&ctx)
}

func testSet2377DescribePHostImage00(ctx *utest.TestContext) {
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
	}

	ctx.Vars["ImageID1"] = ctx.Must(utest.GetValue(resp, "ImageSet.0.ImageId"))
}

func testSet2377DescribePHostMachineType01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuphostClient.NewDescribePHostMachineTypeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuphostClient.DescribePHostMachineType(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostMachineTypeResponse", "str_eq"),
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

func testSet2377DescribePHostResourceInfo02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuphostClient.NewDescribePHostResourceInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuphostClient.DescribePHostResourceInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResourceInfoResponse", "str_eq"),
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

func testSet2377CreatePHost03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewCreatePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID1")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))
	ctx.NoError(utest.SetReqValue(req, "Type", ctx.GetVar("Type")))

	ctx.NoError(utest.SetReqValue(req, "Raid", ctx.GetVar("Raid")))

	ctx.NoError(utest.SetReqValue(req, "Cluster", ctx.GetVar("Cluster")))

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
	}

	ctx.Vars["phostId"] = ctx.Must(utest.GetValue(resp, "PHostId.0"))
}

func testSet2377DescribePHost04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377GetPHostKVMInfo05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuphostClient.NewGetPHostKVMInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuphostClient.GetPHostKVMInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetPHostKVMInfoResponse", "str_eq"),
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

func testSet2377ModifyPHostInfo06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewModifyPHostInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("ModifyName")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("ModifyRemark")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.ModifyPHostInfo(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ModifyPHostInfoResponse", "str_eq"),
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

func testSet2377DescribePHost07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.Name", ctx.GetVar("ModifyName"), "str_eq"),
			ctx.NewValidator("PHostSet.0.Remark", ctx.GetVar("$ModifyRemark"), "str_eq"),
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

func testSet2377StopPHost08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iuphostClient.NewStopPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuphostClient.StopPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopPHostResponse", "str_eq"),
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

func testSet2377DescribePHost09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Stopped", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377ReinstallPHost10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewReinstallPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID1")))

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
	}

}

func testSet2377DescribePHost11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377PoweroffPHost12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewPoweroffPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

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
	}

}

func testSet2377DescribePHost13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Stopped", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377StartPHost14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := uphostClient.NewStartPHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.StartPHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StartPHostResponse", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377DescribePHost15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := uphostClient.NewDescribePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uphostClient.DescribePHost(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribePHostResponse", "str_eq"),
			ctx.NewValidator("PHostSet.0.PMStatus", "Running", "str_eq"),
		},
		MaxRetries:    120,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2377TerminatePHost16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uphostClient.NewTerminatePHostRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "PHostId", ctx.GetVar("phostId")))

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
	}

}
