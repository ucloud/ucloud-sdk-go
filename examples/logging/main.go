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
	req := uhostClient.NewDescribeUHostInstanceRequest()
	req.Region = ucloud.String(region)

	// excepted logging
	uhostClient.DescribeUHostInstance(req)
	cfg.SetActionLevel("DescribeImage", log.WarnLevel)

	noLogReq := uhostClient.NewDescribeImageRequest()
	noLogReq.Region = ucloud.String(region)

	// unexcepted logging
	uhostClient.DescribeImage(noLogReq)
}
