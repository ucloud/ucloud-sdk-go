package main

import (
	"os"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/waiter"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
)

const region = "cn-bj2"
const zone = "cn-bj2-05"
const imageID = "uimage-kg0w4u"

var uhostClient *uhost.UHostClient

func init() {
	cfg := ucloud.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Region = region
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	uhostClient = uhost.NewClient(&cfg, &credential)

	log.Info("setup clients ...")
}

func main() {
	uhostID, err := createUHost("sdk-example-wait")
	if err != nil {
		log.Error(err)
	}

	w := waiter.StateWaiter{
		Pending: []string{"pending"},
		Target:  []string{"available"},
		Refresh: func() (interface{}, string, error) {
			inst, err := describeUHostByID(uhostID)
			if err != nil {
				return nil, "", err
			}

			if inst == nil || inst.State != "Running" {
				return nil, "pending", nil
			}

			return inst, "available", nil
		},
		Timeout: 5 * time.Minute,
	}

	if resp, err := w.Wait(); err != nil {
		log.Error(err)
	} else {
		log.Infof("%#v", resp)
	}
}

func describeUHostByID(uhostID string) (*uhost.UHostInstanceSet, error) {
	req := uhostClient.NewDescribeUHostInstanceRequest()
	req.UHostIds = []string{uhostID}

	resp, err := uhostClient.DescribeUHostInstance(req)
	if err != nil {
		return nil, err
	}
	if len(resp.UHostSet) < 1 {
		return nil, nil
	}

	return &resp.UHostSet[0], nil
}

func createUHost(name string) (string, error) {
	req := uhostClient.NewCreateUHostInstanceRequest()
	req.Name = ucloud.String(name)
	req.Zone = ucloud.String(zone)       // TODO: use random zone
	req.ImageId = ucloud.String(imageID) // TODO: use random image
	req.LoginMode = ucloud.String("Password")
	req.Password = ucloud.String("somePassword_")
	req.ChargeType = ucloud.String("Dynamic")
	req.CPU = ucloud.Int(1)
	req.Memory = ucloud.Int(1024)
	req.Tag = ucloud.String("sdk-example")

	resp, err := uhostClient.CreateUHostInstance(req)
	if err != nil {
		return "", err
	}

	return resp.UHostIds[0], nil
}
