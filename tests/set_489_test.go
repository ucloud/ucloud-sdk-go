package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet489(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DBTypeId", "mysql-5.6")
	ctx.SetVar("InstanceMode", "HA")
	ctx.SetVar("InstanceType", "SATA_SSD")
	ctx.SetVar("Port", "3306")
	ctx.SetVar("MemoryLimit", "1000")
	ctx.SetVar("DiskSpace", "20")
	ctx.SetVar("DBName", "auto_hassd_")
	ctx.SetVar("BackupName", "handbf_test_02")
	ctx.SetVar("UseSSD", "true")

	testSet489DescribeUDBType00(&ctx)
	testSet489DescribeUDBParamGroup01(&ctx)
	testSet489DescribeUDBInstancePrice02(&ctx)
	testSet489CheckUDBInstanceAllowance03(&ctx)
	testSet489CreateUDBInstance04(&ctx)
	testSet489DescribeUDBInstance05(&ctx)
	testSet489DescribeUDBInstanceState06(&ctx)
	testSet489UpdateUDBInstanceBackupStrategy07(&ctx)
	testSet489EditUDBBackupBlacklist08(&ctx)
	testSet489DescribeUDBBackupBlacklist09(&ctx)
	testSet489BackupUDBInstance10(&ctx)
	testSet489DescribeUDBBackup11(&ctx)
	testSet489DescribeUDBInstanceBackupState12(&ctx)
	testSet489CheckUDBInstanceAllowance13(&ctx)
	testSet489DescribeUDBInstanceBackupURL14(&ctx)
	testSet489CheckUDBInstanceAllowance15(&ctx)
	testSet489CreateUDBInstance16(&ctx)
	testSet489DescribeUDBInstanceState17(&ctx)
	testSet489DescribeUDBInstance18(&ctx)
	testSet489StopUDBInstance19(&ctx)
	testSet489DescribeUDBInstance20(&ctx)
	testSet489FetchUDBInstanceEarliestRecoverTime21(&ctx)
	testSet489CheckUDBInstanceAllowance22(&ctx)
	testSet489CreateUDBInstanceByRecovery23(&ctx)
	testSet489DescribeUDBInstanceState24(&ctx)
	testSet489DescribeUDBInstance25(&ctx)
	testSet489StopUDBInstance26(&ctx)
	testSet489DescribeUDBInstance27(&ctx)
	testSet489DeleteUDBBackup28(&ctx)
	testSet489StopUDBInstance29(&ctx)
	testSet489DescribeUDBInstance30(&ctx)
	testSet489DeleteUDBInstance31(&ctx)
	testSet489DeleteUDBInstance32(&ctx)
	testSet489DeleteUDBInstance33(&ctx)
}

func testSet489DescribeUDBType00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBTypeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBType(req)
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

func testSet489DescribeUDBParamGroup01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBParamGroup(req)
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

	ctx.Vars["DataSet_paramGroup"] = ctx.Must(utest.GetValue(resp, "DataSet"))
}

func testSet489DescribeUDBInstancePrice02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstancePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))

	ctx.NoError(utest.SetReqValue(req, "DBTypeId", ctx.GetVar("DBTypeId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstancePrice(req)
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

func testSet489CheckUDBInstanceAllowance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "SQL"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))
	ctx.NoError(utest.SetReqValue(req, "UseSSD", ctx.GetVar("UseSSD")))

	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.CheckUDBInstanceAllowance(req)
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

func testSet489CreateUDBInstance04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewCreateUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat(ctx.GetVar("DBName"), ctx.GetVar("DBTypeId")))))

	ctx.NoError(utest.SetReqValue(req, "AdminPassword", "guanliyuanmima"))
	ctx.NoError(utest.SetReqValue(req, "DBTypeId", ctx.GetVar("DBTypeId")))
	ctx.NoError(utest.SetReqValue(req, "Port", ctx.GetVar("Port")))
	ctx.NoError(utest.SetReqValue(req, "ParamGroupId", ctx.Must(utest.SearchValue(ctx.GetVar("DataSet_paramGroup"), "DBTypeId", ctx.GetVar("DBTypeId"), "GroupId"))))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))

	ctx.NoError(utest.SetReqValue(req, "InstanceMode", ctx.GetVar("InstanceMode")))

	ctx.NoError(utest.SetReqValue(req, "InstanceType", ctx.GetVar("InstanceType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.CreateUDBInstance(req)
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

	ctx.Vars["DBId"] = ctx.Must(utest.GetValue(resp, "DBId"))
}

func testSet489DescribeUDBInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat(ctx.GetVar("DBName"), ctx.GetVar("DBTypeId"))), "str_eq"),
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
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

func testSet489DescribeUDBInstanceState06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("State", "Running", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489UpdateUDBInstanceBackupStrategy07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewUpdateUDBInstanceBackupStrategyRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "BackupTime", "12"))
	ctx.NoError(utest.SetReqValue(req, "BackupDate", "1111110"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.UpdateUDBInstanceBackupStrategy(req)
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

func testSet489EditUDBBackupBlacklist08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewEditUDBBackupBlacklistRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "Blacklist", "test.%"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.EditUDBBackupBlacklist(req)
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

func testSet489DescribeUDBBackupBlacklist09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBBackupBlacklistRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBBackupBlacklist(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Blacklist", "test.%", "str_eq"),
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

func testSet489BackupUDBInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(15) * time.Second)

	req := udbClient.NewBackupUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "BackupName", ctx.GetVar("BackupName")))
	ctx.NoError(utest.SetReqValue(req, "UseBlacklist", "true"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.BackupUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    2,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489DescribeUDBBackup11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBBackupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))
	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBBackup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["BackupId"] = ctx.Must(utest.GetValue(resp, "DataSet.0.BackupId"))
}

func testSet489DescribeUDBInstanceBackupState12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceBackupStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("BackupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceBackupState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("State", "Success", "str_eq"),
			ctx.NewValidator("BackupSize", "0", "ne"),
			ctx.NewValidator("BackupEndTime", "0", "ne"),
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

func testSet489CheckUDBInstanceAllowance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "SQL"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))
	ctx.NoError(utest.SetReqValue(req, "UseSSD", ctx.GetVar("UseSSD")))

	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.CheckUDBInstanceAllowance(req)
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

func testSet489DescribeUDBInstanceBackupURL14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewDescribeUDBInstanceBackupURLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("BackupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceBackupURL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("BackupPath", "http://", "contains"),
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

func testSet489CheckUDBInstanceAllowance15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "SQL"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))
	ctx.NoError(utest.SetReqValue(req, "UseSSD", ctx.GetVar("UseSSD")))

	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.CheckUDBInstanceAllowance(req)
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

func testSet489CreateUDBInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewCreateUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))
	ctx.NoError(utest.SetReqValue(req, "Name", "frombf-hassd-56"))

	ctx.NoError(utest.SetReqValue(req, "AdminPassword", "guanliyuanmima"))
	ctx.NoError(utest.SetReqValue(req, "DBTypeId", "mysql-5.6"))
	ctx.NoError(utest.SetReqValue(req, "Port", "3306"))
	ctx.NoError(utest.SetReqValue(req, "ParamGroupId", ctx.Must(utest.SearchValue(ctx.GetVar("DataSet_paramGroup"), "DBTypeId", "mysql-5.6", "GroupId"))))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))

	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("BackupId")))

	ctx.NoError(utest.SetReqValue(req, "InstanceMode", ctx.GetVar("InstanceMode")))

	ctx.NoError(utest.SetReqValue(req, "InstanceType", ctx.GetVar("InstanceType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.CreateUDBInstance(req)
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

	ctx.Vars["DBId_frombf_56"] = ctx.Must(utest.GetValue(resp, "DBId"))
}

func testSet489DescribeUDBInstanceState17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId_frombf_56")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("State", "Running", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489DescribeUDBInstance18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId_frombf_56")))
	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstanceResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 5 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489StopUDBInstance19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewStopUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId_frombf_56")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.StopUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopUDBInstanceResponse", "str_eq"),
		},
		MaxRetries:    2,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489DescribeUDBInstance20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId_frombf_56")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstanceResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Shutoff", "str_eq"),
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

func testSet489FetchUDBInstanceEarliestRecoverTime21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewFetchUDBInstanceEarliestRecoverTimeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.FetchUDBInstanceEarliestRecoverTime(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("EarliestTime", "0", "ne"),
		},
		MaxRetries:    20,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["EarliestTime"] = ctx.Must(utest.GetValue(resp, "EarliestTime"))
}

func testSet489CheckUDBInstanceAllowance22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.GetVar("MemoryLimit")))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.GetVar("DiskSpace")))
	ctx.NoError(utest.SetReqValue(req, "UseSSD", ctx.GetVar("UseSSD")))

	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.CheckUDBInstanceAllowance(req)
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

func testSet489CreateUDBInstanceByRecovery23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := udbClient.NewCreateUDBInstanceByRecoveryRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "56-hd-frombf"))
	ctx.NoError(utest.SetReqValue(req, "SrcDBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "RecoveryTime", ctx.GetVar("EarliestTime")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.CreateUDBInstanceByRecovery(req)
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

	ctx.Vars["HD_DBId"] = ctx.Must(utest.GetValue(resp, "DBId"))
}

func testSet489DescribeUDBInstanceState24(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("HD_DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("State", "Running", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489DescribeUDBInstance25(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("HD_DBId")))
	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstanceResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
		},
		MaxRetries:    100,
		RetryInterval: 5 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet489StopUDBInstance26(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewStopUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("HD_DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.StopUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopUDBInstanceResponse", "str_eq"),
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

func testSet489DescribeUDBInstance27(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("HD_DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstanceResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Shutoff", "str_eq"),
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

func testSet489DeleteUDBBackup28(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udbClient.NewDeleteUDBBackupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("BackupId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBBackup(req)
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

func testSet489StopUDBInstance29(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewStopUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.StopUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "StopUDBInstanceResponse", "str_eq"),
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

func testSet489DescribeUDBInstance30(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstanceResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Shutoff", "str_eq"),
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

func testSet489DeleteUDBInstance31(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId_frombf_56")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
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

func testSet489DeleteUDBInstance32(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("HD_DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
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

func testSet489DeleteUDBInstance33(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
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
