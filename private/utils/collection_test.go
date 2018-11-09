package utils

import (
	"reflect"
	"testing"
)

func TestMergeMap(t *testing.T) {
	type args struct {
		args []map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			"ok",
			args{
				[]map[string]string{
					{"a": "1"},
					{"b": "2"},
				},
			},
			map[string]string{"a": "1", "b": "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMap(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMapIfNotExists(t *testing.T) {
	type args struct {
		m map[string]string
		k string
		v string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"not_exists", args{map[string]string{"a": "1"}, "b", "2"}, map[string]string{"a": "1", "b": "2"}},
		{"not_exists", args{map[string]string{"a": "1"}, "a", "2"}, map[string]string{"a": "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMapIfNotExists(tt.args.m, tt.args.k, tt.args.v)
			if SetMapIfNotExists(tt.args.m, tt.args.k, tt.args.v); !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("MergeMap() = %v, want %v", tt.args.m, tt.want)
			}
		})
	}
}

func TestCheckStringIn(t *testing.T) {
	type args struct {
		val        string
		availables []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{"get", []string{"get", "post"}}, false},
		{"err", args{"-", []string{"get", "post"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckStringIn(tt.args.val, tt.args.availables); (err != nil) != tt.wantErr {
				t.Errorf("CheckStringIn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
