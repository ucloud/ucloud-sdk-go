package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UDBClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UDBClient {
	client := ucloud.NewClient(config, credential)
	return &UDBClient{
		client: client,
	}
}
