package external

import (
	"fmt"
	"reflect"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

var configLoaders = []ConfigLoader{
	LoadEnvConfig,
	LoadSharedConfig,
}

// Config is the configuration block from any config loader
type Config interface{}

// ConfigLoader is the external configuration loader to load credential and client config from external storag
type ConfigLoader func(...Config) (ConfigProvider, error)

// ConfigProvider is the provider to store and provide config/credential instance
type ConfigProvider interface {
	Credential() *auth.Credential

	Config() *ucloud.Config
}

// ProfileProvider is the interface to provide named profile
type ProfileProvider interface {
	SharedConfigProfile() string
}

// ConfigFileProvider is the
type ConfigFileProvider interface {
	SharedConfigFilename() string

	SharedCredentialFilename() string
}

// LoadDefaultUCloudConfig is the default loader to load config
func LoadDefaultUCloudConfig() (ConfigProvider, error) {
	cfgs := make([]Config, len(configLoaders))
	for _, load := range configLoaders {
		cfg, err := load(cfgs...)
		if err != nil {
			return nil, err
		}

		if !reflect.ValueOf(cfg).IsNil() {
			cfgs = append(cfgs, cfg)
		}
	}

	if len(cfgs) == 0 {
		return NewNullConfig(), nil
	}

	if cfg, ok := cfgs[len(cfgs)-1].(ConfigProvider); !ok {
		return nil, fmt.Errorf("excepted the configuration from external storage should has .Client() and .Credential() method")
	} else {
		return cfg, nil
	}
}
