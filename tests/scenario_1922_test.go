// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario1922(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "1922",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Password":                      "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"SnapshotSysName":               "snapshot-ARK-SYS-01",
				"SnapshotSysDesc":               "snapshot-ARK-SYS-01-desc",
				"SnapDiskType":                  "LocalBoot",
				"SnapshotDataNameModify":        "snapshot-ARK-DATA-01-modify",
				"SnapshotDataDescModify":        "snapshot-ARK-DATA-01-desc-Modify",
				"UhostName":                     "ARK-local-1",
				"SnapshotDataName":              "snapshot-ARK-DATA-01",
				"SnapshotDataDesc":              "snapshot-ARK-DATA-01-desc",
				"CreateFromTimeMachinePassword": "Z3VhbmxpeXVhbm1pbWExMjMhQCM=",
				"ImageID":                       "#{u_get_image_resource($Region,$Zone)}",
				"BootSize":                      20,
				"BootType":                      "LOCAL_NORMAL",
				"DiskSize":                      20,
				"DiskType":                      "LOCAL_NORMAL",
				"UHostType":                     "N2",
				"BootBackup":                    "NONE",
				"DiskBackup":                    "NONE",
				"Region":                        "cn-bj2",
				"Zone":                          "cn-bj2-03",
			}
		},
		Owners: []string{"maggie.an@ucloud.cn"},
		Title:  "主机功能之快照-方舟-N2-LOCAL_NORMAL-LOCAL_NORMAL-1",
		Steps: []*driver.Step{
			testStep1922DescribeImage01,
			testStep1922CreateUHostInstance02,
			testStep1922DescribeUHostInstance03,
			testStep1922StopUHostInstance04,
			testStep1922DescribeUHostInstance05,
			testStep1922UpgradeToArkUHostInstance06,
			testStep1922DescribeUHostInstance07,
			testStep1922StartUHostInstance08,
			testStep1922DescribeUHostInstance09,
			testStep1922DescribeUhostTmMeta10,
			testStep1922DescribeVDiskTmList11,
			testStep1922DescribeVDiskTmList12,
			testStep1922CreateUserVDiskSnapshot13,
			testStep1922CreateUserVDiskSnapshot14,
			testStep1922DescribeSnapshot15,
			testStep1922DescribeSnapshot16,
			testStep1922ModifySnapshot17,
			testStep1922DescribeSnapshot18,
			testStep1922DeleteSnapshot19,
			testStep1922DeleteSnapshot20,
			testStep1922DescribeSnapshot21,
			testStep1922StopUHostInstance22,
			testStep1922DescribeUHostInstance23,
			testStep1922TerminateUHostInstance24,
		},
	})
}

var testStep1922DescribeImage01 = &driver.Step{
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

var testStep1922CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"UHostType": step.Scenario.GetVar("UHostType"),
			"SetId":     step.Scenario.GetVar("SetId"),
			"Region":    step.Scenario.GetVar("Region"),
			"Password":  "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":      step.Scenario.GetVar("UhostName"),
			"Memory":    1024,
			"LoginMode": "Password",
			"ImageId":   step.Scenario.GetVar("ImageID"),
			"HostIp":    step.Scenario.GetVar("HostIp"),
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
			"CPU": 1,
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

var testStep1922DescribeUHostInstance03 = &driver.Step{
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
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.TimemachineFeature", "no", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    100,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1922StopUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.StopUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "StopUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep1922DescribeUHostInstance05 = &driver.Step{
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
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1922UpgradeToArkUHostInstance06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewUpgradeToArkUHostInstanceRequest()
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

		resp, err := client.UpgradeToArkUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "UpgradeToArkUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "普通升级为方舟机型",
	FastFail:      true,
}

var testStep1922DescribeUHostInstance07 = &driver.Step{
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
			validation.Builtins.NewValidator("UHostSet.0.TimemachineFeature", "yes", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.BootDiskState", "Normal", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.BackupType", "DATAARK", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(100) * time.Second,
	MaxRetries:    120,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1922StartUHostInstance08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStartUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.StartUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "StartUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "启动主机",
	FastFail:      true,
}

var testStep1922DescribeUHostInstance09 = &driver.Step{
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
			validation.Builtins.NewValidator("UHostSet.0.State", "Running", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.TimemachineFeature", "yes", "str_eq"),
			validation.Builtins.NewValidator("UHostSet.0.DiskSet.0.BackupType", "DATAARK", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 30 * time.Second,
	Title:         "获取主机信息",
	FastFail:      true,
}

var testStep1922DescribeUhostTmMeta10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUhostTmMeta")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UhostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VdiskIdSys", step.Must(utils.GetValue(resp, "DataSet.0.VdiskId")))
		step.Scenario.SetVar("VdiskIdData", step.Must(utils.GetValue(resp, "DataSet.1.VdiskId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUhostTmMetaResponse", "str_eq"),
			validation.Builtins.NewValidator("UtmStatus", "normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(100) * time.Second,
	MaxRetries:    60,
	RetryInterval: 60 * time.Second,
	Title:         "查询主机方舟状态",
	FastFail:      true,
}

var testStep1922DescribeVDiskTmList11 = &driver.Step{
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
			"VDiskId":      step.Scenario.GetVar("VdiskIdSys"),
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
	FastFail:      true,
}

var testStep1922DescribeVDiskTmList12 = &driver.Step{
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
			"VDiskId":      step.Scenario.GetVar("VdiskIdData"),
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
	FastFail:      true,
}

var testStep1922CreateUserVDiskSnapshot13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("CreateUserVDiskSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"VDiskId": step.Scenario.GetVar("VdiskIdSys"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    step.Scenario.GetVar("SnapshotSysName"),
			"Comment": step.Scenario.GetVar("SnapshotSysDesc"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VdiskSnapIDSys", step.Must(utils.GetValue(resp, "SnapshotId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUserVDiskSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建快照",
	FastFail:      true,
}

var testStep1922CreateUserVDiskSnapshot14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("CreateUserVDiskSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"VDiskId": step.Scenario.GetVar("VdiskIdData"),
			"Region":  step.Scenario.GetVar("Region"),
			"Name":    step.Scenario.GetVar("SnapshotDataName"),
			"Comment": step.Scenario.GetVar("SnapshotDataDesc"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VdiskSnapIDData", step.Must(utils.GetValue(resp, "SnapshotId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUserVDiskSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建快照",
	FastFail:      true,
}

var testStep1922DescribeSnapshot15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"SnapshotIds": []interface{}{
				step.Scenario.GetVar("VdiskSnapIDSys"),
			},
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeSnapshotResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotId", step.Scenario.GetVar("VdiskSnapIDSys"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotName", step.Scenario.GetVar("SnapshotSysName"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.DiskType", "LocalBoot", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.ResourceName", step.Scenario.GetVar("UhostName"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotDescription", step.Scenario.GetVar("SnapshotSysDesc"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    40,
	RetryInterval: 30 * time.Second,
	Title:         "描述快照",
	FastFail:      true,
}

var testStep1922DescribeSnapshot16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"SnapshotIds": []interface{}{
				step.Scenario.GetVar("VdiskSnapIDData"),
			},
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeSnapshotResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotId", step.Scenario.GetVar("VdiskSnapIDData"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotName", step.Scenario.GetVar("SnapshotDataName"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.DiskType", "LocalData", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.ResourceName", step.Scenario.GetVar("UhostName"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotDescription", step.Scenario.GetVar("SnapshotDataDesc"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      true,
}

var testStep1922ModifySnapshot17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("ModifySnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":                step.Scenario.GetVar("Zone"),
			"SnapshotName":        step.Scenario.GetVar("SnapshotDataNameModify"),
			"SnapshotId":          step.Scenario.GetVar("VdiskSnapIDData"),
			"SnapshotDescription": step.Scenario.GetVar("SnapshotDataDescModify"),
			"Region":              step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "ModifySnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "修改快照",
	FastFail:      true,
}

var testStep1922DescribeSnapshot18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"SnapshotIds": []interface{}{
				step.Scenario.GetVar("VdiskSnapIDData"),
			},
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeSnapshotResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotId", step.Scenario.GetVar("VdiskSnapIDData"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotName", step.Scenario.GetVar("SnapshotDataNameModify"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.DiskType", "LocalData", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.ResourceName", step.Scenario.GetVar("UhostName"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.SnapshotDescription", step.Scenario.GetVar("SnapshotDataDescModify"), "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 10 * time.Second,
	Title:         "描述快照",
	FastFail:      true,
}

var testStep1922DeleteSnapshot19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DeleteSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"SnapshotId": step.Scenario.GetVar("VdiskSnapIDSys"),
			"Region":     step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DeleteSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除快照",
	FastFail:      true,
}

var testStep1922DeleteSnapshot20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DeleteSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":       step.Scenario.GetVar("Zone"),
			"SnapshotId": step.Scenario.GetVar("VdiskSnapIDData"),
			"Region":     step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DeleteSnapshotResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除快照",
	FastFail:      true,
}

var testStep1922DescribeSnapshot21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeSnapshot")
		err = req.SetPayload(map[string]interface{}{
			"Zone":   step.Scenario.GetVar("Zone"),
			"Region": step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeSnapshotResponse", "str_eq"),
			validation.Builtins.NewValidator("UHostSnapshotSet", step.Scenario.GetVar("VdiskSnapIDSys"), "object_not_contains"),
			validation.Builtins.NewValidator("UHostSnapshotSet", step.Scenario.GetVar("VdiskSnapIDData"), "object_not_contains"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 30 * time.Second,
	Title:         "描述快照",
	FastFail:      true,
}

var testStep1922StopUHostInstance22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewStopUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("hostId"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.StopUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "StopUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "关闭主机",
	FastFail:      true,
}

var testStep1922DescribeUHostInstance23 = &driver.Step{
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

var testStep1922TerminateUHostInstance24 = &driver.Step{
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "TerminateUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      true,
}
