package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet293(t *testing.T) {
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

	testSet293DescribeUDiskPrice00(&ctx)
	testSet293CheckUDiskAllowance01(&ctx)
	testSet293CreateUDisk02(&ctx)
	testSet293DescribeUDisk03(&ctx)
	testSet293SetUDiskUDataArkMode04(&ctx)
	testSet293DescribeUDisk05(&ctx)
	testSet293CreateUDiskSnapshot06(&ctx)
	testSet293DescribeUDisk07(&ctx)
	testSet293DescribeSnapshot08(&ctx)
	testSet293CloneUDiskSnapshot09(&ctx)
	testSet293DescribeUDisk10(&ctx)
	testSet293CloneUDiskSnapshot11(&ctx)
	testSet293DescribeUDisk12(&ctx)
	testSet293RestoreUHostDisk13(&ctx)
	testSet293DescribeUDisk14(&ctx)
	testSet293DeleteSnapshot15(&ctx)
	testSet293DescribeSnapshot16(&ctx)
	testSet293DescribeUDisk17(&ctx)
	testSet293DeleteUDisk18(&ctx)
	testSet293DescribeUDisk19(&ctx)
	testSet293DeleteUDisk20(&ctx)
	testSet293DescribeUDisk21(&ctx)
	testSet293DeleteUDisk22(&ctx)
}

func testSet293DescribeUDiskPrice00(ctx *utest.TestContext) {
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

func testSet293CheckUDiskAllowance01(ctx *utest.TestContext) {
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

func testSet293CreateUDisk02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udiskClient.NewCreateUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat("auto_", ctx.GetVar("Name")))))

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

func testSet293DescribeUDisk03(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
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

func testSet293SetUDiskUDataArkMode04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewSetUDiskUDataArkModeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", "Yes"))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.SetUDiskUDataArkMode(req)
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

func testSet293DescribeUDisk05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

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
			ctx.NewValidator("DataSet.0.UDataArkMode", "Yes", "str_eq"),
		},
		MaxRetries:    5,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet293CreateUDiskSnapshot06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(300) * time.Second)

	req := iudiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_01_fz"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01_fz"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudiskClient.CreateUDiskSnapshot(req)
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

	ctx.Vars["snapshot_id"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet293DescribeUDisk07(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.SnapshotLimit", "3", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			ctx.NewValidator("DataSet.0.SnapshotCount", "1", "str_eq"),
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

func testSet293DescribeSnapshot08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := iuhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_nofz_id"), "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
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

func testSet293CloneUDiskSnapshot09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudiskClient.NewCloneUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "clone_from_kz_fz"))
	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("snapshot_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", "Yes"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudiskClient.CloneUDiskSnapshot(req)
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

	ctx.Vars["udisk_id_fromkz_fz"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet293DescribeUDisk10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_fz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Name", "clone_from_kz_fz", "str_eq"),
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

func testSet293CloneUDiskSnapshot11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iudiskClient.NewCloneUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "clone_from_kz_nofz"))
	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("snapshot_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", "No"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudiskClient.CloneUDiskSnapshot(req)
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

	ctx.Vars["udisk_id_fromkz_nofz"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet293DescribeUDisk12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_nofz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Name", "clone_from_kz_nofz", "str_eq"),
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

func testSet293RestoreUHostDisk13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := iuhostClient.NewRestoreUHostDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.RestoreUHostDisk(req)
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

func testSet293DescribeUDisk14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

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
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat("auto_", ctx.GetVar("Name"))), "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
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

func testSet293DeleteSnapshot15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot_id"), "str_eq"),
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

func testSet293DescribeSnapshot16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := iuhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "0", "str_eq"),
			ctx.NewValidator("PerDiskQuota", "3", "str_eq"),
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

func testSet293DescribeUDisk17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet293DeleteUDisk18(ctx *utest.TestContext) {
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

func testSet293DescribeUDisk19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_nofz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet293DeleteUDisk20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_nofz")))

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

func testSet293DescribeUDisk21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_fz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet293DeleteUDisk22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_id_fromkz_fz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
