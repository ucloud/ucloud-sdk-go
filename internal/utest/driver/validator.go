package driver

import (
	"github.com/pkg/errors"
	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

// TestValidator is the validator function
type TestValidator func(interface{}) error

// NewValidator will return new validator to validate value is expected
func NewValidator(valuePath string, expected interface{}, comparator string) TestValidator {
	comparators := utest.NewComparators()
	comparatorFunc := comparators.Get(comparator)
	return NewValidatorFromComparator(valuePath, expected, comparatorFunc)
}

func NewValidatorFromComparator(valuePath string, expected interface{}, comparator utest.CompareFunc) TestValidator {
	return func(resp interface{}) error {
		v, err := utest.GetValue(resp, valuePath)
		if err != nil {
			return errors.Errorf("cannot get value from %s, %s", valuePath, err)
		}

		ok, err := comparator(v, expected)
		if err != nil {
			return errors.Errorf("want %s: %#v %s %#v, %s", valuePath, v, comparator, expected, err)
		}

		if !ok {
			return errors.Errorf("want %s %#v %s %#v, but false", valuePath, v, comparator, expected)
		}

		return nil
	}
}

func JSONEqual(value interface{}, expected interface{}) (bool, error) {
	return false, nil
}
