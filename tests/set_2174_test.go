package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2174(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("ConfigId", "03f58ca9-b64d-4bdd-abc7-c6b9a46fd801")
	ctx.SetVar("Password", "Z3VhbmxpeXVhbm1pbWE=")
	ctx.SetVar("HighAvailability", "disable")
	ctx.SetVar("Version", "3.2")
	ctx.SetVar("Protocol", "redis")
	ctx.SetVar("ResourceType", "single")
	ctx.SetVar("Name", "single_redis")

	testSet2174CreateURedisGroup00(&ctx)
	testSet2174DescribeUMem01(&ctx)
	testSet2174CheckURedisAllowance02(&ctx)
	testSet2174CreateURedisBackup03(&ctx)
	testSet2174DescribeUMem04(&ctx)
	testSet2174DescribeURedisBackupState05(&ctx)
	testSet2174DescribeURedisBackupURL06(&ctx)
	testSet2174DeleteURedisGroup07(&ctx)
	testSet2174DescribeURedisGroup08(&ctx)
}

func testSet2174CreateURedisGroup00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewCreateURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))
	ctx.NoError(utest.SetReqValue(req, "AutoBackup", "enable"))
	ctx.NoError(utest.SetReqValue(req, "BackupTime", "3"))
	ctx.NoError(utest.SetReqValue(req, "ConfigId", ctx.GetVar("ConfigId")))
	ctx.NoError(utest.SetReqValue(req, "HighAvailability", ctx.GetVar("HighAvailability")))
	ctx.NoError(utest.SetReqValue(req, "Version", ctx.GetVar("Version")))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "1"))

	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.CreateURedisGroup(req)
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

	ctx.Vars["group_id"] = ctx.Must(utest.GetValue(resp, "GroupId"))
}

func testSet2174DescribeUMem01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

	ctx.NoError(utest.SetReqValue(req, "ResourceType", ctx.GetVar("ResourceType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
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

func testSet2174CheckURedisAllowance02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iumemClient.NewCheckURedisAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", "1"))
	ctx.NoError(utest.SetReqValue(req, "Count", "1"))

	ctx.NoError(utest.SetReqValue(req, "SlaveZone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.CheckURedisAllowance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CheckURedisAllowanceResponse", "str_eq"),
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

func testSet2174CreateURedisBackup03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumemClient.NewCreateURedisBackupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "BackupName", "backup_Redis"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.CreateURedisBackup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "CreateURedisBackupResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["backup_id"] = ctx.Must(utest.GetValue(resp, "BackupId"))
}

func testSet2174DescribeUMem04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

	ctx.NoError(utest.SetReqValue(req, "ResourceType", ctx.GetVar("ResourceType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeUMem(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.State", "Running", "str_eq"),
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

func testSet2174DescribeURedisBackupState05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := iumemClient.NewDescribeURedisBackupStateRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("backup_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.DescribeURedisBackupState(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeURedisBackupStateResponse", "str_eq"),
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

func testSet2174DescribeURedisBackupURL06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := umemClient.NewDescribeURedisBackupURLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "BackupId", ctx.GetVar("backup_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeURedisBackupURL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeURedisBackupURLResponse", "str_eq"),
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

func testSet2174DeleteURedisGroup07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := umemClient.NewDeleteURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DeleteURedisGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DeleteURedisGroupResponse", "str_eq"),
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

func testSet2174DescribeURedisGroup08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "100"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeURedisGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeURedisGroupResponse", "str_eq"),
			ctx.NewValidator("DataSet", ctx.GetVar("group_id"), "object_not_contains"),
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
