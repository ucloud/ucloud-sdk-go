// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"

	"github.com/ucloud/ucloud-sdk-go/services/unet"
)

func TestScenario405(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "405",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Image_Id": "#{u_get_image_resource($Region,$Zone)}",
				"Region":   "tw-tp",
				"Zone":     "tw-tp-01",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "eip自动化回归-基本操作-01-国际线路",
		Steps: []*driver.Step{
			testStep405DescribeImage01,
			testStep405CreateUHostInstance02,
			testStep405AllocateEIP03,
			testStep405DescribeEIP04,
			testStep405GetEIPPrice05,
			testStep405UpdateEIPAttribute06,
			testStep405BindEIP07,
			testStep405ModifyEIPBandwidth08,
			testStep405DescribeEIP09,
			testStep405SetEIPPayMode10,
			testStep405ModifyEIPWeight11,
			testStep405UnBindEIP12,
			testStep405ReleaseEIP13,
			testStep405DescribeEIP14,
			testStep405PoweroffUHostInstance15,
			testStep405TerminateUHostInstance16,
		},
	})
}

var testStep405DescribeImage01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"Region":    step.Scenario.GetVar("Region"),
			"OsType":    "Linux",
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("Image_Id", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      false,
}

var testStep405CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"Tag":         "Default",
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "eip-s1-bgp",
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id"),
			"Disks": []map[string]interface{}{
				{
					"IsBoot": "True",
					"Size":   20,
					"Type":   "LOCAL_NORMAL",
				},
			},
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("UHostId", step.Must(utils.GetValue(resp, "UHostIds.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostIds", 0, "len_ge"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep405AllocateEIP03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewAllocateEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag":          "Default",
			"Remark":       "International",
			"Region":       step.Scenario.GetVar("Region"),
			"PayMode":      "Bandwidth",
			"OperatorName": "International",
			"Name":         "eip-International-01",
			"Bandwidth":    2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("EIPId_01", step.Must(utils.GetValue(resp, "EIPSet.0.EIPId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "申请弹性IP",
	FastFail:      false,
}

var testStep405DescribeEIP04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId_01"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("EIPSet.0.EIPId", step.Scenario.GetVar("EIPId_01"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取弹性IP信息",
	FastFail:      false,
}

var testStep405GetEIPPrice05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewGetEIPPriceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":       step.Scenario.GetVar("Region"),
			"OperatorName": "International",
			"Bandwidth":    2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetEIPPrice(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取弹性IP价格",
	FastFail:      false,
}

var testStep405UpdateEIPAttribute06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewUpdateEIPAttributeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Tag":    "huangchao",
			"Remark": "test-gai",
			"Region": step.Scenario.GetVar("Region"),
			"Name":   "eip-International-gai",
			"EIPId":  step.Scenario.GetVar("EIPId_01"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateEIPAttribute(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "更新弹性IP属性",
	FastFail:      false,
}

var testStep405BindEIP07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewBindEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceType": "uhost",
			"ResourceId":   step.Scenario.GetVar("UHostId"),
			"Region":       step.Scenario.GetVar("Region"),
			"EIPId":        step.Scenario.GetVar("EIPId_01"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.BindEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "绑定弹性IP",
	FastFail:      false,
}

var testStep405ModifyEIPBandwidth08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewModifyEIPBandwidthRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":    step.Scenario.GetVar("Region"),
			"EIPId":     step.Scenario.GetVar("EIPId_01"),
			"Bandwidth": 3,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ModifyEIPBandwidth(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "调整弹性IP带宽",
	FastFail:      false,
}

var testStep405DescribeEIP09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId_01"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("EIPSet.0.Resource.ResourceID", step.Scenario.GetVar("UHostId"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取弹性IP信息",
	FastFail:      false,
}

var testStep405SetEIPPayMode10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewSetEIPPayModeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":    step.Scenario.GetVar("Region"),
			"PayMode":   "Traffic",
			"EIPId":     step.Scenario.GetVar("EIPId_01"),
			"Bandwidth": 2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.SetEIPPayMode(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "设置弹性IP计费方式",
	FastFail:      false,
}

var testStep405ModifyEIPWeight11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewModifyEIPWeightRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Weight": 100,
			"Region": step.Scenario.GetVar("Region"),
			"EIPId":  step.Scenario.GetVar("EIPId_01"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ModifyEIPWeight(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "修改弹性IP出口权重",
	FastFail:      false,
}

var testStep405UnBindEIP12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewUnBindEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceType": "uhost",
			"ResourceId":   step.Scenario.GetVar("UHostId"),
			"Region":       step.Scenario.GetVar("Region"),
			"EIPId":        step.Scenario.GetVar("EIPId_01"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UnBindEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "解绑弹性IP",
	FastFail:      false,
}

var testStep405ReleaseEIP13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewReleaseEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPId":  step.Scenario.GetVar("EIPId_01"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ReleaseEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 8039, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "释放弹性IP",
	FastFail:      false,
}

var testStep405DescribeEIP14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPIds": []interface{}{
				step.Scenario.GetVar("EIPId_01"),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeEIP(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("TotalCount", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取弹性IP信息",
	FastFail:      false,
}

var testStep405PoweroffUHostInstance15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.PoweroffUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep405TerminateUHostInstance16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}
