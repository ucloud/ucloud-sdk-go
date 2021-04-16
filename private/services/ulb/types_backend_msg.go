package ulb

/*
   BackendMsg -
*/
type BackendMsg struct {

	// rs的资源ID
	BackendId string `required:"true"`
	// 修改rs返回的消息
	SubMessage string `required:"true"`
	// 修改rs的返回值
	SubRetCode int `required:"true"`
}
