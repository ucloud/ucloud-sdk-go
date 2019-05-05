package uphost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UPHostClient is the client of UPHost
type UPHostClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UPHostClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UPHostClient {
	client := ucloud.NewClient(config, credential)
	return &UPHostClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UPHostClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UPHostClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UPHostClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UPHostClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
