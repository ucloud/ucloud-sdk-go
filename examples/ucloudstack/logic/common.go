package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// DescribeMetric 查询监控信息
func DescribeMetric(resourceID, resourceType, metricName string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewDescribeMetricRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceID = ucloud.String(resourceID)
	metricReq.ResourceType = ucloud.String(resourceType)
	metricReq.MetricName = []string{metricName}
	metricReq.BeginTime = ucloud.String("1576214967")
	metricReq.EndTime = ucloud.String("1576218567")

	// send request
	metricResp, err := ucloudstackClient.DescribeMetric(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return
	}

	if metricResp.TotalCount > 0 && len(metricResp.Infos[0].Infos) > 0 {
		fmt.Printf("value of %s at %d: %f\n",
			metricResp.Infos[0].MetricName,
			metricResp.Infos[0].Infos[0].Timestamp,
			metricResp.Infos[0].Infos[0].Value)
	}

}

// BindAlarmTemplate 绑定告警模板
func BindAlarmTemplate(resourceID, resourceType, alarmTemplateID string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewBindAlarmTemplateRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceIDs = []string{resourceID}
	metricReq.ResourceType = ucloud.String(resourceType)
	metricReq.AlarmTemplateID = ucloud.String(alarmTemplateID)

	// send request
	resp, err := ucloudstackClient.BindAlarmTemplate(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("BindAlarmTemplate Success")
}

// UnbindAlarmTemplate 解绑告警模板
func UnbindAlarmTemplate(resourceID, resourceType string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewUnbindAlarmTemplateRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceIDs = []string{resourceID}
	metricReq.ResourceType = ucloud.String(resourceType)

	// send request
	resp, err := ucloudstackClient.UnbindAlarmTemplate(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("UnbindAlarmTemplate Success")
}

// ModifyNameAndRemark 修改资源名称和说明
func ModifyNameAndRemark(resourceID, name, remark string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewModifyNameAndRemarkRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceID = ucloud.String(resourceID)
	metricReq.Name = ucloud.String(name)
	metricReq.Remark = ucloud.String(remark)

	// send request
	resp, err := ucloudstackClient.ModifyNameAndRemark(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %+v\n", resp)
		return
	}

	fmt.Printf("ModifyNameAndRemark Success")
}
