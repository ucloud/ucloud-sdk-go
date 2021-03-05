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
		{"interfaceIn", args{[]interface{}{"test"}, "test"}, false, false},
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

func TestComparators(t *testing.T) {
	c := NewComparators()
	type args struct {
		name     string
		value    interface{}
		expected interface{}
	}
	tests := []struct {
		name               string
		args               args
		isNotExpectedError bool
	}{
		{"eq_ok", args{"eq", "a", "a"}, false},
		{"ne_ok", args{"ne", "a", "b"}, false},
		{"abs_eq_ok", args{"abs_eq", "0.1", "0.1"}, false},
		{"lt_ok", args{"lt", "0.1", "0.2"}, false},
		{"le_ok", args{"le", "0.1", "0.1"}, false},
		{"gt_ok", args{"gt", "0.2", "0.1"}, false},
		{"ge_ok", args{"ge", "0.2", "0.2"}, false},
		{"str_eq_ok", args{"str_eq", "a", "a"}, false},
		{"float_eq_ok", args{"float_eq", "0.1", "0.1"}, false},
		{"len_eq_ok", args{"len_eq", "abc", 3}, false},
		{"len_gt_ok", args{"len_gt", "abc", 1}, false},
		{"len_ge_ok", args{"len_ge", "abc", 3}, false},
		{"len_lt_ok", args{"len_lt", "a", 3}, false},
		{"len_le_ok", args{"len_le", "a", 3}, false},
		{"contains_ok", args{"contains", "abc", "a"}, false},
		{"contained_by_ok", args{"contained_by", "a", "abc"}, false},
		{"type_ok", args{"type", "a", "string"}, false},
		{"regex_ok", args{"regex", "a", "a"}, false},
		{"startswith_ok", args{"startswith", "abc", "a"}, false},
		{"endswith_ok", args{"endswith", "abc", "c"}, false},
		{"object_contains_ok", args{"object_contains", "abc", "a"}, false},
		{"object_not_contains_ok", args{"object_not_contains", "abc", "d"}, false},
		{"eq_no", args{"eq", "a", "b"}, true},
		{"ne_no", args{"ne", "a", "a"}, true},
		{"abs_eq_no", args{"abs_eq", "0.1", "0.2"}, true},
		{"lt_no", args{"lt", "0.2", "0.1"}, true},
		{"le_no", args{"le", "0.2", "0.1"}, true},
		{"gt_no", args{"gt", "0.1", "0.2"}, true},
		{"ge_no", args{"ge", "0.1", "0.2"}, true},
		{"str_eq_no", args{"str_eq", "a", "b"}, true},
		{"float_eq_no", args{"float_eq", "0.1", "0.2"}, true},
		{"len_eq_no", args{"len_eq", "abc", 1}, true},
		{"len_gt_no", args{"len_gt", "abc", 3}, true},
		{"len_ge_no", args{"len_ge", "abc", 4}, true},
		{"len_lt_no", args{"len_lt", "a", 1}, true},
		{"len_le_no", args{"len_le", "abc", 1}, true},
		{"contains_no", args{"contains", "abc", "d"}, true},
		{"contained_by_no", args{"contained_by", "a", "def"}, true},
		{"type_no", args{"type", "a", "int"}, true},
		{"regex_no", args{"regex", "a", "b"}, true},
		{"startswith_no", args{"startswith", "abc", "c"}, true},
		{"endswith_no", args{"endswith", "abc", "a"}, true},
		{"object_contains_no", args{"object_contains", "abc", "d"}, true},
		{"object_not_contains_no", args{"object_not_contains", "abc", "a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := c.Get(tt.args.name)
			err := f(tt.args.value, tt.args.expected)
			if (err != nil) != tt.isNotExpectedError {
				t.Errorf("Comparators.%s(%v, %v) error = %v, wantErr %v", tt.args.name, tt.args.value, tt.args.expected, err, tt.isNotExpectedError)
				return
			}
		})
	}

}
