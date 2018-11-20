package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet1201(t *testing.T) {
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

	testSet1201DescribeURedisBackup00(&ctx)
	testSet1201CreateURedisGroup01(&ctx)
	testSet1201DescribeUMem02(&ctx)
	testSet1201ModifyURedisGroupName03(&ctx)
	testSet1201DescribeUMem04(&ctx)
	testSet1201ResizeURedisGroup05(&ctx)
	testSet1201DescribeUMem06(&ctx)
	testSet1201ModifyURedisGroupPassword07(&ctx)
	testSet1201DescribeUMem08(&ctx)
	testSet1201CheckURedisAllowance09(&ctx)
	testSet1201DeleteURedisGroup10(&ctx)
	testSet1201DescribeURedisGroup11(&ctx)
}

func testSet1201DescribeURedisBackup00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := umemClient.NewDescribeURedisBackupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.DescribeURedisBackup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "DescribeURedisBackupResponse", "str_eq"),
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

func testSet1201CreateURedisGroup01(ctx *utest.TestContext) {
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

func testSet1201DescribeUMem02(ctx *utest.TestContext) {
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

func testSet1201ModifyURedisGroupName03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := umemClient.NewModifyURedisGroupNameRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Name", "Redis_Change"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ModifyURedisGroupName(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ModifyURedisGroupNameResponse", "str_eq"),
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

func testSet1201DescribeUMem04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

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

func testSet1201ResizeURedisGroup05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := umemClient.NewResizeURedisGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2"))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return umemClient.ResizeURedisGroup(req)
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

func testSet1201DescribeUMem06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

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

func testSet1201ModifyURedisGroupPassword07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := iumemClient.NewModifyURedisGroupPasswordRequest()

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("group_id")))

	ctx.NoError(utest.SetReqValue(req, "Password", ctx.GetVar("Password")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", ctx.GetVar("ResourceType")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iumemClient.ModifyURedisGroupPassword(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet1201DescribeUMem08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(300) * time.Second)

	req := iumemClient.NewDescribeUMemRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))
	ctx.NoError(utest.SetReqValue(req, "Limit", "1000"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("group_id")))
	ctx.NoError(utest.SetReqValue(req, "Protocol", ctx.GetVar("Protocol")))

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

func testSet1201CheckURedisAllowance09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

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

func testSet1201DeleteURedisGroup10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(50) * time.Second)

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

func testSet1201DescribeURedisGroup11(ctx *utest.TestContext) {
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
