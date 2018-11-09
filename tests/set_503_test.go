package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet503(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DBTypeId", "mysql-5.7")
	ctx.SetVar("InstanceMode", "HA")
	ctx.SetVar("InstanceType", "Normal")
	ctx.SetVar("Port", "3306")
	ctx.SetVar("MemoryLimit", "1000")
	ctx.SetVar("DiskSpace", "20")
	ctx.SetVar("DBName", "auto_habz_")
	ctx.SetVar("UseSSD", "false")

	testSet503DescribeUDBType00(&ctx)
	testSet503DescribeUDBParamGroup01(&ctx)
	testSet503DescribeUDBInstancePrice02(&ctx)
	testSet503CheckUDBInstanceAllowance03(&ctx)
	testSet503CreateUDBInstance04(&ctx)
	testSet503DescribeUDBInstance05(&ctx)
	testSet503DescribeUDBInstanceState06(&ctx)
	testSet503DescribeUDBInstanceUpgradePrice07(&ctx)
	testSet503CheckUDBInstanceAllowance08(&ctx)
	testSet503ResizeUDBInstance09(&ctx)
	testSet503DescribeUDBInstance10(&ctx)
	testSet503StopUDBInstance11(&ctx)
	testSet503DescribeUDBInstance12(&ctx)
	testSet503DeleteUDBInstance13(&ctx)
}

func testSet503DescribeUDBType00(ctx *utest.TestContext) {
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

func testSet503DescribeUDBParamGroup01(ctx *utest.TestContext) {
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

func testSet503DescribeUDBInstancePrice02(ctx *utest.TestContext) {
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

func testSet503CheckUDBInstanceAllowance03(ctx *utest.TestContext) {
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

func testSet503CreateUDBInstance04(ctx *utest.TestContext) {
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

func testSet503DescribeUDBInstance05(ctx *utest.TestContext) {
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

func testSet503DescribeUDBInstanceState06(ctx *utest.TestContext) {
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

func testSet503DescribeUDBInstanceUpgradePrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDescribeUDBInstanceUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.Must(utest.Calculate("+", ctx.GetVar("MemoryLimit"), "1"))))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.Must(utest.Calculate("+", ctx.GetVar("DiskSpace"), "1"))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DescribeUDBInstanceUpgradePrice(req)
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

func testSet503CheckUDBInstanceAllowance08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "SQL"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.Must(utest.Calculate("+", ctx.GetVar("MemoryLimit"), "1"))))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.Must(utest.Calculate("+", ctx.GetVar("DiskSpace"), "1"))))
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

func testSet503ResizeUDBInstance09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := udbClient.NewResizeUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", ctx.Must(utest.Calculate("+", ctx.GetVar("MemoryLimit"), "1"))))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", ctx.Must(utest.Calculate("+", ctx.GetVar("DiskSpace"), "1"))))

	ctx.NoError(utest.SetReqValue(req, "UseSSD", ctx.GetVar("UseSSD")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.ResizeUDBInstance(req)
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

func testSet503DescribeUDBInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(240) * time.Second)

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
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
			ctx.NewValidator("DataSet.0.MemoryLimit", ctx.Must(utest.Calculate("+", ctx.GetVar("MemoryLimit"), "1")), "str_eq"),
			ctx.NewValidator("DataSet.0.DiskSpace", ctx.Must(utest.Calculate("+", ctx.GetVar("DiskSpace"), "10")), "str_eq"),
		},
		MaxRetries:    90,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet503StopUDBInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

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

func testSet503DescribeUDBInstance12(ctx *utest.TestContext) {
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

func testSet503DeleteUDBInstance13(ctx *utest.TestContext) {
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
