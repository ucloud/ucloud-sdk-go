// Code is generated by ucloud-model, DO NOT EDIT IT.

package uphone

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UPhoneClient is the client of UPhone
type UPhoneClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UPhoneClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UPhoneClient {
	meta := ucloud.ClientMeta{Product: "UPhone"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UPhoneClient{
		client,
	}
}