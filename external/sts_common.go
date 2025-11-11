package external

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/metadata"
)

// STS related constants
const (
	// internalBaseUrl is the base URL for UCloud internal API endpoint
	internalBaseUrl = "http://api.service.ucloud.cn"

	// Metadata service paths
	metadataRegionPath = "/meta-data/latest/region"
	metadataZonePath   = "/meta-data/latest/availability-zone"

	// STS API paths
	stsV1BasePath = "/meta-data/v1/uam/security-credentials"
	stsV2BasePath = "/meta-data/v2/iam/security-credentials"
)

// metadataProvider is a common interface for metadata clients
// Both metadata.DefaultClient and mock clients implement this interface
type metadataProvider interface {
	SendRequest(string) (string, error)
	SetHttpClient(http.Client) error
}

// createMetadataClient creates a new metadata client with HTTP client configured
func createMetadataClient() (*metadata.DefaultClient, error) {
	httpClient := http.NewHttpClient()
	client := &metadata.DefaultClient{}
	err := client.SetHttpClient(&httpClient)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set HTTP client for metadata client")
	}
	return client, nil
}

// fetchRegionAndZone retrieves region and zone information from metadata service
// Accepts both metadataProvider interface and *metadata.DefaultClient
func fetchRegionAndZone(client interface{}) (region, zone string, err error) {
	var regionResp, zoneResp string

	// Type switch to handle different client types
	switch c := client.(type) {
	case metadataProvider:
		regionResp, err = c.SendRequest(metadataRegionPath)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to get region from metadata")
		}

		zoneResp, err = c.SendRequest(metadataZonePath)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to get zone from metadata")
		}
	case *metadata.DefaultClient:
		regionResp, err = c.SendRequest(metadataRegionPath)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to get region from metadata")
		}

		zoneResp, err = c.SendRequest(metadataZonePath)
		if err != nil {
			return "", "", errors.Wrap(err, "failed to get zone from metadata")
		}
	default:
		return "", "", errors.New("unsupported client type")
	}

	return strings.TrimSpace(regionResp), strings.TrimSpace(zoneResp), nil
}
