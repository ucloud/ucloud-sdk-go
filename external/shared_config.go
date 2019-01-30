package external

import (
	"fmt"
)

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

func loadSharedConfigFile(cfgFile string, credFile string, profile string) (*sharedConfig, *sharedCredential, error) {
	var err error

	cfgMaps := make([]*sharedConfig, 0)
	credMaps := make([]*sharedCredential, 0)

	// load config file
	err = loadJSONFile(cfgFile, cfgMaps)
	if err != nil {
		return nil, nil, err
	}

	// load credential file
	err = loadJSONFile(credFile, credMaps)
	if err != nil {
		return nil, nil, err
	}

	// validate config file
	activeProfile, err := validateConfig(cfgMaps)
	if err != nil {
		return nil, nil, err
	}

	// load configured profile
	if len(profile) == 0 {
		profile = activeProfile
	}

	var cfg *sharedConfig
	for _, ptr := range cfgMaps {
		if ptr.Profile == profile {
			cfg = ptr
		}
	}

	var cred *sharedCredential
	for _, ptr := range credMaps {
		if ptr.Profile == profile {
			cred = ptr
		}
	}

	return cfg, cred, nil
}

func validateConfig(cfgMaps []*sharedConfig) (string, error) {
	profiles := newSet(hashString, nil)
	var active *sharedConfig
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
		return "", fmt.Errorf("excepted only one named profile is actived, got %d", activeCount)
	}

	if activeCount == 0 {
		return "", fmt.Errorf("excepted exactly one named profile is actived, but not found")
	}

	return active.Profile, nil
}
