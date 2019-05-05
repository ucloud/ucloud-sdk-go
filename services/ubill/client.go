package ubill

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UBillClient is the client of UBill
type UBillClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UBillClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UBillClient {
	client := ucloud.NewClient(config, credential)
	return &UBillClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UBillClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UBillClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UBillClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UBillClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
