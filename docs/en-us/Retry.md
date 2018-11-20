# Retry Policy

## How to enable retrying

Retrying feature is disabled by default.

hen retrying is enabled, such as client-level max retries is greater than 0 (enable retrying for all request), eg.

```go
cfg.MaxRetries = 3
```

or request-level max retries is greater than 0 (enable retrying for current request), eg.

```go
req.WithRetry(3)
```

## Retrying and side-effect

SDK has default retry policy to prevent retrying with side-effect. It will prevent most of action when follow term: ``create``, ``allocate``, ``clone``, ``copy``, ``upload``, is contained.

this actions will be set request-level retryable is false as default,

```go
req.SetRetryable(false)
```

you can set it to true to re-enable it simply.

```
req.SetRetryable(true)
```

Here is a table to show which action will prevent retrying by default:

| Action                  | Is Prevented |
| ----------------------- | ------------ |
| CreateUHostInstance     | √            |
| CopyCustomImage         | √            |
| CreateCustomImage       | √            |
| AllocateEIP             | √            |
| AllocateVIP             | √            |
| CreateBandwidthPackage  | √            |
| AllocateShareBandwidth  | √            |
| CreateFirewall          | √            |
| CreateULB               | √            |
| CreateVServer           | √            |
| CreatePolicy            | √            |
| CreateSSL               | √            |
| CreateVPC               | √            |
| AddVPCNetwork           | √            |
| CreateSubnet            | √            |
| CreateVPCIntercom       | √            |
| CreateRouteTable        | √            |
| CloneRouteTable         | √            |
| CreateUDisk             | √            |
| CloneUDisk              | √            |
| CreateProject           | √            |
| CreateGlobalSSHInstance | √            |
| CreateUDBInstance       | √            |
| BackupUDBInstance        | √            |
| UploadUDBParamGroup         | √            |
| CreateUDBParamGroup             | √            |
| CreateUDBSlave              | √            |
| CreateUDBReplicationInstance           | √            |
| CreateUDBRouteInstance | √            |
| CreateUDBInstanceByRecovery       | √            |
| CreateUMemSpace        | √            |
| CreateUMemcacheGroup         | √            |
| CreateURedisGroup             | √            |
| AllocateUDPN              | √            |

## Example

You can see also [retry example](../examples/retry) for a full example.
