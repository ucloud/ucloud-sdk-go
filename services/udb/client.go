package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDBClient is the client of UDB
type UDBClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDBClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDBClient {
	client := ucloud.NewClient(config, credential)
	return &UDBClient{
		client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UDBClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UDBClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UDBClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UDBClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
