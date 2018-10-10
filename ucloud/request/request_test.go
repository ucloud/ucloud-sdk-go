package request

import (
	"reflect"
	"testing"
)

func TestToQueryMap(t *testing.T) {
	type Composite struct {
		Region string

		unexported string
	}

	type args struct {
		req Common
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"Ok",
			args{
				req: &struct {
					CommonBase
					Id      int
					Name    string
					IsValid bool
					Ips     []string
				}{
					Id:      1,
					Name:    "lilei",
					IsValid: true,
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]string{
				"Id":      "1",
				"Name":    "lilei",
				"IsValid": "true",
				"Ips.0":   "127.0.0.1",
				"Ips.1":   "192.168.1.1",
			},
			false,
		},
		{
			"Partial",
			args{
				req: &struct {
					CommonBase
					Id      int
					Name    string
					IsValid bool
					Ips     []string
				}{
					Id:      1,
					Name:    "",
					IsValid: true,
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]string{
				"Id":      "1",
				"IsValid": "true",
				"Ips.0":   "127.0.0.1",
				"Ips.1":   "192.168.1.1",
			},
			false,
		},
		{
			"IsComposited",
			args{
				req: &struct {
					CommonBase
					Composite
				}{
					CommonBase{},
					Composite{Region: "cn-bj-2"},
				},
			},
			map[string]string{"Region": "cn-bj-2"},
			false,
		},
		{
			"IsCompositedByArray",
			args{
				req: &struct {
					CommonBase
					Arr []Composite
				}{
					CommonBase{},
					[]Composite{
						{Region: "cn-bj2"},
						{Region: "cn-sh1"},
					},
				},
			},
			map[string]string{
				"Arr.0.Region": "cn-bj2",
				"Arr.1.Region": "cn-sh1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToQueryMap(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
