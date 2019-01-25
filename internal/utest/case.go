package utest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
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

		isFailed := false
		for _, validator := range t.Validators {
			if err := validator(resp, err); err != nil {
				isFailed = true
				if i != t.MaxRetries {
					t.T.Logf("skip validate error for retring, %s", err)
					continue
				}
				assert.NoError(t.T, err)
			}
		}

		if i == t.MaxRetries || !isFailed {
			if err != nil && !uerr.IsCodeError(err) {
				assert.NoError(t.T, err)
				return resp, err
			}
			return resp, nil
		}

		<-time.After(t.RetryInterval)
	}

	// never reached here
	return nil, nil
}
