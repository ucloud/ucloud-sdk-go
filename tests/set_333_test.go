package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet333(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "2012_UClou")
	ctx.SetVar("ImageName", "ImageTest")
	ctx.SetVar("TargetImageName", "ImageCopyTest")
	ctx.SetVar("TargetRegion", "cn-sh2")
	ctx.SetVar("TargetZone", "cn-sh2-02")
	ctx.SetVar("myImage", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet333CreateUHostInstance00(&ctx)
	testSet333DescribeUHostInstance01(&ctx)
	testSet333StopUHostInstance02(&ctx)
	testSet333DescribeUHostInstance03(&ctx)
	testSet333CreateCustomImage04(&ctx)
	testSet333DescribeImage05(&ctx)
	testSet333CreateUHostInstance06(&ctx)
	testSet333DescribeUHostInstance07(&ctx)
	testSet333GetProjectList08(&ctx)
	testSet333CopyCustomImage09(&ctx)
	testSet333GetImageCopyProgress10(&ctx)
	testSet333DescribeImage11(&ctx)
	testSet333CreateUHostInstance12(&ctx)
	testSet333DescribeUHostInstance13(&ctx)
	testSet333TerminateCustomImage14(&ctx)
	testSet333StopUHostInstance15(&ctx)
	testSet333StopUHostInstance16(&ctx)
	testSet333DescribeUHostInstance17(&ctx)
	testSet333TerminateUHostInstance18(&ctx)
	testSet333TerminateUHostInstance19(&ctx)
	testSet333TerminateCustomImage20(&ctx)
	testSet333StopUHostInstance21(&ctx)
	testSet333DescribeUHostInstance22(&ctx)
	testSet333TerminateUHostInstance23(&ctx)
}

func testSet333CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("myImage")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

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

func testSet333DescribeUHostInstance01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(400) * time.Second)

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
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet333StopUHostInstance02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

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
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet333DescribeUHostInstance03(ctx *utest.TestContext) {
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
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
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

func testSet333CreateCustomImage04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewCreateCustomImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "ImageName", ctx.GetVar("ImageName")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateCustomImage(req)
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

	ctx.Vars["newImageId"] = ctx.Must(utest.GetValue(resp, "ImageId"))
}

func testSet333DescribeImage05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := uhostClient.NewDescribeImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("newImageId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeImage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeImageResponse", "str_eq"),
			ctx.NewValidator("ImageSet.0.State", "Available", "str_eq"),
			ctx.NewValidator("ImageSet.0.ImageId", ctx.GetVar("newImageId"), "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet333CreateUHostInstance06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("newImageId")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

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

	ctx.Vars["hostId2"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet333DescribeUHostInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet333GetProjectList08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uaccountClient.NewGetProjectListRequest()

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uaccountClient.GetProjectList(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetProjectListResponse", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["TargetProjectID"] = ctx.Must(utest.GetValue(resp, "ProjectSet.0.ProjectId"))
}

func testSet333CopyCustomImage09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewCopyCustomImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SourceImageId", ctx.GetVar("newImageId")))
	ctx.NoError(utest.SetReqValue(req, "TargetRegion", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "TargetProjectId", ctx.GetVar("TargetProjectID")))
	ctx.NoError(utest.SetReqValue(req, "TargetImageName", ctx.GetVar("TargetImageName")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CopyCustomImage(req)
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

	ctx.Vars["cpImageId"] = ctx.Must(utest.GetValue(resp, "TargetImageId"))
}

func testSet333GetImageCopyProgress10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iuhostClient.NewGetImageCopyProgressRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("cpImageId")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.GetImageCopyProgress(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("ProgressInfo", "100.00", "float_eq"),
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

func testSet333DescribeImage11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("cpImageId")))

	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeImage(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeImageResponse", "str_eq"),
			ctx.NewValidator("ImageSet.0.State", "Available", "str_eq"),
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

func testSet333CreateUHostInstance12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("cpImageId")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

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

	ctx.Vars["hostId_new"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet333DescribeUHostInstance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId_new")))

	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet333TerminateCustomImage14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateCustomImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("newImageId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateCustomImage(req)
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

func testSet333StopUHostInstance15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet333StopUHostInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.StopUHostInstance(req)
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

func testSet333DescribeUHostInstance17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId"), ctx.GetVar("hostId2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.1.State", "Stopped", "str_eq"),
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

func testSet333TerminateUHostInstance18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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

func testSet333TerminateUHostInstance19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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

func testSet333TerminateCustomImage20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

	req := uhostClient.NewTerminateCustomImageRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("cpImageId")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateCustomImage(req)
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

func testSet333StopUHostInstance21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId_new")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.StopUHostInstance(req)
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

func testSet333DescribeUHostInstance22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId_new")))

	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
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

func testSet333TerminateUHostInstance23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("TargetRegion")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("TargetZone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId_new")))
	ctx.NoError(utest.SetReqValue(req, "ProjectId", ctx.GetVar("TargetProjectID")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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
