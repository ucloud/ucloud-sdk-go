package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// ULBClient is the client of ULB
type ULBClient struct {
	*ucloud.Client
}

// NewClient will return a instance of ULBClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *ULBClient {
	meta := ucloud.ClientMeta{Product: "ULB"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &ULBClient{
		client,
	}
}
