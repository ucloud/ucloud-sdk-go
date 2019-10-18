package external

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"strings"
	"time"
)

type AssumeRoleRequest struct {
	RoleName string
}

type AssumeRoleResponse struct {
	Expiration    int
	PrivateKey    string
	ProjectID     string
	PublicKey     string
	RoleName      string
	SecurityToken string
	UHostID       string
}

func AssumeRole(req AssumeRoleRequest) (cred auth.Credential, err error) {
	return defaultClient{}.AssumeRole(req)
}

type metadataResolver interface {
	SendRequest(path string) (string, error)
}

type defaultClient struct {
	resolver metadataResolver
}

func (client defaultClient) AssumeRole(req AssumeRoleRequest) (cred auth.Credential, err error) {
	path := "/iam/token"
	if len(req.RoleName) != 0 {
		path += fmt.Sprintf("%s/%s", path, req.RoleName)
	}

	resp, err := client.resolver.SendRequest(path)
	if err != nil {
		return cred, err
	}

	var assumeRole AssumeRoleResponse
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&assumeRole); err != nil {
		return cred, errors.Errorf("failed to decode sts credential, %s", err)
	}

	cred.CanExpire = true
	cred.Expires = time.Unix(int64(assumeRole.Expiration), 0)
	cred.PrivateKey = assumeRole.PrivateKey
	cred.PublicKey = assumeRole.PublicKey
	cred.SecurityToken = assumeRole.SecurityToken
	return cred, nil
}
