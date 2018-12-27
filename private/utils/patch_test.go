package utils

import (
	"reflect"
	"testing"
)

func TestPatcher_Patch(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		p    *RegexpPatcher
		args args
		want []byte
	}{
		{
			"RetCodeMiss", RetCodePatcher,
			args{[]byte(`"RetCode": 100,`)},
			[]byte(`"RetCode": 100,`),
		},
		{
			"RetCodeHit", RetCodePatcher,
			args{[]byte(`"RetCode": "100",`)},
			[]byte(`"RetCode": 100,`),
		},
		{
			"PortMiss", PortPatcher,
			args{[]byte(`"Port": 100,`)},
			[]byte(`"Port": 100,`),
		},
		{
			"PortHit", PortPatcher,
			args{[]byte(`"Port": "100",`)},
			[]byte(`"Port": 100,`),
		},
		{
			"NoNameMiss", RetCodePatcher,
			args{[]byte(`"Port": "100",`)},
			[]byte(`"Port": "100",`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Patch(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Patcher.Patch() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
