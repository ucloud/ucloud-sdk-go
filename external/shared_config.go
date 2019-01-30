package external

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// DefaultSharedConfigFile will return the default shared config filename
func DefaultSharedConfigFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "config.json")
}

// DefaultSharedCredentialsFile will return the default shared credential filename
func DefaultSharedCredentialsFile() string {
	return filepath.Join(userHomeDir(), ".ucloud", "credential.json")
}

// SharedConfig will read configuration
type SharedConfig struct {
	// Named profile
	Profile string

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

// LoadSharedConfig will return a shared configuration
func LoadSharedConfig(cfgs ...Config) (ConfigProvider, error) {
	scf := newSharedConfigFile(cfgs...)
	return loadSharedConfigFile(scf)
}

// LoadSharedConfigFile will return a shared configuration
func LoadSharedConfigFile(configFile, credentialFile, profile string) (ConfigProvider, error) {
	scf := &sharedConfigFile{
		Profile:        profile,
		ConfigFile:     configFile,
		CredentialFile: credentialFile,
	}
	return loadSharedConfigFile(scf)
}

func loadSharedConfigFile(scf *sharedConfigFile) (*SharedConfig, error) {
	cfg, cred, err := scf.Load()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	sc := &SharedConfig{}
	sc.Profile = scf.Profile
	sc.Region = cfg.Region
	sc.Zone = cfg.Zone
	sc.ProjectId = cfg.ProjectID
	sc.BaseUrl = cfg.BaseURL
	sc.Timeout = time.Duration(cfg.Timeout) * time.Second
	sc.PublicKey = cred.PublicKey
	sc.PrivateKey = cred.PrivateKey
	return sc, nil
}

// Config is the configuration of ucloud client
func (sc *SharedConfig) Config() *ucloud.Config {
	return &ucloud.Config{
		ProjectId: sc.ProjectId,
		Zone:      sc.Zone,
		Region:    sc.Region,
		BaseUrl:   sc.BaseUrl,
		Timeout:   sc.Timeout,
	}
}

// Credential is the configuration of ucloud authorization information
func (sc *SharedConfig) Credential() *auth.Credential {
	return &auth.Credential{
		PublicKey:  sc.PublicKey,
		PrivateKey: sc.PrivateKey,
	}
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

type sharedConfigFile struct {
	Profile        string
	ConfigFile     string
	CredentialFile string

	configMaps []sharedConfig
	credMaps   []sharedCredential
}

func (scf *sharedConfigFile) Load() (*sharedConfig, *sharedCredential, error) {
	var err error

	// load config file
	err = loadJSONFile(scf.ConfigFile, &scf.configMaps)
	if err != nil {
		return nil, nil, err
	}

	// load credential file
	err = loadJSONFile(scf.CredentialFile, &scf.credMaps)
	if err != nil {
		return nil, nil, err
	}

	// validate config file
	profile, err := scf.LoadProfile()
	if err != nil {
		return nil, nil, err
	}

	cfg, err := scf.GetSharedConfig(profile)
	if err != nil {
		return nil, nil, err
	}

	cred, err := scf.GetSharedCredential(profile)
	if err != nil {
		return nil, nil, err
	}

	return cfg, cred, nil
}

func (scf *sharedConfigFile) LoadProfile() (string, error) {
	profiles := newSet(hashString, nil)
	var active sharedConfig
	var activeCount int
	for _, cfg := range scf.configMaps {
		if cfg.Active {
			activeCount++
			active = cfg
		}
		profiles.Add(cfg.Profile)
	}

	if profiles.Len() != len(scf.configMaps) {
		return "", fmt.Errorf("excepted the named profile is unique")
	}

	if activeCount > 1 {
		return "", fmt.Errorf("excepted only one named profile is actived, got %d", activeCount)
	}

	if activeCount == 0 {
		return "", fmt.Errorf("excepted exactly one named profile is actived, but not found")
	}

	// load configured profile
	profile := active.Profile
	if len(scf.Profile) != 0 {
		profile = scf.Profile
	}
	return profile, nil
}

func (scf *sharedConfigFile) GetSharedConfig(profile string) (*sharedConfig, error) {
	for _, cfg := range scf.configMaps {
		if cfg.Profile == profile {
			return &cfg, nil
		}
	}
	return nil, fmt.Errorf("excepted shared config with %s profile, but not found", profile)
}

func (scf *sharedConfigFile) GetSharedCredential(profile string) (*sharedCredential, error) {
	for _, cred := range scf.credMaps {
		if cred.Profile == profile {
			return &cred, nil
		}
	}
	return nil, fmt.Errorf("excepted shared credential with %s profile, but not found", profile)
}

func newSharedConfigFile(cfgs ...Config) *sharedConfigFile {
	scf := &sharedConfigFile{}

	// get default configurations
	for _, cfg := range cfgs {
		if v, ok := cfg.(ProfileProvider); ok {
			setStringify(&scf.Profile, v.SharedConfigProfile())
		}

		if v, ok := cfg.(ConfigFileProvider); ok {
			setStringify(&scf.ConfigFile, v.SharedConfigFilename())
			setStringify(&scf.CredentialFile, v.SharedCredentialFilename())
		}
	}

	// if attributes still empty, set the default values
	if len(scf.ConfigFile) == 0 {
		scf.ConfigFile = DefaultSharedConfigFile()
	}

	if len(scf.CredentialFile) == 0 {
		scf.CredentialFile = DefaultSharedCredentialsFile()
	}

	scf.configMaps = make([]sharedConfig, 0)
	scf.credMaps = make([]sharedCredential, 0)
	return scf
}

func setStringify(p *string, s string) {
	if len(s) != 0 {
		*p = s
	}
}
