// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/ufs"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario4458(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "4458",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"ufs_size":      100,
				"storage_type":  "Advanced",
				"charge_type":   "Month",
				"protocol_type": "NFSv4",
				"volume_name1":  "ufs-advanced",
				"remark1":       "ufs-advanced_remark1",
				"volume_name2":  "ufs-advanced_rename",
				"remark2":       "ufs-advanced_remark2",
				"Region":        "cn-bj2",
			}
		},
		Owners: []string{"chenoa.chen@ucloud.cn"},
		Title:  "UFS-SSD性能型（国内机房）",
		Steps: []*driver.Step{
			testStep4458DescribeUFSVolumePrice01,
			testStep4458CreateUFSVolume02,
			testStep4458DescribeUFSVolume03,
			testStep4458DescribeUFSVolume204,
			testStep4458DescribeUFSVPCSubnet05,
			testStep4458AddUFSVolumeMountPoint06,
			testStep4458DescribeUFSVolume207,
			testStep4458DescribeUFSVolumeMountpoint08,
			testStep4458DescribeUFSProductPolicy09,
			testStep4458DescribeUFSVolumeUpgradePrice10,
			testStep4458ExtendUFSVolume11,
			testStep4458DescribeUFSVolume212,
			testStep4458RemoveUFSVolumeMountPoint13,
			testStep4458DescribeUFSVolumeMountpoint14,
			testStep4458UpdateUFSVolumeInfo15,
			testStep4458DescribeUFSVolume216,
			testStep4458RemoveUFSVolume17,
		},
	})
}

var testStep4458DescribeUFSVolumePrice01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVolumePrice")
		err = req.SetPayload(map[string]interface{}{
			"StorageType": step.Scenario.GetVar("storage_type"),
			"Size":        step.Scenario.GetVar("ufs_size"),
			"Region":      step.Scenario.GetVar("Region"),
			"Quantity":    1,
			"ChargeType":  step.Scenario.GetVar("charge_type"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSVolumePriceResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ChargeType", step.Scenario.GetVar("charge_type"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Price", step.Must(functions.Calculate("*", step.Must(functions.Calculate("*", 0.8, 100)), step.Scenario.GetVar("ufs_size"))), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统价格",
	FastFail:      false,
}

var testStep4458CreateUFSVolume02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewCreateUFSVolumeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeName":   step.Scenario.GetVar("volume_name1"),
			"StorageType":  step.Scenario.GetVar("storage_type"),
			"Size":         step.Scenario.GetVar("ufs_size"),
			"Remark":       step.Scenario.GetVar("remark1"),
			"Region":       step.Scenario.GetVar("Region"),
			"ProtocolType": step.Scenario.GetVar("protocol_type"),
			"ChargeType":   step.Scenario.GetVar("charge_type"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUFSVolume(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("volume_id", step.Must(utils.GetValue(resp, "VolumeId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUFSVolumeResponse", "str_eq"),
			validation.Builtins.NewValidator("VolumeStatus", "UnInitialized", "str_eq"),
			validation.Builtins.NewValidator("VolumeName", step.Scenario.GetVar("volume_name1"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建文件系统",
	FastFail:      true,
}

var testStep4458DescribeUFSVolume03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVolume")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSVolumeResponse", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeName", step.Scenario.GetVar("volume_name1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeId", step.Scenario.GetVar("volume_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeStatus", "Initialized", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Size", step.Scenario.GetVar("ufs_size"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统列表",
	FastFail:      false,
}

var testStep4458DescribeUFSVolume204 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewDescribeUFSVolume2Request()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUFSVolume2(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUFSVolume2Response", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeName", step.Scenario.GetVar("volume_name1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeId", step.Scenario.GetVar("volume_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ProtocolType", step.Scenario.GetVar("protocol_type"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Remark", step.Scenario.GetVar("remark1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Size", step.Scenario.GetVar("ufs_size"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统列表",
	FastFail:      false,
}

var testStep4458DescribeUFSVPCSubnet05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVPCSubnet")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("vpc_id_0", step.Must(utils.GetValue(resp, "DataSet.0.VpcId")))
		step.Scenario.SetVar("subnet_id_0", step.Must(utils.GetValue(resp, "DataSet.0.SubnetIds.0.SubnetId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUFSVPCSubnetResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "展示可供文件系统挂载的VPC和SUBNET",
	FastFail:      false,
}

var testStep4458AddUFSVolumeMountPoint06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("AddUFSVolumeMountPoint")
		err = req.SetPayload(map[string]interface{}{
			"VpcId":          step.Scenario.GetVar("vpc_id_0"),
			"VolumeId":       step.Scenario.GetVar("volume_id"),
			"SubnetId":       step.Scenario.GetVar("subnet_id_0"),
			"Region":         step.Scenario.GetVar("Region"),
			"MountPointName": step.Must(functions.Concat("mntpoint-", step.Scenario.GetVar("volume_name1"))),
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
			validation.Builtins.NewValidator("Action", "AddUFSVolumeMountPointResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "添加文件系统挂载点",
	FastFail:      false,
}

var testStep4458DescribeUFSVolume207 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewDescribeUFSVolume2Request()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUFSVolume2(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUFSVolume2Response", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeName", step.Scenario.GetVar("volume_name1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeId", step.Scenario.GetVar("volume_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ProtocolType", step.Scenario.GetVar("protocol_type"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Size", step.Scenario.GetVar("ufs_size"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.TotalMountPointNum", 1, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.MaxMountPointNum", 5, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(2) * time.Second,
	MaxRetries:    60,
	RetryInterval: 2 * time.Second,
	Title:         "获取文件系统列表",
	FastFail:      false,
}

var testStep4458DescribeUFSVolumeMountpoint08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVolumeMountpoint")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSVolumeMountpointResponse", "str_eq"),
			validation.Builtins.NewValidator("TotalMountPointNum", 1, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VpcId", step.Scenario.GetVar("vpc_id_0"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.SubnetId", step.Scenario.GetVar("subnet_id_0"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统挂载点信息",
	FastFail:      false,
}

var testStep4458DescribeUFSProductPolicy09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSProductPolicy")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
			"Product":  step.Scenario.GetVar("storage_type"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSProductPolicyResponse", "str_eq"),
			validation.Builtins.NewValidator("Protocol.0", "NFSv4", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取UFS产品策略",
	FastFail:      false,
}

var testStep4458DescribeUFSVolumeUpgradePrice10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVolumeUpgradePrice")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Size":     step.Must(functions.Calculate("*", 2, step.Scenario.GetVar("ufs_size"))),
			"Region":   step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSVolumeUpgradePriceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "文件系统扩容价格",
	FastFail:      false,
}

var testStep4458ExtendUFSVolume11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewExtendUFSVolumeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Size":     step.Must(functions.Calculate("*", 2, step.Scenario.GetVar("ufs_size"))),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ExtendUFSVolume(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "ExtendUFSVolumeResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "文件系统扩容",
	FastFail:      false,
}

var testStep4458DescribeUFSVolume212 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewDescribeUFSVolume2Request()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUFSVolume2(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUFSVolume2Response", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeName", step.Scenario.GetVar("volume_name1"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeId", step.Scenario.GetVar("volume_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ProtocolType", step.Scenario.GetVar("protocol_type"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Size", step.Must(functions.Calculate("*", 2, step.Scenario.GetVar("ufs_size"))), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.TotalMountPointNum", 1, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.MaxMountPointNum", 5, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    30,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统列表",
	FastFail:      false,
}

var testStep4458RemoveUFSVolumeMountPoint13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("RemoveUFSVolumeMountPoint")
		err = req.SetPayload(map[string]interface{}{
			"VpcId":    step.Scenario.GetVar("vpc_id_0"),
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"SubnetId": step.Scenario.GetVar("subnet_id_0"),
			"Region":   step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "RemoveUFSVolumeMountPointResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    3,
	RetryInterval: 2 * time.Second,
	Title:         "删除文件系统挂载点",
	FastFail:      false,
}

var testStep4458DescribeUFSVolumeMountpoint14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeUFSVolumeMountpoint")
		err = req.SetPayload(map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
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
			validation.Builtins.NewValidator("Action", "DescribeUFSVolumeMountpointResponse", "str_eq"),
			validation.Builtins.NewValidator("TotalMountPointNum", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统挂载点信息",
	FastFail:      false,
}

var testStep4458UpdateUFSVolumeInfo15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("UpdateUFSVolumeInfo")
		err = req.SetPayload(map[string]interface{}{
			"VolumeName": step.Scenario.GetVar("volume_name2"),
			"VolumeId":   step.Scenario.GetVar("volume_id"),
			"Remark":     step.Scenario.GetVar("remark2"),
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
			validation.Builtins.NewValidator("Action", "UpdateUFSVolumeInfoResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "更改文件系统相关信息（名称／备注）",
	FastFail:      false,
}

var testStep4458DescribeUFSVolume216 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewDescribeUFSVolume2Request()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeUFSVolume2(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeUFSVolume2Response", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VolumeName", step.Scenario.GetVar("volume_name2"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Remark", step.Scenario.GetVar("remark2"), "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.TotalMountPointNum", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取文件系统列表",
	FastFail:      false,
}

var testStep4458RemoveUFSVolume17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UFS")
		if err != nil {
			return nil, err
		}
		client := c.(*ufs.UFSClient)

		req := client.NewRemoveUFSVolumeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VolumeId": step.Scenario.GetVar("volume_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.RemoveUFSVolume(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "RemoveUFSVolumeResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除文件系统",
	FastFail:      false,
}
