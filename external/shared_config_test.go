package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
