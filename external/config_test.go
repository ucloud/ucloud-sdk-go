package external

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	TestValueUCloudPublicKey            = "f05816ca02feec1b3fc38b80a1c450cc"
	TestValueUCloudPrivateKey           = "c45f9bec5fa4c6c47fd871fadd97dd2e"
	TestValueUCloudProjectId            = "org-3kopqz"
	TestValueUCloudRegion               = "cn-bj2"
	TestValueUCloudZone                 = "cn-bj2-02"
	TestValueSharedProfile              = "default"
	TestValueUCloudBaseUrl              = "https://api.ucloud.cn"
	TestValueUCloudTimeout              = time.Duration(15) * time.Second
	TestValueUCloudSharedConfigFile     = filepath.Join("test-fixtures", "config.json")
	TestValueUCloudSharedCredentialFile = filepath.Join("test-fixtures", "credential.json")
)

func TestLoadConfig(t *testing.T) {
	setTestEnv()
	c, err := LoadDefaultUCloudConfig()
	assert.NoError(t, err)

	cfg, cred := c.Config(), c.Credential()
	assert.Equal(t, TestValueUCloudPublicKey, cred.PublicKey)
	assert.Equal(t, TestValueUCloudPrivateKey, cred.PrivateKey)

	assert.Equal(t, TestValueUCloudProjectId, cfg.ProjectId)
	assert.Equal(t, TestValueUCloudRegion, cfg.Region)
	assert.Equal(t, time.Duration(15)*time.Second, cfg.Timeout)
	assert.Equal(t, TestValueUCloudBaseUrl, cfg.BaseUrl)
	assert.Equal(t, TestValueUCloudZone, cfg.Zone)
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
