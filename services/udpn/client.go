package udpn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDPNClient is the client of UDPN
type UDPNClient struct {
	client *ucloud.Client
}

// NewClient will return a instance of UDPNClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDPNClient {
	client := ucloud.NewClient(config, credential)
	return &UDPNClient{
		client: client,
	}
}

// AddHttpRequestHandler will append a response handler to client
func (c *UDPNClient) AddHttpRequestHandler(h ucloud.HttpRequestHandler) error {
	return c.client.AddHttpRequestHandler(h)
}

// AddRequestHandler will append a response handler to client
func (c *UDPNClient) AddRequestHandler(h ucloud.RequestHandler) error {
	return c.client.AddRequestHandler(h)
}

// AddHttpResponseHandler will append a http response handler to client
func (c *UDPNClient) AddHttpResponseHandler(h ucloud.HttpResponseHandler) error {
	return c.client.AddHttpResponseHandler(h)
}

// AddResponseHandler will append a response handler to client
func (c *UDPNClient) AddResponseHandler(h ucloud.ResponseHandler) error {
	return c.client.AddResponseHandler(h)
}
