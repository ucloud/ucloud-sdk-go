package external

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// DefaultSharedConfigFile will return the default shared config filename
func DefaultSharedConfigFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "config.json")
}

// DefaultSharedCredentialsFile will return the default shared credential filename
func DefaultSharedCredentialsFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "credential.json")
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

func loadFile(cfgFile, credFile string) ([]sharedConfig, []sharedCredential, error) {
	var err error
	cfgMaps := make([]sharedConfig, 0)
	credMaps := make([]sharedCredential, 0)

	// load config file
	err = loadJSONFile(cfgFile, &cfgMaps)
	if err != nil {
		return nil, nil, err
	}

	// load credential file
	err = loadJSONFile(credFile, &credMaps)
	if err != nil {
		return nil, nil, err
	}

	return cfgMaps, credMaps, nil
}

func loadSharedConfigFile(cfgFile, credFile, profile string) (*config, error) {
	cfgMaps, credMaps, err := loadFile(cfgFile, credFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	// validate config file
	activeProfile, err := validateConfig(cfgMaps)
	if err != nil {
		return nil, err
	}

	// load configured profile
	if len(profile) == 0 {
		profile = activeProfile
	}

	var cfg *sharedConfig
	for i := 0; i < len(cfgMaps); i++ {
		if cfgMaps[i].Profile == profile {
			cfg = &cfgMaps[i]
		}
	}

	var cred *sharedCredential
	for i := 0; i < len(credMaps); i++ {
		if credMaps[i].Profile == profile {
			cred = &credMaps[i]
		}
	}

	if cfg == nil {
		return nil, fmt.Errorf("excepted shared config with %s profile, but not found", profile)
	}

	if cred == nil {
		return nil, fmt.Errorf("excepted shared credential with %s profile, but not found", profile)
	}

	return &config{
		PublicKey:            cred.PublicKey,
		PrivateKey:           cred.PrivateKey,
		ProjectId:            cfg.ProjectID,
		Region:               cfg.Region,
		Zone:                 cfg.Zone,
		BaseUrl:              cfg.BaseURL,
		Timeout:              time.Duration(cfg.Timeout) * time.Second,
		Profile:              profile,
		SharedConfigFile:     cfgFile,
		SharedCredentialFile: credFile,
	}, nil
}

func validateConfig(cfgMaps []sharedConfig) (string, error) {
	profiles := newSet(hashString, nil)
	var active sharedConfig
	var activeCount int
	for _, cfg := range cfgMaps {
		if cfg.Active {
			activeCount++
			active = cfg
		}
		profiles.Add(cfg.Profile)
	}

	if profiles.Len() != len(cfgMaps) {
		return "", fmt.Errorf("excepted the named profile is unique")
	}

	if activeCount > 1 {
		return "", fmt.Errorf("excepted exactly one named profile is actived, got %d", activeCount)
	}

	if activeCount == 0 {
		return "", fmt.Errorf("excepted exactly one named profile is actived, but not found")
	}

	return active.Profile, nil
}
