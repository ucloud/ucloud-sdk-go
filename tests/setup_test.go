package tests

import (
	"os"
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	iudataark "github.com/ucloud/ucloud-sdk-go/internal/services/udataark"
	iudisk "github.com/ucloud/ucloud-sdk-go/internal/services/udisk"
	iuhost "github.com/ucloud/ucloud-sdk-go/internal/services/uhost"
	iulb "github.com/ucloud/ucloud-sdk-go/internal/services/ulb"
	iumon "github.com/ucloud/ucloud-sdk-go/internal/services/umon"
	iunet "github.com/ucloud/ucloud-sdk-go/internal/services/unet"

	"github.com/ucloud/ucloud-sdk-go/services/pathx"
	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"
)

var client *ucloud.Client
var uhostClient *uhost.UHostClient
var unetClient *unet.UNetClient
var ulbClient *ulb.ULBClient
var vpcClient *vpc.VPCClient
var uaccountClient *uaccount.UAccountClient
var pathxClient *pathx.PathXClient
var udiskClient *udisk.UDiskClient

var iuhostClient *iuhost.UHostClient
var iunetClient *iunet.UNetClient
var iulbClient *iulb.ULBClient
var iudiskClient *iudisk.UDiskClient
var iumonClient *iumon.UMonClient
var iudataarkClient *iudataark.UDataArkClient

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	testTeardown()
	os.Exit(exitCode)
}

func testSetup() {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = os.Getenv("UCLOUD_REGION")
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	client = ucloud.NewClient(&cfg, &credential)
	uhostClient = uhost.NewClient(&cfg, &credential)
	unetClient = unet.NewClient(&cfg, &credential)
	ulbClient = ulb.NewClient(&cfg, &credential)
	vpcClient = vpc.NewClient(&cfg, &credential)
	uaccountClient = uaccount.NewClient(&cfg, &credential)
	pathxClient = pathx.NewClient(&cfg, &credential)
	udiskClient = udisk.NewClient(&cfg, &credential)

	iudataarkClient = iudataark.NewClient(&cfg, &credential)
	iudiskClient = iudisk.NewClient(&cfg, &credential)
	iuhostClient = iuhost.NewClient(&cfg, &credential)
	iulbClient = iulb.NewClient(&cfg, &credential)
	iumonClient = iumon.NewClient(&cfg, &credential)
	iunetClient = iunet.NewClient(&cfg, &credential)

	log.Info("setup test fixtures ...")
}

func testTeardown() {}
