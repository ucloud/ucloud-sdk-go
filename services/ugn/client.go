// Code is generated by ucloud-model, DO NOT EDIT IT.

package ugn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UGNClient is the client of UGN
type UGNClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UGNClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UGNClient {
	meta := ucloud.ClientMeta{Product: "UGN"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UGNClient{
		client,
	}
}
