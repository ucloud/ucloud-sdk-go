package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"

	"github.com/pkg/errors"
)

// GetValue will return the value of an object by path
func GetValue(obj interface{}, path string) (interface{}, error) {
	if obj == nil {
		return "", nil
	}

	if path == "" {
		return obj, nil
	}

	if r, ok := obj.(response.GenericResponse); ok {
		obj = r.GetPayload()
	}

	v, err := valueAtPath(obj, path)
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
			return errors.Errorf("request ptr is nil")
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return errors.Errorf("request expected type Struct, got type %s,", rv.Kind().String())
	}

	for k, v := range payload {
		f := rv.FieldByName(k)
		if !f.IsValid() {
			return errors.Errorf("error on setting field %q of request, field is invalid", k)
		}

		if !f.CanSet() {
			return errors.Errorf("error on setting field %q of request, field cannot be set", k)
		}
		if err := setValue(f, v); err != nil {
			return errors.Errorf("error on setting field %q of request, %s", k, err)
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
			return errors.Errorf("convert %q to Int failed, %s", value, err)
		}
		fv = reflect.ValueOf(request.Int(int(intValue)))
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.Errorf("convert %q to Float failed, %s", value, err)
		}
		fv = reflect.ValueOf(request.Float64(floatValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return errors.Errorf("convert %q to Bool failed, %s", value, err)
		}
		fv = reflect.ValueOf(request.Bool(boolValue))
	}

	f.Set(fv)

	return nil
}

// valueAtPath will get struct attribute value by recursive
func valueAtPath(v interface{}, path string) (interface{}, error) {
	path = strings.TrimSpace(path)
	components := strings.Split(path, ".")

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, errors.Errorf("object %#v is nil", v)
		}
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		i, err := strconv.Atoi(components[0])
		if err != nil {
			return nil, errors.Errorf("path %s is invalid at index of array", path)
		}

		length := rv.Len()
		if i >= length {
			return nil, errors.Errorf("path %s is invalid, array has length %v, but got %v", path, length, i)
		}

		itemV := rv.Index(i)
		if !itemV.IsValid() {
			return nil, errors.Errorf("path %s is invalid for map", path)
		}

		if len(components) > 1 {
			return valueAtPath(itemV.Interface(), strings.Join(components[1:], "."))
		}

		return itemV.Interface(), nil
	}

	if rv.Kind() == reflect.Map && !rv.IsNil() {
		itemV := rv.MapIndex(reflect.ValueOf(components[0]))
		if !itemV.IsValid() {
			return nil, errors.Errorf("path %s is invalid for map", path)
		}

		if len(components) > 1 {
			return valueAtPath(itemV.Interface(), strings.Join(components[1:], "."))
		}

		return itemV.Interface(), nil
	}

	if rv.Kind() == reflect.Struct {
		itemV := rv.FieldByNameFunc(func(s string) bool {
			return strings.ToLower(s) == strings.ToLower(components[0])
		})

		if !itemV.IsValid() {
			return nil, errors.Errorf("path %s is invalid for struct", path)
		}

		if len(components) > 1 {
			return valueAtPath(itemV.Interface(), strings.Join(components[1:], "."))
		}

		return itemV.Interface(), nil
	}

	return nil, errors.Errorf("object %#v is invalid, need map or struct", v)
}
