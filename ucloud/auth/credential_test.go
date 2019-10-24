package auth

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
	"time"
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
			args{testCredentialCreateSignQuery00},
			"4f9ef5df2abab2c6fccd1e9515cb7e2df8c6bb65",
		},
		{
			"unorder",
			fields{"ucloudsomeone@example.com1296235120854146120", "46f09bb9fab4f12dfc160dae12273d5332b5debe"},
			args{testCredentialCreateSignQuery01},
			"4f9ef5df2abab2c6fccd1e9515cb7e2df8c6bb65",
		},
		{
			"noPublicKey",
			fields{"ucloudsomeone@example.com1296235120854146120", "46f09bb9fab4f12dfc160dae12273d5332b5debe"},
			args{testCredentialCreateSignQuery02},
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

func TestCredential_BuildCredentialedQuery(t *testing.T) {
	type fields struct {
		PublicKey  string
		PrivateKey string
	}
	type args struct {
		params map[string]string
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
			args{testCredentialBuildCredentialedQuery01},
			"4f9ef5df2abab2c6fccd1e9515cb7e2df8c6bb65",
		},
		{
			"longArray",
			fields{"ucloudsomeone@example.com1296235120854146120", "46f09bb9fab4f12dfc160dae12273d5332b5debe"},
			args{testCredentialBuildCredentialedQuery02},
			"6bc32642779f8cb0bb3d3f128b1c0688abf9a1a3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCredential()
			c.PublicKey = tt.fields.PublicKey
			c.PrivateKey = tt.fields.PrivateKey

			if got := c.BuildCredentialedQuery(tt.args.params); !strings.Contains(got, tt.want) {
				t.Errorf("Credential.BuildCredentialedQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_STSCredential(t *testing.T) {
	c := Credential{
		PublicKey:     "ucloudsomeone@example.com1296235120854146120",
		PrivateKey:    "46f09bb9fab4f12dfc160dae12273d5332b5debe",
		SecurityToken: "some_stoken",
		CanExpire:     true,
		Expires:       time.Time{},
	}
	query := c.BuildCredentialedQuery(testCredentialBuildCredentialedQuery01)
	values, err := url.ParseQuery(query)
	assert.NoError(t, err)
	assert.Equal(t, "170c480ad176a247b324eb92a2cfe536aacfbd04", values.Get("Signature"))
	assert.True(t, c.IsExpired())
}

var testCredentialCreateSignQuery00 = strings.Join(strings.Split(`Action=CreateUHostInstance
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

var testCredentialCreateSignQuery01 = strings.Join(strings.Split(`Action=CreateUHostInstance
&CPU=2
&Region=cn-bj2
&DiskSpace=10
&LoginMode=Password
&ChargeType=Month
&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2
&PublicKey=ucloudsomeone%40example.com1296235120854146120
&Memory=2048
&Password=VUNsb3VkLmNu
&Name=Host01
&Quantity=1
&Zone=cn-bj2-04`, "\n"), "")

var testCredentialCreateSignQuery02 = strings.Join(strings.Split(`Action=CreateUHostInstance
&CPU=2
&ChargeType=Month
&DiskSpace=10
&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2
&LoginMode=Password
&Memory=2048
&Name=Host01
&Password=VUNsb3VkLmNu
&Quantity=1
&Region=cn-bj2
&Zone=cn-bj2-04`, "\n"), "")

var testCredentialBuildCredentialedQuery01 = map[string]string{
	"Action":     "CreateUHostInstance",
	"CPU":        "2",
	"ChargeType": "Month",
	"DiskSpace":  "10",
	"ImageId":    "f43736e1-65a5-4bea-ad2e-8a46e18883c2",
	"LoginMode":  "Password",
	"Memory":     "2048",
	"Name":       "Host01",
	"Password":   "VUNsb3VkLmNu",
	"Quantity":   "1",
	"Region":     "cn-bj2",
	"Zone":       "cn-bj2-04",
}

var testCredentialBuildCredentialedQuery02 = map[string]string{
	"Action":      "AllocateBackendBatch",
	"Backends.0":  "foo",
	"Backends.1":  "bar",
	"Backends.2":  "42",
	"Backends.3":  "foo",
	"Backends.4":  "bar",
	"Backends.5":  "42",
	"Backends.6":  "foo",
	"Backends.7":  "bar",
	"Backends.8":  "42",
	"Backends.9":  "foo",
	"Backends.10": "bar",
	"Backends.11": "42",
}
