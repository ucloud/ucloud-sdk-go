package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet181(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DBTypeId", "mongodb-3.0")
	ctx.SetVar("InstanceMode", "Normal")
	ctx.SetVar("InstanceType", "Normal")
	ctx.SetVar("Port", "27017")
	ctx.SetVar("MemoryLimit", "600")
	ctx.SetVar("DiskSpace", "20")
	ctx.SetVar("DBName", "AUTO")
	ctx.SetVar("UseSSD", "False")
	ctx.SetVar("GroupName", "mongodb3.0默认MMAPv1-shardsvr配置")

	testSet181DescribeUDBType00(&ctx)
	testSet181DescribeUDBParamGroup01(&ctx)
	testSet181CheckUDBInstanceAllowance02(&ctx)
	testSet181CreateUDBInstance03(&ctx)
	testSet181CreateUDBReplicationInstance04(&ctx)
	testSet181StopUDBInstance05(&ctx)
	testSet181StopUDBInstance06(&ctx)
	testSet181DeleteUDBInstance07(&ctx)
	testSet181DeleteUDBInstance08(&ctx)
}

func testSet181DescribeUDBType00(ctx *utest.TestContext) {
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

func testSet181DescribeUDBParamGroup01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

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
			ctx.NewValidator("Action", "DescribeUDBParamGroupResponse", "str_eq"),
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

func testSet181CheckUDBInstanceAllowance02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := iudbClient.NewCheckUDBInstanceAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ClassType", "sql"))
	ctx.NoError(utest.SetReqValue(req, "MemoryLimit", "6000"))
	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "200"))
	ctx.NoError(utest.SetReqValue(req, "UseSSD", "True"))

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

func testSet181CreateUDBInstance03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := udbClient.NewCreateUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat(ctx.GetVar("DBName"), ctx.GetVar("DBTypeId")))))
	ctx.NoError(utest.SetReqValue(req, "AdminUser", "root"))
	ctx.NoError(utest.SetReqValue(req, "AdminPassword", "guanliyuanmima"))
	ctx.NoError(utest.SetReqValue(req, "DBTypeId", ctx.GetVar("DBTypeId")))
	ctx.NoError(utest.SetReqValue(req, "Port", ctx.GetVar("Port")))
	ctx.NoError(utest.SetReqValue(req, "ParamGroupId", ctx.Must(utest.SearchValue(ctx.GetVar("DataSet_paramGroup"), "GroupName", ctx.GetVar("GroupName"), "GroupId"))))
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
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["DBId"] = ctx.Must(utest.GetValue(resp, "DBId"))
}

func testSet181CreateUDBReplicationInstance04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := udbClient.NewCreateUDBReplicationInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SrcId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "Name", "second-auto"))
	ctx.NoError(utest.SetReqValue(req, "Port", ctx.GetVar("Port")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.CreateUDBReplicationInstance(req)
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

	ctx.Vars["secondid"] = ctx.Must(utest.GetValue(resp, "DBId"))
}

func testSet181StopUDBInstance05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

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

func testSet181StopUDBInstance06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := udbClient.NewStopUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("secondid")))

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

func testSet181DeleteUDBInstance07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("secondid")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet181DeleteUDBInstance08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := udbClient.NewDeleteUDBInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
