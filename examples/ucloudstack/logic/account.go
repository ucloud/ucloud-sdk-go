package logic

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// LoginByPassword 通过密码登录
func LoginByPassword(email, password string) {
	cfg, credential := LoadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	loginByPasswordReq := ucloudstackClient.NewLoginByPasswordRequest()
	loginByPasswordReq.UserEmail = ucloud.String(email)
	loginByPasswordReq.Password = ucloud.String(password)

	loginByPasswordResp, err := ucloudstackClient.LoginByPassword(loginByPasswordReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n, message: %s", err, loginByPasswordResp.Message)
		return
	}

	fmt.Printf("Recharge success")
}
