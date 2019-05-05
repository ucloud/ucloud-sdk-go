package http

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPRequestAccessor(t *testing.T) {
	var err error
	req := NewHttpRequest()

	// header getter/setter
	k, v := "Content-Type", string(mimeJSON)
	err = req.SetHeader(k, v)
	assert.NoError(t, err)
	assert.Equal(t, req.GetHeaderMap()[k], v)

	// method getter/setter
	method := "GET"
	err = req.SetMethod(method)
	assert.NoError(t, err)
	assert.Equal(t, req.GetMethod(), method)

	// query getter/setter
	k, v = "foo", "bar"
	err = req.SetQuery(k, v)
	assert.NoError(t, err)
	assert.Equal(t, req.GetQueryMap()[k], v)

	err = req.SetQueryString("foo1=bar")
	assert.NoError(t, err)
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar")

	err = req.SetQueryString("foo1=bar1")
	assert.NoError(t, err)
	assert.Equal(t, req.GetQueryMap()["foo1"], "bar1")

	// timeout getter/setter
	timeout := 1 * time.Second
	err = req.SetTimeout(timeout)
	assert.NoError(t, err)
	assert.Equal(t, req.GetTimeout(), timeout)

	// request getter/setter
	body := []byte("data")
	err = req.SetRequestBody(body)
	assert.NoError(t, err)
	assert.Equal(t, req.GetRequestBody(), body)

	// query builder
	err = req.SetURL("http://api.ucloud.cn?foo2=bar")
	assert.NoError(t, err)
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn")
	assert.Equal(t, req.GetQuery("foo2"), "bar")

	err = req.SetURL("http://api.ucloud.cn:12345?foo2=bar")
	assert.NoError(t, err)
	assert.Equal(t, req.GetQueryMap()["foo2"], "bar")
	assert.Equal(t, req.GetURL(), "http://api.ucloud.cn:12345")
	assert.Equal(t, req.GetQuery("foo2"), "bar")
}

func TestHTTPRequestBuilder(t *testing.T) {
	var err error
	// http request builder
	req := NewHttpRequest()
	err = req.SetQuery("key1", "value1")
	assert.NoError(t, err)

	err = req.SetQuery("key2", "value2")
	assert.NoError(t, err)

	qs := "key1=value1&key2=value2"

	// http request builder - url query with form url mime type
	// query should in body
	err = req.SetHeader("Content-Type", string(mimeFormURLEncoded))
	assert.NoError(t, err)

	err = req.SetRequestBody(nil)
	assert.NoError(t, err)

	httpReq, err := req.buildHTTPRequest()
	body, err := ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(qs))
	assert.NotContains(t, httpReq.URL.String(), qs)

	// http request builder - url query, body with form url mime type
	// query should in url, and use user's body
	err = req.SetHeader("Content-Type", string(mimeFormURLEncoded))
	assert.NoError(t, err)

	err = req.SetRequestBody([]byte("content=1"))
	assert.NoError(t, err)

	httpReq, err = req.buildHTTPRequest()
	body, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte("content=1"))
	assert.Contains(t, httpReq.URL.String(), qs)

	// http request builder - url query, body with json mime type
	// query should in url, and use user's body
	err = req.SetHeader("Content-Type", string(mimeJSON))
	assert.NoError(t, err)

	err = req.SetRequestBody([]byte(`{"content": 1}`))
	assert.NoError(t, err)

	httpReq, err = req.buildHTTPRequest()
	body, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(`{"content": 1}`))
	assert.Contains(t, httpReq.URL.String(), qs)
}

func TestClientSend(t *testing.T) {
	var err error
	req := NewHttpRequest()

	err = req.SetMethod("GET")
	assert.NoError(t, err)

	err = req.SetURL("https://httpbin.org/get")
	assert.NoError(t, err)

	client := NewHttpClient()
	resp, err := client.Send(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetStatusCode())
	assert.NotEmpty(t, resp.GetBody())

	headers := resp.GetHeaders()
	header := headers.Get("Server")
	assert.NotZero(t, header)
	noHeader := headers.Get("not_exists")
	assert.NotZero(t, noHeader)
}

func TestHTTPStatusError(t *testing.T) {
	var err error
	req := NewHttpRequest()

	err = req.SetMethod("GET")
	assert.NoError(t, err)

	err = req.SetURL("https://httpbin.org/status/503")
	assert.NoError(t, err)

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
