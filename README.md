# UCloud Go SDK

[![GitHub (pre-)release](https://img.shields.io/github/release/ucloud/ucloud-sdk-go/all.svg)](https://github.com/ucloud/ucloud-sdk-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ucloud/ucloud-sdk-go)](https://goreportcard.com/report/github.com/ucloud/ucloud-sdk-go)
[![Build Status](https://travis-ci.org/ucloud/ucloud-sdk-go.svg?branch=master)](https://travis-ci.org/ucloud/ucloud-sdk-go)
[![SonarQube](https://sonarcloud.io/api/project_badges/measure?project=ucloud-sdk-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=ucloud-sdk-go)
[![GoDoc](https://godoc.org/github.com/ucloud/ucloud-sdk-go?status.svg)](https://godoc.org/github.com/ucloud/ucloud-sdk-go)
[![GitHub](https://img.shields.io/github/license/ucloud/ucloud-sdk-go.svg)](http://www.apache.org/licenses/LICENSE-2.0)

- Website: https://www.ucloud.cn/
- Mailing list: [Google Group](https://groups.google.com/forum/#!forum/ucloud-sdk-go)

![UCloud Logo](http://cli-ucloud-logo.sg.ufileos.com/ucloud.png)

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

func loadConfig() (*ucloud.Config, *auth.Credential) {
    cfg := ucloud.NewConfig()

    credential := auth.NewCredential()
    credential.PrivateKey ="my_privatekey"
    credential.PublicKey = "my_publickey"

    return &cfg, &credential
}

func main() {
    cfg, credentail := loadConfig()
    uhostClient := uhost.NewClient(cfg, credential)
}
```

## Quick Start

To create a new uhost:

```go
// build Request
req := uhostClient.NewCreateUHostInstanceRequest()
req.Name       = ucloud.String("sdk-example-uhost")
req.Zone       = ucloud.String("cn-bj2-05")
req.ImageId    = ucloud.String("uimage-ixczxu")
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
}

fmt.Printf("resource id of the uhost: %s\n", newUHost.UHostIds[0])
```

**Note**

UHost created above cannot be accessed via Internet unless an EIP is created and bind to the UHost.

### Complex structure in query

You can also set array as query, such as:

```go
req.UHostIds = []string{"uhost-xxx", "uhost-yyy"}
```

will encoded as `UHostIds.0=uhost-xxx&UHostIds.1=uhost-yyy`

```go
dataDisk := uhost.UHostDisk{
    Size: ucloud.Int(20),
    Type: ucloud.String("CLOUD_NORMAL"),
    IsBoot: ucloud.Bool(false),
}
req.Disks = []uhost.UHostDisk{dataDisk}
```

will encoded as `Disks.0.Size=20&Disks.0.Type=CLOUD_NORMAL&Disks.0.IsBoot=false`

## Docs

- [Configure Document](./docs/en-us/Configure.md)
- [Usage Document](./docs/en-us/Usage.md)
- [Retry Policy Document](./docs/en-us/Retry.md)
- [State Waiter Document](./docs/en-us/Wait.md)

## Feedback & Contribution

- [Issue](https://github.com/ucloud/ucloud-sdk-go/issues)
- [Pull Request](https://github.com/ucloud/ucloud-sdk-go/pulls)
