package auth

import (
	"strings"
	"testing"
)

func TestCredential_CreateSign(t *testing.T) {
	type fields struct {
		PublicKey  string
		PrivateKey string
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"standard",
			fields{"ucloudsomeone@example.com1296235120854146120", "46f09bb9fab4f12dfc160dae12273d5332b5debe"},
			args{testCredential_CreateSign_query00},
			"4f9ef5df2abab2c6fccd1e9515cb7e2df8c6bb65",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Credential{
				PublicKey:  tt.fields.PublicKey,
				PrivateKey: tt.fields.PrivateKey,
			}
			if got := c.CreateSign(tt.args.query); got != tt.want {
				t.Errorf("Credential.CreateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testCredential_CreateSign_query00 string = strings.Join(strings.Split(`Action=CreateUHostInstance
&CPU=2
&ChargeType=Month
&DiskSpace=10
&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2
&LoginMode=Password
&Memory=2048
&Name=Host01
&Password=VUNsb3VkLmNu
&PublicKey=ucloudsomeone%40example.com1296235120854146120
&Quantity=1
&Region=cn-bj2
&Zone=cn-bj2-04`, "\n"), "")
