package external

import (
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

const (
	UCloudPublicKeyEnvVar = "UCLOUD_PUBLIC_KEY"

	UCloudPrivateKeyEnvVar = "UCLOUD_PRIVATE_KEY"

	UCloudProjectIdEnvVar = "UCLOUD_PROJECT_ID"

	UCloudRegionEnvVar = "UCLOUD_REGION"

	UCloudSharedProfileEnvVar = "UCLOUD_PROFILE"

	UCloudSharedConfigFileEnvVar = "UCLOUD_SHARED_CONFIG_FILE"

	UCloudSharedCredentialFileEnvVar = "UCLOUD_SHARED_CREDENTIAL_FILE"
)

// EnvConfig is the environment configuration
type EnvConfig struct {
	PublicKey            string
	PrivateKey           string
	ProjectId            string
	Region               string
	Profile              string
	SharedConfigFile     string
	SharedCredentialFile string
}

// LoadEnvConfig will return a default environment configuration
func LoadEnvConfig(cfgs ...Config) (ConfigProvider, error) {
	ec := &EnvConfig{}

	ec.PublicKey = os.Getenv(UCloudPublicKeyEnvVar)
	ec.PrivateKey = os.Getenv(UCloudPrivateKeyEnvVar)
	ec.ProjectId = os.Getenv(UCloudProjectIdEnvVar)
	ec.Region = os.Getenv(UCloudRegionEnvVar)
	ec.Profile = os.Getenv(UCloudSharedProfileEnvVar)
	ec.SharedConfigFile = os.Getenv(UCloudSharedConfigFileEnvVar)
	ec.SharedCredentialFile = os.Getenv(UCloudSharedCredentialFileEnvVar)

	return ec, nil
}

// Credential is the configuration of ucloud authorization information
func (ec *EnvConfig) Credential() *auth.Credential {
	return &auth.Credential{
		PublicKey:  ec.PublicKey,
		PrivateKey: ec.PrivateKey,
	}
}

// Config is the configuration of ucloud client
func (ec *EnvConfig) Config() *ucloud.Config {
	return &ucloud.Config{
		Region:    ec.Region,
		ProjectId: ec.ProjectId,
	}
}

// SharedConfigProfile will return the profile name
func (ec *EnvConfig) SharedConfigProfile() string {
	return ec.Profile
}

// SharedConfigFilename will return the file name of shared config
func (ec *EnvConfig) SharedConfigFilename() string {
	return ec.SharedConfigFile
}

// SharedCredentialFilename will return the file name of shared credential
func (ec *EnvConfig) SharedCredentialFilename() string {
	return ec.SharedCredentialFile
}
