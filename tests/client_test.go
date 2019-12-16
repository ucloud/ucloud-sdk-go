package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

func TestInvokeHandler(t *testing.T) {
	var isHandled bool
	fn := func(c *ucloud.Client, req request.Common, resp response.Common, err error) (response.Common, error) {
		isHandled = true
		return nil, nil
	}

	err := uhostClient.Client.AddResponseHandler(fn)
	assert.NoError(t, err)

	req := uhostClient.NewDescribeImageRequest()
	_, err = uhostClient.DescribeImage(req)
	assert.NoError(t, err)
	assert.True(t, isHandled)
}
