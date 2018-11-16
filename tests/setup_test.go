package tests

import (
	"os"
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/pathx"
	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/ubill"
	"github.com/ucloud/ucloud-sdk-go/services/udb"
	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/services/udpn"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/services/umem"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"

	iubill "github.com/ucloud/ucloud-sdk-go/internal/services/ubill"
	iudataark "github.com/ucloud/ucloud-sdk-go/internal/services/udataark"
	iudb "github.com/ucloud/ucloud-sdk-go/internal/services/udb"
	iudisk "github.com/ucloud/ucloud-sdk-go/internal/services/udisk"
	iudpn "github.com/ucloud/ucloud-sdk-go/internal/services/udpn"
	iuhost "github.com/ucloud/ucloud-sdk-go/internal/services/uhost"
	iulb "github.com/ucloud/ucloud-sdk-go/internal/services/ulb"
	iumem "github.com/ucloud/ucloud-sdk-go/internal/services/umem"
	iumon "github.com/ucloud/ucloud-sdk-go/internal/services/umon"
	iunet "github.com/ucloud/ucloud-sdk-go/internal/services/unet"
)

var config *ucloud.Config
var client *ucloud.Client
var uhostClient *uhost.UHostClient
var unetClient *unet.UNetClient
var ulbClient *ulb.ULBClient
var vpcClient *vpc.VPCClient
var uaccountClient *uaccount.UAccountClient
var pathxClient *pathx.PathXClient
var udiskClient *udisk.UDiskClient
var udbClient *udb.UDBClient
var umemClient *umem.UMemClient
var udpnClient *udpn.UDPNClient
var ubillClient *ubill.UBillClient

var iuhostClient *iuhost.UHostClient
var iunetClient *iunet.UNetClient
var iulbClient *iulb.ULBClient
var iudiskClient *iudisk.UDiskClient
var iumonClient *iumon.UMonClient
var iudataarkClient *iudataark.UDataArkClient
var iudbClient *iudb.UDBClient
var iumemClient *iumem.UMemClient
var iudpnClient *iudpn.UDPNClient
var iubillClient *iubill.UBillClient

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	testTeardown()
	os.Exit(exitCode)
}

func testSetup() {
	cfg := ucloud.NewConfig()
	cfg.MaxRetries = 1
	cfg.LogLevel = log.DebugLevel
	cfg.Region = os.Getenv("UCLOUD_REGION")
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	config = &cfg

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
	udbClient = udb.NewClient(&cfg, &credential)
	umemClient = umem.NewClient(&cfg, &credential)
	udpnClient = udpn.NewClient(&cfg, &credential)
	ubillClient = ubill.NewClient(&cfg, &credential)

	iudataarkClient = iudataark.NewClient(&cfg, &credential)
	iudiskClient = iudisk.NewClient(&cfg, &credential)
	iuhostClient = iuhost.NewClient(&cfg, &credential)
	iulbClient = iulb.NewClient(&cfg, &credential)
	iumonClient = iumon.NewClient(&cfg, &credential)
	iunetClient = iunet.NewClient(&cfg, &credential)
	iudbClient = iudb.NewClient(&cfg, &credential)
	iumemClient = iumem.NewClient(&cfg, &credential)
	iudpnClient = iudpn.NewClient(&cfg, &credential)
	iubillClient = iubill.NewClient(&cfg, &credential)

	log.Info("setup test fixtures ...")
}

func testTeardown() {}
