// Code is generated by ucloud-model, DO NOT EDIT IT.

package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// PathXClient is the client of PathX
type PathXClient struct {
	*ucloud.Client
}

// NewClient will return a instance of PathXClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *PathXClient {
	meta := ucloud.ClientMeta{Product: "PathX"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &PathXClient{
		client,
	}
}
