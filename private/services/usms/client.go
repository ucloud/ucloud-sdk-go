package usms

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// USMSClient is the client of USMS
type USMSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of USMSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *USMSClient {
	client := ucloud.NewClient(config, credential)
	return &USMSClient{
		client,
	}
}
