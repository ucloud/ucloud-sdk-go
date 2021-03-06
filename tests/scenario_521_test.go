// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario521(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "521",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"DiskType":       "DataDisk",
				"Size":           1,
				"UDataArkMode":   "Yes",
				"Name":           "udisk_fz",
				"imageid":        "#{u_get_image_resource($Region,$Zone)}",
				"miaojifangzhou": scenario.Must(functions.Calculate("+", scenario.Must(functions.GetTimestamp(10)), 200)),
				"Region":         "cn-bj2",
				"Zone":           "cn-bj2-02",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "UDisk-普通方舟盘_01",
		Steps: []*driver.Step{
			testStep521DescribeImage01,
			testStep521DescribeUDiskPrice02,
			testStep521CheckUDiskAllowance03,
			testStep521CreateUDisk04,
			testStep521DescribeUDisk05,
			testStep521RenameUDisk06,
			testStep521DescribeUDisk07,
			testStep521DescribeImage08,
			testStep521CreateUHostInstance09,
			testStep521DescribeUHostInstance10,
			testStep521DescribeUHostLite11,
			testStep521AttachUDisk12,
			testStep521DescribeUDisk13,
			testStep521DescribeVDiskTmList14,
			testStep521CloneUDiskUDataArk15,
			testStep521DescribeUDisk16,
			testStep521DeleteUDisk17,
			testStep521DetachUDisk18,
			testStep521DescribeUDisk19,
			testStep521RestoreUDisk20,
			testStep521DescribeUDisk21,
			testStep521DeleteUDisk22,
			testStep521PoweroffUHostInstance23,
			testStep521DescribeUHostInstance24,
			testStep521TerminateUHostInstance25,
		},
	})
}

var testStep521DescribeImage01 = &driver.Step{
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

		step.Scenario.SetVar("imageid", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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

var testStep521DescribeUDiskPrice02 = &driver.Step{
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

var testStep521CheckUDiskAllowance03 = &driver.Step{
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

var testStep521CreateUDisk04 = &driver.Step{
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

		step.Scenario.SetVar("udisk_fz_id", step.Must(utils.GetValue(resp, "UDiskId.0")))
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

var testStep521DescribeUDisk05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Tag", "Default", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep521RenameUDisk06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewRenameUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"UDiskName": step.Must(functions.Concat("auto_", step.Scenario.GetVar("Name"))),
			"UDiskId":   step.Scenario.GetVar("udisk_fz_id"),
			"Region":    step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.RenameUDisk(req)
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
	Title:         "重命名云硬盘",
	FastFail:      false,
}

var testStep521DescribeUDisk07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
			validation.Builtins.NewValidator("DataSet.0.Name", step.Must(functions.Concat("auto_", step.Scenario.GetVar("Name"))), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep521DescribeImage08 = &driver.Step{
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

		step.Scenario.SetVar("imageid", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
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

var testStep521CreateUHostInstance09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":               step.Scenario.GetVar("Zone"),
			"UHostType":          "Normal",
			"TimemachineFeature": "No",
			"StorageType":        "LocalDisk",
			"Region":             step.Scenario.GetVar("Region"),
			"Quantity":           0,
			"Password":           "VXFhNzg5VGVzdCFAIyQ7LA==",
			"NetCapability":      "Normal",
			"Name":               "auto_uhost-同AZ",
			"Memory":             1024,
			"LoginMode":          "Password",
			"ImageId":            step.Scenario.GetVar("imageid"),
			"HotplugFeature":     "false",
			"HostType":           "N2",
			"DiskSpace":          0,
			"ChargeType":         "Month",
			"CPU":                1,
			"BootDiskSpace":      20,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("uhost_id", step.Must(utils.GetValue(resp, "UHostIds.0")))
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
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep521DescribeUHostInstance10 = &driver.Step{
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
				step.Scenario.GetVar("uhost_id"),
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
			validation.Builtins.NewValidator("Action", "DescribeUHostInstanceResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.UHostId", step.Scenario.GetVar("uhost_id"), "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    30,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep521DescribeUHostLite11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUHostLite")
		err = req.SetPayload(map[string]interface{}{
			"Zone":   step.Scenario.GetVar("Zone"),
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  60,
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
			validation.Builtins.NewValidator("TotalCount", 0, "ne"),
		}
	},
	StartupDelay:  time.Duration(1) * time.Second,
	MaxRetries:    3,
	RetryInterval: 5 * time.Second,
	Title:         "内部调用，列出UHost实例",
	FastFail:      false,
}

var testStep521AttachUDisk12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewAttachUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AttachUDisk(req)
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
	StartupDelay:  time.Duration(110) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "挂载云硬盘",
	FastFail:      true,
}

var testStep521DescribeUDisk13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
			validation.Builtins.NewValidator("DataSet.0.Status", "InUse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(20) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep521DescribeVDiskTmList14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeVDiskTmList")
		err = req.SetPayload(map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"VDiskId":      step.Scenario.GetVar("udisk_fz_id"),
			"SnapshotType": "all",
			"Region":       step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeVDiskTmListResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "查询磁盘备份链",
	FastFail:      false,
}

var testStep521CloneUDiskUDataArk15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewCloneUDiskUDataArkRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDiskId":      step.Scenario.GetVar("udisk_fz_id"),
			"UDataArkMode": "Yes",
			"SnapshotTime": step.Scenario.GetVar("miaojifangzhou"),
			"Region":       step.Scenario.GetVar("Region"),
			"Name":         "miaojifangzhou",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CloneUDiskUDataArk(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("udisk_fzclone_id", step.Must(utils.GetValue(resp, "UDiskId.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CloneUDiskUDataArkResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "从数据方舟的备份创建UDisk",
	FastFail:      true,
}

var testStep521DescribeUDisk16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fzclone_id"),
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
			validation.Builtins.NewValidator("Action", "DescribeUDiskResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(180) * time.Second,
	MaxRetries:    60,
	RetryInterval: 20 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep521DeleteUDisk17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fzclone_id"),
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
			validation.Builtins.NewValidator("Action", "DeleteUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}

var testStep521DetachUDisk18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDetachUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DetachUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("UDiskId", step.Scenario.GetVar("udisk_fz_id"), "str_eq"),
			validation.Builtins.NewValidator("UHostId", step.Scenario.GetVar("uhost_id"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "卸载云硬盘",
	FastFail:      false,
}

var testStep521DescribeUDisk19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    120,
	RetryInterval: 3 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      true,
}

var testStep521RestoreUDisk20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewRestoreUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":         step.Scenario.GetVar("Zone"),
			"UDiskId":      step.Scenario.GetVar("udisk_fz_id"),
			"SnapshotTime": step.Scenario.GetVar("miaojifangzhou"),
			"Region":       step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.RestoreUDisk(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "RestoreUDiskResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "从备份恢复数据至UDisk",
	FastFail:      false,
}

var testStep521DescribeUDisk21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDescribeUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
		}
	},
	StartupDelay:  time.Duration(160) * time.Second,
	MaxRetries:    40,
	RetryInterval: 10 * time.Second,
	Title:         "获取云硬盘列表",
	FastFail:      false,
}

var testStep521DeleteUDisk22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UDisk")
		if err != nil {
			return nil, err
		}
		client := c.(*udisk.UDiskClient)

		req := client.NewDeleteUDiskRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UDiskId": step.Scenario.GetVar("udisk_fz_id"),
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
	MaxRetries:    30,
	RetryInterval: 5 * time.Second,
	Title:         "删除云硬盘",
	FastFail:      false,
}

var testStep521PoweroffUHostInstance23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
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
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep521DescribeUHostInstance24 = &driver.Step{
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
				step.Scenario.GetVar("uhost_id"),
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
	StartupDelay:  time.Duration(20) * time.Second,
	MaxRetries:    0,
	RetryInterval: 0 * time.Second,
	Title:         "获取主机信息",
	FastFail:      false,
}

var testStep521TerminateUHostInstance25 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}
