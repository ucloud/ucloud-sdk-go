package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet113(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("Zone", "cn-bj2-03")

	testSet113CreateUFSVolume00(&ctx)
	testSet113DescribeUFSVolume01(&ctx)
	testSet113DescribeUHostLite02(&ctx)
	testSet113DescribeUFSVolume03(&ctx)
	testSet113ClearUFSVolumeWhiteList04(&ctx)
	testSet113DescribeUFSVolume05(&ctx)
	testSet113ExtendUFSVolume06(&ctx)
	testSet113DescribeUFSVolume07(&ctx)
	testSet113RemoveUFSVolume08(&ctx)
}

func testSet113CreateUFSVolume00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := pufsClient.NewCreateUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Size", "1024"))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	//Manually add
	ctx.NoError(utest.SetReqValue(req, "StorageType", "Basic"))
	ctx.NoError(utest.SetReqValue(req, "ProtocolType", "NFSv3"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.CreateUFSVolume(req)
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

		ctx.T.Error(err)

	}

	ctx.Vars["Volume_Id"] = ctx.Must(utest.GetValue(resp, "VolumeId"))
	ctx.Vars["Volume_Name"] = ctx.Must(utest.GetValue(resp, "VolumeName"))
}

func testSet113DescribeUFSVolume01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pufsClient.NewDescribeUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VolumeId", ctx.GetVar("Volume_Id")))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.DescribeUFSVolume(req)
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

		ctx.T.Error(err)

	}

}

func testSet113DescribeUHostLite02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iuhostClient.NewDescribeUHostLiteRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iuhostClient.DescribeUHostLite(req)
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

		ctx.T.Error(err)

	}

}

func testSet113DescribeUFSVolume03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pufsClient.NewDescribeUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VolumeId", "Volume_Id"))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.DescribeUFSVolume(req)
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

		ctx.T.Error(err)

	}

}

func testSet113ClearUFSVolumeWhiteList04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iufsClient.NewClearUFSVolumeWhiteListRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VolumeId", ctx.GetVar("Volume_Id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iufsClient.ClearUFSVolumeWhiteList(req)
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

		ctx.T.Error(err)

	}

}

func testSet113DescribeUFSVolume05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pufsClient.NewDescribeUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VolumeId", "Volume_Id"))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.DescribeUFSVolume(req)
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

		ctx.T.Error(err)

	}

}

func testSet113ExtendUFSVolume06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pufsClient.NewExtendUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VolumeId", ctx.GetVar("Volume_Id")))
	ctx.NoError(utest.SetReqValue(req, "Size", "2048"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.ExtendUFSVolume(req)
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

		ctx.T.Error(err)

	}

}

func testSet113DescribeUFSVolume07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pufsClient.NewDescribeUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "VolumeId", "Volume_Id"))

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.DescribeUFSVolume(req)
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

		ctx.T.Error(err)

	}

}

func testSet113RemoveUFSVolume08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := pufsClient.NewRemoveUFSVolumeRequest()

	ctx.NoError(utest.SetReqValue(req, "VolumeId", ctx.GetVar("Volume_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pufsClient.RemoveUFSVolume(req)
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

		ctx.T.Error(err)

	}

}
