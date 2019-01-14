package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet2457(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	testSet2457CreateVPC00(&ctx)
	testSet2457CreateSubnet01(&ctx)
	testSet2457CreateULB02(&ctx)
	testSet2457CreateVServer03(&ctx)
	testSet2457CreateSSL04(&ctx)
	testSet2457CreateSSL05(&ctx)
	testSet2457DescribeSSL06(&ctx)
	testSet2457DescribeSSL07(&ctx)
	testSet2457BindSSL08(&ctx)
	testSet2457UpdateSSLBinding09(&ctx)
	testSet2457UnbindSSL10(&ctx)
	testSet2457DeleteSSL11(&ctx)
	testSet2457DeleteSSL12(&ctx)
	testSet2457DeleteVServer13(&ctx)
	testSet2457DeleteULB14(&ctx)
	testSet2457DeleteSubnet15(&ctx)
	testSet2457DeleteVPC16(&ctx)
}

func testSet2457CreateVPC00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "ulb-ssl-vpc"))
	ctx.NoError(utest.SetReqValue(req, "Network", "192.168.0.0/16"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["vpc_id"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet2457CreateSubnet01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "Subnet", "192.168.111.0"))

	ctx.NoError(utest.SetReqValue(req, "SubnetName", "ulb-ssl-subnet"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["subnet_id"] = ctx.Must(utest.GetValue(resp, "SubnetId"))
}

func testSet2457CreateULB02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewCreateULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBName", "ulb-ssl-test"))
	ctx.NoError(utest.SetReqValue(req, "Tag", "Default"))

	ctx.NoError(utest.SetReqValue(req, "InnerMode", "No"))
	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Dynamic"))
	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))
	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("subnet_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["ULBId"] = ctx.Must(utest.GetValue(resp, "ULBId"))
}

func testSet2457CreateVServer03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := ulbClient.NewCreateVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerName", "vserver-test"))
	ctx.NoError(utest.SetReqValue(req, "ListenType", "RequestProxy"))
	ctx.NoError(utest.SetReqValue(req, "Protocol", "HTTPS"))
	ctx.NoError(utest.SetReqValue(req, "FrontendPort", "443"))
	ctx.NoError(utest.SetReqValue(req, "Method", "Roundrobin"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceType", "UserDefined"))
	ctx.NoError(utest.SetReqValue(req, "PersistenceInfo", "huangchao"))
	ctx.NoError(utest.SetReqValue(req, "ClientTimeout", "60"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["VServerId"] = ctx.Must(utest.GetValue(resp, "VServerId"))
}

func testSet2457CreateSSL04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewCreateSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SSLName", "证书-1"))

	ctx.NoError(utest.SetReqValue(req, "UserCert", `
-----BEGIN CERTIFICATE-----
MIIFzTCCBLWgAwIBAgIQQ8IswmAhEIKfNhrKqb0F3DANBgkqhkiG9w0BAQsFADCB
lzELMAkGA1UEBhMCQ04xJTAjBgNVBAoTHFRydXN0QXNpYSBUZWNobm9sb2dpZXMs
IEluYy4xHzAdBgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxHTAbBgNVBAsT
FERvbWFpbiBWYWxpZGF0ZWQgU1NMMSEwHwYDVQQDExhUcnVzdEFzaWEgRFYgU1NM
IENBIC0gRzUwHhcNMTYxMjA2MDAwMDAwWhcNMTcxMjA2MjM1OTU5WjAgMR4wHAYD
VQQDDBVtLmVjb2xvZ3ktZW1vYmlsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDxBsuwGdCZdEUs40SQcvUt+9hlmLTgcfkq/h9f1QVPxLq/PC+O
sG76hOgy6N8f7k7x5XgtPKi9O4ydFl8ViYhEXRjYQcUrTm3lu7s9UT2AIUmK0dI+
PZgFU5gDwh8fQLoL24T2lPfkD9TngCnDanfo3xbx/e9hsJkf7hKWix8zrxtYYCUT
t96pTpQeWjr7ggl2bDEfTayJNM+i5xoGBPiQFdxPnKWCjNmXi2dws0d2whi1euRW
gI5wIXji5WKfUf6EvzG0Uzz6i8vsSLGv8pL7C0AuUI4MrPNDesFeA2LEYclQkpHE
E49BkpQvCokCW9d8/r5ASUry+7SrJIncU6FxAgMBAAGjggKJMIIChTAgBgNVHREE
GTAXghVtLmVjb2xvZ3ktZW1vYmlsZS5jb20wCQYDVR0TBAIwADBhBgNVHSAEWjBY
MFYGBmeBDAECATBMMCMGCCsGAQUFBwIBFhdodHRwczovL2Quc3ltY2IuY29tL2Nw
czAlBggrBgEFBQcCAjAZDBdodHRwczovL2Quc3ltY2IuY29tL3JwYTAfBgNVHSME
GDAWgBRtWMd/GufhPy6mjJc1Qrv00zisPzAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0l
BBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMIGbBggrBgEFBQcBAQSBjjCBizA8Bggr
BgEFBQcwAYYwaHR0cDovL3RydXN0YXNpYTItb2NzcC5kaWdpdGFsY2VydHZhbGlk
YXRpb24uY29tMEsGCCsGAQUFBzAChj9odHRwOi8vdHJ1c3Rhc2lhMi1haWEuZGln
aXRhbGNlcnR2YWxpZGF0aW9uLmNvbS90cnVzdGFzaWFnNS5jcnQwggEDBgorBgEE
AdZ5AgQCBIH0BIHxAO8AdQDd6x0reg1PpiCLga2BaHB+Lo6dAdVciI09EcTNtuy+
zAAAAVjT7zdSAAAEAwBGMEQCIDCzWufc1q7hjmrrCetGyoA8EsEqpRSIhmZXStX5
8b7zAiA6x5aAaDK+yMyeAgw71yi3tRVrWayHN+W0+4BxC8u5UQB2AO5Lvbd1zmC6
4UJpH6vhnmajD35fsHLYgwDEe4l6qP3LAAABWNPvN4kAAAQDAEcwRQIgZ/LNgg7n
7AE4O2yZkrXNcqAOmJ3NU2nT6zcnBxPFTTsCIQCjyPbMfWMZTD3kxgxPQ1COw5zJ
sM0dfNmSr3MiU7EhqDANBgkqhkiG9w0BAQsFAAOCAQEAeyfgUhg9ZWVCaz0f+BQU
6fMMfmQ1BDzvVFu+ORoAqyJQogxwIdfjrlz/63YFee5qpUsW/aaz4ma3bb4dpE1K
GsgYe5N3o0xybYlOj+KB61sufYkzQS3HgDevCwjfUlGEbNl4dpO2xh5s5AANXlnz
s/X0+AJ33/bm+fWIjAbIjluaEoM6GETHTXi4Tlxy0j3nsXsB9tIIUibAdTtButef
JJRnikGRN+eHjrsLYe0RUmdKOQz1ik6teHt0MQX0aCe8OlXeyGDd9m8u7+y0nAnH
TVaNuT7vXMWyyXLVUcV898wkBo3Bo3hUiaw0QR0ttgDrf5ZwqPfqpytRW2K5GMZT
uw==
-----END CERTIFICATE-----


-----BEGIN CERTIFICATE-----
MIIFZTCCBE2gAwIBAgIQOhAOfxCeGsWcxf/2QNXkQjANBgkqhkiG9w0BAQsFADCB
yjELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJp
U2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxW
ZXJpU2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0
aG9yaXR5IC0gRzUwHhcNMTYwODExMDAwMDAwWhcNMjYwODEwMjM1OTU5WjCBlzEL
MAkGA1UEBhMCQ04xJTAjBgNVBAoTHFRydXN0QXNpYSBUZWNobm9sb2dpZXMsIElu
Yy4xHzAdBgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxHTAbBgNVBAsTFERv
bWFpbiBWYWxpZGF0ZWQgU1NMMSEwHwYDVQQDExhUcnVzdEFzaWEgRFYgU1NMIENB
IC0gRzUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC39aSJZG/97x3a
6Qmuc9+MubagegRAVUmFYHTYTs8IKB2pM7wXN7W8mekdZaEgUjDFxvRBK/DhTb7U
8ONLsKKdT86aOhzbz2noCTn9wPWnGwkg+/4YKg/dPQQdV9tMsSu0cwqInWHxSAkm
AI1hYFC9D7Sf7Hp/5cRcD+dK454YMRzNOGLQnCVI8JEqrz6o9SOvQNTqTcfqt6DC
0UlXG+MPD1eNPjlzf1Vwaab+VSTgySoC+Ikbq2VsdykeOiGXW/OIiASH7+2LcR05
PmQ7GEOlM8yzoVojFpM8sHz+WxI05ZOPri5+vX3HhHHjWr5432G0dVmgohnZvlVZ
oy8XrlbpAgMBAAGjggF2MIIBcjASBgNVHRMBAf8ECDAGAQH/AgEAMC8GA1UdHwQo
MCYwJKAioCCGHmh0dHA6Ly9zLnN5bWNiLmNvbS9wY2EzLWc1LmNybDAOBgNVHQ8B
Af8EBAMCAQYwLgYIKwYBBQUHAQEEIjAgMB4GCCsGAQUFBzABhhJodHRwOi8vcy5z
eW1jZC5jb20wYQYDVR0gBFowWDBWBgZngQwBAgEwTDAjBggrBgEFBQcCARYXaHR0
cHM6Ly9kLnN5bWNiLmNvbS9jcHMwJQYIKwYBBQUHAgIwGRoXaHR0cHM6Ly9kLnN5
bWNiLmNvbS9ycGEwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMCkGA1Ud
EQQiMCCkHjAcMRowGAYDVQQDExFTeW1hbnRlY1BLSS0yLTYwMTAdBgNVHQ4EFgQU
bVjHfxrn4T8upoyXNUK79NM4rD8wHwYDVR0jBBgwFoAUf9Nlp8Ld7LvwMAnzQzn6
Aq8zMTMwDQYJKoZIhvcNAQELBQADggEBABUphhBbeG7scE3EveIN0dOjXPgwgQi8
I2ZAKYm6DawoGz1lEJVdvFmkyMbP973X80b7mKmn0nNbe1kjA4M0O0hHaMM1ZaEv
7e9vHEAoGyysMO6HzPWYMkyNxcCV7Nos2Uv4RvLDpQHh7P4Kt6fUU13ipcynrtQD
1lFUM0yoTzwwFsPu3Pk+94hL58ErqwqJQwxoHMgLIQeMVHeNKcWFy1bddSbIbCWU
Zs6cMxhrra062ZCpDCbxyEaFNGAtYQMqNz55Z/14XgSUONZ/cJTns6QKhpcgTOwB
fnNzRnk+aWreP7osKhXlz4zs+llP7goBDKFOMMtoEXx3YjJCKgpqmBU=
-----END CERTIFICATE-----
	`))
	ctx.NoError(utest.SetReqValue(req, "PrivateKey", "abc"))
	ctx.NoError(utest.SetReqValue(req, "CaCert", `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA8QbLsBnQmXRFLONEkHL1LfvYZZi04HH5Kv4fX9UFT8S6vzwv
jrBu+oToMujfH+5O8eV4LTyovTuMnRZfFYmIRF0Y2EHFK05t5bu7PVE9gCFJitHS
Pj2YBVOYA8IfH0C6C9uE9pT35A/U54Apw2p36N8W8f3vYbCZH+4SlosfM68bWGAl
E7feqU6UHlo6+4IJdmwxH02siTTPoucaBgT4kBXcT5ylgozZl4tncLNHdsIYtXrk
VoCOcCF44uVin1H+hL8xtFM8+ovL7Eixr/KS+wtALlCODKzzQ3rBXgNixGHJUJKR
xBOPQZKULwqJAlvXfP6+QElK8vu0qySJ3FOhcQIDAQABAoIBAAPvZnfzk/JNcauv
8jihh9s+V2QhQCLB+Z14FK8N3U5WGe5xXx1nSAiTDu912d69l1BfvLyQVvjv9fXC
nb7ORglHs9YkDMIOP8EWdZIkt2pWIMtBbbtSah78JGk7TCLIfcEfzmXwPLPehk1Z
TFVCcb69lbRRvwzLQ1TAIFGQ5+uCEkW02KAl6kx+JnVpsE8/BjqZKG1Ne+sM6dOC
GRd44hgiNHKUT3Xtbw6jttiUFDLKYMYtb7PpRAkZFM8tgnBV6dWWJ3xTYW9kOjPh
XnScNARfphUZVibRhA04og5p1q/MUz9Sz9g2DURuSlo/MP3WZMbVRvZiUN1xhz5v
2WhsddkCgYEA+gWPFo0TbVbZXUrx9J/ptI9NXNx5zjyUrv87MDt1pnmMDgWrsCEI
RqQR4Lp2G11GA7IudiA/ipcZqgcRIIFvb+gu1kObox3BGGs59x+DqFeAPXt6dFG2
W10f9k96/tcbdursurqwd3Zv3cqQqRTKgaP4xHFmexlcwGCF5YwewWMCgYEA9sos
2acNINXwcNRUPnpg82DOrG9Zjr1aiNo9PDJmwGEdC9QMOUWM85dq0M9g388ttiLU
Wr/U4r5yDuqWJPcKtff2BaxSsZpcQ4Id9eddD9L+sxaBGyD23RtOC+IOlkG6WS4g
iUYulQvW69tBHWiwxQu7YMSIE2B3EuySPOQYlBsCgYEAxNwvqB/4lfT2PUDPdj+b
cnILBf0LY1nL8GZCol2O6z91CW1pm8rGi2iQMxRd/nnYsPxRHO2TWnpS2M+rqp5/
settRYQCPdMlwSZcg7oqnhgXf1GEP6Y/IX0Xt4cpXxLcKywarYRlggqdVlMyyA74
zE7hhzuK5442u7rEctN7O+UCgYAoM78ipafp1XAZsT0YAG+Stg504J7CNe5tpL+c
8sjyRd+pcZ2cJsxTUjNAWMf7LZDQvtPBBMb1OPjznRtgYi4IfqBBRFUkQXUOOkAP
MuViEokTO3NErBYK5svL+8NMjuCAbpc2RYyJEyiru0fcNpW1Q7f+h4VzQp+jIY6h
BLdMSQKBgGauU7OQksZCEY2MVAcD5dShYYvWLxOkj4dVVwISN1M6ImCAHwXZ6Nak
6YlzCGT+NbRJbB2cPfsrKXtAJVX15I3iDCKAoGkb+9kiHnPj7Q71KVuWQE6BQx7E
vE88TSsshwtX1s+qU9UWUrMPodK32q5nO3p8N033NvS9wLNfbcdc
-----END RSA PRIVATE KEY-----
	`))
	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["SSLId_01"] = ctx.Must(utest.GetValue(resp, "SSLId"))
}

func testSet2457CreateSSL05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewCreateSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SSLName", "证书-2"))

	ctx.NoError(utest.SetReqValue(req, "UserCert", `
-----BEGIN CERTIFICATE-----
MIIFzTCCBLWgAwIBAgIQQ8IswmAhEIKfNhrKqb0F3DANBgkqhkiG9w0BAQsFADCB
lzELMAkGA1UEBhMCQ04xJTAjBgNVBAoTHFRydXN0QXNpYSBUZWNobm9sb2dpZXMs
IEluYy4xHzAdBgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxHTAbBgNVBAsT
FERvbWFpbiBWYWxpZGF0ZWQgU1NMMSEwHwYDVQQDExhUcnVzdEFzaWEgRFYgU1NM
IENBIC0gRzUwHhcNMTYxMjA2MDAwMDAwWhcNMTcxMjA2MjM1OTU5WjAgMR4wHAYD
VQQDDBVtLmVjb2xvZ3ktZW1vYmlsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDxBsuwGdCZdEUs40SQcvUt+9hlmLTgcfkq/h9f1QVPxLq/PC+O
sG76hOgy6N8f7k7x5XgtPKi9O4ydFl8ViYhEXRjYQcUrTm3lu7s9UT2AIUmK0dI+
PZgFU5gDwh8fQLoL24T2lPfkD9TngCnDanfo3xbx/e9hsJkf7hKWix8zrxtYYCUT
t96pTpQeWjr7ggl2bDEfTayJNM+i5xoGBPiQFdxPnKWCjNmXi2dws0d2whi1euRW
gI5wIXji5WKfUf6EvzG0Uzz6i8vsSLGv8pL7C0AuUI4MrPNDesFeA2LEYclQkpHE
E49BkpQvCokCW9d8/r5ASUry+7SrJIncU6FxAgMBAAGjggKJMIIChTAgBgNVHREE
GTAXghVtLmVjb2xvZ3ktZW1vYmlsZS5jb20wCQYDVR0TBAIwADBhBgNVHSAEWjBY
MFYGBmeBDAECATBMMCMGCCsGAQUFBwIBFhdodHRwczovL2Quc3ltY2IuY29tL2Nw
czAlBggrBgEFBQcCAjAZDBdodHRwczovL2Quc3ltY2IuY29tL3JwYTAfBgNVHSME
GDAWgBRtWMd/GufhPy6mjJc1Qrv00zisPzAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0l
BBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMIGbBggrBgEFBQcBAQSBjjCBizA8Bggr
BgEFBQcwAYYwaHR0cDovL3RydXN0YXNpYTItb2NzcC5kaWdpdGFsY2VydHZhbGlk
YXRpb24uY29tMEsGCCsGAQUFBzAChj9odHRwOi8vdHJ1c3Rhc2lhMi1haWEuZGln
aXRhbGNlcnR2YWxpZGF0aW9uLmNvbS90cnVzdGFzaWFnNS5jcnQwggEDBgorBgEE
AdZ5AgQCBIH0BIHxAO8AdQDd6x0reg1PpiCLga2BaHB+Lo6dAdVciI09EcTNtuy+
zAAAAVjT7zdSAAAEAwBGMEQCIDCzWufc1q7hjmrrCetGyoA8EsEqpRSIhmZXStX5
8b7zAiA6x5aAaDK+yMyeAgw71yi3tRVrWayHN+W0+4BxC8u5UQB2AO5Lvbd1zmC6
4UJpH6vhnmajD35fsHLYgwDEe4l6qP3LAAABWNPvN4kAAAQDAEcwRQIgZ/LNgg7n
7AE4O2yZkrXNcqAOmJ3NU2nT6zcnBxPFTTsCIQCjyPbMfWMZTD3kxgxPQ1COw5zJ
sM0dfNmSr3MiU7EhqDANBgkqhkiG9w0BAQsFAAOCAQEAeyfgUhg9ZWVCaz0f+BQU
6fMMfmQ1BDzvVFu+ORoAqyJQogxwIdfjrlz/63YFee5qpUsW/aaz4ma3bb4dpE1K
GsgYe5N3o0xybYlOj+KB61sufYkzQS3HgDevCwjfUlGEbNl4dpO2xh5s5AANXlnz
s/X0+AJ33/bm+fWIjAbIjluaEoM6GETHTXi4Tlxy0j3nsXsB9tIIUibAdTtButef
JJRnikGRN+eHjrsLYe0RUmdKOQz1ik6teHt0MQX0aCe8OlXeyGDd9m8u7+y0nAnH
TVaNuT7vXMWyyXLVUcV898wkBo3Bo3hUiaw0QR0ttgDrf5ZwqPfqpytRW2K5GMZT
uw==
-----END CERTIFICATE-----


-----BEGIN CERTIFICATE-----
MIIFZTCCBE2gAwIBAgIQOhAOfxCeGsWcxf/2QNXkQjANBgkqhkiG9w0BAQsFADCB
yjELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTowOAYDVQQLEzEoYykgMjAwNiBWZXJp
U2lnbiwgSW5jLiAtIEZvciBhdXRob3JpemVkIHVzZSBvbmx5MUUwQwYDVQQDEzxW
ZXJpU2lnbiBDbGFzcyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0
aG9yaXR5IC0gRzUwHhcNMTYwODExMDAwMDAwWhcNMjYwODEwMjM1OTU5WjCBlzEL
MAkGA1UEBhMCQ04xJTAjBgNVBAoTHFRydXN0QXNpYSBUZWNobm9sb2dpZXMsIElu
Yy4xHzAdBgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxHTAbBgNVBAsTFERv
bWFpbiBWYWxpZGF0ZWQgU1NMMSEwHwYDVQQDExhUcnVzdEFzaWEgRFYgU1NMIENB
IC0gRzUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC39aSJZG/97x3a
6Qmuc9+MubagegRAVUmFYHTYTs8IKB2pM7wXN7W8mekdZaEgUjDFxvRBK/DhTb7U
8ONLsKKdT86aOhzbz2noCTn9wPWnGwkg+/4YKg/dPQQdV9tMsSu0cwqInWHxSAkm
AI1hYFC9D7Sf7Hp/5cRcD+dK454YMRzNOGLQnCVI8JEqrz6o9SOvQNTqTcfqt6DC
0UlXG+MPD1eNPjlzf1Vwaab+VSTgySoC+Ikbq2VsdykeOiGXW/OIiASH7+2LcR05
PmQ7GEOlM8yzoVojFpM8sHz+WxI05ZOPri5+vX3HhHHjWr5432G0dVmgohnZvlVZ
oy8XrlbpAgMBAAGjggF2MIIBcjASBgNVHRMBAf8ECDAGAQH/AgEAMC8GA1UdHwQo
MCYwJKAioCCGHmh0dHA6Ly9zLnN5bWNiLmNvbS9wY2EzLWc1LmNybDAOBgNVHQ8B
Af8EBAMCAQYwLgYIKwYBBQUHAQEEIjAgMB4GCCsGAQUFBzABhhJodHRwOi8vcy5z
eW1jZC5jb20wYQYDVR0gBFowWDBWBgZngQwBAgEwTDAjBggrBgEFBQcCARYXaHR0
cHM6Ly9kLnN5bWNiLmNvbS9jcHMwJQYIKwYBBQUHAgIwGRoXaHR0cHM6Ly9kLnN5
bWNiLmNvbS9ycGEwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMCkGA1Ud
EQQiMCCkHjAcMRowGAYDVQQDExFTeW1hbnRlY1BLSS0yLTYwMTAdBgNVHQ4EFgQU
bVjHfxrn4T8upoyXNUK79NM4rD8wHwYDVR0jBBgwFoAUf9Nlp8Ld7LvwMAnzQzn6
Aq8zMTMwDQYJKoZIhvcNAQELBQADggEBABUphhBbeG7scE3EveIN0dOjXPgwgQi8
I2ZAKYm6DawoGz1lEJVdvFmkyMbP973X80b7mKmn0nNbe1kjA4M0O0hHaMM1ZaEv
7e9vHEAoGyysMO6HzPWYMkyNxcCV7Nos2Uv4RvLDpQHh7P4Kt6fUU13ipcynrtQD
1lFUM0yoTzwwFsPu3Pk+94hL58ErqwqJQwxoHMgLIQeMVHeNKcWFy1bddSbIbCWU
Zs6cMxhrra062ZCpDCbxyEaFNGAtYQMqNz55Z/14XgSUONZ/cJTns6QKhpcgTOwB
fnNzRnk+aWreP7osKhXlz4zs+llP7goBDKFOMMtoEXx3YjJCKgpqmBU=
-----END CERTIFICATE-----
	`))
	ctx.NoError(utest.SetReqValue(req, "PrivateKey", "abc"))
	ctx.NoError(utest.SetReqValue(req, "CaCert", `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA8QbLsBnQmXRFLONEkHL1LfvYZZi04HH5Kv4fX9UFT8S6vzwv
jrBu+oToMujfH+5O8eV4LTyovTuMnRZfFYmIRF0Y2EHFK05t5bu7PVE9gCFJitHS
Pj2YBVOYA8IfH0C6C9uE9pT35A/U54Apw2p36N8W8f3vYbCZH+4SlosfM68bWGAl
E7feqU6UHlo6+4IJdmwxH02siTTPoucaBgT4kBXcT5ylgozZl4tncLNHdsIYtXrk
VoCOcCF44uVin1H+hL8xtFM8+ovL7Eixr/KS+wtALlCODKzzQ3rBXgNixGHJUJKR
xBOPQZKULwqJAlvXfP6+QElK8vu0qySJ3FOhcQIDAQABAoIBAAPvZnfzk/JNcauv
8jihh9s+V2QhQCLB+Z14FK8N3U5WGe5xXx1nSAiTDu912d69l1BfvLyQVvjv9fXC
nb7ORglHs9YkDMIOP8EWdZIkt2pWIMtBbbtSah78JGk7TCLIfcEfzmXwPLPehk1Z
TFVCcb69lbRRvwzLQ1TAIFGQ5+uCEkW02KAl6kx+JnVpsE8/BjqZKG1Ne+sM6dOC
GRd44hgiNHKUT3Xtbw6jttiUFDLKYMYtb7PpRAkZFM8tgnBV6dWWJ3xTYW9kOjPh
XnScNARfphUZVibRhA04og5p1q/MUz9Sz9g2DURuSlo/MP3WZMbVRvZiUN1xhz5v
2WhsddkCgYEA+gWPFo0TbVbZXUrx9J/ptI9NXNx5zjyUrv87MDt1pnmMDgWrsCEI
RqQR4Lp2G11GA7IudiA/ipcZqgcRIIFvb+gu1kObox3BGGs59x+DqFeAPXt6dFG2
W10f9k96/tcbdursurqwd3Zv3cqQqRTKgaP4xHFmexlcwGCF5YwewWMCgYEA9sos
2acNINXwcNRUPnpg82DOrG9Zjr1aiNo9PDJmwGEdC9QMOUWM85dq0M9g388ttiLU
Wr/U4r5yDuqWJPcKtff2BaxSsZpcQ4Id9eddD9L+sxaBGyD23RtOC+IOlkG6WS4g
iUYulQvW69tBHWiwxQu7YMSIE2B3EuySPOQYlBsCgYEAxNwvqB/4lfT2PUDPdj+b
cnILBf0LY1nL8GZCol2O6z91CW1pm8rGi2iQMxRd/nnYsPxRHO2TWnpS2M+rqp5/
settRYQCPdMlwSZcg7oqnhgXf1GEP6Y/IX0Xt4cpXxLcKywarYRlggqdVlMyyA74
zE7hhzuK5442u7rEctN7O+UCgYAoM78ipafp1XAZsT0YAG+Stg504J7CNe5tpL+c
8sjyRd+pcZ2cJsxTUjNAWMf7LZDQvtPBBMb1OPjznRtgYi4IfqBBRFUkQXUOOkAP
MuViEokTO3NErBYK5svL+8NMjuCAbpc2RYyJEyiru0fcNpW1Q7f+h4VzQp+jIY6h
BLdMSQKBgGauU7OQksZCEY2MVAcD5dShYYvWLxOkj4dVVwISN1M6ImCAHwXZ6Nak
6YlzCGT+NbRJbB2cPfsrKXtAJVX15I3iDCKAoGkb+9kiHnPj7Q71KVuWQE6BQx7E
vE88TSsshwtX1s+qU9UWUrMPodK32q5nO3p8N033NvS9wLNfbcdc
-----END RSA PRIVATE KEY-----
	`))
	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.CreateSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

	ctx.Vars["SSLId_02"] = ctx.Must(utest.GetValue(resp, "SSLId"))
}

func testSet2457DescribeSSL06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.SSLId", ctx.GetVar("SSLId_01"), "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DescribeSSL07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDescribeSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_02")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DescribeSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
			ctx.NewValidator("DataSet.0.SSLId", ctx.GetVar("SSLId_02"), "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457BindSSL08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewBindSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.BindSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457UpdateSSLBinding09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := pulbClient.NewUpdateSSLBindingRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "OldSSLId", ctx.GetVar("SSLId_01")))
	ctx.NoError(utest.SetReqValue(req, "NewSSLId", ctx.GetVar("SSLId_02")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pulbClient.UpdateSSLBinding(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457UnbindSSL10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewUnbindSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_02")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.UnbindSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteSSL11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewDeleteSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	// ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	// ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))
	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_01")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteSSL12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewDeleteSSLRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SSLId", ctx.GetVar("SSLId_02")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteSSL(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteVServer13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ulbClient.NewDeleteVServerRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))
	ctx.NoError(utest.SetReqValue(req, "VServerId", ctx.GetVar("VServerId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteVServer(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteULB14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ulbClient.NewDeleteULBRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "ULBId", ctx.GetVar("ULBId")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ulbClient.DeleteULB(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteSubnet15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := vpcClient.NewDeleteSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("subnet_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}

func testSet2457DeleteVPC16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", "0", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {
		ctx.T.Fatal(err)
	}

}
