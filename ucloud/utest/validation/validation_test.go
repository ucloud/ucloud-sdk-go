package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidation(t *testing.T) {
	type TestStruct struct {
		RetCode int
	}
	resp := TestStruct{RetCode: 0}
	testComparator := NewTestContext(map[string]CompareFunc{
		"eq_bool": func(a, b interface{}) error {
			return nil
		},
	})
	assert.NoError(t, testComparator.NewValidator("RetCode", 0, "eq_bool")(resp))
	assert.NoError(t, testComparator.NewValidator("RetCode", 0, "str_eq")(resp))
	assert.Error(t, testComparator.NewValidator("RetCode", 0, "comparator_err")(resp))
	assert.Error(t, testComparator.NewValidator("TestErr", 0, "str_eq")(resp))
	assert.Error(t, testComparator.NewValidator("RetCode", 1, "str_eq")(resp))
	assert.Error(t, testComparator.NewValidator("RetCode", "test_err", "str_eq")(resp))
}
