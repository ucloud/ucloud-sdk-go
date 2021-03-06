// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/ubill"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario3608(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "3608",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Password":           "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"ChargeType":         "Month",
				"CreateCPU":          1,
				"CreateMem":          1024,
				"Name":               "hot-local-linux",
				"ImageID":            "#{u_get_image_resource($Region,$Zone)}",
				"UpgradeCPU":         2,
				"BootSize":           20,
				"BootType":           "LOCAL_NORMAL",
				"DiskSize":           20,
				"DiskType":           "LOCAL_NORMAL",
				"BootBackup":         "NONE",
				"DiskBackup":         "NONE",
				"UHostType":          "N2",
				"MinimalCpuPlatform": "Intel/Auto",
				"MachineType":        "N",
				"Region":             "cn-bj2",
				"Zone":               "cn-bj2-02",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "Auto-N-LOCAL_NORMAL-LOCAL_NORMAL-获取热升级价格-linux",
		Steps: []*driver.Step{
			testStep3608DescribeImage01,
			testStep3608DescribeImage02,
			testStep3608CreateUHostInstance03,
			testStep3608DescribeUHostInstance04,
			testStep3608GetUHostUpgradePrice05,
			testStep3608UpgradeUHostInstance06,
			testStep3608DescribeUHostInstance07,
			testStep3608DescribeOrderDetailInfo08,
			testStep3608PoweroffUHostInstance09,
			testStep3608DescribeUHostInstance10,
			testStep3608TerminateUHostInstance11,
		},
	})
}

var testStep3608DescribeImage01 = &driver.Step{
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
			"Limit":     200,
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ImageSet", step.Must(utils.GetValue(resp, "ImageSet")))
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
	FastFail:      true,
}

var testStep3608DescribeImage02 = &driver.Step{
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
			"ImageId":   step.Must(functions.SearchValue(step.Scenario.GetVar("ImageSet"), "OsName", "CentOS 7.2 64位", "ImageId")),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ImageID", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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
	FastFail:      true,
}

var testStep3608CreateUHostInstance03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("Zone"),
			"Region":             step.Scenario.GetVar("Region"),
			"Quantity":           1,
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":               step.Scenario.GetVar("Name"),
			"MinimalCpuPlatform": step.Scenario.GetVar("MinimalCpuPlatform"),
			"Memory":             step.Scenario.GetVar("CreateMem"),
			"MachineType":        step.Scenario.GetVar("MachineType"),
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("ImageID"),
			"HotplugFeature":     "true",
			"Disks": []map[string]interface{}{
				{
					"BackupType": step.Scenario.GetVar("BootBackup"),
					"IsBoot":     "True",
					"Size":       step.Scenario.GetVar("BootSize"),
					"Type":       step.Scenario.GetVar("BootType"),
				},
				{
					"BackupType": step.Scenario.GetVar("DiskBackup"),
					"IsBoot":     "False",
					"Size":       step.Scenario.GetVar("DiskSize"),
					"Type":       step.Scenario.GetVar("DiskType"),
				},
			},
			"ChargeType": step.Scenario.GetVar("ChargeType"),
			"CPU":        step.Scenario.GetVar("CreateCPU"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("hostId", step.Must(utils.GetValue(resp, "UHostIds.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建云主机",
	FastFail:      true,
}

var testStep3608DescribeUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("CreateTime", step.Must(utils.GetValue(resp, "UHostSet.0.CreateTime")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.CPU", step.Scenario.GetVar("CreateCPU"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Memory", step.Scenario.GetVar("CreateMem"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("hostId"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Name", step.Scenario.GetVar("Name"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BasicImageId", step.Scenario.GetVar("ImageID"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.IsBoot", "True", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.IsBoot", "False", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.DiskType", step.Scenario.GetVar("BootType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.DiskType", step.Scenario.GetVar("DiskType"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.Size", step.Scenario.GetVar("BootSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.1.Size", step.Scenario.GetVar("DiskSize"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.MachineType", step.Scenario.GetVar("MachineType"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    200,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3608GetUHostUpgradePrice05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewGetUHostUpgradePriceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
			"Memory":  2048,
			"CPU":     2,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetUHostUpgradePrice(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "GetUHostUpgradePriceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取主机规格调整差价",
	FastFail:      true,
}

var testStep3608UpgradeUHostInstance06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("UpgradeUHostInstance")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
			"Memory":  2048,
			"CPU":     2,
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "UpgradeUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "在线更新UHost实例的资源配置",
	FastFail:      true,
}

var testStep3608DescribeUHostInstance07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("hostId"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Name", step.Scenario.GetVar("Name"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.CPU", 2, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.Memory", 2048, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    10,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3608DescribeOrderDetailInfo08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UBill")
		if err != nil {
			return nil, err
		}
		client := c.(*ubill.UBillClient)

		req := client.NewDescribeOrderDetailInfoRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ResourceIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"OrderTypes": []interface{}{
				"OT_UPGRADE",
			},
			"OrderStates": []interface{}{
				"OS_FINISHED",
			},
			"Offset":    0,
			"Limit":     25,
			"EndTime":   step.Must(functions.Calculate("+", step.Must(functions.GetTimestamp(10)), 100)),
			"BeginTime": step.Must(functions.Calculate("-", step.Scenario.GetVar("CreateTime"), 2592000)),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeOrderDetailInfo(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeOrderDetailInfoResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取订单信息",
	FastFail:      true,
}

var testStep3608PoweroffUHostInstance09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    5,
	RetryInterval: 30 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      true,
}

var testStep3608DescribeUHostInstance10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone": step.Scenario.GetVar("Zone"),
			"UHostIds": []interface{}{
				step.Scenario.GetVar("hostId"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Stopped", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep3608TerminateUHostInstance11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
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
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}
