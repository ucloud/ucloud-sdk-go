package external

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

type MockedSTSV2Case struct {
	MockedVector        string
	MockedError         error
	MockedRegion        string
	MockedZone          string
	Do                  func(*metadata.DefaultClient) (interface{}, error)
	Golden              interface{}
	GoldenError         bool
	ValidateHeadersFunc func(*http.HttpRequest) error
}

func RunMockedMetadataV2Case(t *testing.T, mockedCase MockedSTSV2Case) {
	httpClient := mock.NewHttpClient()
	err := httpClient.MockHTTP(func(request *http.HttpRequest, response *http.HttpResponse) error {
		if mockedCase.MockedError != nil {
			return mockedCase.MockedError
		}

		// Route based on path
		path := request.GetURL()
		if path == "http://100.80.80.80"+metadataRegionPath {
			region := mockedCase.MockedRegion
			if region == "" {
				region = "cn-bj2"
			}
			return response.SetBody([]byte(region))
		} else if path == "http://100.80.80.80"+metadataZonePath {
			zone := mockedCase.MockedZone
			if zone == "" {
				zone = "cn-bj2-05"
			}
			return response.SetBody([]byte(zone))
		}

		// Validate custom headers only for STS requests (not for region/zone)
		if mockedCase.ValidateHeadersFunc != nil {
			if err := mockedCase.ValidateHeadersFunc(request); err != nil {
				return err
			}
		}

		return response.SetBody([]byte(mockedCase.MockedVector))
	})
	assert.NoError(t, err)

	mdProvider := metadata.DefaultClient{}
	_ = mdProvider.SetHttpClient(httpClient)

	resp, err := mockedCase.Do(&mdProvider)
	assert.Equal(t, mockedCase.Golden, resp)
	if mockedCase.GoldenError {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
}

// TestAssumeRoleV2_BasicRequest tests V2 API with only RoleName parameter
func TestAssumeRoleV2_BasicRequest(t *testing.T) {
	cases := []MockedSTSV2Case{
		{
			MockedVector: mockedAssumeRoleV2,
			Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
				req := AssumeRoleV2Request{RoleName: "ServiceRoleForUK8S"}
				cfg, err := loadSTSConfigV2(req, mdProvider)
				if err != nil {
					return nil, err
				}
				return cfg.Credential(), nil
			},
			Golden:      goldenAssumeRoleV2,
			GoldenError: false,
		},
		{
			MockedVector: mockedAssumeRoleV2,
			Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
				req := AssumeRoleV2Request{} // Empty RoleName
				cfg, err := loadSTSConfigV2(req, mdProvider)
				if err != nil {
					return nil, err
				}
				return cfg.Credential(), nil
			},
			Golden:      goldenAssumeRoleV2,
			GoldenError: false,
		},
	}

	for _, mockedCase := range cases {
		RunMockedMetadataV2Case(t, mockedCase)
	}
}

// TestAssumeRoleV2_WithDurationSeconds tests DurationSeconds parameter
func TestAssumeRoleV2_WithDurationSeconds(t *testing.T) {
	mockedCase := MockedSTSV2Case{
		MockedVector: mockedAssumeRoleV2,
		Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
			req := AssumeRoleV2Request{
				RoleName:        "ServiceRoleForUK8S",
				DurationSeconds: 3600,
			}
			cfg, err := loadSTSConfigV2(req, mdProvider)
			if err != nil {
				return nil, err
			}
			return cfg.Credential(), nil
		},
		Golden:      goldenAssumeRoleV2,
		GoldenError: false,
		ValidateHeadersFunc: func(request *http.HttpRequest) error {
			// Verify Duration-Seconds header is set
			headers := request.GetHeaderMap()
			assert.Equal(t, "3600", headers["Duration-Seconds"])
			return nil
		},
	}

	RunMockedMetadataV2Case(t, mockedCase)
}

// TestAssumeRoleV2_WithSessionName tests SessionName parameter
func TestAssumeRoleV2_WithSessionName(t *testing.T) {
	mockedCase := MockedSTSV2Case{
		MockedVector: mockedAssumeRoleV2,
		Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
			req := AssumeRoleV2Request{
				RoleName:    "ServiceRoleForUK8S",
				SessionName: "test-session-123",
			}
			cfg, err := loadSTSConfigV2(req, mdProvider)
			if err != nil {
				return nil, err
			}
			return cfg.Credential(), nil
		},
		Golden:      goldenAssumeRoleV2,
		GoldenError: false,
		ValidateHeadersFunc: func(request *http.HttpRequest) error {
			// Verify Session-Name header is set
			headers := request.GetHeaderMap()
			assert.Equal(t, "test-session-123", headers["Session-Name"])
			return nil
		},
	}

	RunMockedMetadataV2Case(t, mockedCase)
}

// TestAssumeRoleV2_WithRolePolicy tests RolePolicy parameter (base64 encoded)
func TestAssumeRoleV2_WithRolePolicy(t *testing.T) {
	policyJSON := `{"Version":"1","Statement":[{"Effect":"Allow","Action":["uhost:DescribeUHostInstance"],"Resource":["*"]}]}`
	policyBase64 := base64.StdEncoding.EncodeToString([]byte(policyJSON))

	mockedCase := MockedSTSV2Case{
		MockedVector: mockedAssumeRoleV2,
		Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
			req := AssumeRoleV2Request{
				RoleName:   "ServiceRoleForUK8S",
				RolePolicy: policyBase64,
			}
			cfg, err := loadSTSConfigV2(req, mdProvider)
			if err != nil {
				return nil, err
			}
			return cfg.Credential(), nil
		},
		Golden:      goldenAssumeRoleV2,
		GoldenError: false,
		ValidateHeadersFunc: func(request *http.HttpRequest) error {
			// Verify Role-Policy header is set
			headers := request.GetHeaderMap()
			assert.Equal(t, policyBase64, headers["Role-Policy"])
			return nil
		},
	}

	RunMockedMetadataV2Case(t, mockedCase)
}

// TestAssumeRoleV2_AllParameters tests all parameters combined
func TestAssumeRoleV2_AllParameters(t *testing.T) {
	policyJSON := `{"Version":"1","Statement":[{"Effect":"Allow","Action":["uhost:*"],"Resource":["*"]}]}`
	policyBase64 := base64.StdEncoding.EncodeToString([]byte(policyJSON))

	mockedCase := MockedSTSV2Case{
		MockedVector: mockedAssumeRoleV2,
		Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
			req := AssumeRoleV2Request{
				RoleName:        "ServiceRoleForUK8S",
				DurationSeconds: 3600,
				SessionName:     "full-test-session",
				RolePolicy:      policyBase64,
			}
			cfg, err := loadSTSConfigV2(req, mdProvider)
			if err != nil {
				return nil, err
			}
			return cfg.Credential(), nil
		},
		Golden:      goldenAssumeRoleV2,
		GoldenError: false,
		ValidateHeadersFunc: func(request *http.HttpRequest) error {
			// Verify all headers are set
			headers := request.GetHeaderMap()
			assert.Equal(t, "3600", headers["Duration-Seconds"])
			assert.Equal(t, "full-test-session", headers["Session-Name"])
			assert.Equal(t, policyBase64, headers["Role-Policy"])
			return nil
		},
	}

	RunMockedMetadataV2Case(t, mockedCase)
}

// TestParseISO8601Time tests ISO 8601 time parsing
func TestParseISO8601Time(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "RFC3339 format",
			input:    "2025-10-20T06:51:25Z",
			expected: time.Date(2025, 10, 20, 6, 51, 25, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC3339 with timezone",
			input:    "2025-10-20T06:51:25+08:00",
			expected: time.Date(2025, 10, 20, 6, 51, 25, 0, time.FixedZone("CST", 8*3600)),
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "2025/10/20 06:51:25",
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseISO8601Time(tt.input)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, result.Equal(tt.expected), "Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestAssumeRoleV2_ErrorResponse tests error handling
func TestAssumeRoleV2_ErrorResponse(t *testing.T) {
	mockedCase := MockedSTSV2Case{
		MockedVector: mockedAssumeRoleV2Error,
		Do: func(mdProvider *metadata.DefaultClient) (i interface{}, e error) {
			req := AssumeRoleV2Request{RoleName: "InvalidRole"}
			_, err := loadSTSConfigV2(req, mdProvider)
			return nil, err
		},
		Golden:      nil,
		GoldenError: true,
	}

	RunMockedMetadataV2Case(t, mockedCase)
}

var goldenAssumeRoleV2 = &auth.Credential{
	PublicKey:     "STS.3iXw5eCvtFQBwXmVVypLUa",
	PrivateKey:    "PmZfAbCdEfGhIjKlMnOpQrStUvWxYzDiHw",
	SecurityToken: "AAE7TokenString",
	CanExpire:     true,
	Expires:       time.Date(2025, 11, 11, 4, 32, 54, 0, time.UTC),
}

const mockedAssumeRoleV2 = `
{
    "RetCode": 0,
    "Message": "",
    "Data": {
        "AccessKeyId": "STS.3iXw5eCvtFQBwXmVVypLUa",
        "AccessKeySecret": "PmZfAbCdEfGhIjKlMnOpQrStUvWxYzDiHw",
        "SecurityToken": "AAE7TokenString",
        "Expiration": "2025-11-11T04:32:54Z",
        "CharacterName": "ucs:iam::60036:role/ucs-service-role/ServiceRoleForUK8S",
        "ProjectID": "org-hlrtnn",
        "HostID": "uhost-abc123",
        "UHostID": "uhost-abc123"
    }
}
`

const mockedAssumeRoleV2Error = `
{
    "RetCode": 161,
    "Message": "Role not found"
}
`
