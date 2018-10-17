package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1839(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("CreateCPU", "1")
	ctx.SetVar("CreateBootDisk", "20")
	ctx.SetVar("ChargeType", "Month")
	ctx.SetVar("CreateMem", "1024")
	ctx.SetVar("CreateDiskspace", "0")
	ctx.SetVar("ModifyTag", "newTag")
	ctx.SetVar("ModifyName", "new-uhost-api-status-test")
	ctx.SetVar("ModifyRemark", "newRemark")
	ctx.SetVar("Name", "uhost-api-status-test")
	ctx.SetVar("ImageID", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet1839GetUImagePrice00(&ctx)
	testSet1839GetRestResource01(&ctx)
	testSet1839CreateUHostInstance02(&ctx)
	testSet1839DescribeUHostInstance03(&ctx)
	testSet1839DescribeUHostLite04(&ctx)
	testSet1839DescribeUHostRecycle05(&ctx)
	testSet1839ModifyUHostInstanceTag06(&ctx)
	testSet1839DescribeUHostInstance07(&ctx)
	testSet1839ModifyUHostInstanceName08(&ctx)
	testSet1839DescribeUHostInstance09(&ctx)
	testSet1839ModifyUHostInstanceRemark10(&ctx)
	testSet1839DescribeUHostInstance11(&ctx)
	testSet1839DescribeUHostTags12(&ctx)
	testSet1839GetUHostInstancePrice13(&ctx)
	testSet1839GetUHostUpgradePrice14(&ctx)
	testSet1839GetUHostRenewPrice15(&ctx)
	testSet1839GetUHostInstanceVncInfo16(&ctx)
	testSet1839DescribeResourceMetric17(&ctx)
	testSet1839GetMetric18(&ctx)
	testSet1839DescribeSecurityGroup19(&ctx)
	testSet1839StopUHostInstance20(&ctx)
	testSet1839DescribeUHostInstance21(&ctx)
	testSet1839StartUHostInstance22(&ctx)
	testSet1839DescribeUHostInstance23(&ctx)
	testSet1839PoweroffUHostInstance24(&ctx)
	testSet1839DescribeUHostInstance25(&ctx)
	testSet1839TerminateUHostInstance26(&ctx)
}

func testSet1839GetUImagePrice00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewGetUImagePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Count", "1"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.GetUImagePrice(req)
		},
		Validators: []utest.TestValidator{
			// ctx.NewValidator("PriceSet.0.Price", "0", "ne"),
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

func testSet1839GetRestResource01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewGetRestResourceRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceType", "GPU"))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.GetRestResource(req)
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

func testSet1839CreateUHostInstance02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("CreateDiskspace")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	ctx.NoError(utest.SetReqValue(req, "BootDiskSpace", ctx.GetVar("CreateBootDisk")))

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

func testSet1839DescribeUHostInstance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

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
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId"), "str_eq"),
			ctx.NewValidator("UHostSet.0.TotalDiskSpace", ctx.GetVar("CreateDiskspace"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Name", ctx.GetVar("Name"), "str_eq"),
			// ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
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

func testSet1839DescribeUHostLite04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewDescribeUHostLiteRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeUHostLite(req)
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

func testSet1839DescribeUHostRecycle05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewDescribeUHostRecycleRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeUHostRecycle(req)
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

func testSet1839ModifyUHostInstanceTag06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewModifyUHostInstanceTagRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("ModifyTag")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.ModifyUHostInstanceTag(req)
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

func testSet1839DescribeUHostInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId"), "str_eq"),
			ctx.NewValidator("UHostSet.0.TotalDiskSpace", ctx.GetVar("CreateDiskspace"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Tag", ctx.GetVar("ModifyTag"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Name", ctx.GetVar("Name"), "str_eq"),
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

func testSet1839ModifyUHostInstanceName08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewModifyUHostInstanceNameRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("ModifyName")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.ModifyUHostInstanceName(req)
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

func testSet1839DescribeUHostInstance09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId"), "str_eq"),
			ctx.NewValidator("UHostSet.0.TotalDiskSpace", ctx.GetVar("CreateDiskspace"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Tag", ctx.GetVar("ModifyTag"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Name", ctx.GetVar("ModifyName"), "str_eq"),
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

func testSet1839ModifyUHostInstanceRemark10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewModifyUHostInstanceRemarkRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("ModifyRemark")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.ModifyUHostInstanceRemark(req)
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

func testSet1839DescribeUHostInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId"), "str_eq"),
			ctx.NewValidator("UHostSet.0.TotalDiskSpace", ctx.GetVar("CreateDiskspace"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Tag", ctx.GetVar("ModifyTag"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Name", ctx.GetVar("ModifyName"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Remark", ctx.GetVar("ModifyRemark"), "str_eq"),
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

func testSet1839DescribeUHostTags12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeUHostTagsRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostTags(req)
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

func testSet1839GetUHostInstancePrice13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewGetUHostInstancePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))
	ctx.NoError(utest.SetReqValue(req, "Count", "1"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Dynamic"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.GetUHostInstancePrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("PriceSet.0.Price", "0", "ne"),
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

func testSet1839GetUHostUpgradePrice14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewGetUHostUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "CPU", "8"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "8192"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.GetUHostUpgradePrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Price", "0", "ne"),
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

func testSet1839GetUHostRenewPrice15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewGetUHostRenewPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.GetUHostRenewPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("PriceSet.0.Price", "0", "ne"),
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

func testSet1839GetUHostInstanceVncInfo16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewGetUHostInstanceVncInfoRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.GetUHostInstanceVncInfo(req)
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

func testSet1839DescribeResourceMetric17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumonClient.NewDescribeResourceMetricRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceType", "uhost"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumonClient.DescribeResourceMetric(req)
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

func testSet1839GetMetric18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumonClient.NewGetMetricRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceType", "uhost"))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "MetricName", "CPUUtilization"))

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumonClient.GetMetric(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "GetMetricResponse", "str_eq"),
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

func testSet1839DescribeSecurityGroup19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iunetClient.NewDescribeSecurityGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iunetClient.DescribeSecurityGroup(req)
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

func testSet1839StopUHostInstance20(ctx *utest.TestContext) {
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

func testSet1839DescribeUHostInstance21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet1839StartUHostInstance22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStartUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.StartUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StartUHostInstanceResponse", "str_eq"),
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

func testSet1839DescribeUHostInstance23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
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

func testSet1839PoweroffUHostInstance24(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

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
			ctx.NewValidator("Action", "PoweroffUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1839DescribeUHostInstance25(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet1839TerminateUHostInstance26(ctx *utest.TestContext) {
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
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
