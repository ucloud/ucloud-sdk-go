package uaccount

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UAccountClient is the client of UAccount
type UAccountClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UAccountClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UAccountClient {
	client := ucloud.NewClient(config, credential)
	return &UAccountClient{
		client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UAccountClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UAccountClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UAccountClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UAccountClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
