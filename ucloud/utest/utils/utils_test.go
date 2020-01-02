package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
	"reflect"
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
)

func TestGetValue(t *testing.T) {
	type testSet struct {
		Name string
	}

	type test struct {
		RetCode int
		Message string
		DataSet []testSet
	}

	type args struct {
		obj  interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"ok_name",
			args{
				test{1, "some error", []testSet{{"ok_name"}}},
				"DataSet.0.Name",
			},
			"ok_name",
			false,
		},
		{
			"err_path",
			args{
				test{1, "some error", []testSet{{"err_name"}}},
				"DataSet.0.ErrPath",
			},
			"",
			true,
		},

		{
			"nil_path",
			args{
				test{1, "some error", []testSet{{"nil_path"}}},
				"",
			},
			test{1, "some error", []testSet{{"nil_path"}}},
			false,
		},

		{
			"nil_obj",
			args{
				nil,
				"nil_obj",
			},
			"",
			false,
		},

		{
			"ok_retCode",
			args{
				test{1, "some error", []testSet{{"ok_retCode"}}},
				"RetCode",
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValue(tt.args.obj, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueGeneric(t *testing.T) {
	type args struct {
		obj  map[string]interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"ok_name",
			args{
				map[string]interface{}{
					"RetCode": 1.0,
					"Message": "some error",
					"DataSet": []map[string]interface{}{
						{"Name": "ok_name"},
					},
				},
				"DataSet.0.Name",
			},
			"ok_name",
			false,
		},
		{
			"ok_retCode",
			args{
				map[string]interface{}{
					"RetCode": 1.0,
					"Message": "some error",
					"DataSet": []map[string]interface{}{
						{"Name": "ok_retCode"},
					},
				},

				"RetCode",
			},
			float64(1),
			false,
		},
	}

	resp := &response.BaseGenericResponse{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, resp.SetPayload(tt.args.obj))
			got, err := GetValue(resp, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRequest(t *testing.T) {
	type testNest2 struct {
		Str *string
		Int *int
	}
	type testNest struct {
		Nest2 *testNest2
	}
	type test struct {
		request.Common
		StrArr  []string
		Int     *int
		Str     *string
		Bool    *bool
		Float64 *float64
		Nest    []testNest
		NestPtr []*testNest
	}

	type args struct {
		req     request.Common
		payload map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"map_ok", args{&test{}, map[string]interface{}{
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
				"Float64": 0.1,
				"StrArr":  []interface{}{"str", "str"},
			}}, false},
		{
			"type_ok", args{&test{}, map[string]interface{}{
				"Str":  "str",
				"Int":  "1",
				"Bool": "true",
			}}, false},

		{
			"type_no", args{&test{}, map[string]interface{}{
				"Str":  "str",
				"Int":  "1",
				"Bool": "str",
			}}, true},

		{
			"map_type_no", args{&test{}, map[string]interface{}{
				"Nest": []map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": "str",
							"Str": 111,
						},
					},
				},
			}}, true},

		{
			"not_arg", args{&test{}, map[string]interface{}{
				"ErrArg": "str",
			}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetRequest(tt.args.req, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("SetRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_valueAtPath(t *testing.T) {
	type args struct {
		v    interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"map.map_ok", args{map[string]map[string]int{"Test": {"Name": 42}}, "Test.Name"}, 42, false},
		{"map.map_ptr_ok", args{&map[string]map[string]int{"Test": {"Name": 42}}, "Test.Name"}, 42, false},
		{"nil_value_no", args{nil, "Test.Name"}, nil, true},
		{"component_no", args{&map[string]map[string][]int{"Test": {"Name": {42, 11}}}, "Test.Name.3"}, nil, true},
		{"map.map_no", args{map[string]map[string]int{"Test": {"Name": 42}}, "Test.Test"}, nil, true},
		{"map.struct", args{map[string]struct{ Name int }{"Test": {Name: 42}}, "Test.Name"}, 42, false},
		{"struct.map", args{struct{ Test map[string]int }{Test: map[string]int{"Name": 42}}, "Test.Name"}, 42, false},
		{"struct.struct", args{struct{ Test struct{ Name int } }{Test: struct{ Name int }{Name: 42}}, "Test.Name"}, 42, false},
		{"struct.slice", args{struct{ Name []int }{[]int{42}}, "Name.0"}, 42, false},
		{"slice.struct", args{[]struct{ Name int }{{42}}, "0.Name"}, 42, false},
		{"slice.struct", args{[]struct{ ResourceId string }{{"foo"}}, "0.ResourceID"}, "foo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueAtPath(tt.args.v, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValueAtPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueAtPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
