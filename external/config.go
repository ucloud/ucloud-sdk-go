package external

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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

// DefaultSharedConfigFile will return the default shared config filename
func DefaultSharedConfigFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "config.json")
}

// DefaultSharedCredentialsFile will return the default shared credential filename
func DefaultSharedCredentialsFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "credential.json")
}

// config will read configuration
type config struct {
	// Named profile
	profile              string
	sharedConfigFile     string
	sharedCredentialFile string

	// Credential configuration
	PublicKey  string
	PrivateKey string

	// Client configuration
	ProjectId string
	Zone      string
	Region    string
	BaseUrl   string
	Timeout   time.Duration
}

func newConfig() *config {
	return &config{}
}

func (c *config) loadEnv() error {
	c.PublicKey = os.Getenv(UCloudPublicKeyEnvVar)
	c.PrivateKey = os.Getenv(UCloudPrivateKeyEnvVar)
	c.ProjectId = os.Getenv(UCloudProjectIdEnvVar)
	c.Region = os.Getenv(UCloudRegionEnvVar)

	c.profile = os.Getenv(UCloudSharedProfileEnvVar)
	c.sharedConfigFile = os.Getenv(UCloudSharedConfigFileEnvVar)
	c.sharedCredentialFile = os.Getenv(UCloudSharedCredentialFileEnvVar)
	return nil
}

func (c *config) loadFile() error {
	return nil
}

// Config is the configuration of ucloud client
func (c *config) Config() *ucloud.Config {
	return &ucloud.Config{
		ProjectId: c.ProjectId,
		Zone:      c.Zone,
		Region:    c.Region,
		BaseUrl:   c.BaseUrl,
		Timeout:   c.Timeout,
	}
}

// Credential is the configuration of ucloud authorization information
func (c *config) Credential() *auth.Credential {
	return &auth.Credential{
		PublicKey:  c.PublicKey,
		PrivateKey: c.PrivateKey,
	}
}

// LoadDefaultUCloudConfig is the default loader to load config
func LoadDefaultUCloudConfig() (ConfigProvider, error) {
	cfg := newConfig()

	if err := cfg.loadEnv(); err != nil {
		return nil, fmt.Errorf("error on loading env, %s", err)
	}

	if err := cfg.loadFile(); err != nil {
		return nil, fmt.Errorf("error on loading shared config file, %s", err)
	}

	return cfg, nil
}

func setStringify(p *string, s string) {
	if len(s) != 0 {
		*p = s
	}
}
