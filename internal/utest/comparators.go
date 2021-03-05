package utest

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// CompareFunc is the function definition of test comparator without type-system
type CompareFunc func(interface{}, interface{}) (bool, error)

// Comparators is a collection of compare functions
type Comparators struct {
	comparators map[string]CompareFunc
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

func eq(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return aVal == bVal, nil })
}

func ne(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return aVal != bVal, nil })
}

func absEq(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal == bVal, nil })
}

func lt(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal < bVal, nil })
}

func le(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal <= bVal, nil })
}

func gt(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal > bVal, nil })
}

func ge(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal >= bVal, nil })
}

func floatEq(a, b interface{}) (bool, error) {
	return checkFloats(a, b, func(aVal, bVal float64) (bool, error) { return aVal == bVal, nil })
}

func strEq(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return aVal == bVal, nil })
}

func lenEq(a, b interface{}) (bool, error) {
	return checkLens(a, b, func(aVal, bVal int) (bool, error) { return aVal == bVal, nil })
}

func lenGt(a, b interface{}) (bool, error) {
	return checkLens(a, b, func(aVal, bVal int) (bool, error) { return aVal > bVal, nil })
}

func lenGe(a, b interface{}) (bool, error) {
	return checkLens(a, b, func(aVal, bVal int) (bool, error) { return aVal >= bVal, nil })
}

func lenLt(a, b interface{}) (bool, error) {
	return checkLens(a, b, func(aVal, bVal int) (bool, error) { return aVal < bVal, nil })
}

func lenLe(a, b interface{}) (bool, error) {
	return checkLens(a, b, func(aVal, bVal int) (bool, error) { return aVal <= bVal, nil })
}

func contains(a, b interface{}) (bool, error) {
	rv := reflect.ValueOf(a)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			isEq, err := strEq(rv.Index(i).String(), b)
			if isEq {
				return true, nil
			}

			if err != nil {
				return false, err
			}
		}
		return false, nil
	case reflect.String:
		return checkStrings(a, b, func(aVal, bVal string) (bool, error) {
			return strings.Contains(aVal, bVal), nil
		})
	default:
		return false, errors.Errorf("val must be contained by an iterable type, got %s", rv.String())
	}
}

func containedBy(a, b interface{}) (bool, error) {
	return contains(b, a)
}

func typeEq(a, b interface{}) (bool, error) {
	typeVal, _ := toString(b)
	return reflect.ValueOf(a).Type().String() == typeVal, nil
}

func regexMatch(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return regexp.MatchString(bVal, aVal) })
}

func startsWith(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return strings.HasPrefix(aVal, bVal), nil })
}

func endsWith(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return strings.HasSuffix(aVal, bVal), nil })
}

func objectContains(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return strings.Contains(aVal, bVal), nil })
}

func objectNotContains(a, b interface{}) (bool, error) {
	return checkStrings(a, b, func(aVal, bVal string) (bool, error) { return !strings.Contains(aVal, bVal), nil })
}
