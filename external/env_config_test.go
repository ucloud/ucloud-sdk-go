package external

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvConfig(t *testing.T) {
	setTestEnv()

	c, err := LoadEnvConfig()
	cfg, cred := c.Config(), c.Credential()
	assert.NoError(t, err)
	assert.Equal(t, TestValueUCloudPublicKey, cred.PublicKey)
	assert.Equal(t, TestValueUCloudPrivateKey, cred.PrivateKey)
	assert.Equal(t, TestValueUCloudProjectId, cfg.ProjectId)
	assert.Equal(t, TestValueUCloudRegion, cfg.Region)

	cfProvider, ok := c.(ConfigFileProvider)
	assert.Equal(t, TestValueUCloudSharedConfigFile, cfProvider.SharedConfigFilename())
	assert.Equal(t, TestValueUCloudSharedCredentialFile, cfProvider.SharedCredentialFilename())

	profileProvider, ok := c.(ProfileProvider)
	assert.True(t, ok)
	assert.Equal(t, profileProvider.SharedConfigProfile(), TestValueSharedProfile)
}

func setTestEnv() {
	os.Setenv(UCloudPublicKeyEnvVar, TestValueUCloudPublicKey)
	os.Setenv(UCloudPrivateKeyEnvVar, TestValueUCloudPrivateKey)
	os.Setenv(UCloudProjectIdEnvVar, TestValueUCloudProjectId)
	os.Setenv(UCloudRegionEnvVar, TestValueUCloudRegion)
	os.Setenv(UCloudSharedProfileEnvVar, TestValueSharedProfile)
	os.Setenv(UCloudSharedConfigFileEnvVar, TestValueUCloudSharedConfigFile)
	os.Setenv(UCloudSharedCredentialFileEnvVar, TestValueUCloudSharedCredentialFile)
}
