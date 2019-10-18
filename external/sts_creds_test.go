package external

import (
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
	"testing"
	"time"
)

type MockedMetadataCase struct {
	MockedVector string
	MockedError  error
	Do           func(client defaultClient) (interface{}, error)
	Golden       interface{}
	GoldenError  bool
}

func RunMockedMetadataCase(t *testing.T, mockedCase MockedMetadataCase) {
	httpClient := mock.NewHttpClient()
	err := httpClient.MockHTTP(func(request *http.HttpRequest, response *http.HttpResponse) error {
		if mockedCase.MockedError != nil {
			return mockedCase.MockedError
		}
		return response.SetBody([]byte(mockedCase.MockedVector))
	})
	assert.NoError(t, err)

	metadataResolver := metadata.DefaultClient{}
	_ = metadataResolver.SetHttpClient(httpClient)
	client := defaultClient{resolver: metadataResolver}

	resp, err := mockedCase.Do(client)
	assert.Equal(t, mockedCase.Golden, resp)
	if mockedCase.GoldenError {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
}

func TestAssumeRole(t *testing.T) {
	cases := []MockedMetadataCase{
		{
			MockedVector: mockedAssumeRole,
			Do: func(client defaultClient) (i interface{}, e error) {
				req := AssumeRoleRequest{}
				return AssumeRole(req)
			},
			Golden:      goldenAssumeRole,
			GoldenError: false,
		},
	}

	for _, mockedCase := range cases {
		RunMockedMetadataCase(t, mockedCase)
	}
}

var goldenAssumeRole = auth.Credential{
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
        "PrivateKey": "foo",
        "ProjectID": "org-xp2ucn",
        "PublicKey": "bar",
        "RoleName": "UHostInstance",
        "SecurityToken": "foobar",
        "UHostID": "uhost-hhqeihh5"
    },
    "Message": "",
    "RetCode": 0
}
`
