// Code is generated by ucloud-model, DO NOT EDIT IT.

package uai_modelverse

/*
AppInfo - 应用信息
*/
type AppInfo struct {

	// 应用描述
	AppDes string

	// 应用ID
	AppID string

	// 应用名称
	AppName string

	// 应用状态
	AppState int

	// 应用类型
	AppType int

	// 应用创建时间
	CreateTime int

	// 语言模型ID
	LLMID string

	// 模型采样温度
	LLMTemperature int

	// 模型采样温度的另一种方法（核取样）
	LLMTopP int

	// 应用更新时间
	UpdateTime int
}

/*
LLMInfo - 语言模型列表
*/
type LLMInfo struct {

	// 模型描述
	LLMDes string

	// 模型ID
	LLMID string

	// 模型名称
	LLMName string

	// 模型token单价
	LLMPrice float64

	// 模型类型
	LLMType int
}