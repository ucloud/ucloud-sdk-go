// Code is generated by ucloud-model, DO NOT EDIT IT.



package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UNetClient is the client of UNet
type UNetClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UNetClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UNetClient {
    meta := ucloud.ClientMeta{Product: "UNet"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UNetClient{
		client,
	}
}
