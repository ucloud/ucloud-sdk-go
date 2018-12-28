package http

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPRequest(t *testing.T) {
	req := NewHttpRequest()

	k, v := "Content-Type", "application/json"
	req.SetHeader(k, v)
	assert.Equal(t, req.GetHeaderMap()[k], v)

	method := "GET"
	req.SetMethod(method)
	assert.Equal(t, req.GetMethod(), method)

	k, v = "foo", "bar"
	req.SetQuery(k, v)
	assert.Equal(t, req.GetQueryMap()[k], v)

	req.SetQueryString("foo1=bar")
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar")

	req.SetQueryString("foo1=bar1")
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar1")

	timeout := 1 * time.Second
	req.SetTimeout(timeout)
	assert.Equal(t, req.GetTimeout(), timeout)

	req.SetURL("http://api.ucloud.cn?foo2=bar")
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn")
	assert.Equal(t, req.GetQuery("foo2"), "bar")

	req.SetURL("http://api.ucloud.cn:12345?foo2=bar")
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn:12345")
	assert.Equal(t, req.GetQuery("foo2"), "bar")

	body := []byte("data")
	req.SetRequestBody(body)
	assert.Equal(t, req.GetRequestBody(), body)
}

func TestClientSend(t *testing.T) {
	req := NewHttpRequest()
	req.SetMethod("GET")
	req.SetURL("https://httpbin.org/get")

	client := NewHttpClient()
	resp, err := client.Send(&req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetStatusCode())
	assert.NotEmpty(t, resp.GetBody())
}

func TestHTTPStatusError(t *testing.T) {
	req := NewHttpRequest()
	req.SetMethod("GET")
	req.SetURL("https://httpbin.org/status/503")

	client := NewHttpClient()
	resp, err := client.Send(&req)
	assert.Nil(t, resp)

	statusErr, ok := err.(StatusError)
	assert.Equal(t, ok, true)
	assert.Equal(t, statusErr.StatusCode, 503)
	assert.Equal(t, statusErr.Message, "503 SERVICE UNAVAILABLE")
}

func TestHTTPMock(t *testing.T) {
	// TODO
	type mockResponse struct {
	}

	type mockClient struct {
	}
}
