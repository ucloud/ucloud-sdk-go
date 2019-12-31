package validation

import (
	"github.com/pkg/errors"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
)

var Builtins = NewTestContext(nil)

type TestComparator struct {
	Comparators *Comparators
}

func NewTestContext(m map[string]CompareFunc) TestComparator {
	comparators := NewComparators()
	for k, v := range m {
		comparators.comparators[k] = v
	}

	return TestComparator{
		Comparators: &comparators,
	}
}

// NewValidator will return new validator to validate value is expected
func (ctx *TestComparator) NewValidator(valuePath string, expected interface{}, comparator string) driver.TestValidator {
	comparatorFunc := ctx.Comparators.Get(comparator)
	if comparatorFunc == nil {
		return func(i interface{}) error {
			return errors.Errorf("the NewValidator cannot get comparator func from %s", comparator)
		}
	}

	return func(resp interface{}) error {
		v, err := utils.GetValue(resp, valuePath)
		if err != nil {
			return errors.Errorf("cannot get value from %s, %s", valuePath, err)
		}
		err = comparatorFunc(v, expected)
		if err != nil {
			if IsNotExpectedError(err) {
				return errors.Errorf("want %s %#v %s %#v, but false", valuePath, v, comparator, expected)
			}
			return errors.Errorf("want %s: %#v %s %#v, %s", valuePath, v, comparator, expected, err)
		}
		return nil
	}
}
