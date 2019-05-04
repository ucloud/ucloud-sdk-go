package mock

import (
	"encoding/json"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
)

// Request is the parameters of an action invoking
type Request map[string]interface{}

// Response is the response data of an action invoking
type Response map[string]interface{}

// Func is the handler to process http request and return a http response or error
type Func func(*http.HttpRequest, *http.HttpResponse) error

// DataFunc is the function to resolve the mapping of request and response data
type DataFunc func(Request, Response) error

// HttpClient is the mocked client of ucloud api
type HttpClient struct {
	mockFuncs []Func
}

// NewHttpClient will return a new mocked client instance
func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

// MockHTTP will append the mocking function to the mocked client
func (c *HttpClient) MockHTTP(fn Func) error {
	c.mockFuncs = append(c.mockFuncs, fn)
	return nil
}

// MockData will append the data mocking function to the mocked client
func (c *HttpClient) MockData(fn DataFunc) error {
	c.mockFuncs = append(c.mockFuncs, func(req *http.HttpRequest, resp *http.HttpResponse) error {
		reqData := Request{}
		respData := Response{}

		// reserving request data passed through by previous mocking handler
		for k, v := range req.GetQueryMap() {
			reqData[k] = v
		}

		// reserving response data passed through by previous mocking handler
		if body := resp.GetBody(); len(body) > 0 {
			err := json.Unmarshal(body, &respData)
			if err != nil {
				return err
			}
		}

		// call data mocking handler
		err := fn(reqData, respData)
		if err != nil {
			return err
		}

		// set response data into http response
		body, err := json.Marshal(respData)
		if err != nil {
			return err
		}

		return resp.SetBody(body)
	})
	return nil
}

// Send is the mocked method to handle the request and response data
func (c *HttpClient) Send(req *http.HttpRequest) (*http.HttpResponse, error) {
	resp := http.NewHttpResponse()

	for _, handler := range c.mockFuncs {
		err := handler(req, resp)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}
