package stepflow

/*
WorkflowTemplate - Workflow对象定义

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type WorkflowTemplate struct {

	// 工作流的Activities，详细信息见工作流的Activity定义
	Activites []ActivityTemplate

	// 工作流的输入，详细信息见Param
	Input []Param

	// 工作流的输出，详细信息见Param
	Output []Param
}
