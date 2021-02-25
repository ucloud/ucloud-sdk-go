package auth

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	params := map[string]interface{}{
		"12": 3,
		"4":  "56",
	}

	sign := Sign(params, "789")
	assert.Equal(t, "f7c3bc1d808e04732adf679965ccc34ca7ae3441", sign)
}

func TestVerifySign(t *testing.T) {
	params := map[string]interface{}{
		"12": 3,
		"4":  "56",
	}

	assert.True(t, VerifySign(params, "789", "f7c3bc1d808e04732adf679965ccc34ca7ae3441"))
}

func TestMap2String(t *testing.T) {
	params := map[string]interface{}{
		"12": 3,
		"4":  "56",
		"7":  []int{8, 9, 10},
		"11": map[string]interface{}{"12": 13, "14": 15},
	}
	signedStr := map2String(params)
	t.Log("signed string: ", signedStr)
	assert.Equal(t, "111213141512345678910", signedStr)
}

func Test_reflectArr2String(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	str := reflect2String(reflect.ValueOf(arr))
	fmt.Println(str)
}
