package external

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// ConfigProvider is the provider to store and provide config/credential instance
type ConfigProvider interface {
	Credential() *auth.Credential

	Config() *ucloud.Config
}
