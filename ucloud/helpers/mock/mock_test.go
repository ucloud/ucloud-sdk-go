package mock

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"

	proto "github.com/ucloud/ucloud-sdk-go/private/protocol/http"
)

func newTestClient() *ucloud.Client {
	cfg := ucloud.NewConfig()
	credential := auth.NewCredential()
	return ucloud.NewClient(&cfg, &credential)
}

func newMockedHttpClient() *HttpClient {
	c := NewHttpClient()

	c.MockHTTP(func(req *proto.HttpRequest, resp *proto.HttpResponse) error {
		if action := req.GetQuery("Action"); len(action) != 0 && action == "TestMockHTTP" {
			resp.SetBody([]byte(`{"Action": "TestMockHTTPResponse"}`))
		}
		return nil
	})

	c.MockData(func(reqData Request, respData Response) error {
		if action, ok := reqData["Action"]; ok && action == "TestMockData" {
			respData["Action"] = "TestMockDataResponse"
		}
		return nil
	})

	c.MockData(func(reqData Request, respData Response) error {
		if action, ok := reqData["Action"]; ok && action == "TestMockError" {
			return http.ErrServerClosed
		}
		return nil
	})

	return c
}

func TestHttpClient_Send(t *testing.T) {
	httpClient := newMockedHttpClient()

	type args struct {
		Action string
	}

	tests := []struct {
		name    string
		client  *HttpClient
		args    args
		want    string
		wantErr bool
		err     error
	}{
		{"http", httpClient, args{"TestMockHTTP"}, `{"Action":"TestMockHTTPResponse"}`, false, nil},
		{"data", httpClient, args{"TestMockData"}, `{"Action":"TestMockDataResponse"}`, false, nil},
		{"error", httpClient, args{"TestMockError"}, "", true, http.ErrServerClosed},
	}

	// test for mocked http client
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := proto.NewHttpRequest()
			req.SetQuery("Action", tt.args.Action)
			got, err := tt.client.Send(req)

			if tt.wantErr {
				// assert response error
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				// assert response body
				assert.NoError(t, err)
				assert.Equal(t, tt.want, string(got.GetBody()))
			}
		})
	}
}

func TestMockClient(t *testing.T) {
	client := newTestClient()
	httpClient := newMockedHttpClient()
	client.SetHttpClient(httpClient)

	type args struct {
		Action string
	}

	tests := []struct {
		name    string
		client  *ucloud.Client
		args    args
		want    string
		wantErr bool
	}{
		{"http", client, args{"TestMockHTTP"}, `TestMockHTTPResponse`, false},
		{"data", client, args{"TestMockData"}, `TestMockDataResponse`, false},
		{"error", client, args{"TestMockError"}, "", true},
	}

	// test for mocked ucloud client
	for _, tt := range tests {
		req := &request.CommonBase{}
		resp := &response.CommonBase{}
		t.Run(tt.name, func(t *testing.T) {
			err := client.InvokeAction(tt.args.Action, req, resp)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, resp.Action)
			}
		})
	}
}
