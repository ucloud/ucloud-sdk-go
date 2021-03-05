package validation

import (
	"testing"
)

func Test_checkFloats(t *testing.T) {
	type args struct {
		a         interface{}
		b         interface{}
		checkFunc func(float64, float64) error
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		isNotExpectedError bool
	}{
		{"ok", args{"0.1", "0.1", func(aVal, bVal float64) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"check_false", args{"0.1", "0.2", func(aVal, bVal float64) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, true},
		{"check_round", args{"0.675", "0.676", func(aVal, bVal float64) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},

		{"not_float", args{"0.1", "x", func(aVal, bVal float64) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, false},
		{"not_bool", args{false, true, func(aVal, bVal float64) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkFloats(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFloats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if IsNotExpectedError(err) != tt.isNotExpectedError {
				t.Errorf("checkFloats()IsNotExpectedError = %v, want %v", IsNotExpectedError(err), tt.isNotExpectedError)
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
		checkFunc func(string, string) error
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		isNotExpectedError bool
	}{
		{"ok", args{"a", "a", func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"ok_int", args{1, 1, func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"ok_float", args{1.0, 1.0, func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"check_false", args{"a", "b", func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, true},
		{"check_struct", args{Test{"a"}, Test{"b"}, func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, true},
		{"check_bool", args{"False", "false", func(aVal, bVal string) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkStrings(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if IsNotExpectedError(err) != tt.isNotExpectedError {
				t.Errorf("checkStrings() IsNotExpectedError = %v, want %v", IsNotExpectedError(err), tt.isNotExpectedError)
			}
		})
	}
}

func Test_checkLens(t *testing.T) {
	type args struct {
		a         interface{}
		b         interface{}
		checkFunc func(int, int) error
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		isNotExpectedError bool
	}{
		{"ok", args{[]string{"1"}, 1, func(aVal, bVal int) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"ok_str", args{"1", "1", func(aVal, bVal int) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, false, false},
		{"check_false", args{[]string{"1"}, []string{}, func(aVal, bVal int) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, false},

		{"cannot_get_length", args{"0.1", 1, func(aVal, bVal int) error {
			if aVal != bVal {
				return NewNotExpectedError()
			}
			return nil
		}}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkLens(tt.args.a, tt.args.b, tt.args.checkFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkLens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if IsNotExpectedError(err) != tt.isNotExpectedError {
				t.Errorf("checkLens() IsNotExpectedError = %v, want %v", IsNotExpectedError(err), tt.isNotExpectedError)
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

func Test_toString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"ok", args{"1"}, "1", false},
		{"ok_int", args{1}, "1", false},
		{"ok_float", args{1.0}, "1", false},
		{"ok_struct", args{struct{ Name string }{Name: "1"}}, `{"Name":"1"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
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
		{"not_bool", args{true}, 0.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toFloat(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
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
				t.Errorf("ToLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"ok_int", args{1}, 1, false},
		{"ok_float", args{1.0}, 1, false},
		{"not_str", args{"0.1"}, 0, true},
		{"ok_str", args{"1"}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toInt(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
