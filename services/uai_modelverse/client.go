// Code is generated by ucloud-model, DO NOT EDIT IT.

package uai_modelverse

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UAI_ModelverseClient is the client of UAI_Modelverse
type UAI_ModelverseClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UAI_ModelverseClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UAI_ModelverseClient {
	meta := ucloud.ClientMeta{Product: "UAI_Modelverse"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UAI_ModelverseClient{
		client,
	}
}