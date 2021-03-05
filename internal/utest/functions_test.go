package utest

import (
	"strings"
	"testing"
)

func TestGetZoneImage(t *testing.T) {
	type args struct {
		zone interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{"cn-bj2-02"}, "uimage-ixczxu", false},
		{"invalid_zone", args{"cn-bj100"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetZoneImage(tt.args.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetZoneImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetZoneImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImageResource(t *testing.T) {
	type args struct {
		region interface{}
		zone   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{"", "cn-bj2-02"}, "uimage-ixczxu", false},
		{"invalid_zone", args{"", "cn-bj100"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetImageResource(tt.args.region, tt.args.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImageResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetImageResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegionImage(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{"cn-bj2"}, "uimage-rq2kat", false},
		{"invalid_region", args{"cn-bj100"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegionImage(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegionImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRegionImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTimestamp(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// {"ok", args{"13"}, "", false}, // TODO: how to mock?
		{"invalid_length", args{"20"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimestamp(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatWithVertical(t *testing.T) {
	type args struct {
		input []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{[]interface{}{"foo", 42}}, "foo|42", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConcatWithVertical(tt.args.input...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcatWithVertical() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConcatWithVertical() = %v, want %v", got, tt.want)
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

func TestCalculate(t *testing.T) {
	type args []interface{}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"sum", args{"+", "1", "1"}, 2, false},
		{"sub", args{"-", "2", "1"}, 1, false},
		{"multi", args{"*", "1", "3"}, 3, false},

		{"number", args{"+", "1", 1}, 2, false},
		{"not number", args{"+", "1", "x"}, 0, true},

		{"literal", args{"101010"}, 101010, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args[0], tt.args[1:]...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
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
			"ok",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchValue(tt.args.arr, tt.args.originKey, tt.args.originValue, tt.args.destKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotS, _ := toString(got)
			wantS, _ := toString(tt.want)
			if gotS != wantS {
				t.Errorf("SearchValue() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestGetUUID(t *testing.T) {
	uuidS, _ := GetUUID()
	length := len(strings.Split(uuidS, "-"))
	if length != 4 {
		t.Errorf("GetUUID() length is %v, want %v", length, 4)
	}
}

func TestGetNotEqual(t *testing.T) {
	type args struct {
		v  interface{}
		vL []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"hit", args{1, []interface{}{1, 2}}, "2", false},
		{"miss", args{1, []interface{}{1, 1}}, "1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNotEqual(tt.args.v, tt.args.vL...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNotEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
