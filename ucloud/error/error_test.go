package uerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientError(t *testing.T) {
	originErr := errors.New("o")
	err := NewClientError(ErrUnexcepted, originErr)

	assert.Equal(t, err.Name(), ErrUnexcepted)
	assert.Equal(t, err.Code(), -1)
	assert.Equal(t, err.OriginError(), originErr)
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 0)
}

func TestServerError(t *testing.T) {
	err := NewServerCodeError(160, "o")

	assert.Equal(t, err.Name(), ErrRetCode)
	assert.Equal(t, err.Code(), 160)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 200)

	err = NewServerStatusError(400, "o")

	assert.Equal(t, err.Name(), ErrHTTPStatus)
	assert.Equal(t, err.Code(), -1)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 400)
}

func TestRetryableError(t *testing.T) {
	err := NewRetryableError(errors.New("o"))

	assert.Equal(t, err.Name(), ErrUnexcepted)
	assert.Equal(t, err.Code(), -1)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), true)
	assert.Equal(t, err.StatusCode(), 0)

	err = NewRetryableError(NewServerCodeError(160, "o"))
	assert.Equal(t, err.Retryable(), true)

	err = NewRetryableError(NewServerStatusError(400, "o"))
	assert.Equal(t, err.Retryable(), true)

	err = NewServerStatusError(400, "o")
	assert.Equal(t, err.Retryable(), false)

	err = NewServerStatusError(429, "o")
	assert.Equal(t, err.Retryable(), true)
}

func TestNonRetryableError(t *testing.T) {
	err := NewNonRetryableError(NewRetryableError(errors.New("o")))
	assert.Equal(t, err.Retryable(), false)
}
