package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UMemClient is the client of UMem
type UMemClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UMemClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UMemClient {
	client := ucloud.NewClient(config, credential)
	return &UMemClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UMemClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UMemClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UMemClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UMemClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
