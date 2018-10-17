package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet302(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("tag", "tag_api_test")
	ctx.SetVar("remark", "remark_api_test")
	ctx.SetVar("fw_name_1", "fw_name_A_")
	ctx.SetVar("Region", "cn-bj2")
	ctx.SetVar("fw_rule_1_protocol", "TCP")
	ctx.SetVar("fw_rule_1_port", "1111")
	ctx.SetVar("fw_rule_1_srcip", "0.0.0.0/0")
	ctx.SetVar("fw_rule_1_action", "ACCEPT")
	ctx.SetVar("fw_rule_1_priority", "HIGH")
	ctx.SetVar("Zone", "cn-bj2-05")
	ctx.SetVar("uhost_name_1", "firewall_api_test")
	ctx.SetVar("fw_rule_2_protocol", "UDP")
	ctx.SetVar("fw_rule_2_port", "2222")
	ctx.SetVar("fw_rule_2_srcip", "10.0.0.0/8")
	ctx.SetVar("fw_rule_2_action", "DROP")
	ctx.SetVar("fw_rule_2_priority", "LOW")
	ctx.SetVar("fw_name_2", "fw_name_2_")
	ctx.SetVar("tag_2", "tag_api_test_2")
	ctx.SetVar("remark_2", "remark_api_test_2")
	ctx.SetVar("Image_Id", ctx.Must(utest.GetImageResource(ctx.GetVar("Region"), ctx.GetVar("Zone"))))
	ctx.SetVar("recommend_web", "recommend web")

	testSet302CreateUHostInstance00(&ctx)
	testSet302CreateFirewall01(&ctx)
	testSet302DescribeFirewall02(&ctx)
	testSet302GrantFirewall03(&ctx)
	testSet302DescribeFirewallResource04(&ctx)
	testSet302UpdateFirewall05(&ctx)
	testSet302UpdateFirewallAttribute06(&ctx)
	testSet302DescribeFirewall07(&ctx)
	testSet302DescribeFirewall08(&ctx)
	testSet302GrantFirewall09(&ctx)
	testSet302PoweroffUHostInstance10(&ctx)
	testSet302TerminateUHostInstance11(&ctx)
	testSet302DescribeFirewall12(&ctx)
	testSet302DeleteFirewall13(&ctx)
	testSet302DescribeFirewall14(&ctx)
}

func testSet302CreateUHostInstance00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := uhostClient.NewCreateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "ImageId", ctx.GetVar("Image_Id")))
	ctx.NoError(utest.SetReqValue(req, "LoginMode", "Password"))
	ctx.NoError(utest.SetReqValue(req, "Password", "dWNsb3VkLmNuMA=="))

	ctx.NoError(utest.SetReqValue(req, "CPU", "1"))
	ctx.NoError(utest.SetReqValue(req, "Memory", "1024"))

	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("uhost_name_1")))

	ctx.NoError(utest.SetReqValue(req, "DiskSpace", "0"))

	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("tag")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.CreateUHostInstance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 10 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["uhost_id1"] = ctx.Must(utest.GetValue(resp, "UHostIds.0"))
}

func testSet302CreateFirewall01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(180) * time.Second)

	req := unetClient.NewCreateFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat(ctx.GetVar("fw_name_1"), ctx.GetVar("Zone")))))
	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("tag")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("remark")))
	ctx.NoError(utest.SetReqValue(req, "Rule", ctx.Must(utest.ConcatWithVertical(ctx.GetVar("fw_rule_1_protocol"), ctx.GetVar("fw_rule_1_port"), ctx.GetVar("fw_rule_1_srcip"), ctx.GetVar("fw_rule_1_action"), ctx.GetVar("fw_rule_1_priority")))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.CreateFirewall(req)
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

	ctx.Vars["fw_id1"] = ctx.Must(utest.GetValue(resp, "FWId"))
}

func testSet302DescribeFirewall02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewDescribeFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewall(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.FWId", ctx.GetVar("fw_id1"), "str_eq"),
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat(ctx.GetVar("fw_name_1"), ctx.GetVar("Zone"))), "str_eq"),
			ctx.NewValidator("DataSet.0.Tag", ctx.GetVar("tag"), "str_eq"),
			ctx.NewValidator("DataSet.0.Remark", ctx.GetVar("remark"), "str_eq"),
			ctx.NewValidator("DataSet.0.ResourceCount", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.Type", "user defined", "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.ProtocolType", ctx.GetVar("fw_rule_1_protocol"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.DstPort", ctx.GetVar("fw_rule_1_port"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.SrcIP", ctx.GetVar("fw_rule_1_srcip"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.RuleAction", ctx.GetVar("fw_rule_1_action"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.Priority", ctx.GetVar("fw_rule_1_priority"), "str_eq"),
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

func testSet302GrantFirewall03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewGrantFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "UHost"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("uhost_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GrantFirewall(req)
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

func testSet302DescribeFirewallResource04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewDescribeFirewallResourceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewallResource(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("ResourceSet.0.Name", ctx.GetVar("uhost_name_1"), "str_eq"),
			ctx.NewValidator("ResourceSet.0.ResourceType", "uhost", "str_eq"),
			ctx.NewValidator("ResourceSet.0.ResourceID", ctx.GetVar("uhost_id1"), "str_eq"),
			ctx.NewValidator("TotalCount", "1", "str_eq"),
			ctx.NewValidator("ResourceSet.0.Tag", ctx.GetVar("tag"), "str_eq"),
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

func testSet302UpdateFirewall05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewUpdateFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))
	ctx.NoError(utest.SetReqValue(req, "Rule", ctx.Must(utest.ConcatWithVertical(ctx.GetVar("fw_rule_2_protocol"), ctx.GetVar("fw_rule_2_port"), ctx.GetVar("fw_rule_2_srcip"), ctx.GetVar("fw_rule_2_action"), ctx.GetVar("fw_rule_2_priority")))))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.UpdateFirewall(req)
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

func testSet302UpdateFirewallAttribute06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewUpdateFirewallAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))
	ctx.NoError(utest.SetReqValue(req, "Name", ctx.Must(utest.Concat(ctx.GetVar("fw_name_2"), ctx.GetVar("Zone")))))
	ctx.NoError(utest.SetReqValue(req, "Tag", ctx.GetVar("tag_2")))
	ctx.NoError(utest.SetReqValue(req, "Remark", ctx.GetVar("remark_2")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.UpdateFirewallAttribute(req)
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

func testSet302DescribeFirewall07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(3) * time.Second)

	req := unetClient.NewDescribeFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewall(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.FWId", ctx.GetVar("fw_id1"), "str_eq"),
			ctx.NewValidator("DataSet.0.Name", ctx.Must(utest.Concat(ctx.GetVar("fw_name_2"), ctx.GetVar("Zone"))), "str_eq"),
			ctx.NewValidator("DataSet.0.Tag", ctx.GetVar("tag_2"), "str_eq"),
			ctx.NewValidator("DataSet.0.Remark", ctx.GetVar("remark_2"), "str_eq"),
			ctx.NewValidator("DataSet.0.ResourceCount", "1", "str_eq"),
			ctx.NewValidator("DataSet.0.Type", "user defined", "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.ProtocolType", ctx.GetVar("fw_rule_2_protocol"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.DstPort", ctx.GetVar("fw_rule_2_port"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.SrcIP", ctx.GetVar("fw_rule_2_srcip"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.RuleAction", ctx.GetVar("fw_rule_2_action"), "str_eq"),
			ctx.NewValidator("DataSet.0.Rule.0.Priority", ctx.GetVar("fw_rule_2_priority"), "str_eq"),
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

func testSet302DescribeFirewall08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDescribeFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Limit", "20"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewall(req)
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

	ctx.Vars["fw_dataset"] = ctx.Must(utest.GetValue(resp, "DataSet"))
}

func testSet302GrantFirewall09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewGrantFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.Must(utest.SearchValue(ctx.GetVar("fw_dataset"), "Type", ctx.GetVar("recommend_web"), "FWId"))))
	ctx.NoError(utest.SetReqValue(req, "ResourceType", "UHost"))
	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("uhost_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.GrantFirewall(req)
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

func testSet302PoweroffUHostInstance10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(60) * time.Second)

	req := uhostClient.NewPoweroffUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.PoweroffUHostInstance(req)
		},
		Validators:    []utest.TestValidator{},
		MaxRetries:    3,
		RetryInterval: 60 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet302TerminateUHostInstance11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(90) * time.Second)

	req := uhostClient.NewTerminateUHostInstanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))
	ctx.NoError(utest.SetReqValue(req, "UHostId", ctx.GetVar("uhost_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return uhostClient.TerminateUHostInstance(req)
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

func testSet302DescribeFirewall12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewDescribeFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewall(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.ResourceCount", "0", "str_eq"),
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

func testSet302DeleteFirewall13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := unetClient.NewDeleteFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))
	ctx.NoError(utest.SetReqValue(req, "FWId", ctx.GetVar("fw_id1")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DeleteFirewall(req)
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

func testSet302DescribeFirewall14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewDescribeFirewallRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Limit", "10"))
	ctx.NoError(utest.SetReqValue(req, "Offset", "0"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.DescribeFirewall(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet", ctx.GetVar("fw_id1"), "object_not_contains"),
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
