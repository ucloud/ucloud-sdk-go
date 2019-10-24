package main

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

func main() {
	client := metadata.NewClient()

	// describe instance information
	inst, err := client.GetInstanceIdentityDocument()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] Metadata: %+v\n", inst)

	// describe user data for cloud-init
	userData, err := client.GetUserData()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] UserData: %+v\n", userData)

	// describe vendor data for ucloud vendor scripts
	vendorData, err := client.GetVendorData()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
	}
	fmt.Printf("[INFO] VendorData: %+v\n", vendorData)
}
