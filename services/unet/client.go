package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UNetClient is the client of UNet
type UNetClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UNetClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UNetClient {
	client := ucloud.NewClient(config, credential)
	return &UNetClient{
		client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UNetClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UNetClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UNetClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UNetClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
