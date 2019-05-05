package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDiskClient is the client of UDisk
type UDiskClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UDiskClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDiskClient {
	client := ucloud.NewClient(config, credential)
	return &UDiskClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UDiskClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UDiskClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UDiskClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UDiskClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
