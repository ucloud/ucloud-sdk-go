// Code is generated by ucloud-model, DO NOT EDIT IT.

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDBClient is the client of UDB
type UDBClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDBClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDBClient {
	meta := ucloud.ClientMeta{Product: "UDB"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UDBClient{
		client,
	}
}
