package utest

import (
	"testing"
)

func Test_checkFloats(t *testing.T) {
	type args struct {
		a         interface{}
		b         interface{}
		checkFunc func(float64, float64) (bool, error)
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"ok", args{"0.1", "0.1", func(a, b float64) (bool, error) { return a == b, nil }}, true, false},
		{"check_false", args{"0.1", "0.2", func(a, b float64) (bool, error) { return a == b, nil }}, false, false},
		{"check_round", args{"0.675", "0.676", func(a, b float64) (bool, error) { return a == b, nil }}, true, false},

		{"not_float", args{"0.1", "x", func(a, b float64) (bool, error) { return a == b, nil }}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkFloats(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFloats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"ok", args{"1"}, "1", false},
		{"ok_int", args{1}, "1", false},
		{"ok_float", args{1.0}, "1", false},
		{"ok_struct", args{struct{ Name string }{Name: "1"}}, "{Name:1}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkStrings(t *testing.T) {
	type Test struct {
		Name string
	}

	type args struct {
		a         interface{}
		b         interface{}
		checkFunc func(string, string) (bool, error)
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"ok", args{"a", "a", func(a, b string) (bool, error) { return a == b, nil }}, true, false},
		{"ok_int", args{1, 1, func(a, b string) (bool, error) { return a == b, nil }}, true, false},
		{"ok_float", args{1.0, 1.0, func(a, b string) (bool, error) { return a == b, nil }}, true, false},
		{"check_false", args{"a", "b", func(a, b string) (bool, error) { return a == b, nil }}, false, false},
		{"check_struct", args{Test{"a"}, Test{"b"}, func(a, b string) (bool, error) { return a == b, nil }}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkStrings(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLens(t *testing.T) {
	type args struct {
		a         interface{}
		b         interface{}
		checkFunc func(int, int) (bool, error)
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"ok", args{[]string{"1"}, []string{"1"}, func(a, b int) (bool, error) { return a == b, nil }}, true, false},
		{"ok_str", args{"1", "1", func(a, b int) (bool, error) { return a == b, nil }}, true, false},
		{"check_false", args{[]string{"1"}, []string{}, func(a, b int) (bool, error) { return a == b, nil }}, false, false},

		{"cannot_get_length", args{"0.1", 1, func(a, b int) (bool, error) { return a == b, nil }}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkLens(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkLens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkLens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFloat(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"ok_int", args{1}, 1.0, false},
		{"ok_float", args{1.0}, 1.0, false},
		{"ok_str", args{"0.1"}, 0.1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toFloat(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundTo(t *testing.T) {
	type args struct {
		num  float64
		perc uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"ok", args{1.123, 2}, 1.12},
		{"ok_int", args{1, 2}, 1.00},

		{"ok_to_even", args{0.675, 2}, 0.68},
		{"ok_to_odd", args{0.685, 2}, 0.69},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundTo(tt.args.num, tt.args.perc); got != tt.want {
				t.Errorf("roundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLen(t *testing.T) {
	type test struct {
		Name string
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"str", args{"abc"}, 3, false},
		{"slice", args{[]string{"abc"}}, 1, false},
		{"map", args{map[string]string{"Name": "abc"}}, 1, false},

		{"int", args{123}, 0, true},
		{"float", args{0.675}, 0, true},
		{"nil", args{nil}, 0, true},
		{"struct", args{test{Name: "abc"}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toLen(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
