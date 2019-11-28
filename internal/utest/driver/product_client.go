package driver

import (
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type serviceClient interface {
	AddResponseHandler(ucloud.ResponseHandler) error
}

// ServiceFactory is the validator function
type serviceFactory func(*ucloud.Config, *auth.Credential) serviceClient

func newServiceClient(product string, cfg *ucloud.Config, cred *auth.Credential) serviceClient {
	for k, v := range serviceFactoryMap {
		if k == product {
			c := v(cfg, cred)
			return c
		} else if len(product) == 0 {
			c := ucloud.NewClient(cfg, cred)
			return c
		}
	}

	return nil
}

var serviceFactoryMap = map[string]serviceFactory{
	"UHost": func(cfg *ucloud.Config, cred *auth.Credential) serviceClient {
		return uhost.NewClient(cfg, cred)
	},
}
