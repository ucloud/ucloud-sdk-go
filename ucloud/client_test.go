package ucloud

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	stdhttp "net/http"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

const (
	testDefaultAction = "DescribeUHostInstance"
)

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	testTeardown()
	os.Exit(exitCode)
}

func testSetup() {}

func testTeardown() {}

func TestClientTimeout(t *testing.T) {
	req := &MockRequest{}
	resp := &MockResponse{}

	client := newTestClient()
	client.config.BaseUrl = "https://httpbin.org/delay/2"
	client.config.Timeout = 1 * time.Second
	client.config.MaxRetries = 1
	client.SetupRequest(req)

	err := client.InvokeAction("foo", req, resp)
	uErr, ok := err.(uerr.ClientError)
	assert.True(t, ok)
	assert.Equal(t, uErr.Name(), uerr.ErrNetwork)
	assert.Equal(t, req.GetRetryCount(), 1)
	assert.Equal(t, req.GetMaxretries(), 1)
}

func TestClient_setup(t *testing.T) {
	cfg := NewConfig()
	credential := auth.NewCredential()
	credential.CanExpire = true
	credential.Expires = time.Time{}
	credential.SecurityToken = "foo"

	client := NewClientWithMeta(&cfg, &credential, ClientMeta{Product: "OpenSDK"})
	assert.Equal(t, "OpenSDK", client.GetMeta().Product)
	assert.True(t, credential.IsExpired())

	err := client.InvokeAction("ExpiredCredential", &MockRequest{}, &MockResponse{})
	assert.Error(t, err)
	expiredErr := err.(uerr.ClientError)
	assert.Equal(t, uerr.ErrCredentialExpired, expiredErr.Name())
}

func Test_errorHandler(t *testing.T) {
	req := &MockRequest{}
	resp := &MockResponse{}
	client := newTestClient()

	steps := []struct {
		name string
		step func() error
	}{
		{
			name: "unexpected error",
			step: func() error {
				_, err := errorHandler(&client, req, resp, errors.New("unexpected error"))
				if uErr, ok := err.(uerr.ClientError); !ok || uErr.Name() != uerr.ErrSendRequest {
					return errors.New("unexpected error should be convert to unknown client error")
				}
				return nil
			},
		},
		{
			name: "http status error",
			step: func() error {
				_, err := errorHandler(&client, req, resp, uerr.NewServerStatusError(404, "404 NotFound"))
				if uErr, ok := err.(uerr.ServerError); !ok || uErr.StatusCode() != 404 {
					return errors.New("http status error should be convert to status server error")
				}
				return nil
			},
		},
		{
			name: "server timeout error",
			step: func() error {
				httpClient := &stdhttp.Client{Timeout: time.Duration(1)}
				httpReq, err := stdhttp.NewRequest("GET", "https://httpbin.org/delay/2", nil)
				if err != nil {
					return err
				}
				_, err = httpClient.Do(httpReq)
				_, err = errorHandler(&client, req, resp, err)
				if uErr, ok := err.(uerr.ClientError); !ok || uErr.Name() != uerr.ErrNetwork {
					return errors.New("timeout error should be convert to network client error")
				}
				return nil
			},
		},
		{
			name: "business error",
			step: func() error {
				resp := &response.CommonBase{Message: "Missing Action", RetCode: 160}
				_, err := errorHandler(&client, req, resp, nil)
				if uErr, ok := err.(uerr.ServerError); !ok || uErr.Code() != 160 {
					return errors.New("ucloud error should be raised for non-zero retCode")
				}
				return nil
			},
		},
	}
	for _, step := range steps {
		t.Run(step.name, func(t *testing.T) {
			err := step.step()
			if err != nil {
				t.Errorf("errorHandler() error %s", err)
			}
		})
	}
}

func TestAccClientRequestUUID(t *testing.T) {
	client := newTestClient()
	resp := response.CommonBase{}
	_ = client.InvokeAction("DescribeUHostInstance", &request.CommonBase{}, &resp)
	assert.NotZero(t, resp.GetRequestUUID())
}

func TestLoggingLevel(t *testing.T) {
	client := newTestClient() // level: WarnLevel

	logger := client.GetLogger()
	logger.SetLevel(log.WarnLevel)
	client.SetLogger(logger)
	client.config.LogLevel = log.WarnLevel

	client.logActionWarnf("test", "%s", "foo")

	assert.Equal(t, client.logger.GetLevel(), log.WarnLevel)
	assert.Equal(t, client.config.GetActionLevel(testDefaultAction), log.WarnLevel)

	client.config.SetActionLevel(testDefaultAction, log.InfoLevel)
	assert.Equal(t, client.config.GetActionLevel(testDefaultAction), log.InfoLevel)

	client.config.SetActionLevel(testDefaultAction, log.ErrorLevel)
	assert.Equal(t, client.config.GetActionLevel(testDefaultAction), log.ErrorLevel)
}
