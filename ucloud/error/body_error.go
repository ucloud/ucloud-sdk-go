package uerr

import "fmt"

var (
	ErrResponseBodyError     = "body.ResponseBodyError"
	ErrNullResponseBodyError = "body.NullResponseBodyError"
)

type BodyError struct {
	name      string
	err       error
	body      string
	retryable bool
}

func (e BodyError) Error() string {
	return fmt.Sprintf("response body\n[%s] got error, %s", e.body, e.err)
}

// NewResponseBodyError will create a new response body error
func NewResponseBodyError(err error, body string) BodyError {
	return BodyError{
		name:      ErrResponseBodyError,
		err:       err,
		body:      body,
		retryable: false,
	}
}

// NewResponseBodyError will create a new response body error
func NewNullResponseBodyError() BodyError {
	return BodyError{
		name:      ErrNullResponseBodyError,
		err:       fmt.Errorf("got response body is nil"),
		body:      "",
		retryable: false,
	}
}

// Name will return error name
func (e BodyError) Name() string {
	return e.name
}

// Code will return server code
func (e BodyError) Code() int {
	return -1
}

// StatusCode will return http status code
func (e BodyError) StatusCode() int {
	return 0
}

// Message will return message
func (e BodyError) Message() string {
	return e.err.Error()
}

// OriginError will return the origin error that caused by
func (e BodyError) OriginError() error {
	return e.err
}

// Retryable will return if the error is retryable
func (e BodyError) Retryable() bool {
	return e.retryable
}
