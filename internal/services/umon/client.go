package umon

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UMonClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UMonClient {
	client := ucloud.NewClient(config, credential)
	return &UMonClient{
		client: client,
	}
}
