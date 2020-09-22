package ucloud

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type clientFactory func() *Client
type clientCaseGolden struct {
	Action  string
	RetCode int
}

type ClientTestCase struct {
	Name        string
	InputVector clientFactory          // client factory to create test client
	Mock        map[string]interface{} // mock data for http client, see `ucloud/helpers/mock`
	Golden      *clientCaseGolden      // return data
	Error       string                 // error contained message, if not empty, means expect an error raised
}

func TestClient(t *testing.T) {
	tests := []ClientTestCase{
		{
			Name:        "basic",
			InputVector: newTestClient,
			Mock:        map[string]interface{}{"Action": "Test", "RetCode": 0},
			Golden:      &clientCaseGolden{Action: "Test", RetCode: 0},
		},
		{
			Name:        "notFound",
			InputVector: newTestClient,
			Mock: map[string]interface{}{
				"Action":  "Test",
				"RetCode": 161,
				"Message": "Action [Test] not found",
			},
			Golden: &clientCaseGolden{Action: "Test", RetCode: 161},
			Error:  "not found",
		},
		{
			Name: "httpRequestHandler",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddHttpRequestHandler(func(c *Client, req *http.HttpRequest) (httpRequest *http.HttpRequest, e error) {
					_ = req.SetQuery("Action", "Echo")
					return req, nil
				})
				return client
			},
			Mock:   nil,
			Golden: &clientCaseGolden{Action: "Echo", RetCode: 0},
		},
		{
			Name: "requestHandler",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddRequestHandler(func(c *Client, req request.Common) (common request.Common, e error) {
					_ = req.SetAction("Echo")
					return req, nil
				})
				return client
			},
			Mock:   nil,
			Golden: &clientCaseGolden{Action: "Echo", RetCode: 0},
		},
		{
			Name: "httpResponseHandler",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddHttpResponseHandler(func(c *Client, req *http.HttpRequest, resp *http.HttpResponse, err error) (httpResponse *http.HttpResponse, e error) {
					_ = resp.SetBody([]byte(`{"Action": "Mock", "RetCode": 42}`))
					return resp, nil
				})
				return client
			},
			Mock:   nil,
			Golden: &clientCaseGolden{Action: "Mock", RetCode: 42},
		},
		{
			Name: "responseHandler",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddResponseHandler(func(c *Client, req request.Common, resp response.Common, err error) (common response.Common, e error) {
					rv := reflect.ValueOf(resp).Elem()
					rv.FieldByName("Action").Set(reflect.ValueOf("Mock"))
					rv.FieldByName("RetCode").Set(reflect.ValueOf(42))
					return resp, nil
				})
				return client
			},
			Mock:   nil,
			Golden: &clientCaseGolden{Action: "Mock", RetCode: 42},
		},
		{
			Name: "invalidHTTPRequest",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddHttpRequestHandler(func(c *Client, req *http.HttpRequest) (httpRequest *http.HttpRequest, e error) {
					return req, fmt.Errorf("http query is invalid")
				})
				return client
			},
			Mock:  nil,
			Error: "InvalidRequestError",
		},
		{
			Name: "invalidRequest",
			InputVector: func() *Client {
				client := newTestClient()
				_ = client.AddRequestHandler(func(c *Client, req request.Common) (common request.Common, e error) {
					return req, fmt.Errorf("request is invalid")
				})
				return client
			},
			Mock:  nil,
			Error: "InvalidRequestError",
		},

		{
			Name: "NullCredential",
			InputVector: func() *Client {
				cfg := NewConfig()
				return &Client{
					config:     &cfg,
					credential: nil,
				}
			},
			Mock:  nil,
			Error: "NullCredentialError",
		},

		{
			Name: "NullConfig",
			InputVector: func() *Client {
				cre := auth.NewCredential()
				return &Client{
					config:     nil,
					credential: &cre,
				}
			},
			Mock:  nil,
			Error: "NullConfigError",
		},

		{
			Name: "NullClient",
			InputVector: func() *Client {
				return NewClient(nil, nil)
			},
			Mock:  nil,
			Error: "NullCredentialError",
		},
	}

	for _, test := range tests {
		var err error
		client := test.InputVector()

		err = mockTestClient(client, test.Mock)
		assert.NoError(t, err)

		// send mocked request, assert response value
		req := request.CommonBase{}
		resp := response.CommonBase{}
		err = client.InvokeAction("Test", client.SetupRequest(&req), &resp)
		if test.Error == "" {
			if test.Golden.Action != "" {
				assert.Equal(t, resp.Action, test.Golden.Action)
			}

			assert.Equal(t, resp.RetCode, test.Golden.RetCode)
		} else {
			if !assert.Error(t, err) {
				t.FailNow()
			}
			if !assert.Contains(t, err.Error(), test.Error) {
				t.FailNow()
			}

			if test.Golden != nil {
				retCodeErr, ok := err.(uerr.ServerError)
				if !assert.Equal(t, ok, true) {
					t.FailNow()
				}
				assert.Equal(t, retCodeErr.Code(), test.Golden.RetCode)
			}
		}
	}
}

type httpMockedTest struct {
	InputVector   string
	MockedVector  mock.Func
	Golden        interface{}
	GoldenErr     bool
	GoldenErrName string
}

func TestClient_http_mock(t *testing.T) {
	tests := []httpMockedTest{
		{
			InputVector: "HTTPMockStatus400",
			MockedVector: func(httpRequest *http.HttpRequest, httpResponse *http.HttpResponse) error {
				httpResponse.SetStatusCode(400)
				return http.NewStatusError(400, "Bad Request")
			},
			GoldenErr:     true,
			GoldenErrName: uerr.ErrHTTPStatus,
		},
		{
			InputVector: "HTTPMockStatus400WithRequestUUID",
			MockedVector: func(httpRequest *http.HttpRequest, httpResponse *http.HttpResponse) error {
				httpResponse.GetHeaders().Set(headerKeyRequestUUID, "foo-bar")
				httpResponse.SetStatusCode(400)
				return http.NewStatusError(400, "Bad Request")
			},
			GoldenErr:     true,
			GoldenErrName: uerr.ErrHTTPStatus,
		},
		{
			InputVector: "ResponseBodyError",
			MockedVector: func(httpRequest *http.HttpRequest, httpResponse *http.HttpResponse) error {
				b := `"{}"`
				if err := httpResponse.SetBody([]byte(b)); err != nil {
					return err
				}
				return nil
			},
			GoldenErr:     true,
			GoldenErrName: uerr.ErrResponseBodyError,
		},
		{
			InputVector: "EmptyResponseBodyError",
			MockedVector: func(httpRequest *http.HttpRequest, httpResponse *http.HttpResponse) error {
				if err := httpResponse.SetBody(nil); err != nil {
					return err
				}
				return nil
			},
			GoldenErr:     true,
			GoldenErrName: uerr.ErrEmptyResponseBodyError,
		},
	}
	for _, test := range tests {
		client := newTestClient()
		httpClient := mock.NewHttpClient()
		_ = httpClient.MockHTTP(test.MockedVector)
		_ = client.SetHttpClient(httpClient)
		var resp MockResponse
		err := client.InvokeAction(test.InputVector, &MockRequest{}, &resp)
		if test.GoldenErr {
			if !assert.Error(t, err) {
				t.FailNow()
			}
			if len(test.GoldenErrName) > 0 {
				uErr, ok := err.(uerr.Error)
				if !assert.Equal(t, ok, true) {
					t.FailNow()
				}
				assert.Equal(t, uErr.Name(), test.GoldenErrName)
			}
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.Golden, resp)
		}
	}
}

func newTestClient() *Client {
	cfg := NewConfig()
	cfg.BaseUrl = "https://api.ucloud.cn"
	cfg.Region = "cn-bj2"
	cfg.Zone = "cn-bj2-02"
	cfg.ProjectId = "foo"
	cfg.LogLevel = log.DebugLevel
	cfg.Timeout = 5 * time.Second
	cfg.MaxRetries = 1

	credential := auth.NewCredential()
	return NewClient(&cfg, &credential)
}

func mockTestClient(client *Client, data map[string]interface{}) error {
	var err error
	httpClient := mock.NewHttpClient()

	// mock data with golden
	err = httpClient.MockData(func(requests mock.Request, responses mock.Response) error {
		if action, ok := requests["Action"]; ok && action == "Echo" {
			for k, v := range requests {
				responses[k] = v
			}
		} else {
			for k, v := range data {
				responses[k] = v
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = client.SetHttpClient(httpClient)
	if err != nil {
		return err
	}

	return nil
}

type MockRequest struct {
	request.CommonBase
}

type MockResponse struct {
	response.CommonBase

	TotalCount int
	UHostSet   []map[string]interface{}
}

func TestGenericClient(t *testing.T) {
	tests := []ClientTestCase{
		{
			Name:        "generic_ok",
			InputVector: newTestClient,
			Mock:        map[string]interface{}{"Action": "Test", "Message": "", "RetCode": 0.0},
			Golden:      &clientCaseGolden{Action: "Test", RetCode: 0},
			Error:       "",
		},

		{
			Name:        "generic_no",
			InputVector: newTestClient,
			Mock:        map[string]interface{}{"Action": "Test", "Message": "Action [Test] not found", "RetCode": 161},
			Golden:      &clientCaseGolden{Action: "Test", RetCode: 161},
			Error:       "not found",
		},
	}

	for _, test := range tests {
		var err error
		client := test.InputVector()

		err = mockTestClient(client, test.Mock)
		assert.NoError(t, err)

		// send mocked request, assert response value
		req := client.NewGenericRequest()
		resp, err := client.GenericInvoke(req)
		if test.Error == "" {
			if test.Golden.Action != "" {
				assert.Equal(t, resp.GetAction(), test.Golden.Action)
			}

			assert.Equal(t, resp.GetRetCode(), test.Golden.RetCode)
			assert.Equal(t, test.Mock, resp.GetPayload())
		} else {
			if !assert.Error(t, err) {
				t.FailNow()
			}
			if !assert.Contains(t, err.Error(), test.Error) {
				t.FailNow()
			}
			retCodeErr, ok := err.(uerr.ServerError)
			if !assert.Equal(t, ok, true) {
				t.FailNow()
			}
			assert.Equal(t, retCodeErr.Code(), test.Golden.RetCode)
		}
	}
}
