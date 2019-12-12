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

	password := base64.StdEncoding.EncodeToString([]byte("ucloud_password"))
	req := client.NewGenericRequest()
	//err := req.SetAction("CreateUHostInstance")
	//if err != nil {
	//	panic(err)
	//}
	req.SetPayload(map[string]interface{}{
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

	genericResp, err := client.GenericInvoke(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(genericResp.Payload()["UHostIds"].([]interface{})[0].(string))
	fmt.Println(genericResp.Payload())
	fmt.Println(genericResp.Payload()["Action"])

	type CreateUHostInstanceResponse struct {
		// response.CommonBase
		// UHost实例Id集合
		UHostIds []string

		// IP信息
		IPs []string
	}
	resp := &CreateUHostInstanceResponse{}
	if err := genericResp.Unmarshal(resp); err != nil {
		panic(err)
	}
	fmt.Println(resp.UHostIds[0])
	fmt.Println(resp)
	// fmt.Println(resp.Action)

	r := client.NewGenericRequest()

	r.SetPayload(map[string]interface{}{
		"Action":   "DescribeUHostInstance",
		"Zone":     zone,
		"UHostIds": []string{resp.UHostIds[0]},
		// "UHostIds": []string{genericResp.Payload()["UHostIds"].([]interface{})[0].(string)},
	})

	res, err := client.GenericInvoke(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Payload())
	fmt.Println(res.Payload()["Action"])
	fmt.Println(res.Payload()["UHostSet"].([]interface{})[0].(map[string]interface{})["BasicImageId"].(string))
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
