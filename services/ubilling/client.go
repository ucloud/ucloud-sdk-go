package ubillings

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UHostClient is the client of UHost
type UBillingClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UHostClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UBillingClient {
	client := ucloud.NewClient(config, credential)
	return &UBillingClient{
		client: client,
	}
}