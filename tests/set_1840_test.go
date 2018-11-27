package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1840(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Password", "2012_UClou")
	ctx.SetVar("SnapshotSysName", "snapshot-ARK-SYS-01")
	ctx.SetVar("SnapshotSysDesc", "snapshot-ARK-SYS-01-desc")
	ctx.SetVar("SnapDiskType", "LocalBoot")
	ctx.SetVar("SnapshotDataNameModify", "snapshot-ARK-DATA-01-modify")
	ctx.SetVar("SnapshotDataDescModify", "snapshot-ARK-DATA-01-desc-Modify")
	ctx.SetVar("UhostName", "uhost-snapshot-ARK-auto-api-1")
	ctx.SetVar("SnapshotDataName", "snapshot-ARK-DATA-01")
	ctx.SetVar("SnapshotDataDesc", "snapshot-ARK-DATA-01-desc")
	ctx.SetVar("CreateFromTimeMachinePassword", "Z3VhbmxpeXVhbm1pbWExMjMhQCM=")
	ctx.SetVar("ImageID", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))

	testSet1840CreateUHostInstance00(&ctx)
	testSet1840DescribeUHostInstance01(&ctx)
	testSet1840StopUHostInstance02(&ctx)
	testSet1840DescribeUHostInstance03(&ctx)
	testSet1840UpgradeToArkUHostInstance04(&ctx)
	testSet1840DescribeUHostInstance05(&ctx)
	testSet1840StartUHostInstance06(&ctx)
	testSet1840DescribeUHostInstance07(&ctx)
	testSet1840DescribeUhostTmMeta08(&ctx)
	testSet1840DescribeVDiskTmList09(&ctx)
	testSet1840DescribeVDiskTmList10(&ctx)
	testSet1840StopUHostInstance11(&ctx)
	testSet1840DescribeUHostInstance12(&ctx)
	testSet1840TerminateUHostInstance13(&ctx)
}

func testSet1840CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("ImageID")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))

	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("UhostName")))

	ctx.NoError(utest.SetReqValue(req, "TimemachineFeature", "no"))
	ctx.NoError(utest.SetReqValue(req, "HotplugFeature", "false"))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", 20))

	ctx.NoError(utest.SetReqValue(req, "GPU", 0)) // TODO: check

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

func testSet1840DescribeUHostInstance01(ctx *utest.TestContext) {
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
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.TimemachineFeature", "no", "str_eq"),
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

func testSet1840StopUHostInstance02(ctx *utest.TestContext) {
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

func testSet1840DescribeUHostInstance03(ctx *utest.TestContext) {
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
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1840UpgradeToArkUHostInstance04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := uhostClient.NewUpgradeToArkUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostIds", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.UpgradeToArkUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "UpgradeToArkUHostInstanceResponse", "str_eq"),
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

func testSet1840DescribeUHostInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

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
			ctx.NewValidator("UHostSet.0.TimemachineFeature", "yes", "str_eq"),
			ctx.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.BackupType", "DATAARK", "str_eq"),
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

func testSet1840StartUHostInstance06(ctx *utest.TestContext) {
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

func testSet1840DescribeUHostInstance07(ctx *utest.TestContext) {
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
			ctx.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("UHostSet.0.TimemachineFeature", "yes", "str_eq"),
			ctx.NewValidator("UHostSet.0.DiskSet.0.BackupType", "DATAARK", "str_eq"),
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

func testSet1840DescribeUhostTmMeta08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

	req := iudataarkClient.NewDescribeUhostTmMetaRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UhostId", ctx.GetVar("hostId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudataarkClient.DescribeUhostTmMeta(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUhostTmMetaResponse", "str_eq"),
			ctx.NewValidator("UtmStatus", "normal", "str_eq"),
		},
		MaxRetries:    60,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VdiskIdSys"] = ctx.Must(utest.GetValue(resp, "DataSet.0.VdiskId"))
	ctx.Vars["VdiskIdData"] = ctx.Must(utest.GetValue(resp, "DataSet.1.VdiskId"))
}

func testSet1840DescribeVDiskTmList09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudataarkClient.NewDescribeVDiskTmListRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "VDiskId", ctx.GetVar("VdiskIdSys")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotType", "all"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudataarkClient.DescribeVDiskTmList(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeVDiskTmListResponse", "str_eq"),
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

func testSet1840DescribeVDiskTmList10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudataarkClient.NewDescribeVDiskTmListRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "VDiskId", ctx.GetVar("VdiskIdData")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotType", "all"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudataarkClient.DescribeVDiskTmList(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeVDiskTmListResponse", "str_eq"),
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

func testSet1840StopUHostInstance11(ctx *utest.TestContext) {
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

func testSet1840DescribeUHostInstance12(ctx *utest.TestContext) {
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
		MaxRetries:    30,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1840TerminateUHostInstance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
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
