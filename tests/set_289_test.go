package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet289(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DiskType", "DataDisk")
	ctx.SetVar("Size", "1")
	ctx.SetVar("UDataArkMode", "No")
	ctx.SetVar("Name", "auto_udisk_nofz")

	testSet289DescribeUDiskPrice00(&ctx)
	testSet289CheckUDiskAllowance01(&ctx)
	testSet289CreateUDisk02(&ctx)
	testSet289DescribeUDisk03(&ctx)
	testSet289CreateUDiskSnapshot04(&ctx)
	testSet289DescribeUDisk05(&ctx)
	testSet289DescribeSnapshot06(&ctx)
	testSet289CreateUDiskSnapshot07(&ctx)
	testSet289DescribeSnapshot08(&ctx)
	testSet289CreateUDiskSnapshot09(&ctx)
	testSet289DescribeSnapshot10(&ctx)
	testSet289CreateUDiskSnapshot11(&ctx)
	testSet289CloneUDiskSnapshot12(&ctx)
	testSet289CloneUDiskSnapshot13(&ctx)
	testSet289RestoreUHostDisk14(&ctx)
	testSet289DescribeUDisk15(&ctx)
	testSet289DeleteSnapshot16(&ctx)
	testSet289DescribeSnapshot17(&ctx)
	testSet289DeleteSnapshot18(&ctx)
	testSet289DescribeSnapshot19(&ctx)
	testSet289DeleteSnapshot20(&ctx)
	testSet289DescribeSnapshot21(&ctx)
	testSet289DeleteUDisk22(&ctx)
	testSet289DescribeUDisk23(&ctx)
	testSet289DeleteUDisk24(&ctx)
	testSet289DescribeUDisk25(&ctx)
	testSet289DeleteUDisk26(&ctx)
}

func testSet289DescribeUDiskPrice00(ctx *utest.TestContext) {
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

func testSet289CheckUDiskAllowance01(ctx *utest.TestContext) {
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

func testSet289CreateUDisk02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

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
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["udisk_nofz_id"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet289DescribeUDisk03(ctx *utest.TestContext) {
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

func testSet289CreateUDiskSnapshot04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_01"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
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

func testSet289DescribeUDisk05(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.SnapshotLimit", "3", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			ctx.NewValidator("DataSet.0.SnapshotCount", "1", "str_eq"),
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

func testSet289DescribeSnapshot06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_nofz_id"), "str_eq"), // TODO: udisk nofz id
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		},
		MaxRetries:    50,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289CreateUDiskSnapshot07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_02"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
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

	ctx.Vars["snapshot_id1"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet289DescribeSnapshot08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_nofz_id"), "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		},
		MaxRetries:    6,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289CreateUDiskSnapshot09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_01"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
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

	ctx.Vars["snapshot_id2"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet289DescribeSnapshot10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_nofz_id"), "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		},
		MaxRetries:    6,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289CreateUDiskSnapshot11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_04"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "16999", "gt"),
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

func testSet289CloneUDiskSnapshot12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCloneUDiskSnapshotRequest()

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
			return udiskClient.CloneUDiskSnapshot(req)
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

	ctx.Vars["udisk_id_fromkz_fz"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet289CloneUDiskSnapshot13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewCloneUDiskSnapshotRequest()

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
			return udiskClient.CloneUDiskSnapshot(req)
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

func testSet289RestoreUHostDisk14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewRestoreUHostDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.RestoreUHostDisk(req)
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

func testSet289DescribeUDisk15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(50) * time.Second)

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
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289DeleteSnapshot16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot_id"), "str_eq"),
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

func testSet289DescribeSnapshot17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
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

func testSet289DeleteSnapshot18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot_id1"), "str_eq"),
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

func testSet289DescribeSnapshot19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "0", "str_eq"),
			ctx.NewValidator("PerDiskQuota", "3", "str_eq"),
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

func testSet289DeleteSnapshot20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot_id2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot_id2"), "str_eq"),
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

func testSet289DescribeSnapshot21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot_id2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("TotalCount", "0", "str_eq"),
			ctx.NewValidator("PerDiskQuota", "3", "str_eq"),
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

func testSet289DeleteUDisk22(ctx *utest.TestContext) {
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

func testSet289DescribeUDisk23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(100) * time.Second)

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
		MaxRetries:    3,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289DeleteUDisk24(ctx *utest.TestContext) {
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

func testSet289DescribeUDisk25(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

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
		MaxRetries:    3,
		RetryInterval: 30 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet289DeleteUDisk26(ctx *utest.TestContext) {
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
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
