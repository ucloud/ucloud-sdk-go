// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

/*
MirrorCondition - 回源条件

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type MirrorCondition struct {

	// 回源条件，path以某个prefix为前缀
	Prefix string

	// 状态码 404 表示回源
	StatusCode int
}
