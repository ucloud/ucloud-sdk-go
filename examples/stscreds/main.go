package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/external"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func main() {
	c, err := external.LoadSTSConfig(external.AssumeRoleRequest{RoleName: "Uk8sServiceCharacter"})
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Config: %+v\n", c)

	describeUHost(c.Config(), c.Credential())
}

func describeUHost(cfg *ucloud.Config, cred *auth.Credential) {
	client := uhost.NewClient(cfg, cred)

	req := client.NewDescribeUHostInstanceRequest()
	req.Region = ucloud.String("cn-bj2")

	resp, err := client.DescribeUHostInstance(req)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Response: %+v\n", resp)
}
