package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"

	"github.com/pkg/errors"
	"github.com/ucloud/ucloud-sdk-go/private/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// GetValue will return the value of an object by path
func GetValue(obj interface{}, path string) (interface{}, error) {
	if obj == nil {
		return "", nil
	}

	if r, ok := obj.(response.GenericResponse); ok {
		obj = r.GetPayload()
	}

	v, err := utils.ValueAtPath(obj, path)
	if err != nil {
		return "", err
	}

	return v, nil
}

// SetRequest will set the request by the map input
func SetRequest(req request.Common, payload map[string]interface{}) error {
	rv := reflect.ValueOf(req)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return errors.Errorf("struct is nil")
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return errors.Errorf("got type %s, expected struct", rv.Kind().String())
	}

	for k, v := range payload {
		f := rv.FieldByName(k)
		if !f.IsValid() {
			return fmt.Errorf("struct field %s is invalid", k)
		}

		if !f.CanSet() {
			return fmt.Errorf("cannot set %s, field cannot be set", k)
		}
		if err := setValue(f, v); err != nil {
			return err
		}
	}
	return nil
}

func setValue(rv reflect.Value, value interface{}) error {

	if m, ok := value.(map[string]interface{}); ok {
		if rv.IsNil() {
			zero := reflect.New(rv.Type().Elem())
			rv.Set(zero)
		}

		for k, v := range m {
			f := rv.Elem().FieldByName(k)
			if f.IsNil() {
				zero := reflect.New(f.Type().Elem())
				f.Set(zero)
			}

			if err := setValue(f, v); err != nil {
				return err
			}
		}
	} else if l, ok := value.([]interface{}); ok {
		for _, item := range l {
			rv.Set(reflect.Append(rv, reflect.ValueOf(item)))
		}
	} else if l, ok := value.([]map[string]interface{}); ok {
		for _, item := range l {
			fv := reflect.New(rv.Type().Elem())
			if err := setValue(fv, item); err != nil {
				return err
			}
			rv.Set(reflect.Append(rv, fv.Elem()))
		}
	} else {
		if err := convertValue(rv, value); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func convertValue(f reflect.Value, v interface{}) error {
	value := fmt.Sprint(v)
	fv := reflect.ValueOf(v)

	switch f.Type().Elem().Kind() {
	case reflect.String:
		fv = reflect.ValueOf(&value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		fv = reflect.ValueOf(ucloud.Int(int(intValue)))
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		fv = reflect.ValueOf(ucloud.Float64(floatValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		fv = reflect.ValueOf(ucloud.Bool(boolValue))
	}

	f.Set(fv)

	return nil
}
