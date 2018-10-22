package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UMemClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UMemClient {
	client := ucloud.NewClient(config, credential)
	return &UMemClient{
		client: client,
	}
}
