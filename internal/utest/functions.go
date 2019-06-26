package utest

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var zoneImages = map[string]string{
	"cn-bj1-01":   "uimage-zf2xoa",
	"cn-bj2-02":   "uimage-csarne", // origin is uimage-u0k3m3
	"cn-bj2-03":   "b2689fc412ee5fa108fa5b23ed2e00e6",
	"cn-bj2-04":   "uimage-rq2kat",
	"cn-bj2-05":   "uimage-5rqi0y",
	"cn-sh-01":    "uimage-65fa28",
	"cn-sh2-01":   "uimage-p0c51y",
	"cn-sh2-02":   "uimage-qnrtyc",
	"cn-gd-02":    "uimage-b54e21",
	"hk-01":       "uimage-g3hvlg",
	"hk-02":       "uimage-gcs1cr",
	"us-ca-01":    "uimage-0duw4w",
	"us-ws-01":    "uimage-ggiase",
	"ge-fra-01":   "uimage-unynvz",
	"th-bkk-01":   "uimage-2bsbiy",
	"kr-seoul-01": "uimage-kxrlft",
	"sg-01":       "uimage-oqpggx",
	"tw-kh-01":    "uimage-hwgsqi",
	"rus-mosc-01": "uimage-cgfvqy",
	"jpn-tky-01":  "uimage-jshpqn",
	"tw-tp-01":    "uimage-toxa1t",
	"cn-zj-01":    "uimage-vv0zdq",
}

// GetZoneImage will return the image id for zone
func GetZoneImage(input interface{}) (string, error) {
	zone, err := toString(input)
	if err != nil {
		return "", err
	}
	if img, ok := zoneImages[zone]; ok {
		return img, nil
	}
	return "", errors.Errorf("cannot get zone image, invalid zone %s", zone)
}

// GetImageResource will return the image id for region and zone
func GetImageResource(inputRegion, inputZone interface{}) (string, error) {
	return GetZoneImage(inputZone)
}

var regionImages = map[string]string{
	"cn-bj1":   "uimage-zf2xoa",
	"cn-bj2":   "uimage-rq2kat",
	"cn-sh":    "uimage-65fa28",
	"cn-sh2":   "uimage-p0c51y",
	"cn-gd":    "uimage-b54e21",
	"hk":       "uimage-gcs1cr",
	"us-ca":    "uimage-0duw4w",
	"us-ws":    "uimage-ggiase",
	"ge-fra":   "uimage-unynvz",
	"th-bkk":   "uimage-2bsbiy",
	"kr-seoul": "uimage-kxrlft",
	"sg":       "uimage-oqpggx",
	"tw-kh":    "uimage-hwgsqi",
	"rus-mosc": "uimage-cgfvqy",
	"jpn-tky":  "uimage-jshpqn",
	"tw-tp":    "uimage-toxa1t",
	"cn-zj":    "uimage-vv0zdq",
}

// GetRegionImage will return the image id for region
func GetRegionImage(input interface{}) (string, error) {
	region, err := toString(input)
	if err != nil {
		return "", err
	}
	if img, ok := regionImages[region]; ok {
		return img, nil
	}
	return "", errors.Errorf("cannot get region image, invalid region %s", region)
}

var udpnResource = map[string]string{
	"cn-gd":  "cn-bj2",
	"cn-bj2": "cn-gd",
}

// GetUDPNRegionResource the vpc id for udpn peer destinition
func GetUDPNRegionResource(input interface{}) (string, error) {
	region, err := toString(input)
	if err != nil {
		return "", err
	}
	if img, ok := udpnResource[region]; ok {
		return img, nil
	}
	return "", errors.Errorf("cannot get udpn region resource, invalid region %s", region)
}

// GetNotEqual will return the first item of vL that is not equal to v
func GetNotEqual(v interface{}, vL ...interface{}) (string, error) {
	if len(vL) < 2 {
		return "", errors.Errorf("cannot return value")
	}

	vs, err := toString(v)
	if err != nil {
		return "", err
	}

	target := vs
	for _, c := range vL {
		target, err = toString(c)
		if err != nil {
			return "", err
		}

		if vs != target {
			break
		}
	}

	return target, nil
}

// GetTimestamp will return the timestamp string
func GetTimestamp(input interface{}) (string, error) {
	strLen, err := toInt(input)
	if err != nil {
		return getTimestamp(13)
	}

	return getTimestamp(strLen)
}

func getTimestamp(strLen int) (string, error) {
	if strLen < 0 || 19 < strLen {
		return "", errors.Errorf("timestamp length can only between 0 and 16")
	}
	return strconv.FormatInt(time.Now().UnixNano(), 10)[:strLen], nil
}

// Concat will concat any data as string
func Concat(input ...interface{}) (string, error) {
	return joinAsString("", input...)
}

// ConcatWithVertical will concat any data as string join with '|'
func ConcatWithVertical(input ...interface{}) (string, error) {
	return joinAsString("|", input...)
}

func joinAsString(sep string, input ...interface{}) (string, error) {
	items := []string{}
	for _, item := range input {
		s, err := toString(item)
		if err != nil {
			return "", err
		}
		items = append(items, s)
	}
	return strings.Join(items, sep), nil
}

// Calculate will to calculate two number by operator
func Calculate(op interface{}, rvL ...interface{}) (int, error) {
	l := len(rvL)
	if l == 0 {
		v, err := toInt(op)
		if err != nil {
			return 0, err
		}
		return v, nil
	}

	if l < 2 {
		return 0, errors.Errorf("not enough arguments, expect least two item, got %v", len(rvL))
	}

	vL := make([]int, l)
	for i := 0; i < l; i++ {
		v, err := toInt(rvL[i])
		if err != nil {
			return 0, err
		}
		vL[i] = v
	}

	opSymbol, err := toString(op)
	if err != nil {
		return 0, err
	}

	switch opSymbol {
	case "+":
		return caculateInt(func(a, b int) int { return a + b }, vL...), nil
	case "-":
		return caculateInt(func(a, b int) int { return a - b }, vL...), nil
	case "*":
		return caculateInt(func(a, b int) int { return a * b }, vL...), nil
	default:
		return 0, errors.Errorf("function Calculate has not support %s", op)
	}
}

func caculateInt(fn func(a, b int) int, vL ...int) int {
	r := fn(vL[0], vL[1])
	for _, v := range vL[2:] {
		r = fn(r, v)
	}
	return r
}

// SearchValue will search key/value in an collection and return the value of destination key
func SearchValue(arr interface{}, originKey string, originValue interface{}, destKey string) (interface{}, error) {
	jsonPayload, err := json.Marshal(arr)
	if err != nil {
		return "", nil
	}

	mArr := []map[string]interface{}{}
	err = json.Unmarshal(jsonPayload, &mArr)
	if err != nil {
		return "", err
	}

	for _, m := range mArr {
		if val, ok := m[originKey]; !ok {
			continue
		} else {
			a, _ := toString(originValue)
			b, _ := toString(val)
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

// GetUUID will return uuid string
func GetUUID() (string, error) {
	items := []string{}
	for _, strLen := range []int{8, 4, 4, 16} {
		s, err := getRandomString(strLen)
		if err != nil {
			return "", err
		}
		items = append(items, s)
	}
	return strings.Join(items, "-"), nil
}

func getRandomString(strLen int) (string, error) {
	b := make([]byte, strLen)
	if _, err := rand.Read(b); err != nil {
		return "", errors.Errorf("cannot generate random string")
	}
	return fmt.Sprintf("%X", b), nil
}
