package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateSecurityGroup 创建安全组
func CreateSecurityGroup() {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createReq := ucloudstackClient.NewCreateSecurityGroupRequest()
	createReq.Region = ucloud.String("cn")
	createReq.Zone = ucloud.String("zone-01")
	createReq.Rule = []string{"TCP|12|0.0.0.0/0|ACCEPT|HIGH|1"}
	createReq.Name = ucloud.String("test")

	// send request
	resp, err := ucloudstackClient.CreateSecurityGroup(createReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("CreateSecurityGroup Success, sgID: %s", resp.SGID)
}

// DescribeSecurityGroup 查询安全组
func DescribeSecurityGroup(sgID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeReq := ucloudstackClient.NewDescribeSecurityGroupRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.SGIDs = []string{sgID}
	describeReq.Offset = ucloud.Int(0)
	describeReq.Limit = ucloud.Int(100)

	// send request
	resp, err := ucloudstackClient.DescribeSecurityGroup(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("DescribeSecurityGroup Success, infos: %+v \n", resp.Infos)
}

// DeleteSecurityGroup 删除安全组
func DeleteSecurityGroup(sgID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeReq := ucloudstackClient.NewDeleteSecurityGroupRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.SGID = ucloud.String(sgID)

	// send request
	resp, err := ucloudstackClient.DeleteSecurityGroup(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("DeleteSecurityGroup Success")
}

// CreateSecurityGroupRule 创建安全组规则
func CreateSecurityGroupRule(sgID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	createReq := ucloudstackClient.NewCreateSecurityGroupRuleRequest()
	createReq.Region = ucloud.String("cn")
	createReq.Zone = ucloud.String("zone-01")
	createReq.SGID = ucloud.String(sgID)
	createReq.Rules = []string{"TCP|88|0.0.0.0/0|ACCEPT|HIGH|1"}

	// send request
	resp, err := ucloudstackClient.CreateSecurityGroupRule(createReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("CreateSecurityGroupRule Success, infos: %+v \n", resp)
}

// UpdateSecurityGroupRule 修改安全组规则
func UpdateSecurityGroupRule(sgID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeReq := ucloudstackClient.NewUpdateSecurityGroupRuleRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.SGID = ucloud.String(sgID)
	describeReq.Rules = []string{"TCP|21|0.0.0.0/0|ACCEPT|HIGH|1|sg_rule-xS7i2DfZg"}

	// send request
	resp, err := ucloudstackClient.UpdateSecurityGroupRule(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("UpdateSecurityGroupRule Success")
}

// DeleteSecurityGroupRule  删除安全组规则
func DeleteSecurityGroupRule(sgID, sgRuleID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// request
	describeReq := ucloudstackClient.NewDeleteSecurityGroupRuleRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.SGID = ucloud.String(sgID)
	describeReq.SGRuleID = ucloud.String(sgRuleID)

	// send request
	resp, err := ucloudstackClient.DeleteSecurityGroupRule(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("DeleteSecurityGroupRule Success")
}
