package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet283(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-05")
	ctx.SetVar("DiskType", "DataDisk")
	ctx.SetVar("Size", "1")
	ctx.SetVar("UDataArkMode", "No")
	ctx.SetVar("Name", "udisk_nofz")
	ctx.SetVar("imageid", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet283DescribeUDiskPrice00(&ctx)
	testSet283CheckUDiskAllowance01(&ctx)
	testSet283CreateUDisk02(&ctx)
	testSet283DescribeUDisk03(&ctx)
	testSet283RenameUDisk04(&ctx)
	testSet283DescribeUDisk05(&ctx)
	testSet283DescribeImage06(&ctx)
	testSet283CreateUHostInstance07(&ctx)
	testSet283DescribeUHostInstance08(&ctx)
	testSet283DescribeUHostLite09(&ctx)
	testSet283AttachUDisk10(&ctx)
	testSet283DescribeUDisk11(&ctx)
	testSet283DetachUDisk12(&ctx)
	testSet283DescribeUDisk13(&ctx)
	testSet283DeleteUDisk14(&ctx)
	testSet283PoweroffUHostInstance15(&ctx)
	testSet283DescribeUHostInstance16(&ctx)
	testSet283TerminateUHostInstance17(&ctx)
}

func testSet283DescribeUDiskPrice00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", ctx.GetVar("UDataArkMode")))
	ctx.NoError(utest.SetReqValue(req, "DiskType", ctx.GetVar("DiskType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDiskPrice(req)
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

func testSet283CheckUDiskAllowance01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudiskClient.NewCheckUDiskAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))
	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudiskClient.CheckUDiskAllowance(req)
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

func testSet283CreateUDisk02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udiskClient.NewCreateUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", ctx.GetVar("UDataArkMode")))

	ctx.NoError(utest.SetReqValue(req, "DiskType", ctx.GetVar("DiskType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDisk(req)
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

	ctx.Vars["udisk_nofz_id"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet283DescribeUDisk03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			ctx.NewValidator("DataSet.0.Tag", "Default", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet283RenameUDisk04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewRenameUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "UDiskName", ctx.Must(utest.Concat("auto_", ctx.GetVar("Name")))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.RenameUDisk(req)
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

func testSet283DescribeUDisk05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat("auto_", ctx.GetVar("Name"))), "str_eq"),
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

func testSet283DescribeImage06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageType", "Base"))
	ctx.NoError(utest.SetReqValue(req, "OsType", "Linux"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeImage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["imageid"] = ctx.Must(utest.GetValue(resp, "ImageSet.0.ImageId"))
}

func testSet283CreateUHostInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("imageid")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", "Z3VhbmxpeXVhbm1pbWExMjMhQCM="))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))
	ctx.NoError(utest.SetReqValue(req, "StorageType", "LocalDisk"))
	ctx.NoError(utest.SetReqValue(req, "Name", "auto_uhost-同AZ"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))
	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))

	ctx.NoError(utest.SetReqValue(req, "BootDiskSpace", "20"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))
	// ctx.NoError(utest.SetReqValue(req, "HostType", "N1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "GPU", 0)) // TODO: check

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
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

	ctx.Vars["uhost_id"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet283DescribeUHostInstance08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("uhost_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("uhost_id"), "str_eq"),
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

func testSet283DescribeUHostLite09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := iuhostClient.NewDescribeUHostLiteRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "60"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeUHostLite(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "0", "ne"),
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

func testSet283AttachUDisk10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := udiskClient.NewAttachUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.AttachUDisk(req)
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

func testSet283DescribeUDisk11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "InUse", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet283DetachUDisk12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDetachUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DetachUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UDiskId", ctx.GetVar("udisk_nofz_id"), "str_eq"),
			ctx.NewValidator("UHostId", ctx.GetVar("uhost_id"), "str_eq"),
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

func testSet283DescribeUDisk13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    50,
		RetryInterval: 6 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet283DeleteUDisk14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDisk(req)
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

func testSet283PoweroffUHostInstance15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
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

func testSet283DescribeUHostInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("uhost_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
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

func testSet283TerminateUHostInstance17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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
