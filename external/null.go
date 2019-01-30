package external

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// NullConfig is an empty configuration
type NullConfig struct{}

// NewNullConfig will return an empty null config
func NewNullConfig() *NullConfig {
	return &NullConfig{}
}

// Credential is the configuration of ucloud authorization information
func (nc *NullConfig) Credential() *auth.Credential {
	cred := auth.NewCredential()
	return &cred
}

// Config is the configuration of ucloud client
func (nc *NullConfig) Config() *ucloud.Config {
	cfg := ucloud.NewConfig()
	return &cfg
}
