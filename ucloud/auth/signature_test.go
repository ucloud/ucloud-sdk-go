package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	intValue := int(42)
	int8Value := int8(42)
	int16Value := int16(42)
	int32Value := int32(42)
	int64Value := int64(42)
	uintValue := uint(42)
	uint8Value := uint8(42)
	uint16Value := uint16(42)
	uint32Value := uint32(42)
	uint64Value := uint64(42)

	boolValue := true
	strValue := "foo"
	float32Value := float32(42.1)
	float64Value := float64(42.1)
	float32TValue := float32(42.0)
	float64TValue := float64(42.0)

	params := map[string]interface{}{
		"int":    intValue,
		"int8":   int8Value,
		"int16":  int16Value,
		"int32":  int32Value,
		"int64":  int64Value,
		"uint":   uintValue,
		"uint8":  uint8Value,
		"uint16": uint16Value,
		"uint32": uint32Value,
		"uint64": uint64Value,

		"intPointer":    &intValue,
		"int8Pointer":   &int8Value,
		"int16Pointer":  &int16Value,
		"int32Pointer":  &int32Value,
		"int64Pointer":  &int64Value,
		"uintPointer":   &uintValue,
		"uint8Pointer":  &uint8Value,
		"uint16Pointer": &uint16Value,
		"uint32Pointer": &uint32Value,
		"uint64Pointer": &uint64Value,

		"boolValue":     boolValue,
		"strValue":      strValue,
		"float32Value":  float32Value,
		"float64Value":  float64Value,
		"float32TValue": float32TValue,
		"float64TValue": float64TValue,

		"boolValuePointer":     &boolValue,
		"strValuePointer":      &strValue,
		"float32ValuePointer":  &float32Value,
		"float64ValuePointer":  &float64Value,
		"float32TValuePointer": &float32TValue,
		"float64TValuePointer": &float64TValue,

		"map": map[string]interface{}{
			"foo": "bar",
		},
		"array": []interface{}{"foo", "bar"},
		"arrMap": []interface{}{
			map[string]interface{}{"foo": "bar"},
		},
		"arr2": []interface{}{
			[]interface{}{"foo", "bar"},
		},
		"arrMapMissTyped1": []map[string]string{
			{"foo": "bar"},
		},
		"arrMapMissTyped2": []interface{}{
			map[string]string{"foo": "bar"},
		},
		"mapArrMapMissTyped": map[string][]map[string]interface{}{
			"foo": {
				map[string]interface{}{"foo": "bar"},
			},
		},
	}

	assert.Equal(t, `arr2foobararrMapfoobararrMapMissTyped1arrMapMissTyped2arrayfoobarboolValuetrueboolValuePointertruefloat32TValue42float32TValuePointer42float32Value42.099998474121094float32ValuePointer42.099998474121094float64TValue42float64TValuePointer42float64Value42.1float64ValuePointer42.1int42int1642int16Pointer42int3242int32Pointer42int6442int64Pointer42int842int8Pointer42intPointer42mapfoobarmapArrMapMissTypedstrValuefoostrValuePointerfoouint42uint1642uint16Pointer42uint3242uint32Pointer42uint6442uint64Pointer42uint842uint8Pointer42uintPointer42`, map2String(params))
	sign := sign(params, "foo")
	assert.Equal(t, "682fc4fbb46a71c6c690a1f8cb233f4174662845", sign)
}
