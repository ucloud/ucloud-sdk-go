// Code is generated by ucloud-model, DO NOT EDIT IT.

package unvs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UNVSClient is the client of UNVS
type UNVSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UNVSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UNVSClient {
	meta := ucloud.ClientMeta{Product: "UNVS"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UNVSClient{
		client,
	}
}