package stepflow

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// StepFlowClient is the client of StepFlow
type StepFlowClient struct {
	*ucloud.Client
}

// NewClient will return a instance of StepFlowClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *StepFlowClient {
	client := ucloud.NewClient(config, credential)
	return &StepFlowClient{
		client,
	}
}
