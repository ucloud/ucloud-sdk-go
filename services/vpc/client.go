package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// VPCClient is the client of VPC2.0
type VPCClient struct {
	*ucloud.Client
}

// NewClient will return a instance of VPCClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *VPCClient {
	client := ucloud.NewClient(config, credential)
	return &VPCClient{
		client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *VPCClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *VPCClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *VPCClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *VPCClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
