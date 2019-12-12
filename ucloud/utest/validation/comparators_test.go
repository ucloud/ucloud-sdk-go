package validation

import (
	"testing"
)

func Test_lenEq(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name               string
		args               args
		isNotExpectedError bool
		wantErr            bool
	}{
		{"equal", args{[]map[string]string{{"foo": "bar"}}, 1}, false, false},
		{"not_equal", args{[]map[string]string{{"foo": "bar"}}, 2}, true, true},
		{"error", args{[]map[string]string{{"foo": "bar"}}, "x"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := lenEq(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("lenEq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if IsNotExpectedError(err) != tt.isNotExpectedError {
				t.Errorf("lenEq() IsNotExpectedError = %v, want %v", IsNotExpectedError(err), tt.isNotExpectedError)
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
		name               string
		args               args
		isNotExpectedError bool
		wantErr            bool
	}{
		{"arrIn", args{[]string{"1", "2"}, "1"}, false, false},
		{"arrNotIn", args{[]string{"1", "2"}, "3"}, true, true},
		{"strIn", args{"12", "1"}, false, false},
		{"strNotIn", args{"12", "3"}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := contains(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("contains(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
				return
			}
			if IsNotExpectedError(err) != tt.isNotExpectedError {
				t.Errorf("contains(%v) IsNotExpectedError = %v, want %v", tt.args, IsNotExpectedError(err), tt.isNotExpectedError)
			}
		})
	}
}
