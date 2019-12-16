package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet286(t *testing.T) {
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
	ctx.SetVar("Name", "auto_udisk_nofz")

	testSet286DescribeUDiskPrice00(&ctx)
	testSet286CheckUDiskAllowance01(&ctx)
	testSet286CreateUDisk02(&ctx)
	testSet286DescribeUDisk03(&ctx)
	testSet286CloneUDisk04(&ctx)
	testSet286DescribeUDisk05(&ctx)
	testSet286DescribeUDiskUpgradePrice06(&ctx)
	testSet286ResizeUDisk07(&ctx)
	testSet286DescribeUDisk08(&ctx)
	testSet286DeleteUDisk09(&ctx)
	testSet286DescribeUDisk10(&ctx)
	testSet286DeleteUDisk11(&ctx)
}

func testSet286DescribeUDiskPrice00(ctx *utest.TestContext) {
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

func testSet286CheckUDiskAllowance01(ctx *utest.TestContext) {
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

func testSet286CreateUDisk02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := udiskClient.NewCreateUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("Name")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", ctx.GetVar("UDataArkMode")))
	ctx.NoError(utest.SetReqValue(req, "Tag", "test"))
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

func testSet286DescribeUDisk03(ctx *utest.TestContext) {
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
			ctx.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			ctx.NewValidator("DataSet.0.Tag", "test", "str_eq"),
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

func testSet286CloneUDisk04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewCloneUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Name", "clone_nofz"))
	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("udisk_nofz_id")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))
	ctx.NoError(utest.SetReqValue(req, "Quantity", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CloneUDisk(req)
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

	ctx.Vars["clone_udisk_id_nofz"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet286DescribeUDisk05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_udisk_id_nofz")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
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

func testSet286DescribeUDiskUpgradePrice06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.Must(utest.Calculate("+", ctx.GetVar("Size"), "1"))))
	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("udisk_nofz_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDiskUpgradePrice(req)
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

func testSet286ResizeUDisk07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udiskClient.NewResizeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_nofz_id")))
	ctx.NoError(utest.SetReqValue(req, "Size", ctx.Must(utest.Calculate("+", ctx.GetVar("Size"), "1"))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.ResizeUDisk(req)
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

func testSet286DescribeUDisk08(ctx *utest.TestContext) {
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
			ctx.NewValidator("DataSet.0.Size", ctx.Must(utest.Calculate("+", ctx.GetVar("Size"), "1")), "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet286DeleteUDisk09(ctx *utest.TestContext) {
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

func testSet286DescribeUDisk10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_udisk_id_nofz")))

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

func testSet286DeleteUDisk11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_udisk_id_nofz")))

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
