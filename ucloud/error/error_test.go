package uerr

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientError(t *testing.T) {
	originErr := errors.New("o")
	err := NewClientError(ErrUnexpected, originErr)

	assert.Equal(t, err.Name(), ErrUnexpected)
	assert.Equal(t, err.Code(), -1)
	assert.Equal(t, err.OriginError(), originErr)
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 0)
	assert.NotZero(t, err.Error())
	assert.NotZero(t, err.Message())

	assert.True(t, IsNetworkError(net.UnknownNetworkError("net error")))
	assert.False(t, IsNetworkError(nil))
	assert.True(t, IsNetworkError(errors.New("net/http: request canceled")))

	assert.True(t, isRetryableName(ErrNetwork))
	assert.False(t, isRetryableName(ErrUnexpected))
}

func TestServerError(t *testing.T) {
	err := NewServerCodeError(160, "o")

	assert.NotZero(t, err.Error())
	assert.NotZero(t, err.Message())
	assert.True(t, IsCodeError(err))
	assert.Equal(t, err.Name(), ErrRetCode)
	assert.Equal(t, err.Code(), 160)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 200)

	err = NewServerStatusError(400, "o")

	assert.NotZero(t, err.Error())
	assert.NotZero(t, err.Message())
	assert.False(t, IsCodeError(err))
	assert.Equal(t, err.Name(), ErrHTTPStatus)
	assert.Equal(t, err.Code(), -1)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 400)

	err = NewResponseBodyError(nil, "o")

	assert.NotZero(t, err.Error())
	assert.NotZero(t, err.Message())
	assert.False(t, IsCodeError(err))
	assert.Equal(t, err.Name(), ErrResponseBodyError)
	assert.Equal(t, err.Code(), 0)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 200)

	err = NewEmptyResponseBodyError()

	assert.NotZero(t, err.Error())
	assert.NotZero(t, err.Message())
	assert.False(t, IsCodeError(err))
	assert.Equal(t, err.Name(), ErrEmptyResponseBodyError)
	assert.Equal(t, err.Code(), 0)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), false)
	assert.Equal(t, err.StatusCode(), 200)
}

func TestRetryableError(t *testing.T) {
	err := NewRetryableError(errors.New("o"))

	assert.Equal(t, err.Name(), ErrUnexpected)
	assert.Equal(t, err.Code(), -1)
	assert.Error(t, err.OriginError())
	assert.Equal(t, err.Retryable(), true)
	assert.Equal(t, err.StatusCode(), 0)

	err = NewRetryableError(NewServerCodeError(160, "o"))
	assert.Equal(t, err.Retryable(), true)

	err = NewRetryableError(NewServerStatusError(400, "o"))
	assert.Equal(t, err.Retryable(), true)

	err = NewRetryableError(NewClientError(ErrNetwork, errors.New("o")))
	assert.Equal(t, err.Retryable(), true)

	err = NewServerStatusError(400, "o")
	assert.Equal(t, err.Retryable(), false)

	err = NewServerStatusError(429, "o")
	assert.Equal(t, err.Retryable(), true)
}

func isRetryable(err error) bool {
	if e, ok := err.(Error); ok {
		return e.Retryable()
	}
	return false
}

func TestNonRetryableError(t *testing.T) {
	var err error
	unexpected := errors.New("o")

	err = NewNonRetryableError(NewClientError(ErrUnexpected, unexpected))
	assert.False(t, isRetryable(err))

	err = NewNonRetryableError(NewServerCodeError(1, "foo"))
	assert.False(t, isRetryable(err))

	err = NewNonRetryableError(unexpected)
	assert.False(t, isRetryable(err))
}
