package ufs

/*
VPCInfo - vpc信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn for detail.
*/
type VPCInfo struct {

	// Vpc ID
	VpcId string

	// subnet信息
	SubnetIds []SubnetFrontEnd

	// 是否创建subnet，如果为false，则提示用户创建subnet。如果为true，同时SubnetIds为空，则表明当前vpc下已经没有subnet可以挂载
	SubnetCreated bool
}
