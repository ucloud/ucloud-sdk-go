package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UFSClient is the client of UFS
type UFSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UFSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UFSClient {
	client := ucloud.NewClient(config, credential)
	return &UFSClient{
		client,
	}
}
