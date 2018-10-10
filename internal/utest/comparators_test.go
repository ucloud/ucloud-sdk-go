package utest

import "testing"

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
