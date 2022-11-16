// Code is generated by ucloud-model, DO NOT EDIT IT.



package cube

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// CubeClient is the client of Cube
type CubeClient struct {
	*ucloud.Client
}

// NewClient will return a instance of CubeClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *CubeClient {
    meta := ucloud.ClientMeta{Product: "Cube"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &CubeClient{
		client,
	}
}
