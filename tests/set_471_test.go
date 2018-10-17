package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet471(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("CreateBootDisk", "20")
	ctx.SetVar("ChargeType", "Month")
	ctx.SetVar("CreateCPU", "1")
	ctx.SetVar("CreateMem", "1024")
	ctx.SetVar("CreateDiskspace", "0")
	ctx.SetVar("NewPassword", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("Name", "uhost-basic-api-ResetPassword")
	ctx.SetVar("ImageID", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet471CreateUHostInstance00(&ctx)
	testSet471DescribeUHostInstance01(&ctx)
	testSet471StopUHostInstance02(&ctx)
	testSet471DescribeUHostInstance03(&ctx)
	testSet471ResetUHostInstancePassword04(&ctx)
	testSet471DescribeUHostInstance05(&ctx)
	testSet471PoweroffUHostInstance06(&ctx)
	testSet471DescribeUHostInstance07(&ctx)
	testSet471TerminateUHostInstance08(&ctx)
}

func testSet471CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))
	ctx.NoError(utest.SetReqValue(req, "StorageType", "LocalDisk"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("CreateDiskspace")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "BootDiskSpace", ctx.GetVar("CreateBootDisk")))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "GPU", 0))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
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

	ctx.Vars["hostId"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet471DescribeUHostInstance01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Name", ctx.GetVar("Name"), "str_eq"),
			ctx.NewValidator("UHostSet.0.TotalDiskSpace", ctx.GetVar("CreateDiskspace"), "str_eq"),
			ctx.NewValidator("UHostSet.0.StorageType", "LocalDisk", "str_eq"),
			// ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
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

func testSet471StopUHostInstance02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.StopUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    5,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet471DescribeUHostInstance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet471ResetUHostInstancePassword04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := uhostClient.NewResetUHostInstancePasswordRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("NewPassword")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.ResetUHostInstancePassword(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet471DescribeUHostInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet471PoweroffUHostInstance06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet471DescribeUHostInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
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

func testSet471TerminateUHostInstance08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

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
