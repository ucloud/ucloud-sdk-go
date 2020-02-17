package validation

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

func checkFloats(a, b interface{}, checkFunc func(float64, float64) error) error {
	aVal, bVal, err := toFloats(a, b)
	if err != nil {
		return err
	}
	return checkFunc(aVal, bVal)
}

func checkStrings(a, b interface{}, checkFunc func(string, string) error) error {
	aVal, bVal, err := toStrings(a, b)
	if err != nil {
		return err
	}
	if strings.ToLower(aVal) == "false" {
		aVal = "false"
	}
	if strings.ToLower(bVal) == "false" {
		bVal = "false"
	}
	return checkFunc(aVal, bVal)
}

func checkLens(a, b interface{}, checkFunc func(int, int) error) error {
	aVal, err := toLen(a)
	if err != nil {
		return err
	}

	bVal, err := toInt(b)
	if err != nil {
		return err
	}

	return checkFunc(aVal, bVal)
}

func toFloats(a, b interface{}) (float64, float64, error) {
	aVal, err := toFloat(a)
	if err != nil {
		return 0.0, 0.0, err
	}

	bVal, err := toFloat(b)
	if err != nil {
		return 0.0, 0.0, err
	}

	return roundTo(aVal, 2), roundTo(bVal, 2), nil
}

func roundTo(num float64, perc uint) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(perc))+"f", num), 64)
	return val
}

func toStrings(a, b interface{}) (string, string, error) {
	aVal, err := toString(a)
	if err != nil {
		return "", "", err
	}

	bVal, err := toString(b)
	if err != nil {
		return "", "", err
	}

	return aVal, bVal, nil
}

func toFloat(v interface{}) (float64, error) {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), nil
	case reflect.Float32, reflect.Float64:
		return rv.Float(), nil
	case reflect.String:
		val, err := strconv.ParseFloat(rv.String(), 64)
		if err != nil {
			return 0.0, errors.Errorf("value %s cannot convert to float, %s", v, err)
		}
		return val, nil
	default:
		return 0.0, errors.Errorf("value %s cannot convert to float", v)
	}
}

func toInt(v interface{}) (int, error) {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(rv.Int()), nil
	case reflect.Float32, reflect.Float64:
		return int(rv.Float()), nil
	case reflect.String:
		val, err := strconv.Atoi(rv.String())
		if err != nil {
			return 0.0, errors.Errorf("value %s cannot convert to int, %s", v, err)
		}
		return val, nil
	default:
		return 0, errors.Errorf("value %s cannot convert to int", v)
	}
}

func toString(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Struct:
		b, err := json.Marshal(v)
		return string(b), err
	default:
		return fmt.Sprintf("%+v", v), nil
	}
}

var canGetLengthTypes = []reflect.Kind{reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String}

func toLen(v interface{}) (int, error) {
	rv := reflect.ValueOf(v)
	for _, t := range canGetLengthTypes {
		if rv.Kind() == t {
			return rv.Len(), nil
		}
	}
	return 0, errors.Errorf("value %+v cannot get length", v)
}
