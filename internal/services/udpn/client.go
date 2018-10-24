package udpn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UDPNClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UDPNClient {
	client := ucloud.NewClient(config, credential)
	return &UDPNClient{
		client: client,
	}
}
