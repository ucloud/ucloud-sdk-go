package main

import (
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/ulb"
)

const region = "cn-bj2"

func main() {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = region
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	client := ulb.NewClient(&cfg, &credential)

	req := client.NewDescribeULBRequest()
	req.ULBId = ucloud.String("ulb-not-found-for-retry")
	req.WithRetry(3)

	_, err := client.DescribeULB(req)
	if err != nil {
		log.Error(err)
	}
}
