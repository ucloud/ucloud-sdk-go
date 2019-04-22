package usms

/*
ReceiptPerSession - 每个提交的回执结果集合

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type ReceiptPerSession struct {

	// 发送短信时返回的SessionNo
	SessionNo string

	// 每个手机号的短信回执信息集合
	ReceiptSet []ReceiptPerPhone
}
