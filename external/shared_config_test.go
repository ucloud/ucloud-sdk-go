package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharedConfigFile(t *testing.T) {
	type args struct {
	}

	tests := []struct {
		name    string
		prepare func([]sharedConfig)
		wantErr bool
	}{
		{
			"oneActive",
			func(cfgMaps []sharedConfig) {
				cfgMaps[0].Active = false
				cfgMaps[1].Active = true
			},
			false,
		},
		{
			"multiActive",
			func(cfgMaps []sharedConfig) {
				cfgMaps[0].Active = true
				cfgMaps[1].Active = true
			},
			true,
		},
		{
			"noActive",
			func(cfgMaps []sharedConfig) {
				cfgMaps[0].Active = false
				cfgMaps[1].Active = false
			},
			true,
		},
		{
			"uniqueProfile",
			func(cfgMaps []sharedConfig) {
				cfgMaps[0].Profile = TestValueEnvUCloudProfile
				cfgMaps[1].Profile = TestValueEnvUCloudProfile
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfgMaps, _, err := loadFile(
				TestValueEnvUCloudSharedConfigFile,
				TestValueEnvUCloudSharedCredentialFile,
			)
			assert.NoError(t, err)

			tt.prepare(cfgMaps)

			cfgPath, err := writeTestTempConfigFile(cfgMaps)
			assert.NoError(t, err)

			_, err = loadSharedConfigFile(
				cfgPath,
				TestValueEnvUCloudSharedCredentialFile,
				"",
			)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLoadSharedConfigFile_not_exists(t *testing.T) {
	cfg, err := loadSharedConfigFile(
		"not_exists",
		TestValueEnvUCloudSharedCredentialFile,
		"",
	)
	assert.Nil(t, cfg)
	assert.Nil(t, err)

	_, err = loadSharedConfigFile(
		TestValueEnvUCloudSharedConfigFile,
		"",
		"",
	)
	assert.Nil(t, cfg)
	assert.Nil(t, err)
}
