package external

import "os"

const (
	UCloudPublicKeyEnvVar = "UCLOUD_PUBLIC_KEY"

	UCloudPrivateKeyEnvVar = "UCLOUD_PRIVATE_KEY"

	UCloudProjectIdEnvVar = "UCLOUD_PROJECT_ID"

	UCloudRegionEnvVar = "UCLOUD_REGION"

	UCloudSharedProfileEnvVar = "UCLOUD_PROFILE"

	UCloudSharedConfigFileEnvVar = "UCLOUD_SHARED_CONFIG_FILE"

	UCloudSharedCredentialFileEnvVar = "UCLOUD_SHARED_CREDENTIAL_FILE"
)

func loadEnvConfig() (*config, error) {
	return &config{
		PublicKey:            os.Getenv(UCloudPublicKeyEnvVar),
		PrivateKey:           os.Getenv(UCloudPrivateKeyEnvVar),
		ProjectId:            os.Getenv(UCloudProjectIdEnvVar),
		Region:               os.Getenv(UCloudRegionEnvVar),
		Profile:              os.Getenv(UCloudSharedProfileEnvVar),
		SharedConfigFile:     os.Getenv(UCloudSharedConfigFileEnvVar),
		SharedCredentialFile: os.Getenv(UCloudSharedCredentialFileEnvVar),
	}, nil
}
