package main

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/external"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

func printMetaHandler(c *ucloud.Client, req request.Common, resp response.Common, err error) (response.Common, error) {
	// 静态信息，由 SDK 注入到 client 中
	fmt.Println("[INFO] Product: ", c.GetMeta().Product)

	return resp, err
}

func main() {
	c, err := external.LoadDefaultUCloudConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Config: %+v\n", c)

	cfg, cred := c.Config(), c.Credential()
	cfg.Region = "cn-bj2"

	uhostClient := uhost.NewClient(cfg, cred)
	_ = uhostClient.AddResponseHandler(printMetaHandler)

	req := uhostClient.NewDescribeUHostInstanceRequest()
	resp, err := uhostClient.DescribeUHostInstance(req)
	if err != nil {
		print(err)
	}
	fmt.Printf("%#v\n", resp)
}
