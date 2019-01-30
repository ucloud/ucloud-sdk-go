package http

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPRequestAccessor(t *testing.T) {
	req := NewHttpRequest()

	// header getter/setter
	k, v := "Content-Type", string(mimeJSON)
	req.SetHeader(k, v)
	assert.Equal(t, req.GetHeaderMap()[k], v)

	// method getter/setter
	method := "GET"
	req.SetMethod(method)
	assert.Equal(t, req.GetMethod(), method)

	// query getter/setter
	k, v = "foo", "bar"
	req.SetQuery(k, v)
	assert.Equal(t, req.GetQueryMap()[k], v)

	req.SetQueryString("foo1=bar")
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar")

	req.SetQueryString("foo1=bar1")
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar1")

	// timeout getter/setter
	timeout := 1 * time.Second
	req.SetTimeout(timeout)
	assert.Equal(t, req.GetTimeout(), timeout)

	// request getter/setter
	body := []byte("data")
	req.SetRequestBody(body)
	assert.Equal(t, req.GetRequestBody(), body)

	// query builder
	req.SetURL("http://api.ucloud.cn?foo2=bar")
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn")
	assert.Equal(t, req.GetQuery("foo2"), "bar")

	req.SetURL("http://api.ucloud.cn:12345?foo2=bar")
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn:12345")
	assert.Equal(t, req.GetQuery("foo2"), "bar")
}

func TestHTTPRequestBuilder(t *testing.T) {
	// http request builder
	req := NewHttpRequest()
	req.SetQuery("key1", "value1")
	req.SetQuery("key2", "value2")
	qs := "key1=value1&key2=value2"

	// http request builder - url query with form url mime type
	// query should in body
	req.SetHeader("Content-Type", string(mimeFormURLEncoded))
	req.SetRequestBody(nil)
	httpReq, err := req.buildHTTPRequest()
	body, err := ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(qs))
	assert.NotContains(t, httpReq.URL.String(), qs)

	// http request builder - url query, body with form url mime type
	// query should in url, and use user's body
	req.SetHeader("Content-Type", string(mimeFormURLEncoded))
	req.SetRequestBody([]byte("content=1"))
	httpReq, err = req.buildHTTPRequest()
	body, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte("content=1"))
	assert.Contains(t, httpReq.URL.String(), qs)

	// http request builder - url query, body with json mime type
	// query should in url, and use user's body
	req.SetHeader("Content-Type", string(mimeJSON))
	req.SetRequestBody([]byte(`{"content": 1}`))
	httpReq, err = req.buildHTTPRequest()
	body, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(`{"content": 1}`))
	assert.Contains(t, httpReq.URL.String(), qs)
}

func TestClientSend(t *testing.T) {
	req := NewHttpRequest()
	req.SetMethod("GET")
	req.SetURL("https://httpbin.org/get")

	client := NewHttpClient()
	resp, err := client.Send(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetStatusCode())
	assert.NotEmpty(t, resp.GetBody())
}

func TestHTTPStatusError(t *testing.T) {
	req := NewHttpRequest()
	req.SetMethod("GET")
	req.SetURL("https://httpbin.org/status/503")

	client := NewHttpClient()
	resp, err := client.Send(req)
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
