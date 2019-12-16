package utest

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"github.com/ucloud/ucloud-sdk-go/private/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// GetValue will return the value of an object by path
func GetValue(obj interface{}, path string) (interface{}, error) {
	if obj == nil {
		return "", nil
	}

	v, err := utils.ValueAtPath(obj, path)
	if err != nil {
		return "", err
	}

	return v, nil
}

// SetReqValue will set value into pointer referenced or slice
func SetReqValue(addr interface{}, field string, values ...interface{}) error {
	if len(values) == 0 {
		return errors.Errorf("no values to be set")
	}

	rv := reflect.ValueOf(addr)
	if rv.IsValid() == false {
		return errors.Errorf("struct is invalid")
	}

	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return errors.Errorf("struct is nil")
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return errors.Errorf("got type %s, expected struct", rv.Kind().String())
	}

	rv = rv.FieldByName(field)
	if !rv.IsValid() {
		return errors.Errorf("struct field %s is invalid", field)
	}

	if rv.Kind() != reflect.Ptr && rv.Kind() != reflect.Slice {
		return errors.Errorf("only support pointer and slice type value")
	}

	if !rv.CanSet() {
		return errors.Errorf("cannot set %s, field cannot be set", field)
	}

	rValues := []reflect.Value{}
	for _, value := range values {
		s, _ := toString(value)
		v, err := convertValueWithType(rv.Type().Elem(), s)
		if err != nil {
			return err
		}
		rValues = append(rValues, v)
	}

	if rv.Kind() == reflect.Ptr {
		rv.Set(rValues[0])
	} else if rv.Kind() == reflect.Slice {
		rv.Set(reflect.MakeSlice(rv.Type(), 0, 0))
		for _, willSet := range rValues {
			rv.Set(reflect.Append(rv, willSet.Elem()))
		}
	}

	return nil
}

func convertValueWithType(typ reflect.Type, value string) (reflect.Value, error) {
	v := reflect.Zero(typ)

	switch typ.Kind() {
	case reflect.String:
		v = reflect.ValueOf(&value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return v, err
		}
		v = reflect.ValueOf(ucloud.Int(int(intValue)))
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return v, err
		}
		v = reflect.ValueOf(ucloud.Float64(floatValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return v, err
		}
		v = reflect.ValueOf(ucloud.Bool(boolValue))
	}

	return v, nil
}
