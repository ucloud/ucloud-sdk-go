package utest

import (
	"testing"
)

func Test_lenEq(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"equal", args{[]map[string]string{{"foo": "bar"}}, 1}, true, false},
		{"not_equal", args{[]map[string]string{{"foo": "bar"}}, 2}, false, false},
		{"error", args{[]map[string]string{{"foo": "bar"}}, "x"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lenEq(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("lenEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lenEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"arrIn", args{[]string{"1", "2"}, "1"}, true, false},
		{"arrNotIn", args{[]string{"1", "2"}, "3"}, false, false},
		{"strIn", args{"12", "1"}, true, false},
		{"strNotIn", args{"12", "3"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := contains(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("contains(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("contains(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
