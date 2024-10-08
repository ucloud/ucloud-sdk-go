// Code is generated by ucloud-model, DO NOT EDIT IT.

package udns

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDNSClient is the client of UDNS
type UDNSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDNSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDNSClient {
	meta := ucloud.ClientMeta{Product: "UDNS"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UDNSClient{
		client,
	}
}
