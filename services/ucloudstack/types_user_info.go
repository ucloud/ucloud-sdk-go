// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

/*
UserInfo - 租户信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type UserInfo struct {

	// 账户余额
	Amount float64

	// 账户创建时间。时间戳
	CreateTime int

	// 租户名称
	Name string

	// 私钥
	PrivateKey string

	// 公钥
	PublicKey string

	// 用户状态。USER_STATUS_AVAILABLE：正常，USER_STATUS_FREEZE：冻结
	Status string

	// 更新时间。时间戳
	UpdateTime int

	// 租户ID.
	UserID int
}
