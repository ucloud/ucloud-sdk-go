# State Waiter

SDK also provide state waiter helper to improver usage experience.

## Why use it?

Waiter can wait for remote state is achieved to target state. such as,

- create and wait for resource state is completed.
- invoke/start/stop a resource and wait for it is finished.
- custom retry policies and wait for retrying is finished.

## Usage

By the default, waiter will use exponential back-off delay between twice request. you can see [advanced options](#Advanced Options) for detail.

```go
waiter.StateWaiter{
    Pending: []string{"pending"},
    Target:  []string{"available"},
    Refresh: CustomRefreshFn,
    Timeout: 5 * time.Minute,
}
```

You can custom you own state machine from your own function. custom state refresh function should matched ``waiter.RefreshFunc``:

```go
type RefreshFunc func() (result interface{}, state string, err error)
```

seems like:

```go
func CustomRefreshFn() (interface{}, string, error) {
    inst, err := describeUHostByID(uhostID)
    if err != nil {
        return nil, "", err
    }

    if inst == nil || inst.State != "Running" {
        return nil, "pending", nil
    }

    return inst, "available", nil
}
```

## Advanced Options

### Timeout

Timeout option will wait and refresh until target state or timeout is reached. State waiter use exponential back-off interval. 

Option ``Timeout`` is required. By the default, max interval is 10s, if no other options is set, the refresh interval is:

```
0 100ms 200ms 400ms 800ms 1.6s 3.2s 6.4s 10s 10s
```

if ``MinTimeout`` is set, the refresh interval will start by a minimal value:

```
0 3s 6s 10s 10s ... (100ms < MinTimeout <= 10s, eg. 3s)

0 11s 10s 10s 10s ... (10s < MinTimeout, eg. 11s)
```

if ``PollInterval`` is set, the exponential back-off interval will be disabled. it will use simple poll interval instead.

```
0 3s 3s 3s 3s ... (PollInterval, eg. 3s)
```

### Delay

Delay option will wait until delay time is passed before first state is refreshed.

```go
waiter.StateWaiter{
    Pending: []string{"pending"},
    Target:  []string{"available"},
    Refresh: CustomRefreshFn,
    Timeout: 5 * time.Minute,
    Delay: 30 * time.Second,
}
```

It is usually used in wait a little second after resource is created.

## Example

You can see also [wait for remote state example](../examples/wait) for a full example.
