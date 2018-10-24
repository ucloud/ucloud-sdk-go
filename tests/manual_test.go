package tests

import (
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

func TestVPCSubnetUpdate(t *testing.T) {
	req := vpcClient.NewUpdateSubnetAttributeRequest()
	req.SubnetId = ucloud.String("uvnet-e3mi5v")
	req.Name = ucloud.String("hello")

	resp, err := vpcClient.UpdateSubnetAttribute(req)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", resp)
}
