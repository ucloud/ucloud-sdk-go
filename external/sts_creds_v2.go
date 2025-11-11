package external

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

// AssumeRoleV2Request is the request for STS V2 API
// Supports V2 API parameters via HTTP headers
type AssumeRoleV2Request struct {
	// RoleName is the name of the role to assume (required)
	RoleName string

	// DurationSeconds specifies the duration of the STS token in seconds
	// Default is 3600 seconds (1 hour) if not specified
	// Valid range: 900-3600 (15 minutes to 1 hour)
	DurationSeconds int

	// SessionName is the name of the session (optional)
	// Used for auditing and tracking purposes
	SessionName string

	// RolePolicy is an additional policy to further restrict the STS token permissions
	// Should be base64 encoded JSON string (optional)
	RolePolicy string
}

// LoadSTSConfigV2 loads STS V2 configuration from UCloud metadata service
func LoadSTSConfigV2(req AssumeRoleV2Request) (ConfigProvider, error) {
	client, err := createMetadataClient()
	if err != nil {
		return nil, err
	}
	return loadSTSConfigV2(req, client)
}

// assumeRoleV2Data represents the response data from STS V2 API
type assumeRoleV2Data struct {
	// AccessKeyId is the temporary access key ID
	AccessKeyId string

	// AccessKeySecret is the temporary access key secret
	AccessKeySecret string

	// SecurityToken is the security token for temporary credentials
	SecurityToken string

	// Expiration is the expiration time in ISO 8601 format (e.g., "2025-10-20T06:51:25Z")
	Expiration string

	// CharacterName is the role URN (e.g., "ucs:iam::60036:role/ucs-service-role/ServiceRoleForUK8S")
	CharacterName string

	// ProjectID is the project ID
	ProjectID string

	// HostID is the host ID (may be UHostID or UPHostID)
	HostID string

	// UHostID is the UHost instance ID
	UHostID string

	// UPHostId is the UPHost instance ID (physical host)
	UPHostId string
}

// assumeRoleV2Response represents the full response from STS V2 API
type assumeRoleV2Response struct {
	RetCode int
	Message string
	Data    assumeRoleV2Data
}

func loadSTSConfigV2(req AssumeRoleV2Request, client *metadata.DefaultClient) (ConfigProvider, error) {
	// Build API path for V2
	path := stsV2BasePath
	if len(req.RoleName) != 0 {
		path += "/" + req.RoleName
	}

	// Prepare headers for V2 API
	headers := make(map[string]string)
	if req.DurationSeconds > 0 {
		headers["Duration-Seconds"] = fmt.Sprintf("%d", req.DurationSeconds)
	}
	if len(req.SessionName) > 0 {
		headers["Session-Name"] = req.SessionName
	}
	if len(req.RolePolicy) > 0 {
		headers["Role-Policy"] = req.RolePolicy
	}

	// Send request with headers
	var resp string
	var err error

	// Use SendRequestWithHeaders if headers are provided
	if len(headers) > 0 {
		resp, err = client.SendRequestWithHeaders(path, headers)
	} else {
		// Fallback to regular SendRequest if no headers needed
		resp, err = client.SendRequest(path)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to send STS V2 request")
	}

	// Parse response
	var roleResp assumeRoleV2Response
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&roleResp); err != nil {
		return nil, errors.Errorf("failed to decode STS V2 credential, %s", err)
	}

	if roleResp.RetCode != 0 {
		return nil, errors.Errorf("STS V2 API returned error: RetCode=%d, Message=%s", roleResp.RetCode, roleResp.Message)
	}

	// Fetch region and zone from metadata
	region, zone, err := fetchRegionAndZone(client)
	if err != nil {
		return nil, err
	}

	// Parse expiration time (ISO 8601 format)
	roleData := roleResp.Data
	expirationTime, err := parseISO8601Time(roleData.Expiration)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse expiration time: %s", roleData.Expiration)
	}

	// Map V2 API fields to SDK config
	// AccessKeyId -> PublicKey
	// AccessKeySecret -> PrivateKey
	stsConfig := &config{
		CanExpire:     true,
		Expires:       expirationTime,
		PublicKey:     roleData.AccessKeyId,     // AccessKeyId maps to PublicKey
		PrivateKey:    roleData.AccessKeySecret, // AccessKeySecret maps to PrivateKey
		SecurityToken: roleData.SecurityToken,
		ProjectId:     roleData.ProjectID,
		Region:        region,
		Zone:          zone,
		BaseUrl:       internalBaseUrl,
	}

	return stsConfig, nil
}

// parseISO8601Time parses ISO 8601 time format to time.Time
// Supports formats: "2025-10-20T06:51:25Z", "2025-10-20T06:51:25+08:00"
func parseISO8601Time(timeStr string) (time.Time, error) {
	// Try RFC3339 format first (standard ISO 8601)
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t, nil
	}

	// Try alternative formats
	formats := []string{
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05Z07:00",
	}

	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.Errorf("unable to parse time string: %s", timeStr)
}
