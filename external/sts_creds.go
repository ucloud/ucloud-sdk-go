package external

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type AssumeRoleRequest struct {
	RoleName string
}

func LoadSTSConfig(req AssumeRoleRequest) (ConfigProvider, error) {
	client, err := createMetadataClient()
	if err != nil {
		return nil, err
	}
	return loadSTSConfig(req, client)
}

type assumeRoleData struct {
	Expiration    int
	PrivateKey    string
	ProjectID     string
	PublicKey     string
	CharacterName string
	SecurityToken string
	UHostID       string
	UPHostId      string
}

type assumeRoleResponse struct {
	RetCode int
	Message string
	Data    assumeRoleData
}

func loadSTSConfig(req AssumeRoleRequest, client metadataProvider) (ConfigProvider, error) {
	// Build API path for V1
	path := stsV1BasePath
	if len(req.RoleName) != 0 {
		path += "/" + req.RoleName
	}

	// Request STS credentials
	resp, err := client.SendRequest(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send STS V1 request")
	}

	// Parse response
	var roleResp assumeRoleResponse
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&roleResp); err != nil {
		return nil, errors.Errorf("failed to decode sts credential, %s", err)
	}

	// Fetch region and zone from metadata
	region, zone, err := fetchRegionAndZone(client)
	if err != nil {
		return nil, err
	}

	// Build STS config
	roleData := roleResp.Data
	stsConfig := &config{
		CanExpire:     true,
		Expires:       time.Unix(int64(roleData.Expiration), 0),
		PrivateKey:    roleData.PrivateKey,
		PublicKey:     roleData.PublicKey,
		SecurityToken: roleData.SecurityToken,
		ProjectId:     roleData.ProjectID,
		Region:        region,
		Zone:          zone,
		BaseUrl:       internalBaseUrl,
	}
	return stsConfig, nil
}
