package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
)

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	saveReport()
	os.Exit(exitCode)
}

var ucloudstackClient *ucloudstack.UCloudStackClient
var spec = &driver.Specification{}

func testSetup() {

	cfg := ucloud.NewConfig()
	cfg.MaxRetries = 1
	cfg.LogLevel = log.DebugLevel
	cfg.Region = os.Getenv("UCLOUD_REGION")
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	// cfg.SetActionLevel("DescribeImage", log.WarnLevel)
	cfg.SetActionLevel("GetRegion", log.WarnLevel)
	cfg.SetActionLevel("DescribeUDBParamGroup", log.WarnLevel)

	// setup Specification
	spec.Config = &cfg
	spec.Credential = &credential

	log.Info("setup test fixtures ...")

	ustackCfg := ucloud.NewConfig()
	ustackCfg.BaseUrl = "http://console.pre.ucloudstack.com/api"
	ustackCfg.MaxRetries = 1
	ustackCfg.LogLevel = log.DebugLevel
	ustackCfg.Region = "cn"

	ustackCred := auth.NewCredential()
	ustackCred.PrivateKey = os.Getenv("UCLOUDSTACK_PRIVATE_KEY")
	ustackCred.PublicKey = os.Getenv("UCLOUDSTACK_PUBLIC_KEY")

	ucloudstackClient = ucloudstack.NewClient(&ustackCfg, &ustackCred)
}

func saveReport() {
	data, err := json.Marshal(spec.Report())
	if err != nil {
		log.Warn("cannot mae")
	}

	err = ioutil.WriteFile("./report.json", data, 0755)
	if err != nil {
		panic(err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("UCLOUD_PUBLIC_KEY"); v == "" {
		t.Fatal("UCLOUD_PUBLIC_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("UCLOUD_PRIVATE_KEY"); v == "" {
		t.Fatal("UCLOUD_PRIVATE_KEY must be set for acceptance tests")
	}

	if !driver.IsAcc() {
		t.Skip(driver.ACC_SKIP_REASON)
	}

}
