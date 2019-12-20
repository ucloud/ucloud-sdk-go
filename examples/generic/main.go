package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
)

const region = "cn-sh2"
const zone = "cn-sh2-02"
const imageID = "uimage-dm50wf"

func main() {
	cfg, credential := loadConfig()
	client := ucloud.NewClient(cfg, credential)

	password := base64.StdEncoding.EncodeToString([]byte("ucloud_password_test"))
	reqCreate := client.NewGenericRequest()
	reqCreate.SetPayload(map[string]interface{}{
		"Action":  "CreateUHostInstance",
		"Zone":    zone,
		"ImageId": imageID,
		"NetworkInterface": []map[string]interface{}{
			{
				"EIP": map[string]interface{}{
					"Bandwidth":    1,
					"OperatorName": "Bgp",
					"PayMode":      "Bandwidth",
				},
			},
		},

		"LoginMode":  "Password",
		"Password":   password,
		"ChargeType": "Dynamic",
		"CPU":        1,
		"Memory":     2048,
		"Disks": []map[string]interface{}{
			{
				"Size":   20,
				"Type":   "LOCAL_NORMAL",
				"IsBoot": "true",
			},
		},
	})

	genericResp, err := client.GenericInvoke(reqCreate)
	if err != nil {
		panic(err)
	}

	type CreateUHostInstanceResponse struct {
		// UHost实例Id集合
		UHostIds []string

		// IP信息
		IPs []string
	}
	respCreate := &CreateUHostInstanceResponse{}
	if err := genericResp.Unmarshal(respCreate); err != nil {
		panic(err)
	}

	reqDescribe := client.NewGenericRequest()
	reqDescribe.SetPayload(map[string]interface{}{
		"Action":   "DescribeUHostInstance",
		"Zone":     zone,
		"UHostIds": []string{respCreate.UHostIds[0]},
	})

	respDescribe, err := client.GenericInvoke(reqDescribe)
	if err != nil {
		panic(err)
	}

	fmt.Println(respDescribe.GetPayload()["UHostSet"].([]interface{})[0].(map[string]interface{})["State"].(string))
}

func loadConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = region
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	log.Info("setup clients ...")

	return &cfg, &credential
}
