# Manual Proxy and Interceptor Implementation

While the Parsley CLI can automatically generate proxies for your services, you can also implement them manually using the `pkg/features` package. This provides maximum flexibility for custom interception logic.

## Interceptors

Interceptors allow you to run code before and after a method call, or when an error occurs. To create an interceptor, you can embed `features.InterceptorBase` and implement the `features.MethodInterceptor` interface.

### The MethodInterceptor Interface

The `MethodInterceptor` interface defines the hooks for method interception:

```golang
type MethodInterceptor interface {
    Interceptor
    Enter(target any, methodName string, parameters []ParameterInfo)
    Exit(target any, methodName string, returnValues []ReturnValueInfo)
    OnError(target any, methodName string, err error)
}
```

### Implementing an Interceptor

Use `features.NewInterceptorBase` to initialize the base interceptor with a name and execution position.

```golang
type loggingInterceptor struct {
    features.InterceptorBase
}

func NewLoggingInterceptor() features.MethodInterceptor {
    return &loggingInterceptor{
        InterceptorBase: features.NewInterceptorBase("logging", 0),
    }
}

func (i *loggingInterceptor) Enter(target any, methodName string, parameters []features.ParameterInfo) {
    fmt.Printf("Entering method: %s\n", methodName)
}

func (i *loggingInterceptor) Exit(target any, methodName string, returnValues []features.ReturnValueInfo) {
    fmt.Printf("Exiting method: %s\n", methodName)
}

func (i *loggingInterceptor) OnError(target any, methodName string, err error) {
    fmt.Printf("Error in method %s: %v\n", methodName, err)
}
```

## Proxies

A proxy wraps a target object and invokes registered interceptors during method execution. You can create a proxy by embedding `features.ProxyBase`.

### Creating a Proxy

Use `features.NewProxyBase` to create a new proxy instance for a target object and a set of interceptors.

```golang
func NewProxyBase[T any](target T, interceptors []MethodInterceptor) ProxyBase
```

### Manual Proxy Example

To manually implement a proxy for a service, you must create a new type that implements the service interface and delegates calls to the target while invoking the proxy base methods.

```golang
type greeterProxy struct {
    features.ProxyBase
    target Greeter
}

func NewGreeterProxy(target Greeter, interceptors []features.MethodInterceptor) Greeter {
    return &greeterProxy{
        ProxyBase: features.NewProxyBase(target, interceptors),
        target:    target,
    }
}

func (p *greeterProxy) Greet() string {
    // 1. Create call context
    callCtx := features.NewMethodCallContext("Greet", nil, nil, "string")
    
    // 2. Invoke 'Enter' interceptors
    p.InvokeEnterMethodInterceptors(callCtx)
    
    // 3. Call the actual target
    result := p.target.Greet()
    
    // 4. Invoke 'Exit' interceptors
    p.InvokeExitMethodInterceptors(callCtx)
    
    return result
}
```

## How It Works

The `ProxyBase` provides helper methods to trigger interceptors:

* `InvokeEnterMethodInterceptors`: Calls `Enter` on all registered interceptors.
* `InvokeExitMethodInterceptors`: Calls `Exit` on all registered interceptors.
* `InvokeMethodErrorInterceptors`: Calls `OnError` on all registered interceptors if an error is detected in the return values.

By manually implementing these calls, you have full control over which methods are intercepted and how the call context is constructed.
