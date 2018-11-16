# UCloud Go SDK

[![GoDoc](https://godoc.org/github.com/ucloud/ucloud-sdk-go?status.svg)](https://godoc.org/github.com/ucloud/ucloud-sdk-go)
[![Google Groups](https://img.shields.io/badge/chat-google%20groups-brightgreen.svg)](https://groups.google.com/forum/#!forum/ucloud-sdk-go)
[![GitHub (pre-)release](https://img.shields.io/github/release/ucloud/ucloud-sdk-go/all.svg)](https://github.com/ucloud/ucloud-sdk-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ucloud/ucloud-sdk-go)](https://goreportcard.com/report/github.com/ucloud/ucloud-sdk-go)
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

```
dep ensure -add github.com/ucloud/ucloud-sdk-go
```
## Usage

```golang
import {
   "github.com/ucloud/ucloud-sdk-go"
} 
```

Create a new UCloud client of each service to access the different parts of the UCloud API.

## Authentication

Currently, user public key & private key is the only method of authenticating with the API. You could get your keys here: [Key Generation](https://console.ucloud.cn/uapi/apikey)

You can then use your keys to create a new client of uhost service: 

```go
package main

import (
    "github.com/ucloud/ucloud-sdk-go/sdk"
    "github.com/ucloud/ucloud-sdk-go/sdk/auth"
    "github.com/ucloud/ucloud-sdk-go/services/uhost"
)

func loadConfig()(*sdk.Config, *auth.Credential){
    cfg := sdk.NewConfig()

    credential := auth.NewCredential()
    credential.PrivateKey ="my_privatekey"
    credential.PublicKey = "my_publickey"

    return &cfg, &credential
}

func main(){
    cfg, credentail := loadConfig()
    uhostClient := uhost.NewClient(cfg, credential)
}
```

## Quick Start

To create a new uhost:

```go
// build Request
req := uhostClient.NewCreateUHostInstanceRequest()
req.Name       = sdk.String("sdk-example-uhost")
req.Zone       = sdk.String("cn-bj2-05")
req.ImageId    = sdk.String("uimage-ixczxu")
req.LoginMode  = sdk.String("Password")
req.Password   = sdk.String("my_uhost_password")
req.ChargeType = sdk.String("Dynamic")
req.CPU        = sdk.Int(1)
req.Memory     = sdk.Int(1024)
req.Tag        = sdk.String("sdk-example")

// send request
newUHost,err := uhostClient.CreateUHostInstance(req)

if err != nil{
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
    Size: sdk.Int(20),
    Type: sdk.String("CLOUD_NORMAL"),
    IsBoot: sdk.Bool(false),
}
req.Disks = []uhost.UHostDisk{dataDisk}
```

will encoded as `Disks.0.Size=20&Disks.0.Type=CLOUD_NORMAL&Disks.0.IsBoot=false`

## Docs

- [Configure Document](./docs/Configure.md)
- [Usage Document](./docs/Usage.md)
- [Retry Policy Document](./docs/Retry.md)
- [State Waiter Document](./docs/Wait.md)

## Feedback & Contribution

- [Issue](https://github.com/ucloud/ucloud-sdk-go/issues)
- [Pull Request](https://github.com/ucloud/ucloud-sdk-go/pulls)
