// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario285(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "285",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"eip_bw":           2,
				"bwp1_bw":          3,
				"eip_operatorname": "International",
				"bwp2_bw":          4,
				"Image_Id":         "#{u_get_image_resource($Region,$Zone)}",
				"Region":           "tw-tp",
				"Zone":             "tw-tp-01",
			}
		},
		Owners: []string{"chenoa.chen@ucloud.cn"},
		Title:  "BandwidthPackage自动化回归-带宽包-基础-02",
		Steps: []*driver.Step{
			testStep285DescribeImage01,
			testStep285AllocateEIP02,
			testStep285CreateVPC03,
			testStep285CreateSubnet04,
			testStep285CreateUHostInstance05,
			testStep285BindEIP06,
			testStep285CreateBandwidthPackage07,
			testStep285DescribeBandwidthPackage08,
			testStep285GetBandwidthSum09,
			testStep285DeleteBandwidthPackage10,
			testStep285CreateBandwidthPackage11,
			testStep285DescribeBandwidthPackage12,
			testStep285DeleteBandwidthPackage13,
			testStep285UnBindEIP14,
			testStep285ReleaseEIP15,
			testStep285PoweroffUHostInstance16,
			testStep285TerminateUHostInstance17,
			testStep285DeleteSubnet18,
			testStep285DeleteVPC19,
		},
	})
}

var testStep285DescribeImage01 = &driver.Step{
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

var testStep285AllocateEIP02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewAllocateEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":       step.Scenario.GetVar("Region"),
			"PayMode":      "Bandwidth",
			"OperatorName": step.Scenario.GetVar("eip_operatorname"),
			"Name":         "eip_auto",
			"ChargeType":   "Month",
			"Bandwidth":    step.Scenario.GetVar("eip_bw"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateEIP(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("eip_id", step.Must(utils.GetValue(resp, "EIPSet.0.EIPId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "申请弹性IP",
	FastFail:      false,
}

var testStep285CreateVPC03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Network": []interface{}{
				"172.16.0.0/16",
			},
			"Name": "bandewithtest",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVPC(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VPCId", step.Must(utils.GetValue(resp, "VPCId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateVPCResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建VPC",
	FastFail:      false,
}

var testStep285CreateSubnet04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":   step.Scenario.GetVar("VPCId"),
			"Subnet":  "172.16.11.0",
			"Region":  step.Scenario.GetVar("Region"),
			"Netmask": 24,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateSubnet(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("SubnetId", step.Must(utils.GetValue(resp, "SubnetId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateSubnetResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建子网",
	FastFail:      false,
}

var testStep285CreateUHostInstance05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"VPCId":       step.Scenario.GetVar("VPCId"),
			"Tag":         "Default",
			"SubnetId":    step.Scenario.GetVar("SubnetId"),
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "packet-s1-internationnal",
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

var testStep285BindEIP06 = &driver.Step{
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
			"EIPId":        step.Scenario.GetVar("eip_id"),
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
	StartupDelay:  time.Duration(180) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "绑定弹性IP",
	FastFail:      false,
}

var testStep285CreateBandwidthPackage07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewCreateBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"TimeRange":  2,
			"Region":     step.Scenario.GetVar("Region"),
			"EnableTime": step.Must(functions.Calculate("+", 30, step.Must(functions.GetTimestamp(10)))),
			"EIPId":      step.Scenario.GetVar("eip_id"),
			"Bandwidth":  step.Scenario.GetVar("bwp1_bw"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateBandwidthPackage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("bwp_id1", step.Must(utils.GetValue(resp, "BandwidthPackageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建带宽包",
	FastFail:      false,
}

var testStep285DescribeBandwidthPackage08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeBandwidthPackage(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.EIPId", step.Scenario.GetVar("eip_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.Bandwidth", step.Scenario.GetVar("bwp1_bw"), "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.BandwidthPackageId", step.Scenario.GetVar("bwp_id1"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取带宽包信息",
	FastFail:      false,
}

var testStep285GetBandwidthSum09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("GetBandwidthSum")
		err = req.SetPayload(map[string]interface{}{
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
			validation.Builtins.NewValidator("Info.BwOutSum", 0, "gt"),
			validation.Builtins.NewValidator("Info.BwOutPackageSum", 0, "gt"),
			validation.Builtins.NewValidator("Info.BwOutAll", 0, "gt"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取带宽总量信息",
	FastFail:      false,
}

var testStep285DeleteBandwidthPackage10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDeleteBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":             step.Scenario.GetVar("Region"),
			"BandwidthPackageId": step.Scenario.GetVar("bwp_id1"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteBandwidthPackage(req)
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
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除带宽包",
	FastFail:      false,
}

var testStep285CreateBandwidthPackage11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewCreateBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"TimeRange": 1,
			"Region":    step.Scenario.GetVar("Region"),
			"EIPId":     step.Scenario.GetVar("eip_id"),
			"Bandwidth": step.Scenario.GetVar("bwp2_bw"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateBandwidthPackage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("bwp_id2", step.Must(utils.GetValue(resp, "BandwidthPackageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建带宽包",
	FastFail:      false,
}

var testStep285DescribeBandwidthPackage12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDescribeBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeBandwidthPackage(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.EIPId", step.Scenario.GetVar("eip_id"), "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.Bandwidth", step.Scenario.GetVar("bwp2_bw"), "str_eq"),
			validation.Builtins.NewValidator("DataSets.0.BandwidthPackageId", step.Scenario.GetVar("bwp_id2"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取带宽包信息",
	FastFail:      false,
}

var testStep285DeleteBandwidthPackage13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewDeleteBandwidthPackageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":             step.Scenario.GetVar("Region"),
			"BandwidthPackageId": step.Scenario.GetVar("bwp_id2"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteBandwidthPackage(req)
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
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除带宽包",
	FastFail:      false,
}

var testStep285UnBindEIP14 = &driver.Step{
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
			"EIPId":        step.Scenario.GetVar("eip_id"),
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
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "解绑弹性IP",
	FastFail:      false,
}

var testStep285ReleaseEIP15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UNet")
		if err != nil {
			return nil, err
		}
		client := c.(*unet.UNetClient)

		req := client.NewReleaseEIPRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"EIPId":  step.Scenario.GetVar("eip_id"),
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
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "释放弹性IP",
	FastFail:      false,
}

var testStep285PoweroffUHostInstance16 = &driver.Step{
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
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep285TerminateUHostInstance17 = &driver.Step{
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
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}

var testStep285DeleteSubnet18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetId": step.Scenario.GetVar("SubnetId"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteSubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DeleteSubnetResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    5,
	RetryInterval: 300 * time.Second,
	Title:         "删除子网",
	FastFail:      false,
}

var testStep285DeleteVPC19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":  step.Scenario.GetVar("VPCId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteVPC(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DeleteVPCResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除VPC",
	FastFail:      false,
}
