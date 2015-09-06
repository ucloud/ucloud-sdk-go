package main

import (
	"fmt"

	"github.com/xiaohui/goucloud/ucloud"
	"github.com/xiaohui/goucloud/service/uhost"
	"github.com/xiaohui/goucloud/ucloud/auth"
	"github.com/xiaohui/goucloud/ucloud/utils"
)



func main() {

	hostsvc := uhost.New(&ucloud.Config{
		Credentials:  &auth.KeyPair{
			PublicKey: "ucloudsomeone@example.com1296235120854146120",
			PrivateKey: "46f09bb9fab4f12dfc160dae12273d5332b5debe",
		},
		Region: "cn-north-01",
		ProjectID: "",
	})

	createUhostParams := uhost.CreateUHostInstanceParams{

		Region: "cn-north-03",
		ImageId: "uimage-j4fbrn",
		LoginMode: "Password",
		Password: "UGFzc3dvcmQx",
		CPU: 1,
		Memory:2048,
		Quantity:1,
		Count: 1,
	}


	response, err := hostsvc.CreateUHostInstance(&createUhostParams)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)


	// describeimage
	imageparams := uhost.DescribeImageParams{
		Region: "cn-north-03",
	}


	imageresp, err := hostsvc.DescribeImage(&imageparams)
	if err != nil {
		fmt.Println(err)
	}

	utils.Prettify(imageresp)
}