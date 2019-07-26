package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseAccessor(t *testing.T) {
	resp := CommonBase{
		Action:  "Test",
		RetCode: 42,
		Message: "foo",
	}
	assert.Equal(t, "Test", resp.GetAction())
	assert.Equal(t, 42, resp.GetRetCode())
	assert.Equal(t, "foo", resp.GetMessage())

	resp.SetRequestUUID("foo")
	assert.Equal(t, "foo", resp.GetRequestUUID())

	resp.SetRequest(nil)
	assert.Equal(t, nil, resp.GetRequest())
}
