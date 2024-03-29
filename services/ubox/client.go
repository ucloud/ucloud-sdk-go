// Code is generated by ucloud-model, DO NOT EDIT IT.

package ubox

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UBoxClient is the client of UBox
type UBoxClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UBoxClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UBoxClient {
	meta := ucloud.ClientMeta{Product: "UBox"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UBoxClient{
		client,
	}
}
