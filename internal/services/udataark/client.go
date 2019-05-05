package udataark

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDataArkClient is the client of UDataArk
type UDataArkClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDataArkClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDataArkClient {
	client := ucloud.NewClient(config, credential)
	return &UDataArkClient{
		client,
	}
}
