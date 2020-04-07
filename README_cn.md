[English](README.md) | 简体中文

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

## 安装

### 使用 go get 安装

```bash
go get github.com/ucloud/ucloud-sdk-go
```

**Note** 如果遇到网络不稳定，可以使用代理服务器来加速下载，例如使用 GOPROXY 加速：

```go
export GOPROXY=https://goproxy.cn
```

再次执行安装命令即可。

### 使用 go mod 安装

在任意代码中添加

```go
import _ "github.com/ucloud/ucloud-sdk-go"
```

之后在项目根目录执行：

```bash
go mod init
go mod tidy
```

**Note**：如同时使用 go mod 和 Goland IDE，请在 Settings 中搜索 vgo，并启用 vgo 支持。

**Note**：如同时使用 go mod 和 GOPATH，注意 go mod init/tidy 不能在 GOPATH 下执行，把项目从 GOPATH 下移出即可。

### 使用 dep 安装

```bash
dep ensure -add github.com/ucloud/ucloud-sdk-go
```

## 初次使用

目前，Go SDK 使用 PublicKey/PrivateKey 作为唯一的鉴权方式，该公私钥可以从以下途径获取：

- [UAPI 密钥管理](https://console.ucloud.cn/uapi/apikey)

下面提供一个简单的示例：

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

将上述代码中 client 相关配置，以及主机的 image id 等，替换成自己的配置，即可创建一台云主机。

在该示例中，使用 Go SDK 完成了一个创建云主机的请求。至此，已经涵盖了 SDK 的基本核心用法，可以构建自己的脚本啦！

SDK 中的每一个 api 调用都有详细的注释文档，可以通过 Editor/IDE 跳转到具体的方法中查看（也可以[查看接口文档](https://godoc.org/github.com/ucloud/ucloud-sdk-go)），并根据 IDE 自动补全和报错信息继续探索 SDK 的用法。

如果需要了解这段代码提及但未完全覆盖的使用技巧，请参考：

- [通用配置](https://docs.ucloud.cn/opensdk-go/configure)，了解如何配置 SDK，如日志、重试、服务访问端点（公有云、专有云）等。
- [错误处理](https://docs.ucloud.cn/opensdk-go/error)，了解如何处理不同类型的 SDK 异常，包括参数错误，RetCode 不为 0 的业务异常等。
- [类型系统](https://docs.ucloud.cn/opensdk-go/typesystem)，了解 SDK 如何校验参数，并规范化 API 的返回值。
- [请求中间件](https://docs.ucloud.cn/opensdk-go/middleware)，了解如何拦截 SDK 发起的请求，并统一添加额外的逻辑。
- [工具箱](https://docs.ucloud.cn/opensdk-go/helpers)，SDK 提供的额外支持，如状态轮询函数等。

## 获取更多示例

### 基于场景的示例

SDK 提供了部分基于场景的示例，并提供了对应的资源销毁逻辑，可以点击以下链接查看源码：

- [客户端配置](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/configure)，介绍各种配置项的作用
- [读取外部配置](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/external)，使用外部配置文件，如 UCloud CLI 配置文件
- [日志](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/logging)，配置日志选项
- [重试](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/retry)，演示自动重试的功能
- [工具箱-状态轮询](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/wait)，演示如何优雅地等待主机开机 
- [批量创建云主机](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/uhost)，演示如何批量创建云主机
- [创建基于负载均衡器的两层架构](https://github.com/ucloud/ucloud-sdk-go/tree/master/examples/two-tier)，ULB + UHost

### 基于请求的示例

控制台 UAPI 产品提供了基于请求的示例，可以在控制台填写请求参数，自动生成 SDK 样例，可以直接拷贝使用，详情请见：

- [打开 UAPI](https://console.ucloud.cn/uapi/ucloudapi)
- 选择希望调用的 API，如 [CreateUHostInstance](https://console.ucloud.cn/uapi/detail?id=CreateUHostInstance)
- 填写参数，拷贝界面右侧 Go SDK 的示例代码
- 保存请求代码为 `main.go`
- 执行 `go mod init main`
- 执行 `go mod tidy`
- `go run ./main.go`

**Note** 如果遇到网络不稳定，可以使用代理服务器来加速下载，例如使用 GOPROXY 加速：

```go
export GOPROXY=https://goproxy.cn
```

**Note**：如同时使用 go mod 和 Goland IDE，请在 Settings 中搜索 vgo，并启用 vgo 支持。

**Note**：如同时使用 go mod 和 GOPATH，注意 go mod init/tidy 不能在 GOPATH 下执行，把项目从 GOPATH 下移出即可。

## 反馈 & 贡献

- [Issue](https://github.com/ucloud/ucloud-sdk-go/issues)
- [Pull Request](https://github.com/ucloud/ucloud-sdk-go/pulls)
