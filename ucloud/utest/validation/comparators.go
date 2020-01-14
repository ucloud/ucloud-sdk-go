package validation

import (
	"math"
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// CompareFunc is the function definition of test comparator without type-system
type CompareFunc func(value interface{}, expected interface{}) error

// Comparators is a collection of compare functions
type Comparators struct {
	comparators map[string]CompareFunc
}

type NotExpectedError struct {
	message string
}

func (e *NotExpectedError) Error() string {
	return e.message
}

func NewNotExpectedError() error {
	return &NotExpectedError{"is not expected error"}
}

func IsNotExpectedError(err error) bool {
	if _, ok := err.(*NotExpectedError); ok {
		return true
	}
	return false
}

// NewComparators will return Comparators
func NewComparators() Comparators {
	comparators := map[string]CompareFunc{
		"eq":     eq,
		"equals": eq,
		"==":     eq,

		"abs_eq":     absEq,
		"abs_equals": absEq,

		"lt":        lt,
		"less_than": lt,

		"le":                  le,
		"less_than_or_equals": le,

		"gt":           gt,
		"greater_than": gt,

		"ge":                     ge,
		"greater_than_or_equals": ge,

		"ne":         ne,
		"not_equals": ne,

		"str_eq":        strEq,
		"string_equals": strEq,

		"float_eq":     floatEq,
		"float_equals": floatEq,

		"len_eq":        lenEq,
		"length_equals": lenEq,
		"count_eq":      lenEq,

		"len_gt":              lenGt,
		"count_gt":            lenGt,
		"length_greater_than": lenGt,
		"count_greater_than":  lenGt,

		"len_ge":                        lenGe,
		"count_ge":                      lenGe,
		"length_greater_than_or_equals": lenGe,
		"count_greater_than_or_equals":  lenGe,

		"len_lt":           lenLt,
		"count_lt":         lenLt,
		"length_less_than": lenLt,
		"count_less_than":  lenLt,

		"len_le":                     lenLe,
		"count_le":                   lenLe,
		"length_less_than_or_equals": lenLe,
		"count_less_than_or_equals":  lenLe,

		"contains":            contains,
		"contained_by":        containedBy,
		"type":                typeEq,
		"regex":               regexMatch,
		"startswith":          startsWith,
		"endswith":            endsWith,
		"object_contains":     objectContains,
		"object_not_contains": objectNotContains,
	}
	return Comparators{comparators}
}

// Get will return a comparator function by name
func (c *Comparators) Get(name string) CompareFunc {
	if fn, ok := c.comparators[name]; ok {
		return fn
	}
	return nil
}

func eq(a, b interface{}) error {
	return strEq(a, b)
}

func ne(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if aVal == bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func absEq(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if math.Abs(aVal) != math.Abs(bVal) {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lt(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if aVal >= bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func le(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if aVal > bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func gt(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if aVal <= bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func ge(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if aVal < bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func floatEq(a, b interface{}) error {
	return checkFloats(a, b, func(aVal, bVal float64) error {
		if aVal != bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func strEq(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if aVal != bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lenEq(a, b interface{}) error {
	return checkLens(a, b, func(aVal, bVal int) error {
		if aVal != bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lenGt(a, b interface{}) error {
	return checkLens(a, b, func(aVal, bVal int) error {
		if aVal <= bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lenGe(a, b interface{}) error {
	return checkLens(a, b, func(aVal, bVal int) error {
		if aVal < bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lenLt(a, b interface{}) error {
	return checkLens(a, b, func(aVal, bVal int) error {
		if aVal >= bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func lenLe(a, b interface{}) error {
	return checkLens(a, b, func(aVal, bVal int) error {
		if aVal > bVal {
			return NewNotExpectedError()
		}
		return nil
	})
}

func contains(a, b interface{}) error {
	f := reflect.ValueOf(a)
	switch f.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < f.Len(); i++ {
			item := f.Index(i)
			for item.Kind() == reflect.Ptr || item.Kind() == reflect.Interface {
				if f.IsNil() {
					break
				}
				item = item.Elem()
			}
			err := strEq(item.String(), b)
			if err != nil {
				if !IsNotExpectedError(err) {
					return err
				} else {
					continue
				}
			}
			return nil
		}
		return NewNotExpectedError()
	case reflect.String:
		return checkStrings(a, b, func(aVal, bVal string) error {
			if !strings.Contains(aVal, bVal) {
				return NewNotExpectedError()
			}
			return nil
		})
	default:
		return errors.Errorf("val must be contained by an iterable type, got %s", f.String())
	}
}

func containedBy(a, b interface{}) error {
	return contains(b, a)
}

func typeEq(a, b interface{}) error {
	typeVal, _ := toString(b)
	if reflect.ValueOf(a).Type().String() != typeVal {
		return NewNotExpectedError()
	}
	return nil
}

func regexMatch(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		matched, err := regexp.MatchString(bVal, aVal)
		if err != nil {
			return err
		}

		if !matched {
			return NewNotExpectedError()
		}

		return nil
	})
}

func startsWith(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if !strings.HasPrefix(aVal, bVal) {
			return NewNotExpectedError()
		}
		return nil
	})
}

func endsWith(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if !strings.HasSuffix(aVal, bVal) {
			return NewNotExpectedError()
		}
		return nil
	})
}

func objectContains(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if !strings.Contains(aVal, bVal) {
			return NewNotExpectedError()
		}
		return nil
	})
}

func objectNotContains(a, b interface{}) error {
	return checkStrings(a, b, func(aVal, bVal string) error {
		if strings.Contains(aVal, bVal) {
			return NewNotExpectedError()
		}
		return nil
	})
}

var DefaultComparators = NewComparators()
