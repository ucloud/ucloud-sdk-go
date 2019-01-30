package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultTestSharedConfigFile() *sharedConfigFile {
	return &sharedConfigFile{
		ConfigFile:     TestValueUCloudSharedConfigFile,
		CredentialFile: TestValueUCloudSharedCredentialFile,
	}
}

func TestSharedConfigFile_active(t *testing.T) {
	scf := defaultTestSharedConfigFile()
	scf.Profile = TestValueSharedProfile

	cfg, cred, err := scf.Load()
	assert.NoError(t, err)
	assert.Equal(t, TestValueUCloudPublicKey, cred.PublicKey)
	assert.Equal(t, TestValueUCloudPrivateKey, cred.PrivateKey)
	assert.Equal(t, TestValueUCloudProjectId, cfg.ProjectID)
	assert.Equal(t, TestValueUCloudRegion, cfg.Region)
	assert.Equal(t, TestValueSharedProfile, cfg.Profile)
	assert.Equal(t, 15, cfg.Timeout)
	assert.Equal(t, TestValueUCloudBaseUrl, cfg.BaseURL)
	assert.Equal(t, TestValueUCloudZone, cfg.Zone)

	profile, err := scf.LoadProfile()
	assert.NoError(t, err)
	assert.Equal(t, TestValueSharedProfile, profile)
}

func TestSharedConfigFile_multi_active(t *testing.T) {
	// load default config file
	scf := defaultTestSharedConfigFile()
	_, _, err := scf.Load()
	assert.NoError(t, err)

	// overwrite config with multiple active configuration
	scf.configMaps[0].Active = true
	scf.configMaps[1].Active = true
	cfgPath, err := writeTestTempConfigFile(scf.configMaps)
	assert.NoError(t, err)

	// validate with error
	scf.ConfigFile = cfgPath
	_, _, err = scf.Load()
	assert.Error(t, err)
}

func TestSharedConfigFile_no_active(t *testing.T) {
	// load default config file
	scf := defaultTestSharedConfigFile()
	_, _, err := scf.Load()
	assert.NoError(t, err)

	// overwrite config with multiple active configuration
	scf.configMaps[0].Active = false
	scf.configMaps[1].Active = false
	cfgPath, err := writeTestTempConfigFile(scf.configMaps)
	assert.NoError(t, err)

	// validate with error
	scf.ConfigFile = cfgPath
	_, _, err = scf.Load()
	assert.Error(t, err)
}

func TestSharedConfigFile_unique_profile(t *testing.T) {
	// load default config file
	scf := defaultTestSharedConfigFile()
	_, _, err := scf.Load()
	assert.NoError(t, err)

	// overwrite config with multiple active configuration
	scf.configMaps[0].Profile = TestValueSharedProfile
	scf.configMaps[1].Profile = TestValueSharedProfile
	cfgPath, err := writeTestTempConfigFile(scf.configMaps)
	assert.NoError(t, err)

	// validate with error
	scf.ConfigFile = cfgPath
	_, _, err = scf.Load()
	assert.Error(t, err)
}

func TestLoadSharedConfigFile_not_exists(t *testing.T) {
	// load default config file
	scf := defaultTestSharedConfigFile()
	scf.ConfigFile = "not_exists"
	_, _, err := scf.Load()
	assert.Error(t, err)

	p, err := LoadSharedConfigFile(scf.ConfigFile, scf.CredentialFile, scf.Profile)
	assert.NoError(t, err)
	assert.Nil(t, p)
}
