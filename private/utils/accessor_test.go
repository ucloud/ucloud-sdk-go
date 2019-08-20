package utils

import (
	"reflect"
	"testing"
)

func TestValueAtPath(t *testing.T) {
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
		{"map.map", args{map[string]map[string]int{"Test": {"Name": 42}}, "Test.Name"}, 42, false},
		{"map.struct", args{map[string]struct{ Name int }{"Test": {Name: 42}}, "Test.Name"}, 42, false},
		{"struct.map", args{struct{ Test map[string]int }{Test: map[string]int{"Name": 42}}, "Test.Name"}, 42, false},
		{"struct.struct", args{struct{ Test struct{ Name int } }{Test: struct{ Name int }{Name: 42}}, "Test.Name"}, 42, false},
		{"struct.slice", args{struct{ Name []int }{[]int{42}}, "Name.0"}, 42, false},
		{"slice.struct", args{[]struct{ Name int }{{42}}, "0.Name"}, 42, false},
		{"slice.struct", args{[]struct{ ResourceId string }{{"foo"}}, "0.ResourceID"}, "foo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValueAtPath(tt.args.v, tt.args.path)
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
