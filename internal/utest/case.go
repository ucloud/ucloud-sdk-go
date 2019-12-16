package utest

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/ucloud/error"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestValidator is the validator function
type TestValidator func(interface{}, error) error

// TestCase is the case definition of test case
type TestCase struct {
	Invoker       func() (interface{}, error)
	Validators    []TestValidator
	MaxRetries    int
	RetryInterval time.Duration
	T             *testing.T
}

// Run will run the test case with retry
func (t *TestCase) Run() (interface{}, error) {
	for i := 0; i < t.MaxRetries+1; i++ {
		resp, err := t.Invoker()
		if e, ok := err.(uerr.Error); ok && e.Name() == uerr.ErrSendRequest {
			assert.NoError(t.T, err)
			assert.FailNow(t.T, "got unexcepted error")
			return nil, nil
		}

		isValidateFailed := false
		for _, validator := range t.Validators {
			if err := validator(resp, err); err != nil {
				isValidateFailed = true
				if i != t.MaxRetries {
					fmt.Printf("skip validate error for retring, %s\n", err)
					continue
				}
				assert.NoError(t.T, err)
			}
		}

		if !isValidateFailed {
			return resp, nil
		}

		if i == t.MaxRetries && isValidateFailed {
			if err != nil {
				// log and fail now
				t.T.Fatal(err)
			}
			t.T.FailNow()
		}

		<-time.After(t.RetryInterval)
	}

	// never reached here
	return nil, nil
}
