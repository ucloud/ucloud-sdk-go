package utest

import (
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

func TestGetValue(t *testing.T) {
	type testSet struct {
		Name string
	}

	type test struct {
		RetCode int
		Message string
		DataSet []testSet
	}

	type args struct {
		obj  interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"ok",
			args{
				test{1, "some error", []testSet{{"1"}}},
				"DataSet.0.Name",
			},
			"1",
			false,
		},
		{
			"ok",
			args{
				test{1, "some error", []testSet{{"1"}}},
				"RetCode",
			},
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValue(tt.args.obj, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetReqValue(t *testing.T) {
	type Test struct {
		Name *string
		CPU  *int
		IPs  []string
	}
	testObj := Test{}

	SetReqValue(&testObj, "Name", "name")
	if *testObj.Name != "name" {
		t.Errorf("SetReqValue() = %#v, want %v", ucloud.StringValue(testObj.Name), "name")
	}

	SetReqValue(&testObj, "CPU", "1")
	if *testObj.CPU != 1 {
		t.Errorf("SetReqValue() = %#v, want %v", ucloud.IntValue(testObj.CPU), 1)
	}

	SetReqValue(&testObj, "IPs", "192.168.0.1")
	if testObj.IPs[0] != "192.168.0.1" {
		t.Errorf("SetReqValue() = %#v, want %v", testObj.IPs[0], "192.168.0.1")
	}

	SetReqValue(&testObj, "IPs", "192.168.0.1")
	if testObj.IPs[0] != "192.168.0.1" {
		t.Errorf("SetReqValue() = %#v, want %v", testObj.IPs[0], "192.168.0.1")
	}

	SetReqValue(&testObj, "IPs", "192.168.0.1", "192.168.0.2")
	if testObj.IPs[0] != "192.168.0.1" || testObj.IPs[1] != "192.168.0.2" {
		t.Errorf("SetReqValue() = %#v, want %v", testObj.IPs, []string{"192.168.0.1", "192.168.0.2"})
	}
}
