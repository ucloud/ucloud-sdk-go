package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func main() {
	cred := &auth.Credential{
		PublicKey:  "ucloudsomeone@example.com1296235120854146120",
		PrivateKey: "46f09bb9fab4f12dfc160dae12273d5332b5debe",
	}
	d := "Action=DescribeUHostInstance&Limit=10&Region=cn-bj2"
	fmt.Println(cred.CreateSign(d))
}
