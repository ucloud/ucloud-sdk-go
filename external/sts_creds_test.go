package external

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

type MockedSTSCase struct {
	MockedVector string
	MockedError  error
	Do           func(metadataProvider) (interface{}, error)
	Golden       interface{}
	GoldenError  bool
}

func RunMockedMetadataCase(t *testing.T, mockedCase MockedSTSCase) {
	httpClient := mock.NewHttpClient()
	err := httpClient.MockHTTP(func(request *http.HttpRequest, response *http.HttpResponse) error {
		if mockedCase.MockedError != nil {
			return mockedCase.MockedError
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

func TestAssumeRole(t *testing.T) {
	cases := []MockedSTSCase{
		{
			MockedVector: mockedAssumeRole,
			Do: func(mdProvider metadataProvider) (i interface{}, e error) {
				req := AssumeRoleRequest{}
				cfg, err := loadSTSConfig(req, mdProvider)
				if err != nil {
					return nil, err
				}
				return cfg.Credential(), nil
			},
			Golden:      goldenAssumeRole,
			GoldenError: false,
		},
		{
			MockedVector: mockedAssumeRole,
			Do: func(mdProvider metadataProvider) (i interface{}, e error) {
				req := AssumeRoleRequest{RoleName: "UHostInstance"}
				cfg, err := loadSTSConfig(req, mdProvider)
				if err != nil {
					return nil, err
				}
				return cfg.Credential(), nil
			},
			Golden:      goldenAssumeRole,
			GoldenError: false,
		},
	}

	for _, mockedCase := range cases {
		RunMockedMetadataCase(t, mockedCase)
	}
}

var goldenAssumeRole = &auth.Credential{
	PublicKey:     "foo",
	PrivateKey:    "bar",
	SecurityToken: "foobar",
	CanExpire:     true,
	Expires:       time.Unix(1571218055, 0),
}

const mockedAssumeRole = `
{
    "Data": {
        "Expiration": 1571218055,
        "PrivateKey": "bar",
        "ProjectID": "org-xp2ucn",
        "PublicKey": "foo",
        "RoleName": "UHostInstance",
        "SecurityToken": "foobar",
        "UHostID": "uhost-hhqeihh5"
    },
    "Message": "",
    "RetCode": 0
}
`
