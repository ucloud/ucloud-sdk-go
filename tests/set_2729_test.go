package tests

import (
	"strconv"
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2729(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("GroupName", "CreateIgName1")
	ctx.SetVar("Password", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("ChargeType", "Month")
	ctx.SetVar("CreateCPU", "1")
	ctx.SetVar("CreateMem", "1024")
	ctx.SetVar("NewPassword", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("Name", "uhost-ig-api-N2-Normal-LocalDisk")
	ctx.SetVar("ImageID", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))
	ctx.SetVar("BootSize", "20")
	ctx.SetVar("BootType", "LOCAL_NORMAL")
	ctx.SetVar("DiskSize", "20")
	ctx.SetVar("DiskType", "LOCAL_NORMAL")
	ctx.SetVar("BootBackup", "NONE")
	ctx.SetVar("DiskBackup", "NONE")
	ctx.SetVar("UHostType", "N2")

	testSet2729DescribeImage00(&ctx)
	testSet2729CreateIsolationGroup01(&ctx)
	testSet2729DescribeIsolationGroup02(&ctx)
	testSet2729CreateUHostInstance03(&ctx)
	testSet2729CreateUHostInstance04(&ctx)
	testSet2729CreateUHostInstance05(&ctx)
	testSet2729CreateUHostInstance06(&ctx)
	testSet2729CreateUHostInstance07(&ctx)
	testSet2729CreateUHostInstance08(&ctx)
	testSet2729CreateUHostInstance09(&ctx)
	testSet2729CreateUHostInstance10(&ctx)
	testSet2729DescribeUHostInstance11(&ctx)
	testSet2729LeaveIsolationGroup12(&ctx)
	testSet2729DescribeUHostInstance13(&ctx)
	testSet2729DescribeUHostInstance14(&ctx)
	testSet2729DescribeIsolationGroup15(&ctx)
	testSet2729CreateUHostInstance16(&ctx)
	testSet2729DescribeIsolationGroup17(&ctx)
	testSet2729DescribeUHostInstance18(&ctx)
	testSet2729DeleteIsolationGroup19(&ctx)
	testSet2729DescribeUHostInstance20(&ctx)
	testSet2729DescribeUHostInstance21(&ctx)
	testSet2729StopUHostInstance22(&ctx)
	testSet2729StopUHostInstance23(&ctx)
	testSet2729StopUHostInstance24(&ctx)
	testSet2729StopUHostInstance25(&ctx)
	testSet2729StopUHostInstance26(&ctx)
	testSet2729StopUHostInstance27(&ctx)
	testSet2729StopUHostInstance28(&ctx)
	testSet2729StopUHostInstance29(&ctx)
	testSet2729DescribeUHostInstance30(&ctx)
	testSet2729TerminateUHostInstance31(&ctx)
	testSet2729TerminateUHostInstance32(&ctx)
	testSet2729TerminateUHostInstance33(&ctx)
	testSet2729TerminateUHostInstance34(&ctx)
	testSet2729TerminateUHostInstance35(&ctx)
	testSet2729TerminateUHostInstance36(&ctx)
	testSet2729TerminateUHostInstance37(&ctx)
	testSet2729TerminateUHostInstance38(&ctx)
}

func testSet2729DescribeImage00(ctx *utest.TestContext) {
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

		ctx.T.Error(err)

	}

	ctx.Vars["ImageID"] = ctx.Must(utest.GetValue(resp, "ImageSet.0.ImageId"))
}

func testSet2729CreateIsolationGroup01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupName", ctx.GetVar("GroupName")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateIsolationGroupResponse", "str_eq"),
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

	ctx.Vars["GroupId"] = ctx.Must(utest.GetValue(resp, "GroupId"))
}

func testSet2729DescribeIsolationGroup02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("GroupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeIsolationGroupResponse", "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupName", ctx.GetVar("GroupName"), "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupId", ctx.GetVar("GroupId"), "str_eq"),
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

func testSet2729CreateUHostInstance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))
	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId1"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId2"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId3"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))
	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId4"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId5"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId6"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
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

	ctx.Vars["hostId7"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729CreateUHostInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "ne"),
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

func testSet2729DescribeUHostInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId1"), ctx.GetVar("hostId2"), ctx.GetVar("hostId3"), ctx.GetVar("hostId4"), ctx.GetVar("hostId5"), ctx.GetVar("hostId6"), ctx.GetVar("hostId7")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.1.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.2.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.3.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.4.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.5.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.6.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.1.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.2.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.3.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.4.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.5.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.6.State", "Running", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2729LeaveIsolationGroup12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewLeaveIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("GroupId")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.LeaveIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "LeaveIsolationGroupResponse", "str_eq"),
			ctx.NewValidator("UHostId", ctx.GetVar("hostId1"), "str_eq"),
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

func testSet2729DescribeUHostInstance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.UHostId", ctx.GetVar("hostId1"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.IsolationGroup", "0", "len_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729DescribeUHostInstance14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId2"), ctx.GetVar("hostId3"), ctx.GetVar("hostId4"), ctx.GetVar("hostId5"), ctx.GetVar("hostId6"), ctx.GetVar("hostId7")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.1.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.2.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.3.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.4.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("UHostSet.5.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729DescribeIsolationGroup15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("GroupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeIsolationGroupResponse", "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupName", ctx.GetVar("GroupName"), "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupId", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.SpreadInfoSet.0.UHostCount", "6", "str_eq"),
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

func testSet2729CreateUHostInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", ctx.GetVar("ChargeType")))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "UHostType", ctx.GetVar("UHostType")))
	ctx.NoError(utest.SetReqValue(req, "CPU", ctx.GetVar("CreateCPU")))
	ctx.NoError(utest.SetReqValue(req, "Memory", ctx.GetVar("CreateMem")))

	ctx.NoError(utest.SetReqValue(req, "NetCapability", "Normal"))
	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "No"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))

	ctx.NoError(utest.SetReqValue(req, "IsolationGroup", ctx.GetVar("GroupId")))

	bootSize, _ := strconv.Atoi(ctx.GetVar("BootSize").(string))
	diskSize, _ := strconv.Atoi(ctx.GetVar("DiskSize").(string))
	req.Disks = []uhost.UHostDisk{
		{
			Size:       ucloud.Int(bootSize),
			IsBoot:     ucloud.String("True"),
			Type:       ucloud.String(ctx.GetVar("BootType").(string)),
			BackupType: ucloud.String(ctx.GetVar("BootBackup").(string)),
		},
		{
			Size:       ucloud.Int(diskSize),
			IsBoot:     ucloud.String("False"),
			Type:       ucloud.String(ctx.GetVar("DiskType").(string)),
			BackupType: ucloud.String(ctx.GetVar("DiskBackup").(string)),
		},
	}

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
		ctx.T.FailNow()

	}

	ctx.Vars["hostId8"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet2729DescribeIsolationGroup17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("GroupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeIsolationGroupResponse", "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupName", ctx.GetVar("GroupName"), "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.GroupId", ctx.GetVar("GroupId"), "str_eq"),
			ctx.NewValidator("IsolationGroupSet.0.SpreadInfoSet.0.UHostCount", "7", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729DescribeUHostInstance18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId8")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.IsolationGroup", ctx.GetVar("GroupId"), "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2729DeleteIsolationGroup19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDeleteIsolationGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("GroupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DeleteIsolationGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteIsolationGroupResponse", "str_eq"),
			ctx.NewValidator("GroupId", ctx.GetVar("GroupId"), "str_eq"),
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

func testSet2729DescribeUHostInstance20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId1"), ctx.GetVar("hostId2"), ctx.GetVar("hostId3"), ctx.GetVar("hostId4"), ctx.GetVar("hostId5"), ctx.GetVar("hostId6"), ctx.GetVar("hostId7"), ctx.GetVar("hostId8")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.1.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.2.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.3.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.4.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.5.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.6.IsolationGroup", "0", "len_eq"),
			ctx.NewValidator("UHostSet.7.IsolationGroup", "0", "len_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729DescribeUHostInstance21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId1"), ctx.GetVar("hostId2"), ctx.GetVar("hostId3"), ctx.GetVar("hostId4"), ctx.GetVar("hostId5"), ctx.GetVar("hostId6"), ctx.GetVar("hostId7"), ctx.GetVar("hostId8")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.CPU", ctx.GetVar("CreateCPU"), "str_eq"),
			ctx.NewValidator("UHostSet.0.Memory", ctx.GetVar("CreateMem"), "str_eq"),
			ctx.NewValidator("UHostSet.0.HostType", ctx.GetVar("UHostType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.BasicImageId", ctx.GetVar("ImageID"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.DiskType", ctx.GetVar("BootType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.DiskType", ctx.GetVar("DiskType"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.Size", ctx.GetVar("BootSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.1.Size", ctx.GetVar("DiskSize"), "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.1.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.2.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.3.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.4.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.5.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.6.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.7.State", "Running", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId1")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId2")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance24(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId3")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance25(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId4")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance26(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId5")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance27(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId6")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance28(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId7")))

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
		ctx.T.FailNow()

	}

}

func testSet2729StopUHostInstance29(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewStopUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId8")))

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

		ctx.T.Error(err)

	}

}

func testSet2729DescribeUHostInstance30(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewDescribeUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId1"), ctx.GetVar("hostId2"), ctx.GetVar("hostId3"), ctx.GetVar("hostId4"), ctx.GetVar("hostId5"), ctx.GetVar("hostId6"), ctx.GetVar("hostId7"), ctx.GetVar("hostId8")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.DescribeUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			ctx.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.1.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.2.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.3.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.4.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.5.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.6.State", "Stopped", "str_eq"),
			ctx.NewValidator("UHostSet.7.State", "Stopped", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet2729TerminateUHostInstance31(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
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

func testSet2729TerminateUHostInstance32(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance33(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId3")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance34(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId4")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance35(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId5")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance36(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId6")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance37(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId7")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet2729TerminateUHostInstance38(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("hostId8")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}
