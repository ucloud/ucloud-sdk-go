package stepflow

/*
ActivityTemplate - 工作流的Activity定义

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type ActivityTemplate struct {

	// Activity的输入
	Input *string

	// Activity的名字
	Name string

	// 下一个Activity的名字
	Next string

	// Activity的输出，详见Param
	Output []string

	// Activity的重试次数
	RetryTimes string

	// Activity的超时时间
	Timeout string

	// Activity的类型
	Type string
}
