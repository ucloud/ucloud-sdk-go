package usms

/*
ReceiptPerPhone - 每个目的手机号的发送回执信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ReceiptPerPhone struct {

	// 手机号码
	Phone string

	// 消耗短信条数
	CostCount int

	// 回执结果
	ReceiptResult string

	// 回执结果描述
	ReceiptDesc string

	// 回执返回时间
	ReceiptTime int
}
