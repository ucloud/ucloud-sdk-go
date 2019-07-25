package ucloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
	"testing"
	"time"
)

type clientFactory func() Client
type clientCaseGolden struct {
	Action  string
	RetCode int
}

type ClientTestCase struct {
	Name        string
	InputVector clientFactory          // client factory to create test client
	Mock        map[string]interface{} // mock data for http client, see `ucloud/helpers/mock`
	Golden      clientCaseGolden       // return data
	Error       string                 // error contained message, if not empty, means expect an error raised
}

func TestClient(t *testing.T) {
	tests := []ClientTestCase{
		{
			Name:        "basic",
			InputVector: newTestClient,
			Mock:        map[string]interface{}{"RetCode": "0"},
			Golden:      clientCaseGolden{RetCode: 0},
		},
	}

	for _, test := range tests {
		var err error
		client := test.InputVector()
		httpClient := mock.NewHttpClient()

		err = httpClient.MockData(func(requests mock.Request, responses mock.Response) error {
			for k, v := range test.Mock {
				responses[k] = v
			}
			return nil
		})
		assert.NoError(t, err)

		err = client.SetHttpClient(httpClient)
		if err != nil {
			return
		}

		req := request.CommonBase{}
		resp := response.CommonBase{}
		err = client.InvokeAction("Test", &req, &resp)
		if test.Error != "" {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.Error)

			if retCodeErr, ok := err.(uerr.ServerError); ok && test.Golden.RetCode != 0 {
				assert.Equal(t, test.Golden.RetCode, retCodeErr.Code())
			}
		}
	}
}

func newTestClient() Client {
	cfg := NewConfig()
	cfg.BaseUrl = "https://api.ucloud.cn"
	cfg.Region = "cn-bj2"
	cfg.LogLevel = log.WarnLevel
	cfg.Timeout = 5 * time.Second
	cfg.MaxRetries = 1

	credential := auth.NewCredential()
	return *NewClient(&cfg, &credential)
}

type MockRequest struct {
	request.CommonBase
}

type MockResponse struct {
	response.CommonBase

	TotalCount int
	UHostSet   []map[string]interface{}
}

type mockClient struct{}

func (c *mockClient) Send(req *http.HttpRequest) (*http.HttpResponse, error) {
	resp := &http.HttpResponse{}
	resp.SetBody([]byte(fmt.Sprintf(`{"Action": "%sResponse", "RetCode": 0, "Message": ""}`, testDefaultAction)))
	resp.SetStatusCode(200)
	return resp, nil
}
