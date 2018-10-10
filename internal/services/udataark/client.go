package udataark

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UDataArkClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UDataArkClient {
	client := ucloud.NewClient(config, credential)
	return &UDataArkClient{
		client: client,
	}
}
