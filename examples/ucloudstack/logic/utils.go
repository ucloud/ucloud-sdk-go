package logic

import (
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// LoadUcloudStackConfig 加载配置
func LoadUcloudStackConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.BaseUrl = "http://console.dev.ucloudstack.com/api"

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUDSTACK_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUDSTACK_PUBLIC_KEY")

	return &cfg, &credential
}
