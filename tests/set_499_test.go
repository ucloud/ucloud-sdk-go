package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet499(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")
	ctx.SetVar("DBTypeId", "percona-5.5")
	ctx.SetVar("InstanceMode", "HA")
	ctx.SetVar("InstanceType", "SATA_SSD")
	ctx.SetVar("Port", "3306")
	ctx.SetVar("MemoryLimit", "1000")
	ctx.SetVar("DiskSpace", "20")
	ctx.SetVar("DBName", "auto_hassd_")
	ctx.SetVar("UseSSD", "true")

	testSet499DescribeUDBType00(&ctx)
	testSet499DescribeUDBParamGroup01(&ctx)
	testSet499DescribeUDBInstancePrice02(&ctx)
	testSet499CheckUDBInstanceAllowance03(&ctx)
	testSet499CreateUDBInstance04(&ctx)
	testSet499DescribeUDBInstance05(&ctx)
	testSet499DescribeUDBInstanceState06(&ctx)
	testSet499CreateUDBParamGroup07(&ctx)
	testSet499UpdateUDBParamGroup08(&ctx)
	testSet499UploadUDBParamGroup09(&ctx)
	testSet499ChangeUDBParamGroup10(&ctx)
	testSet499DescribeUDBInstance11(&ctx)
	testSet499RestartUDBInstance12(&ctx)
	testSet499DescribeUDBInstance13(&ctx)
	testSet499DeleteUDBParamGroup14(&ctx)
	testSet499StopUDBInstance15(&ctx)
	testSet499DescribeUDBInstance16(&ctx)
	testSet499DeleteUDBInstance17(&ctx)
	testSet499DeleteUDBParamGroup18(&ctx)
}

func testSet499DescribeUDBType00(ctx *utest.TestContext) {
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

func testSet499DescribeUDBParamGroup01(ctx *utest.TestContext) {
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

func testSet499DescribeUDBInstancePrice02(ctx *utest.TestContext) {
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

func testSet499CheckUDBInstanceAllowance03(ctx *utest.TestContext) {
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

func testSet499CreateUDBInstance04(ctx *utest.TestContext) {
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

func testSet499DescribeUDBInstance05(ctx *utest.TestContext) {
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

func testSet499DescribeUDBInstanceState06(ctx *utest.TestContext) {
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

func testSet499CreateUDBParamGroup07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewCreateUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "GroupName", "auto_config_create"))
	ctx.NoError(utest.SetReqValue(req, "Description", "self_create"))
	ctx.NoError(utest.SetReqValue(req, "SrcGroupId", ctx.Must(utest.SearchValue(ctx.GetVar("DataSet_paramGroup"), "DBTypeId", ctx.GetVar("DBTypeId"), "GroupId"))))
	ctx.NoError(utest.SetReqValue(req, "DBTypeId", ctx.GetVar("DBTypeId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.CreateUDBParamGroup(req)
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

	ctx.Vars["create_config_id"] = ctx.Must(utest.GetValue(resp, "GroupId"))
}

func testSet499UpdateUDBParamGroup08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewUpdateUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("create_config_id")))
	ctx.NoError(utest.SetReqValue(req, "Key", "auto_increment_increment"))
	ctx.NoError(utest.SetReqValue(req, "Value", "2"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.UpdateUDBParamGroup(req)
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

func testSet499UploadUDBParamGroup09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewUploadUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "DBTypeId", ctx.GetVar("DBTypeId")))
	ctx.NoError(utest.SetReqValue(req, "GroupName", "auto_config_upload"))
	ctx.NoError(utest.SetReqValue(req, "Description", "self_upload"))
	ctx.NoError(utest.SetReqValue(req, "Content", "W215c3FsZF0KYXV0b19pbmNyZW1lbnRfaW5jcmVtZW50ID0gMgpiYWNrX2xvZyA9IDIwMDAKYmlubG9nLWZvcm1hdCA9IE1JWEVECmNoYXJhY3Rlcl9zZXRfc2VydmVyID0gdXRmOApldmVudF9zY2hlZHVsZXIgPSBPTgpleHBpcmVfbG9nc19kYXlzID0gNwppbm5vZGJfYnVmZmVyX3Bvb2xfc2l6ZSA9IDUzNjg3MDkxMjAwCmlubm9kYl9maWxlX3Blcl90YWJsZSA9IDEKaW5ub2RiX2ZsdXNoX2xvZ19hdF90cnhfY29tbWl0ID0gMgppbm5vZGJfZmx1c2hfbWV0aG9kID0gT19ESVJFQ1QKaW5ub2RiX2lvX2NhcGFjaXR5ID0gMjAwMAppbm5vZGJfbG9nX2J1ZmZlcl9zaXplID0gODM4ODYwOAppbm5vZGJfbWF4X2RpcnR5X3BhZ2VzX3BjdCA9IDUwCmlubm9kYl9vcGVuX2ZpbGVzID0gMTAyNAppbm5vZGJfcmVhZF9pb190aHJlYWRzID0gOAppbm5vZGJfc29ydF9idWZmZXJfc2l6ZSA9IDEwNDg1NzYKaW5ub2RiX3RocmVhZF9jb25jdXJyZW5jeSA9IDIwCmlubm9kYl93cml0ZV9pb190aHJlYWRzID0gOAprZXlfYnVmZmVyX3NpemUgPSAzMzU1NDQzMgpsb2NhbF9pbmZpbGUgPSAxCmxvZ19iaW5fdHJ1c3RfZnVuY3Rpb25fY3JlYXRvcnMgPSAxCmxvbmdfcXVlcnlfdGltZSA9IDMKbWF4X2FsbG93ZWRfcGFja2V0ID0gMTY3NzcyMTYKbWF4X2Nvbm5lY3RfZXJyb3JzID0gMTAwMDAwMAptYXhfY29ubmVjdGlvbnMgPSAyMDAwCm15aXNhbV9zb3J0X2J1ZmZlcl9zaXplID0gODM4ODYwOApuZXRfYnVmZmVyX2xlbmd0aCA9IDgxOTIKcGVyZm9ybWFuY2Vfc2NoZW1hID0gMApwZXJmb3JtYW5jZV9zY2hlbWFfbWF4X3RhYmxlX2luc3RhbmNlcyA9IDIwMApxdWVyeV9jYWNoZV9zaXplID0gMApxdWVyeV9jYWNoZV90eXBlID0gMApyZWFkX2J1ZmZlcl9zaXplID0gMjYyMTQ0CnJlYWRfcm5kX2J1ZmZlcl9zaXplID0gNTI0Mjg4CnNraXBfbmFtZV9yZXNvbHZlID0gMQpzb3J0X2J1ZmZlcl9zaXplID0gNTI0Mjg4CnN5bmNfYmlubG9nID0gMQp0YWJsZV9vcGVuX2NhY2hlID0gMTI4CnRocmVhZF9jYWNoZV9zaXplID0gNTA="))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.UploadUDBParamGroup(req)
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

	ctx.Vars["upload_config_id"] = ctx.Must(utest.GetValue(resp, "GroupId"))
}

func testSet499ChangeUDBParamGroup10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iudbClient.NewChangeUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "DBId", ctx.GetVar("DBId")))
	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("upload_config_id")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iudbClient.ChangeUDBParamGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("Action", "ChangeUDBParamGroupResponse", "str_eq"),
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

func testSet499DescribeUDBInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("DataSet.0.ParamGroupId", ctx.GetVar("upload_config_id"), "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet499RestartUDBInstance12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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
			ctx.NewValidator("Action", "RestartUDBInstanceResponse", "str_eq"),
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

func testSet499DescribeUDBInstance13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

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

func testSet499DeleteUDBParamGroup14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udbClient.NewDeleteUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("upload_config_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBParamGroup(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "7064", "str_eq"),
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

func testSet499StopUDBInstance15(ctx *utest.TestContext) {
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

func testSet499DescribeUDBInstance16(ctx *utest.TestContext) {
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

func testSet499DeleteUDBInstance17(ctx *utest.TestContext) {
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

func testSet499DeleteUDBParamGroup18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udbClient.NewDeleteUDBParamGroupRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "GroupId", ctx.GetVar("create_config_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udbClient.DeleteUDBParamGroup(req)
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
