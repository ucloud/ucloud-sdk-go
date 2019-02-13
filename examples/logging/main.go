package main

import (
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
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

	log.Info("setup clients ...")

	uhostClient := uhost.NewClient(&cfg, &credential)

	// included logging
	req1 := uhostClient.NewDescribeUHostInstanceRequest()
	req1.Region = ucloud.String(region)

	uhostClient.DescribeUHostInstance(req1)

	// excepted logging
	cfg.SetActionLevel("DescribeImage", log.WarnLevel)
	req2 := uhostClient.NewDescribeImageRequest()
	req2.Region = ucloud.String(region)

	uhostClient.DescribeImage(req2)
}
