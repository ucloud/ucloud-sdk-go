package umon

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UMonClient is the client of UMon
type UMonClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UMonClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UMonClient {
	client := ucloud.NewClient(config, credential)
	return &UMonClient{
		client,
	}
}
