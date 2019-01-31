package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetString(t *testing.T) {
	type args struct {
		data []string
	}

	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{"normal", args{[]string{"foo", "bar", "foobar"}}, 3},
		{"distinct", args{[]string{"foo", "foo", "bar"}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSet(hashString, nil)

			for _, v := range tt.args.data {
				s.Add(v)
			}

			assert.Equal(t, s.Len(), tt.wantLen)
			assert.Subset(t, tt.args.data, s.List())

			for _, v := range s.List() {
				s.Remove(v)
			}

			assert.Equal(t, s.Len(), 0)
		})
	}
}

func TestSetInt(t *testing.T) {
	type args struct {
		data []int
	}

	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{"normal", args{[]int{1, 2, 3}}, 3},
		{"distinct", args{[]int{1, 1, 2}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSet(hashInt, nil)

			for _, v := range tt.args.data {
				s.Add(v)
			}

			assert.Equal(t, s.Len(), tt.wantLen)
			assert.Subset(t, tt.args.data, s.List())

			for _, v := range s.List() {
				s.Remove(v)
			}

			assert.Equal(t, s.Len(), 0)
		})
	}
}
