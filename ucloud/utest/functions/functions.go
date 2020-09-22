package functions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// GetTimestamp will return the timestamp string
func GetTimestamp(strLen int) (int, error) {
	if strLen < 0 || 19 < strLen {
		return 0, errors.Errorf("timestamp length can only between 0 and 16")
	}
	intStr := strconv.FormatInt(time.Now().UnixNano(), 10)[:strLen]
	ts, _ := strconv.Atoi(intStr)
	return ts, nil
}

// TimeDelta given timestamp(10bit) and calculate relative delta time
func TimeDelta(timeStamp int, value int, typ string) (int, error) {
	delta := 0
	if typ == "days" {
		delta = timeStamp + value*86400
	} else if typ == "hours" {
		delta = timeStamp + value*3600
	}
	return delta, nil
}

// Concat will concat any data as string
func Concat(input ...interface{}) (string, error) {
	var vL []string
	for _, v := range input {
		vL = append(vL, fmt.Sprint(v))
	}
	return strings.Join(vL, ""), nil
}

// Calculate will to calculate two number by operator
func Calculate(op string, left, right interface{}) (interface{}, error) {
	l, err := isFloat(left)
	if err != nil {
		return 0, err
	}
	r, err := isFloat(right)
	if err != nil {
		return 0, err
	}

	if l || r {
		leftValue := convertToFloat(left)
		rightValue := convertToFloat(right)
		switch op {
		case "+":
			return leftValue + rightValue, nil
		case "-":
			return leftValue - rightValue, nil
		case "*":
			return leftValue * rightValue, nil
		case "/":
			return leftValue / rightValue, nil
		default:
			return 0, errors.Errorf("function Calculate has not support %s", op)
		}
	}

	leftValue, rightValue := reflect.ValueOf(left).Int(), reflect.ValueOf(right).Int()
	switch op {
	case "+":
		return leftValue + rightValue, nil
	case "-":
		return leftValue - rightValue, nil
	case "*":
		return leftValue * rightValue, nil
	case "/":
		return leftValue / rightValue, nil
	default:
		return 0, errors.Errorf("function Calculate has not support %s", op)
	}
}

func convertToFloat(v interface{}) float64 {
	if is, _ := isFloat(v); is {
		return reflect.ValueOf(v).Float()
	}
	return float64(reflect.ValueOf(v).Int())
}

func isFloat(v interface{}) (bool, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return false, nil
	case reflect.Float32, reflect.Float64:
		return true, nil
	}
	return false, errors.Errorf("function Calculate only support calculate int and float value, got %s", rv.Kind())
}

// SearchValue will search key/value in an collection and return the value of destination key
func SearchValue(arr interface{}, originKey string, originValue interface{}, destKey string) (interface{}, error) {
	jsonPayload, err := json.Marshal(arr)
	if err != nil {
		return "", nil
	}

	var mArr []map[string]interface{}
	err = json.Unmarshal(jsonPayload, &mArr)
	if err != nil {
		return "", err
	}

	for _, m := range mArr {
		if val, ok := m[originKey]; !ok {
			continue
		} else {
			a := fmt.Sprint(originValue)
			b := fmt.Sprint(val)
			if a != b {
				continue
			}
		}

		if val, ok := m[destKey]; ok {
			return val, nil
		}
	}

	return "", errors.Errorf("value of key: %s is not found", destKey)
}
