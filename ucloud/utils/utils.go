package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/xiaohui/goucloud/ucloud/auth"
)

func ConvertParamsToValues(params interface{}, values *url.Values) {

	elem := reflect.ValueOf(params)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	elemType := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		fieldName := elemType.Field(i).Name

		field := elem.Field(i)
		kind := field.Kind()
		if (kind == reflect.Ptr ||
			kind == reflect.Array ||
			kind == reflect.Slice ||
			kind == reflect.Map ||
			kind == reflect.Chan) && field.IsNil() {
			continue

		}

		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}

		var v string
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(field.Int(), 10)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v = strconv.FormatUint(field.Uint(), 10)

		case reflect.Float32:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 32)

		case reflect.Float64:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 64)

		case reflect.Bool:
			v = strconv.FormatBool(field.Bool())

		case reflect.String:
			v = field.String()
		}

		if v != "" {
			name := elemType.Field(i).Tag.Get("ArgName")
			if name == "" {
				name = fieldName
			}

			values.Set(name, v)
		}
	}
}

func UrlWithSignature(values *url.Values, baseUrl, privateKey string) (string, error) {

	urlEncoded, err := url.QueryUnescape(values.Encode())
	if err != nil {
		return "", fmt.Errorf("unescape failed, error: %s", err)
	}

	// replace '&' and '=' in url
	urlEncoded = strings.Replace(urlEncoded, "=", "", -1)
	urlEncoded = strings.Replace(urlEncoded, "&", "", -1)

	signature, err := auth.GenerateSignature(urlEncoded, privateKey)
	if err != nil {
		return "", fmt.Errorf("generate signature error:%s", err)
	}

	return baseUrl + "?" + values.Encode() + "&Signature=" + url.QueryEscape(signature), nil
}

func DumpVal(vals ...interface{}) {
	for _, val := range vals {
		prettyJSON, err := json.MarshalIndent(val, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(prettyJSON))
	}
}
