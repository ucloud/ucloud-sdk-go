package main

import (
	"os"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
)

func main() {
	cfg := ucloud.NewConfig()

	// set common region and project id for any request
	// you can overwrite it by request
	cfg.Region = "cn-bj2"
	cfg.ProjectId = "xxx"

	// enable auto-retry for any request, default is disabled(max retries is 0)
	// it will auto retry for
	//   * network error
	//   * server error with http status 429, 502, 503 and 504
	cfg.MaxRetries = 3

	// disable sdk log, default is log.InfoLevel
	cfg.LogLevel = log.PanicLevel

	// enable sdk debug log level
	// if debug log is enable, it will print all of request params and response body
	cfg.LogLevel = log.DebugLevel

	// set timeout for any request, default is 30s
	cfg.Timeout = 30 * time.Second

	// the followed is used for private user and partner
	// should not be set in general usage

	// custom User-Agent for any request
	cfg.UserAgent = "UCloud-CLI/0.1.0"

	cred := auth.NewCredential()

	// set credential info for any request
	// it is required
	cred.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")
	cred.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
}
