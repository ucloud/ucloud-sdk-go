<p align="center">
    <img src="https://ucloud-sdk.dl.ufileos.com/logos%2Flogo-mini.png" />
</p>

<h1 align="center">UCloud Go SDK</h1>

[![GitHub (pre-)release](https://img.shields.io/github/release/ucloud/ucloud-sdk-go/all.svg)](https://github.com/ucloud/ucloud-sdk-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ucloud/ucloud-sdk-go)](https://goreportcard.com/report/github.com/ucloud/ucloud-sdk-go)
[![Codecov](https://codecov.io/gh/ucloud/ucloud-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/ucloud/ucloud-sdk-go)
[![Build Status](https://travis-ci.org/ucloud/ucloud-sdk-go.svg?branch=master)](https://travis-ci.org/ucloud/ucloud-sdk-go)
[![SonarQube](https://sonarcloud.io/api/project_badges/measure?project=ucloud-sdk-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=ucloud-sdk-go)
[![GoDoc](https://godoc.org/github.com/ucloud/ucloud-sdk-go?status.svg)](https://godoc.org/github.com/ucloud/ucloud-sdk-go)
[![GitHub](https://img.shields.io/github/license/ucloud/ucloud-sdk-go.svg)](http://www.apache.org/licenses/LICENSE-2.0)

- [Website](https://www.ucloud.cn/)
- [Documentation](https://docs.ucloud.cn/opensdk-go/README)
- [Google Group](https://groups.google.com/forum/#!forum/ucloud-sdk-go)

UCloud SDK is a Go client library for accessing the UCloud API.

This client can run on Linux, macOS and Windows. It requires Golang 1.10.x and above. 

## Install

### via source

```go
go get github.com/ucloud/ucloud-sdk-go
```

### via dep

```bash
dep ensure -add github.com/ucloud/ucloud-sdk-go
```

## Usage

```go
import (
   "github.com/ucloud/ucloud-sdk-go"
)
```

Create a new UCloud client of each service to access the different parts of the UCloud API.

## Authentication

Currently, user public key & private key is the only method of authenticating with the API. You could get your keys here: [Key Generation](https://console.ucloud.cn/uapi/apikey)

You can then use your keys to create a new client of uhost service: 

```go
package main

import (
    "github.com/ucloud/ucloud-sdk-go/ucloud"
    "github.com/ucloud/ucloud-sdk-go/ucloud/auth"
    "github.com/ucloud/ucloud-sdk-go/services/uhost"
)

func main() {
    cfg := ucloud.NewConfig()

    credential := auth.NewCredential()
    // replace the public/private key by your own
    credential.PrivateKey ="my_private_key"
    credential.PublicKey = "my_public_key"

    uhostClient := uhost.NewClient(&cfg, &credential)

    // apply a request ...
}
```

## Quick Start

To create a new uhost, here is an example to construct an api request:

```go
// build Request
req := uhostClient.NewCreateUHostInstanceRequest()
req.Name       = ucloud.String("sdk-example-uhost")
req.Zone       = ucloud.String("cn-bj2-05")
req.ImageId    = ucloud.String("uimage-xxx") // you can replace the image with an available id
req.LoginMode  = ucloud.String("Password")
req.Password   = ucloud.String("my_uhost_password")
req.ChargeType = ucloud.String("Dynamic")
req.CPU        = ucloud.Int(1)
req.Memory     = ucloud.Int(1024)
req.Tag        = ucloud.String("sdk-example")
```

### Complex Query

You can set array as query for `DescribeUHostInstance` action, such as:

```go
req.UHostIds = []string{"uhost-xxx", "uhost-yyy"}
```

will encoded as `UHostIds.0=uhost-xxx&UHostIds.1=uhost-yyy`

You can also set the complex struct array for `CreateUHostInstance` action, such as:

```go
dataDisk := uhost.UHostDisk{
    Size: ucloud.Int(40),
    Type: ucloud.String("CLOUD_SSD"),
    IsBoot: ucloud.Bool(true),
}
req.Disks = []uhost.UHostDisk{dataDisk}
```

will encoded as `Disks.0.Size=20&Disks.0.Type=CLOUD_NORMAL&Disks.0.IsBoot=false`.

Then apply the request simply.

```go
// send request
newUHost, err := uhostClient.CreateUHostInstance(req)
if err != nil {
    fmt.Printf("something bad happened: %s\n", err)
} else {
    fmt.Printf("resource id of the uhost: %s\n", newUHost.UHostIds[0])
}
```

There are more examples in the [examples](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples) folder. 
You can copy any of them for your advanced usage. 

## Documentations

If you need more complex topic, please read the docs.

- [Configure Document](./docs/en-us/Configure.md)
- [Usage Document](./docs/en-us/Usage.md)
- [Retry Policy Document](./docs/en-us/Retry.md)
- [State Waiter Document](./docs/en-us/Wait.md)

If you are a chinese user, you can also visit the official website for documentation:

- [Official Website](https://docs.ucloud.cn/opensdk-go/README)

## Feedback & Contribution

- [Issue](https://github.com/ucloud/ucloud-sdk-go/issues)
- [Pull Request](https://github.com/ucloud/ucloud-sdk-go/pulls)
