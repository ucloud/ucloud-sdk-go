package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/external"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
)

func main() {
	c, err := external.LoadDefaultUCloudConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Config: %+v\n", c)

	cfg, cred := c.Config(), c.Credential()

	uhostClient := uhost.NewClient(cfg, cred)

	req := uhostClient.NewDescribeUHostInstanceRequest()
	resp, err := uhostClient.DescribeUHostInstance(req)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Response: %+v\n", resp)
}
