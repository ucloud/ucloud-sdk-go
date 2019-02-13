package external

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	TestValueEnvUCloudPublicKey            = "f05816ca02feec1b3fc38b80a1c450cc"
	TestValueEnvUCloudPrivateKey           = "c45f9bec5fa4c6c47fd871fadd97dd2e"
	TestValueEnvUCloudProjectId            = "org-3kopqz"
	TestValueEnvUCloudRegion               = "cn-bj2"
	TestValueEnvUCloudProfile              = "default"
	TestValueEnvUCloudSharedConfigFile     = filepath.Join("test-fixtures", "config.json")
	TestValueEnvUCloudSharedCredentialFile = filepath.Join("test-fixtures", "credential.json")

	TestValueFileUCloudDefaultPublicKey  = "f05816ca02feec1b3fc38b80a1c450cc"
	TestValueFileUCloudDefaultPrivateKey = "c45f9bec5fa4c6c47fd871fadd97dd2e"
	TestValueFileUCloudPublicKey         = "18ce27e79596c5adc986f66707b3fd55"
	TestValueFileUCloudPrivateKey        = "2c125558f9004b73751e3a4658e2f52b"
	TestValueFileUCloudProjectId         = "org-0qakla"
	TestValueFileUCloudRegion            = "cn-sh2"
	TestValueFileUCloudZone              = "cn-sh2-02"
	TestValueFileUCloudTimeout           = time.Duration(30) * time.Second
	TestValueFileUCloudBaseUrl           = "https://api.ucloud.cn"
)

func TestLoadConfig(t *testing.T) {
	setTestEnv()

	c, err := LoadDefaultUCloudConfig()
	assert.NoError(t, err)

	checkTestConfig(t, c)
}

func TestLoadEnvConfig(t *testing.T) {
	setTestEnv()

	c := &config{}
	c.loadEnv()

	checkTestEnvConfig(t, c)
}

func TestLoadSharedFile(t *testing.T) {
	c := &config{
		SharedConfigFile:     TestValueEnvUCloudSharedConfigFile,
		SharedCredentialFile: TestValueEnvUCloudSharedCredentialFile,
	}
	c.loadFile()

	checkTestConfig(t, c)
}

func checkTestEnvConfig(t *testing.T, c *config) {
	assert.Equal(t, TestValueEnvUCloudPublicKey, c.PublicKey)
	assert.Equal(t, TestValueEnvUCloudPrivateKey, c.PrivateKey)
	assert.Equal(t, TestValueEnvUCloudProjectId, c.ProjectId)
	assert.Equal(t, TestValueEnvUCloudRegion, c.Region)
	assert.Equal(t, TestValueEnvUCloudProfile, c.Profile)
	assert.Equal(t, TestValueEnvUCloudSharedConfigFile, c.SharedConfigFile)
	assert.Equal(t, TestValueEnvUCloudSharedCredentialFile, c.SharedCredentialFile)
}

func checkTestConfig(t *testing.T, c ConfigProvider) {
	cfg, cred := c.Config(), c.Credential()

	checkTestCredential(t, cred)
	checkTestClientConfig(t, cfg)
}

func setTestEnv() {
	os.Setenv(UCloudPublicKeyEnvVar, TestValueEnvUCloudPublicKey)
	os.Setenv(UCloudPrivateKeyEnvVar, TestValueEnvUCloudPrivateKey)
	os.Setenv(UCloudProjectIdEnvVar, TestValueEnvUCloudProjectId)
	os.Setenv(UCloudRegionEnvVar, TestValueEnvUCloudRegion)
	os.Setenv(UCloudSharedProfileEnvVar, TestValueEnvUCloudProfile)
	os.Setenv(UCloudSharedConfigFileEnvVar, TestValueEnvUCloudSharedConfigFile)
	os.Setenv(UCloudSharedCredentialFileEnvVar, TestValueEnvUCloudSharedCredentialFile)
}

func writeTestTempConfigFile(vL []sharedConfig) (string, error) {
	return writeTestTempFile(vL)
}

func writeTestTempCredentialFile(vL []sharedCredential) (string, error) {
	return writeTestTempFile(vL)
}

func writeTestTempFile(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	f, err := ioutil.TempFile("", "ucloud-sdk-go-conf")
	if err != nil {
		return "", err
	}

	_, err = f.Write(bs)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
