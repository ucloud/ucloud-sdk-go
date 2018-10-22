package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UBillClient struct {
	client *ucloud.Client
}

func NewClient(config *ucloud.Config, credential *auth.Credential) *UBillClient {
	client := ucloud.NewClient(config, credential)
	return &UBillClient{
		client: client,
	}
}
