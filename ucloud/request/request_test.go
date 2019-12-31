package request

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestRequestAccessor(t *testing.T) {
	var err error
	req := &CommonBase{}

	s := "foo"
	assert.Equal(t, "", req.GetAction())
	assert.Nil(t, req.GetActionRef())
	err = req.SetActionRef(&s)
	assert.NoError(t, err)
	assert.NotNil(t, req.GetActionRef())

	assert.Equal(t, "", req.GetRegion())
	assert.Nil(t, req.GetRegionRef())
	err = req.SetRegionRef(&s)
	assert.NoError(t, err)
	assert.NotNil(t, req.GetRegionRef())

	assert.Equal(t, "", req.GetZone())
	assert.Nil(t, req.GetZoneRef())
	err = req.SetZoneRef(&s)
	assert.NoError(t, err)
	assert.NotNil(t, req.GetZoneRef())

	assert.Equal(t, "", req.GetProjectId())
	assert.Nil(t, req.GetProjectIdRef())
	err = req.SetProjectIdRef(&s)
	assert.NoError(t, err)
	assert.NotNil(t, req.GetProjectIdRef())

	req.SetRetryCount(42)
	assert.Equal(t, 42, req.GetRetryCount())

	req.SetRetryable(true)
	assert.Equal(t, true, req.GetRetryable())
	req.SetRetryable(false)

	req.WithRetry(42)
	assert.Equal(t, 42, req.GetMaxretries())
	assert.Equal(t, true, req.GetRetryable())

	now := time.Now()
	req.SetRequestTime(now)
	assert.Equal(t, now, req.GetRequestTime())

	req.WithTimeout(1 * time.Second)
	assert.Equal(t, 1*time.Second, req.GetTimeout())

	err = req.SetAction("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.GetAction())
	assert.Equal(t, String("Zm9v"), ToBase64Query(String(req.GetAction())))

	err = req.SetRegion("cn-bj2")
	assert.NoError(t, err)
	assert.Equal(t, "cn-bj2", req.GetRegion())

	err = req.SetZone("cn-bj2-02")
	assert.NoError(t, err)
	assert.Equal(t, "cn-bj2-02", req.GetZone())

	err = req.SetProjectId("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.GetProjectId())
}

func TestRequestAccessorForGeneric(t *testing.T) {
	var err error
	req := &BaseGenericRequest{}

	assert.Equal(t, "", req.GetAction())
	err = req.CommonBase.SetAction("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.GetAction())

	assert.Equal(t, "", req.GetRegion())
	err = req.CommonBase.SetRegion("cn-bj2")
	assert.NoError(t, err)
	assert.Equal(t, "cn-bj2", req.GetRegion())

	assert.Equal(t, "", req.GetZone())
	err = req.CommonBase.SetZone("cn-bj2-02")
	assert.NoError(t, err)
	assert.Equal(t, "cn-bj2-02", req.GetZone())

	assert.Equal(t, "", req.GetProjectId())
	err = req.CommonBase.SetProjectId("bar")
	assert.NoError(t, err)
	assert.Equal(t, "bar", req.GetProjectId())

	assert.Equal(t, map[string]interface{}{
		"Region":    "cn-bj2",
		"Zone":      "cn-bj2-02",
		"Action":    "foo",
		"ProjectId": "bar",
	}, req.GetPayload())

	assert.NoError(t, req.SetPayload(map[string]interface{}{
		"Region":    "cn-sh2",
		"Zone":      "cn-sh2-02",
		"Action":    "bar",
		"ProjectId": "foo",
	}))

	assert.Equal(t, "cn-sh2-02", req.GetZone())
	assert.Equal(t, "cn-sh2", req.GetRegion())
	assert.Equal(t, "foo", req.GetProjectId())
	assert.Equal(t, "bar", req.GetAction())

	assert.Error(t, req.SetPayload(map[string]interface{}{
		"Region": 1,
	}))
	assert.Error(t, req.SetPayload(map[string]interface{}{
		"Zone": true,
	}))
	assert.Error(t, req.SetPayload(map[string]interface{}{
		"Action": 1,
	}))
	assert.Error(t, req.SetPayload(map[string]interface{}{
		"ProjectId": 1,
	}))
}

func TestEncodeOne(t *testing.T) {
	i := 42

	tests := []struct {
		name        string
		inputVector interface{}
		golden      string
		wantErr     bool
	}{
		{"int", 42, "42", false},
		{"uint", uint(42), "42", false},
		{"float", 42.0, "42", false},
		{"float", 42.1, "42.1", false},
		{"bool", true, "true", false},
		{"string", "foo", "foo", false},
		{"pointer", &i, "42", false},
		{"error", struct{}{}, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rv := reflect.ValueOf(test.inputVector)
			got, err := encodeOne(&rv)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.golden, got)
				assert.NoError(t, err)
			}
		})
	}
}

func TestToQueryMap(t *testing.T) {
	type Composite struct {
		Region string

		unexported string
	}

	type compositedRequest struct {
		CommonBase
	}

	type args struct {
		req Common
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"Ok",
			args{
				req: &struct {
					CommonBase
					Id      int
					Name    string
					IsValid bool
					Ips     []string
				}{
					Id:      1,
					Name:    "lilei",
					IsValid: true,
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]string{
				"Id":      "1",
				"Name":    "lilei",
				"IsValid": "true",
				"Ips.0":   "127.0.0.1",
				"Ips.1":   "192.168.1.1",
			},
			false,
		},
		{
			"Partial",
			args{
				req: &struct {
					CommonBase
					Id      int
					Name    string
					IsValid bool
					Ips     []string
				}{
					Id:      1,
					Name:    "",
					IsValid: true,
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]string{
				"Id":      "1",
				"IsValid": "true",
				"Ips.0":   "127.0.0.1",
				"Ips.1":   "192.168.1.1",
			},
			false,
		},
		{
			"IsComposited",
			args{
				req: &struct {
					CommonBase
					Composite
				}{
					CommonBase{},
					Composite{Region: "cn-bj2"},
				},
			},
			map[string]string{"Composite.Region": "cn-bj2"},
			false,
		},
		{
			"IsCompositedByArray",
			args{
				req: &struct {
					CommonBase
					Arr []Composite
				}{
					CommonBase{},
					[]Composite{
						{Region: "cn-bj2"},
						{Region: "cn-sh1"},
					},
				},
			},
			map[string]string{
				"Arr.0.Region": "cn-bj2",
				"Arr.1.Region": "cn-sh1",
			},
			false,
		},
		{
			"PointerStruct",
			args{
				req: &struct {
					CommonBase
					Ptr *Composite
				}{
					CommonBase{},
					&Composite{Region: "cn-bj2"},
				},
			},
			map[string]string{
				"Ptr.Region": "cn-bj2",
			},
			false,
		},
		{
			"CommonRegion",
			args{
				req: func() Common {
					req := &compositedRequest{}
					_ = req.SetRegion("cn-bj2")
					return req
				}(),
			},
			map[string]string{
				"Region": "cn-bj2",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToQueryMap(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToQueryMapForGeneric(t *testing.T) {
	tests := []struct {
		name    string
		payload map[string]interface{}
		want    map[string]string
		wantErr bool
	}{
		{
			"Ok",
			map[string]interface{}{
				"Id":      1,
				"Name":    "lilei",
				"IsValid": true,
				"Ips":     []string{"127.0.0.1", "192.168.1.1"},
			},
			map[string]string{
				"Id":      "1",
				"Name":    "lilei",
				"IsValid": "true",
				"Ips.0":   "127.0.0.1",
				"Ips.1":   "192.168.1.1",
			},
			false,
		},

		{
			"Ok_nest_map",
			map[string]interface{}{
				"Str":  "str",
				"Int":  1,
				"Bool": true,
				"Nest": []map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": 1,
							"Str": "str",
						},
					},
				},
			},
			map[string]string{
				"Str":              "str",
				"Int":              "1",
				"Bool":             "true",
				"Nest.0.Nest2.Int": "1",
				"Nest.0.Nest2.Str": "str",
			},
			false,
		},

		{
			"Ok_ptr_map",
			map[string]interface{}{
				"Nest": []*map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": 1,
							"Str": "str",
						},
					},
				},
			},
			map[string]string{
				"Nest.0.Nest2.Int": "1",
				"Nest.0.Nest2.Str": "str",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &BaseGenericRequest{}
			assert.NoError(t, req.SetPayload(tt.payload))
			got, err := ToQueryMap(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
