English | [简体中文](README_cn.md)

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

## Installation

### Use `go get`

```bash
go get github.com/ucloud/ucloud-sdk-go
```

**Note** if meet network problem, you can use go proxy to speed up the downloaded, eg: use GOPROXY environment variable

```go
export GOPROXY=https://goproxy.io
```

Replay the command to retry installation.

### Use `go mod`

Add the following snippet to any code.

```go
import _ "github.com/ucloud/ucloud-sdk-go"
```

And execute this commands：

```bash
go mod init
go mod tidy
```

**Note**：If using `go mod` and `Goland IDE` together, please search `vgo` on `Settings`, and enable `vgo` support.

**Note**：If using `go mod` 和 `GOPATH`, notice the `go mod init/tidy` can not run with `GOPATH`, please move out current project from `GOPATH`.

### Use `dep`

```bash
dep ensure -add github.com/ucloud/ucloud-sdk-go
```

## First Using

Currently, Go SDK use `PublicKey/PrivateKey` as authentication method, the key can be found from：

- [UAPI Key Generation](https://console.ucloud.cn/uapi/apikey)

Here is a simple example：

```go
package main

import (
    "fmt"

    "github.com/ucloud/ucloud-sdk-go/ucloud"
    "github.com/ucloud/ucloud-sdk-go/ucloud/auth"
    "github.com/ucloud/ucloud-sdk-go/services/uhost"
)

func main() {
    cfg := ucloud.NewConfig()
    cfg.Region = "cn-bj2"

    // replace the public/private key by your own
    credential := auth.NewCredential()
    credential.PrivateKey = "my_private_key"
    credential.PublicKey = "my_public_key"

    uhostClient := uhost.NewClient(&cfg, &credential)

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

    // send request
    newUHost,err := uhostClient.CreateUHostInstance(req)
    if err != nil {
        fmt.Printf("something bad happened: %s\n", err)
    } else {
        fmt.Printf("resource id of the uhost: %s\n", newUHost.UHostIds[0])
    }
}
```

Replace the client configuration and host image id by your custom value, then you can create a cloud host。

At this example, you had already completed a `CreateUHostInstance` request by UCloud Go SDK, it is covered most of feature of sdk, you can write your own script for free!

Each API call in the SDK has a detailed comment and document, You can use Editor/IDE to jump to the specific API method code to view (can also to see [Go Doc](https://godoc.org/github.com/ucloud/ucloud-sdk-go)）, and use completion of IDE and error logging to inspect usage of SDK。

If you are interested the advanced usage about above code example, please see the documentation: 

- [Configuration](https://docs.ucloud.cn/opensdk-go/configure), Learn how to configure SDK, such as Logging, Retrying, Service endpoint(Public Cloud & Private Cloud), and etc.
- [Error Handling](https://docs.ucloud.cn/opensdk-go/error), Learn how to resolve the several exception types, include parameters error, business error for *RetCode is not 0*.
- [Type System](https://docs.ucloud.cn/opensdk-go/typesystem), Learn how to validate parameters, and normalize the response of API.
- [Request Middleware](https://docs.ucloud.cn/opensdk-go/middleware), Learn how to intercept request that applied by SDK, and add common logic to the lifecycle.
- [Toolbox & Helpers](https://docs.ucloud.cn/opensdk-go/helpers), The additional helpers, such as the poll function for waiting resource state.

## More Examples

### Examples based on Environment

SDK provided some example based on environment, and provide resource provisioning logic . You can click the following link to view details：

- [Client Configuration](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/configure), Introducing the usage of configuration
- [Read External Configuration](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/external), Use external configuration, such as configuration from *UCloud CLI*
- [Logging](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/logging), Configure options of logging
- [Retrying](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/retry), Introducing the auto retrying feature
- [Toolbox: State Poll](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/wait), Learn how to wait host opened gracefully 
- [Create Host](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/uhost), Introducing how to create host
- [Create Two-Tier architecture based on Load Balancer](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/two-tier), ULB + UHost

### Examples based on Request

UCloud `UAPI` product provided examples based on request, you can fill request parameters and generate example of SDK, it can be copied immediately:

- [Open UAPI](https://console.ucloud.cn/uapi/ucloudapi)
- Select your interested API, such as [CreateUHostInstance](https://console.ucloud.cn/uapi/detail?id=CreateUHostInstance)
- Fill parameters, copy the sdk example code at right panel
- Save code into `main.go`
- Executing `go mod init main`
- Executing `go mod tidy`
- `go run ./main.go`

**Note** if meet network problem, you can use go proxy to speed up the downloaded, eg: use GOPROXY environment variable

```go
export GOPROXY=https://goproxy.io
```

**Note**：If using `go mod` and `Goland IDE` together, please search `vgo` on `Settings`, and enable `vgo` support.

**Note**：If using `go mod` 和 `GOPATH`, notice the `go mod init/tidy` can not run with `GOPATH`, please move out current project from `GOPATH`.

## Feedback & Contribution

- [Issue](https://github.com/ucloud/ucloud-sdk-go/issues)
- [Pull Request](https://github.com/ucloud/ucloud-sdk-go/pulls)
