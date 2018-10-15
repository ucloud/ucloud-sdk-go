# Configure SDK

## Client Config

client config can control many common behaviors of sdk, it is required for any client.

### Common Request Fields

To set common region and project id for any request, you can also overwrite it by request.

```go
cfg := ucloud.NewConfig()

cfg.Region = "cn-bj2"
cfg.ProjectId = "xxx"
```

### Auto Retry

To enable auto-retry for any request, default is disabled(max retries is 0), it will auto retry for

* network error
* server error with http status 429, 502, 503 and 504

```go
cfg.MaxRetries = 3
```

See [Retry Policy](./Retry.md) for more details.

### Logging

To set log level, you can see log with different level on **stdout**, you can also use it to enable or disable log.

```go
// disable sdk log, default is log.InfoLevel
cfg.LogLevel = log.PanicLevel

// enable sdk debug log level
// if debug log is enable, it will print all of request params and response body
cfg.LogLevel = log.DebugLevel
```

### Timeout

To set timeout for any network request, default is 30s.

```go
cfg.Timeout = 30 * time.Second
```

### Internal Usage

the followed configuration should not be set in general usage.

To set User-Agent, you can append custom UserAgent for any request, it also used for setting up special channel for customer.

```go
cfg.UserAgent = "UCloud-CLI/0.1.0"
```

## Credential Config

To set credential info for any request, credential is required for any client
```go
credential := auth.NewCredential()

credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")
credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
```
