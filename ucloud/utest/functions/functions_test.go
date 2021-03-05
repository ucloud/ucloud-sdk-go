package functions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTimestamp(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid_length", args{20}, true},
		{"valid_length", args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTimestamp(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConcat(t *testing.T) {
	type args struct {
		input []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{[]interface{}{"foo", 42}}, "foo42", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Concat(tt.args.input...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Concat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchValue(t *testing.T) {
	type testStruct struct {
		Name string
		Data int
	}
	type args struct {
		arr         interface{}
		originKey   string
		originValue string
		destKey     string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"ok_foo",
			args{
				[]testStruct{
					{Name: "foo", Data: 42},
					{Name: "bar", Data: 142},
				},
				"Name",
				"foo",
				"Data",
			},
			42,
			false,
		},

		{
			"ok_bar",
			args{
				[]testStruct{
					{Name: "foo", Data: 42},
					{Name: "bar", Data: 142},
				},
				"Name",
				"bar",
				"Data",
			},
			142,
			false,
		},

		{
			"not",
			args{
				[]testStruct{
					{Name: "foo", Data: 42},
					{Name: "bar", Data: 142},
				},
				"Name",
				"foo",
				"Test",
			},
			"",
			true,
		},

		{
			"err_marshal",
			args{
				"test",
				"Name",
				"bar",
				"Data",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchValue(tt.args.arr, tt.args.originKey, tt.args.originValue, tt.args.destKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotS := fmt.Sprint(got)
			wantS := fmt.Sprint(tt.want)
			if gotS != wantS {
				t.Errorf("SearchValue() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestTimeDelta(t *testing.T) {
	type args struct {
		timeStamp int
		value     int
		typ       string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"add_hour", args{100000, 1, "hours"}, 103600, false},
		{"add_day", args{100000, 1, "days"}, 186400, false},
		{"subtract_hour", args{100000, -1, "hours"}, 96400, false},
		{"subtract_day", args{100000, -1, "days"}, 13600, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TimeDelta(tt.args.timeStamp, tt.args.value, tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeDelta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TimeDelta() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	type args struct {
		op    string
		left  interface{}
		right interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"sum", args{"+", 1, 1}, int64(2), false},
		{"sub", args{"-", 2, 1}, int64(1), false},
		{"multi", args{"*", 1, 3}, int64(3), false},
		{"divide", args{"/", 3, 1}, int64(3), false},
		{"invalid", args{"t", 3, 1}, 0, true},

		{"float_sum", args{"+", 1.1, 1}, 2.1, false},
		{"float_sub", args{"-", 2.1, 1}, 1.1, false},
		{"string_sub", args{"-", "2.1", "1"}, 0, true},
		{"float_multi", args{"*", 1, 3.0}, 3.0, false},
		{"float_divide", args{"/", 3.0, 1.0}, 3.0, false},
		{"float_invalid", args{"t", 3.0, 1.0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.op, tt.args.left, tt.args.right)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
