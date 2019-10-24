package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ucloud-sdk-go/private/protocol/http"
	"github.com/ucloud/ucloud-sdk-go/ucloud/helpers/mock"
)

type MockedMetadataCase struct {
	MockedVector string
	MockedError  error
	Do           func(Client) (interface{}, error)
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

	mockedClient := DefaultClient{}
	_ = mockedClient.SetHttpClient(httpClient)

	resp, err := mockedCase.Do(mockedClient)
	if mockedCase.GoldenError {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, mockedCase.Golden, resp)
	}
}

func TestMetadataClient(t *testing.T) {
	cases := []MockedMetadataCase{
		{
			MockedVector: mockedMetadata,
			Do: func(client Client) (i interface{}, e error) {
				return client.GetInstanceIdentityDocument()
			},
			Golden:      goldenMetadata,
			GoldenError: false,
		},
		{
			MockedVector: mockedMetadataError,
			Do: func(client Client) (i interface{}, e error) {
				return client.GetInstanceIdentityDocument()
			},
			Golden:      nil,
			GoldenError: true,
		},
		{
			MockedError: http.NewStatusError(400, "Bad Request"),
			Do: func(client Client) (i interface{}, e error) {
				return client.GetInstanceIdentityDocument()
			},
			Golden:      nil,
			GoldenError: true,
		},
		{
			MockedVector: "metadata",
			Do: func(client Client) (i interface{}, e error) {
				return client.GetMetadata("/cloud-name")
			},
			Golden:      "metadata",
			GoldenError: false,
		},
		{
			MockedVector: "user-data",
			Do: func(client Client) (i interface{}, e error) {
				return client.GetUserData()
			},
			Golden:      "user-data",
			GoldenError: false,
		},
		{
			MockedVector: "vendor-data",
			Do: func(client Client) (i interface{}, e error) {
				return client.GetVendorData()
			},
			Golden:      "vendor-data",
			GoldenError: false,
		},
	}

	for _, mockedCase := range cases {
		RunMockedMetadataCase(t, mockedCase)
	}
}

func TestClientSetup(t *testing.T) {
	client := NewClient()
	assert.NotZero(t, client)
}

var goldenMetadata = Metadata{
	AvailabilityZone: "hk-02",
	CloudName:        "ucloud",
	InstanceId:       "uhost-hhqeihh5",
	LocalHostname:    "10-8-134-210",
	NetworkConfig: MDNetworkConfig{
		Ethernets: map[string]MDEthernet{
			"eth0": {
				Addresses: []string{"10.8.134.210/16"},
				Gateway4:  "10.8.0.1",
				Match:     MDMatch{MacAddress: "52:54:00:71:04:e1"},
				MTU:       1454,
				NameServers: MDNameServers{
					Addresses: []string{"10.8.255.1", "10.8.255.2", "114.114.114.114"},
				},
			},
		},
		Version: 2,
	},
	Platform:      "uhost",
	PublicSSHKeys: []string{},
	Region:        "hk",
	UHost: MDUHost{
		CPU: 1,
		Disks: []MDDisks{
			{
				BackupType: "",
				DiskId:     "bsi-ifmi12",
				DiskType:   "CLOUD_SSD",
				Drive:      "vda",
				Encrypted:  false,
				IsBoot:     true,
				Name:       "系统盘_现网调试勿删",
				Size:       20,
			},
		},
		GPU:            0,
		Hotplug:        false,
		ImageId:        "uimage-iwxgro",
		IsolationGroup: "",
		MachineType:    "N",
		Memory:         1024,
		Name:           "现网调试勿删",
		NetCapability:  "Normal",
		NetworkInterfaces: []MDNetworkInterfaces{
			{
				IPs: []MDIPs{
					{IPAddress: "10.8.134.210", Type: "Private", Bandwidth: 0, IPId: ""},
					{IPAddress: "152.32.135.18", Type: "International", Bandwidth: 11, IPId: "eip-t4frnxjn"},
				},
				Mac:      "52:54:00:71:04:E1",
				SubnetId: "subnet-jodfnu",
				VpcId:    "uvnet-vxil3h",
			},
		},
		OsName:    "CentOS 7.6 64位",
		ProjectId: "org-xp2ucn",
		Region:    "hk",
		Remark:    "",
		Tag:       "Default",
		UHostId:   "uhost-hhqeihh5",
		Zone:      "hk-02",
	},
}

const mockedMetadata = `
{
    "availability-zone": "hk-02",
    "cloud-name": "ucloud",
    "instance-id": "uhost-hhqeihh5",
    "local-hostname": "10-8-134-210",
    "network-config": {
        "ethernets": {
            "eth0": {
                "addresses": [
                    "10.8.134.210/16"
                ],
                "gateway4": "10.8.0.1",
                "match": {
                    "macaddress": "52:54:00:71:04:e1"
                },
                "mtu": 1454,
                "nameservers": {
                    "addresses": [
                        "10.8.255.1",
                        "10.8.255.2",
                        "114.114.114.114"
                    ]
                }
            }
        },
        "version": 2
    },
    "platform": "uhost",
    "public-ssh-keys": [],
    "region": "hk",
    "uhost": {
        "cpu": 1,
        "disks": [
            {
                "backup-type": "",
                "disk-id": "bsi-ifmi12",
                "disk-type": "CLOUD_SSD",
                "drive": "vda",
                "encrypted": false,
                "is-boot": true,
                "name": "系统盘_现网调试勿删",
                "size": 20
            }
        ],
        "gpu": 0,
        "hotplug": false,
        "image-id": "uimage-iwxgro",
        "isolation-group": "",
        "machine-type": "N",
        "memory": 1024,
        "name": "现网调试勿删",
        "net-capability": "Normal",
        "network-interfaces": [
            {
                "ips": [
                    {
                        "ip-address": "10.8.134.210",
                        "type": "Private"
                    },
                    {
                        "bandwidth": 11,
                        "ip-address": "152.32.135.18",
                        "ip-id": "eip-t4frnxjn",
                        "type": "International"
                    }
                ],
                "mac": "52:54:00:71:04:E1",
                "subnet-id": "subnet-jodfnu",
                "vpc-id": "uvnet-vxil3h"
            }
        ],
        "os-name": "CentOS 7.6 64位",
        "project-id": "org-xp2ucn",
        "region": "hk",
        "remark": "",
        "tag": "Default",
        "uhost-id": "uhost-hhqeihh5",
        "zone": "hk-02"
    }
}
`

const mockedMetadataError = `
invalid data
`
