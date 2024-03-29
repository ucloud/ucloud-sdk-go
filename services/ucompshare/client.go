// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucompshare

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UCompShareClient is the client of UCompShare
type UCompShareClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UCompShareClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UCompShareClient {
	meta := ucloud.ClientMeta{Product: "UCompShare"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UCompShareClient{
		client,
	}
}
