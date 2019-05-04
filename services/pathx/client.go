package pathx

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// PathXClient is the client of PathX
type PathXClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of PathXClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *PathXClient {
	client := ucloud.NewClient(config, credential)
	return &PathXClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *PathXClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *PathXClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *PathXClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *PathXClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
