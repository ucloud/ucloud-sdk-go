package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UBillClient is the client of UBill
type UBillClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UBillClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UBillClient {
	meta := ucloud.ClientMeta{Product: "UBill"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UBillClient{
		client,
	}
}
