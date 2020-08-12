// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario523(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "523",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"DiskType":     "DataDisk",
				"Size":         1,
				"UDataArkMode": "Yes",
				"Name":         "auto_udisk_fz",
				"Region":       "cn-bj2",
				"Zone":         "cn-bj2-02",
			}
		},
		Owners: []string{"chenoa.chen@ucloud.cn"},
		Title:  "UDisk-普通方舟盘_02",
		Steps: []*driver.Step{
			testStep523DescribeUDiskPrice01,
			testStep523CheckUDiskAllowance02,
			testStep523CreateUDisk03,
			testStep523DescribeUDisk04,
			testStep523CloneUDisk05,
			testStep523DescribeUDiskUpgradePrice06,
			testStep523ResizeUDisk07,
			testStep523DescribeUDisk08,
			testStep523DeleteUDisk09,
			testStep523DescribeUDisk10,
			testStep523DeleteUDisk11,
		},
	})
}

var testStep523DescribeUDiskPrice01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskPriceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": step.Scenario.GetVar("UDataArkMode"),
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     1,
			"DiskType":     step.Scenario.GetVar("DiskType"),
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDiskPrice(req)
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
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘价格",
	FastFail:      false,
}

var testStep523CheckUDiskAllowance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("CheckUDiskAllowance")
		err = req.SetPayload(map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"Size":     10,
			"Region":   step.Scenario.GetVar("Region"),
			"DiskType": step.Scenario.GetVar("DiskType"),
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
			validation.Builtins.NewValidator("Action", "CheckUDiskAllowanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "检查UDisk资源余量",
	FastFail:      false,
}

var testStep523CreateUDisk03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCreateUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDataArkMode": step.Scenario.GetVar("UDataArkMode"),
			"Tag":          "test",
			"Size":         step.Scenario.GetVar("Size"),
			"Region":       step.Scenario.GetVar("Region"),
			"Quantity":     0,
			"Name":         step.Scenario.GetVar("Name"),
			"DiskType":     step.Scenario.GetVar("DiskType"),
			"ChargeType":   "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUDisk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("udisk_id", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(1) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "创建云硬盘",
	FastFail:      false,
}

var testStep523DescribeUDisk04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"Offset":  0,
			"Limit":   100,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Tag", "test", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep523CloneUDisk05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCloneUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"SourceId":   step.Scenario.GetVar("udisk_id"),
			"Region":     step.Scenario.GetVar("Region"),
			"Quantity":   0,
			"Name":       "clone_fz",
			"ChargeType": "Month",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CloneUDisk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("clone_udisk_id", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 2 * time.Second,
	Title:         "克隆云硬盘",
	FastFail:      false,
}

var testStep523DescribeUDiskUpgradePrice06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskUpgradePriceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":     step.Scenario.GetVar("Zone"),
			"SourceId": step.Scenario.GetVar("udisk_id"),
			"Size":     step.Must(functions.Calculate("+", step.Scenario.GetVar("Size"), 1)),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDiskUpgradePrice(req)
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘升级价格",
	FastFail:      false,
}

var testStep523ResizeUDisk07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewResizeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Size":    step.Must(functions.Calculate("+", step.Scenario.GetVar("Size"), 1)),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ResizeUDisk(req)
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
	StartupDelay:  time.Duration(80) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "调整云硬盘",
	FastFail:      false,
}

var testStep523DescribeUDisk08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Size", step.Must(functions.Calculate("+", step.Scenario.GetVar("Size"), 1)), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep523DeleteUDisk09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteUDisk(req)
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}

var testStep523DescribeUDisk10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("clone_udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    3,
	RetryInterval: 20 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep523DeleteUDisk11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("clone_udisk_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteUDisk(req)
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}
