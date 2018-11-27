package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet488(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DBTypeId", "percona-5.6")
	ctx.SetVar("InstanceMode", "HA")
	ctx.SetVar("InstanceType", "Normal")
	ctx.SetVar("Port", "3306")
	ctx.SetVar("MemoryLimit", "1000")
	ctx.SetVar("DiskSpace", "20")
	ctx.SetVar("DBName", "auto_habz_")
	ctx.SetVar("UseSSD", "false")

	testSet488DescribeUDBInstance00(&ctx)
	testSet488DescribeUDBType01(&ctx)
	testSet488DescribeUDBParamGroup02(&ctx)
	testSet488DescribeUDBInstancePrice03(&ctx)
	testSet488CheckUDBInstanceAllowance04(&ctx)
	testSet488CreateUDBInstance05(&ctx)
	testSet488DescribeUDBInstance06(&ctx)
	testSet488DescribeUDBInstanceState07(&ctx)
	testSet488DescribeUDBInstancePhpMyAdminURL08(&ctx)
	testSet488ModifyUDBInstanceName09(&ctx)
	testSet488DescribeUDBInstance10(&ctx)
	testSet488ModifyUDBInstancePassword11(&ctx)
	testSet488RestartUDBInstance12(&ctx)
	testSet488DescribeUDBInstance13(&ctx)
	testSet488StopUDBInstance14(&ctx)
	testSet488DescribeUDBInstance15(&ctx)
	testSet488StartUDBInstance16(&ctx)
	testSet488DescribeUDBInstance17(&ctx)
	testSet488StopUDBInstance18(&ctx)
	testSet488DescribeUDBInstance19(&ctx)
	testSet488DeleteUDBInstance20(&ctx)
}

func testSet488DescribeUDBInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstance(req)
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

func testSet488DescribeUDBType01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

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

func testSet488DescribeUDBParamGroup02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

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

func testSet488DescribeUDBInstancePrice03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

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

func testSet488CheckUDBInstanceAllowance04(ctx *utest.TestContext) {
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

func testSet488CreateUDBInstance05(ctx *utest.TestContext) {
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

func testSet488DescribeUDBInstance06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
		MaxRetries:    60,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet488DescribeUDBInstanceState07(ctx *utest.TestContext) {
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

func testSet488DescribeUDBInstancePhpMyAdminURL08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewDescribeUDBInstancePhpMyAdminURLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.DescribeUDBInstancePhpMyAdminURL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeUDBInstancePhpMyAdminURLResponse", "str_eq"),
			ctx.NewValidator("PMAPath", "", "ne"),
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

func testSet488ModifyUDBInstanceName09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udbClient.NewModifyUDBInstanceNameRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat("rename_", ctx.GetVar("DBTypeId")))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.ModifyUDBInstanceName(req)
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

func testSet488DescribeUDBInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat("rename_", ctx.GetVar("DBTypeId"))), "str_eq"),
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
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

func testSet488ModifyUDBInstancePassword11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewModifyUDBInstancePasswordRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "Password", "2012_UClou"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.ModifyUDBInstancePassword(req)
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

func testSet488RestartUDBInstance12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udbClient.NewRestartUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.RestartUDBInstance(req)
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

func testSet488DescribeUDBInstance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
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

func testSet488StopUDBInstance14(ctx *utest.TestContext) {
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

func testSet488DescribeUDBInstance15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Shutoff", "str_eq"),
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

func testSet488StartUDBInstance16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewStartUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.StartUDBInstance(req)
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

func testSet488DescribeUDBInstance17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
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

func testSet488StopUDBInstance18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

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

func testSet488DescribeUDBInstance19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.DBTypeId", ctx.GetVar("DBTypeId"), "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Shutoff", "str_eq"),
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

func testSet488DeleteUDBInstance20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
