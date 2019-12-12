// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

func TestSet2117(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: 2117,
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Region":       "cn-sh2",
				"Zone":         "cn-sh2-02",
				"ResourceType": "single",
				"ConfigId":     "9a891891-c245-4b66-bce8-67e59430d67c",
				"Name":         "single_memcache",
				"Protocol":     "memcache",
				"Count":        1,
			}
		},
		Owners: []string{"caroline.yuan@ucloud.cn"},
		Title:  "SingleMemcache",
		Steps: []*driver.Step{
			testStep2117CreateUMemcacheGroup00,
			testStep2117CheckUMemcacheAllowance01,
			testStep2117DescribeUMem02,
			testStep2117DescribeUMemcacheGroup03,
			testStep2117ResizeUMemcacheGroup04,
			testStep2117DescribeUMem05,
			testStep2117RestartUMemcacheGroup06,
			testStep2117DescribeUMem07,
			testStep2117DeleteUMemcacheGroup08,
			testStep2117DescribeUMem09,
		},
	})
}

var testStep2117CreateUMemcacheGroup00 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("CreateUMemcacheGroup")
		req.SetPayload(map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"Size":       1,
			"Region":     step.Scenario.GetVar("Region"),
			"Quantity":   1,
			"Name":       step.Scenario.GetVar("Name"),
			"ConfigId":   step.Scenario.GetVar("ConfigId"),
			"ChargeType": "Month",
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("group_id", step.Must(utils.GetValue(resp, "GroupId")))

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUMemcacheGroupResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建单机Memcache",
	FastFail:      false,
}

var testStep2117CheckUMemcacheAllowance01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("CheckUMemcacheAllowance")
		req.SetPayload(map[string]interface{}{
			"Zone":   step.Scenario.GetVar("Zone"),
			"Size":   1,
			"Region": step.Scenario.GetVar("Region"),
			"Count":  step.Scenario.GetVar("Count"),
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CheckUMemcacheAllowanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "检查UMemcache剩余资源",
	FastFail:      false,
}

var testStep2117DescribeUMem02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUMem")
		req.SetPayload(map[string]interface{}{
			"ResourceId": step.Scenario.GetVar("group_id"),
			"Region":     step.Scenario.GetVar("Region"),
			"Protocol":   step.Scenario.GetVar("Protocol"),
			"Offset":     0,
			"Limit":      1000,
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.State", "Running", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(20) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取UMem列表",
	FastFail:      false,
}

var testStep2117DescribeUMemcacheGroup03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUMemcacheGroup")
		req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"Region":  step.Scenario.GetVar("Region"),
			"Offset":  0,
			"Limit":   100,
			"GroupId": step.Scenario.GetVar("group_id"),
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUMemcacheGroupResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(100) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "显示Memcache",
	FastFail:      true,
}

var testStep2117ResizeUMemcacheGroup04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("ResizeUMemcacheGroup")
		req.SetPayload(map[string]interface{}{
			"Size":    2,
			"Region":  step.Scenario.GetVar("Region"),
			"GroupId": step.Scenario.GetVar("group_id"),
		})

		resp, err := client.GenericInvoke(req)
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
	Title:         "调整容量",
	FastFail:      false,
}

var testStep2117DescribeUMem05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUMem")
		req.SetPayload(map[string]interface{}{
			"ResourceId": step.Scenario.GetVar("group_id"),
			"Region":     step.Scenario.GetVar("Region"),
			"Protocol":   step.Scenario.GetVar("Protocol"),
			"Offset":     0,
			"Limit":      1000,
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.State", "Running", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取UMem列表",
	FastFail:      false,
}

var testStep2117RestartUMemcacheGroup06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("RestartUMemcacheGroup")
		req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"Region":  step.Scenario.GetVar("Region"),
			"GroupId": step.Scenario.GetVar("group_id"),
		})

		resp, err := client.GenericInvoke(req)
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
	Title:         "重启Memcache组",
	FastFail:      false,
}

var testStep2117DescribeUMem07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUMem")
		req.SetPayload(map[string]interface{}{
			"ResourceId": step.Scenario.GetVar("group_id"),
			"Region":     step.Scenario.GetVar("Region"),
			"Protocol":   step.Scenario.GetVar("Protocol"),
			"Offset":     0,
			"Limit":      1000,
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.State", "Running", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取UMem列表",
	FastFail:      false,
}

var testStep2117DeleteUMemcacheGroup08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DeleteUMemcacheGroup")
		req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"Region":  step.Scenario.GetVar("Region"),
			"GroupId": step.Scenario.GetVar("group_id"),
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DeleteUMemcacheGroupResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除单机Memcache",
	FastFail:      false,
}

var testStep2117DescribeUMem09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.NewClient("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)
		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUMem")
		req.SetPayload(map[string]interface{}{
			"ResourceId": step.Scenario.GetVar("group_id"),
			"Region":     step.Scenario.GetVar("Region"),
			"Protocol":   step.Scenario.GetVar("Protocol"),
			"Offset":     0,
			"Limit":      100,
		})

		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUMemResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet", step.Scenario.GetVar("group_id"), "object_not_contains"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取UMem列表",
	FastFail:      false,
}
