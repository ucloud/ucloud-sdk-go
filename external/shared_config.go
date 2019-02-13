package external

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
)

// DefaultProfile is the default named profile for ucloud sdk
const DefaultProfile = "default"

// DefaultSharedConfigFile will return the default shared config filename
func DefaultSharedConfigFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "config.json")
}

// DefaultSharedCredentialsFile will return the default shared credential filename
func DefaultSharedCredentialsFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "credential.json")
}

// LoadUCloudConfigFile will load ucloud client config from config file
func LoadUCloudConfigFile(cfgFile, profile string) (*ucloud.Config, error) {
	if len(profile) == 0 {
		return nil, fmt.Errorf("expected ucloud named profile is not empty")
	}

	cfgMaps, err := loadConfigFile(cfgFile)
	if err != nil {
		return nil, err
	}

	c := getSharedConfig(cfgMaps, profile)
	return c.Config(), nil
}

// LoadUCloudCredentialFile will load ucloud credential config from config file
func LoadUCloudCredentialFile(credFile, profile string) (*auth.Credential, error) {
	if len(profile) == 0 {
		return nil, fmt.Errorf("expected ucloud named profile is not empty")
	}

	credMaps, err := loadCredFile(credFile)
	if err != nil {
		return nil, err
	}

	c := getSharedCredential(credMaps, profile)
	return c.Credential(), nil
}

type sharedConfig struct {
	ProjectID string `json:"project_id"`
	Region    string `json:"region"`
	Zone      string `json:"zone"`
	BaseURL   string `json:"base_url"`
	Timeout   int    `json:"timeout_sec"`
	Profile   string `json:"profile"`
	Active    bool   `json:"active"`
}

type sharedCredential struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Profile    string `json:"profile"`
}

func loadConfigFile(cfgFile string) ([]sharedConfig, error) {
	realCfgFile := cfgFile
	cfgMaps := make([]sharedConfig, 0)

	// try to load default config
	if len(realCfgFile) == 0 {
		realCfgFile = DefaultSharedConfigFile()
	}

	// load config file
	err := loadJSONFile(realCfgFile, &cfgMaps)
	if err != nil {
		// skip error for loading default config
		if len(cfgFile) == 0 {
			log.Debugf("config file is empty")
		} else {
			return nil, err
		}
	}

	return cfgMaps, nil
}

func loadCredFile(credFile string) ([]sharedCredential, error) {
	realCredFile := credFile
	credMaps := make([]sharedCredential, 0)

	// try to load default credential
	if len(credFile) == 0 {
		realCredFile = DefaultSharedCredentialsFile()
	}

	// load credential file
	err := loadJSONFile(realCredFile, &credMaps)
	if err != nil {
		// skip error for loading default credential
		if len(credFile) == 0 {
			log.Debugf("credential file is empty")
		} else {
			return nil, err
		}
	}

	return credMaps, nil
}

func loadSharedConfigFile(cfgFile, credFile, profile string) (*config, error) {
	cfgMaps, err := loadConfigFile(cfgFile)
	if err != nil {
		return nil, err
	}

	credMaps, err := loadCredFile(credFile)
	if err != nil {
		return nil, err
	}

	// load configured profile
	if len(profile) == 0 {
		profile = DefaultProfile
	}

	c := &config{
		Profile:              profile,
		SharedConfigFile:     cfgFile,
		SharedCredentialFile: credFile,
	}
	c.merge(getSharedConfig(cfgMaps, profile))
	c.merge(getSharedCredential(credMaps, profile))

	return c, nil
}

func getSharedConfig(cfgMaps []sharedConfig, profile string) *config {
	cfg := &sharedConfig{}

	for i := 0; i < len(cfgMaps); i++ {
		if cfgMaps[i].Profile == profile {
			cfg = &cfgMaps[i]
		}
	}

	return &config{
		ProjectId: cfg.ProjectID,
		Region:    cfg.Region,
		Zone:      cfg.Zone,
		BaseUrl:   cfg.BaseURL,
		Timeout:   time.Duration(cfg.Timeout) * time.Second,
	}
}

func getSharedCredential(credMaps []sharedCredential, profile string) *config {
	cred := &sharedCredential{}

	for i := 0; i < len(credMaps); i++ {
		if credMaps[i].Profile == profile {
			cred = &credMaps[i]
		}
	}

	return &config{
		PublicKey:  cred.PublicKey,
		PrivateKey: cred.PrivateKey,
	}
}
