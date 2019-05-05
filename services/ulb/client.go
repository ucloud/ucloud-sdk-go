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
	client := ucloud.NewClient(config, credential)
	return &ULBClient{
		client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *ULBClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *ULBClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *ULBClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *ULBClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
