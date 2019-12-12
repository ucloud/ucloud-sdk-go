package utils

import (
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
		StrArr []string
		Int    *int
		Str    *string
		Bool   *bool
		Nest   []testNest
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
			"test", args{&test{}, map[string]interface{}{
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
				"StrArr": []interface{}{"str", "str"},
			}}, false},
		{
			"test", args{&test{}, map[string]interface{}{
				"Str":  "str",
				"Int":  "1",
				"Bool": "true",
			}}, false},

		{
			"test", args{&test{}, map[string]interface{}{
				"Str":  "str",
				"Int":  1,
				"Bool": "str",
			}}, true},

		{
			"test", args{&test{}, map[string]interface{}{
				"Nest": []map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": "str",
							"Str": 111,
						},
					},
				},
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
