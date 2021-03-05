package utest

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// TestContext is the fixture of sdk acceptance testing
type TestContext struct {
	T *testing.T

	Vars        map[string]interface{}
	Comparators *Comparators
}

// NewTestContext will return a new TestContext
func NewTestContext() TestContext {
	comparators := NewComparators()

	return TestContext{
		Vars:        make(map[string]interface{}),
		Comparators: &comparators,
	}
}

// SetVar will set the variable of test context
func (ctx *TestContext) SetVar(name string, value interface{}) {
	ctx.Vars[name] = value
}

// GetVar will return the variable of test context
func (ctx *TestContext) GetVar(name string) interface{} {
	if v, ok := ctx.Vars[name]; ok {
		return v
	}
	return ""
}

// NewValidator will return new validator to validate value is expected
func (ctx *TestContext) NewValidator(valuePath string, expected interface{}, comparator string) TestValidator {
	compratorFunc := ctx.Comparators.Get(comparator)
	return func(resp interface{}, respErr error) error {
		v, err := GetValue(resp, valuePath)
		if err != nil {
			return errors.Errorf("cannot get value from %s, %s", valuePath, err)
		}

		ok, err := compratorFunc(v, expected)
		if err != nil {
			return errors.Errorf("want %s: %#v %s %#v, %s", valuePath, v, comparator, expected, err)
		}

		if !ok {
			return errors.Errorf("want %s %#v %s %#v, but false", valuePath, v, comparator, expected)
		}

		return nil
	}
}

// MustString will check error is nil and return the string value
func (ctx *TestContext) MustString(v string, err error) string {
	assert.NoError(ctx.T, err)
	return v
}

// Must will check error is nil and return the value
func (ctx *TestContext) Must(v interface{}, err error) interface{} {
	assert.NoError(ctx.T, err)
	return v
}

// NoError will check error is nil
func (ctx *TestContext) NoError(err error) {
	assert.NoError(ctx.T, err)
}
