---
meta:
  title: Parsley Docs - "features" package
description: Official documentation of the Parsley´s "features" package
icon: '<svg viewBox="0 0 24 24"><path d="M10.68,10.16c-.28.08-.57.16-.9.25-.16.05-.2.06-.35-.13-.18-.22-.31-.36-.57-.48-.76-.39-1.5-.28-2.18.19-.82.56-1.24,1.38-1.23,2.41.01,1.02.68,1.86,1.63,1.99.82.11,1.51-.19,2.05-.84.11-.14.21-.29.33-.47h-2.33c-.25,0-.31-.17-.23-.38.16-.39.45-1.05.62-1.38.04-.08.12-.2.3-.2h3.88c.17-.58.46-1.13.83-1.65.88-1.22,1.94-1.86,3.38-2.12,1.23-.23,2.39-.1,3.44.65.95.69,1.54,1.61,1.7,2.83.21,1.72-.27,3.11-1.39,4.31-.8.85-1.77,1.39-2.89,1.63-.21.04-.43.06-.64.08-.11.01-.22.02-.33.03-1.1-.03-2.1-.36-2.94-1.12-.59-.54-1-1.21-1.21-1.98-.14.3-.31.59-.51.86-.87,1.21-2,1.96-3.44,2.16-1.18.17-2.28-.08-3.24-.84-.89-.71-1.4-1.65-1.53-2.82-.16-1.39.23-2.63,1.03-3.72.86-1.18,1.99-1.93,3.38-2.2,1.13-.22,2.22-.08,3.2.62.64.44,1.1,1.05,1.4,1.79.07.11.02.18-.12.22-.42.11-.77.21-1.13.31h-.01ZM18.66,11.61v.14c-.06,1.09-.58,1.91-1.53,2.43-.64.34-1.3.38-1.97.08-.87-.41-1.33-1.41-1.11-2.4.27-1.19.99-1.94,2.11-2.21,1.15-.28,2.24.43,2.46,1.69.02.09.02.18.03.28h.01Z" style="fill-rule: evenodd;"/></svg>'
label: features
tags: [ "packages", "features" ]
category:
  - Package Reference
---
# features

```go
import "github.com/matzefriedrich/parsley/pkg/features"
```

## Index

- [func RegisterLazy\[T any\]\(registry types.ServiceRegistry, activatorFunc func\(\) T, \_ types.LifetimeScope\) error](<#RegisterLazy>)
- [func RegisterList\[T any\]\(registry types.ServiceRegistry\) error](<#RegisterList>)
- [func RegisterNamed\[T any\]\(registry types.ServiceRegistry, services ...registration.NamedServiceRegistrationFunc\) error](<#RegisterNamed>)
- [type ArgMatch](<#ArgMatch>)
  - [func Exact\[T comparable\]\(expected T\) ArgMatch](<#Exact>)
  - [func IsAny\(\) ArgMatch](<#IsAny>)
- [type Interceptor](<#Interceptor>)
- [type InterceptorBase](<#InterceptorBase>)
  - [func NewInterceptorBase\(name string, position int\) InterceptorBase](<#NewInterceptorBase>)
  - [func \(i InterceptorBase\) Name\(\) string](<#InterceptorBase.Name>)
  - [func \(i InterceptorBase\) Position\(\) int](<#InterceptorBase.Position>)
- [type Lazy](<#Lazy>)
- [type MethodCallContext](<#MethodCallContext>)
  - [func NewMethodCallContext\(methodName string, parameters map\[string\]interface\{\}\) \*MethodCallContext](<#NewMethodCallContext>)
- [type MethodInterceptor](<#MethodInterceptor>)
- [type MockBase](<#MockBase>)
  - [func NewMockBase\(\) MockBase](<#NewMockBase>)
  - [func \(m \*MockBase\) AddFunction\(name string, signature string\)](<#MockBase.AddFunction>)
  - [func \(m \*MockBase\) TraceMethodCall\(name string, arguments ...any\)](<#MockBase.TraceMethodCall>)
  - [func \(m \*MockBase\) Verify\(name string, times TimesFunc, matches ...ArgMatch\) bool](<#MockBase.Verify>)
- [type MockFunction](<#MockFunction>)
  - [func \(m MockFunction\) String\(\) string](<#MockFunction.String>)
- [type ParameterInfo](<#ParameterInfo>)
  - [func \(p ParameterInfo\) String\(\) string](<#ParameterInfo.String>)
- [type ProxyBase](<#ProxyBase>)
  - [func NewProxyBase\[T any\]\(target T, interceptors \[\]MethodInterceptor\) ProxyBase](<#NewProxyBase>)
  - [func \(p \*ProxyBase\) InvokeEnterMethodInterceptors\(callContext \*MethodCallContext\)](<#ProxyBase.InvokeEnterMethodInterceptors>)
  - [func \(p \*ProxyBase\) InvokeExitMethodInterceptors\(callContext \*MethodCallContext\)](<#ProxyBase.InvokeExitMethodInterceptors>)
  - [func \(p \*ProxyBase\) InvokeMethodErrorInterceptors\(callContext \*MethodCallContext, returnValues ...interface\{\}\)](<#ProxyBase.InvokeMethodErrorInterceptors>)
- [type ReturnValueInfo](<#ReturnValueInfo>)
  - [func \(r ReturnValueInfo\) String\(\) string](<#ReturnValueInfo.String>)
- [type TimesFunc](<#TimesFunc>)
  - [func TimesAtLeastOnce\(\) TimesFunc](<#TimesAtLeastOnce>)
  - [func TimesExactly\(n int\) TimesFunc](<#TimesExactly>)
  - [func TimesNever\(\) TimesFunc](<#TimesNever>)
  - [func TimesOnce\(\) TimesFunc](<#TimesOnce>)


<a name="RegisterLazy"></a>
## func [RegisterLazy](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/lazy_services.go#L41>)

```go
func RegisterLazy[T any](registry types.ServiceRegistry, activatorFunc func() T, _ types.LifetimeScope) error
```

RegisterLazy registers a lazily\-activated service in the service registry using the provided activator function.

<a name="RegisterList"></a>
## func [RegisterList](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/list_services.go#L10>)

```go
func RegisterList[T any](registry types.ServiceRegistry) error
```

RegisterList registers a function that resolves and returns a list of services of type T with the specified registry.

<a name="RegisterNamed"></a>
## func [RegisterNamed](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/named_services.go#L26>)

```go
func RegisterNamed[T any](registry types.ServiceRegistry, services ...registration.NamedServiceRegistrationFunc) error
```

RegisterNamed registers named services with their respective activator functions and lifetime scopes. It supports dependency injection by associating names with service instances.

<a name="ArgMatch"></a>
## type [ArgMatch](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L4>)

ArgMatch is a function type used to match an argument against a certain condition during mock function verification.

```go
type ArgMatch func(actual any) bool
```

<a name="Exact"></a>
### func [Exact](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L14>)

```go
func Exact[T comparable](expected T) ArgMatch
```

Exact returns an ArgMatch that checks if a given argument is exactly equal to the specified expected value.

<a name="IsAny"></a>
### func [IsAny](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L7>)

```go
func IsAny() ArgMatch
```

IsAny always returns true, enabling it to match any given argument during mock function verification.

<a name="Interceptor"></a>
## type [Interceptor](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L10-L13>)

Interceptor is a base interface type for defining interceptors that can be used to monitor or alter the behavior of other components.

```go
type Interceptor interface {
    Name() string
    Position() int
}
```

<a name="InterceptorBase"></a>
## type [InterceptorBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L25-L28>)

InterceptorBase serves as a foundational structure for defining interceptors, managing essential data like name and position.

```go
type InterceptorBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewInterceptorBase"></a>
### func [NewInterceptorBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L41>)

```go
func NewInterceptorBase(name string, position int) InterceptorBase
```

NewInterceptorBase creates a new instance of InterceptorBase with the specified name and position for managing interceptor metadata.

<a name="InterceptorBase.Name"></a>
### func \(InterceptorBase\) [Name](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L31>)

```go
func (i InterceptorBase) Name() string
```

Name retrieves the name of the interceptor, which is useful for identification and debugging purposes.

<a name="InterceptorBase.Position"></a>
### func \(InterceptorBase\) [Position](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L36>)

```go
func (i InterceptorBase) Position() int
```

Position returns the position of the interceptor, helping determine its order in processing flows within a system.

<a name="Lazy"></a>
## type [Lazy](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/lazy_services.go#L34-L36>)

Lazy represents a type whose value is initialized lazily upon first access, typically to improve performance or manage resources.

```go
type Lazy[T any] interface {
    Value() T
}
```

<a name="MethodCallContext"></a>
## type [MethodCallContext](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L49-L53>)

MethodCallContext captures the context of a method call, including method name, parameters, and return values.

```go
type MethodCallContext struct {
    // contains filtered or unexported fields
}
```

<a name="NewMethodCallContext"></a>
### func [NewMethodCallContext](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L83>)

```go
func NewMethodCallContext(methodName string, parameters map[string]interface{}) *MethodCallContext
```

NewMethodCallContext creates a new MethodCallContext instance with the provided method name and parameters.

<a name="MethodInterceptor"></a>
## type [MethodInterceptor](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L17-L22>)

MethodInterceptor provides hooks to intercept method execution on a proxy object. It allows entering before method invocations, exiting after method executions, and handling errors during method execution for monitoring or altering behavior.

```go
type MethodInterceptor interface {
    Interceptor
    Enter(target any, methodName string, parameters []ParameterInfo)
    Exit(target any, methodName string, returnValues []ReturnValueInfo)
    OnError(target any, methodName string, err error)
}
```

<a name="MockBase"></a>
## type [MockBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L5-L7>)

MockBase is used as a foundational struct to track and manage mocked functions and their call history. It helps in testing by allowing function signature tracking and call verification.

```go
type MockBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewMockBase"></a>
### func [NewMockBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L29>)

```go
func NewMockBase() MockBase
```

NewMockBase initializes and returns an instance of MockBase, ideal for setting up and using mock functions in tests.

<a name="MockBase.AddFunction"></a>
### func \(\*MockBase\) [AddFunction](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L36>)

```go
func (m *MockBase) AddFunction(name string, signature string)
```

AddFunction adds a new mock function with the specified name and signature to the MockBase instance.

<a name="MockBase.TraceMethodCall"></a>
### func \(\*MockBase\) [TraceMethodCall](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L45>)

```go
func (m *MockBase) TraceMethodCall(name string, arguments ...any)
```

TraceMethodCall logs the invocation of a mocked function with specified arguments to facilitate function call tracking during testing. Before function calls can be tracked, the function must be registered with the MockBase instance; use AddFunction.

<a name="MockBase.Verify"></a>
### func \(\*MockBase\) [Verify](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L56>)

```go
func (m *MockBase) Verify(name string, times TimesFunc, matches ...ArgMatch) bool
```

Verify checks if a mock function was called a specific number of times, optionally matching provided argument conditions.

<a name="MockFunction"></a>
## type [MockFunction](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L10-L14>)

MockFunction provides a structure to represent a mocked function in test scenarios. It allows tracking its calls and signature.

```go
type MockFunction struct {
    // contains filtered or unexported fields
}
```

<a name="MockFunction.String"></a>
### func \(MockFunction\) [String](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock.go#L17>)

```go
func (m MockFunction) String() string
```

String returns the signature of the mocked function if it exists, otherwise it returns the function's name.

<a name="ParameterInfo"></a>
## type [ParameterInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L57-L61>)

ParameterInfo represents information about a method parameter, including its value, type, and name. It is used in method interception where parameters need to be inspected or logged.

```go
type ParameterInfo struct {
    // contains filtered or unexported fields
}
```

<a name="ParameterInfo.String"></a>
### func \(ParameterInfo\) [String](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L64>)

```go
func (p ParameterInfo) String() string
```

String returns a formatted string representation of the ParameterInfo, useful for logging and debugging purposes.

<a name="ProxyBase"></a>
## type [ProxyBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L93-L96>)

ProxyBase facilitates method interception by allowing the inclusion of multiple interceptors to target method calls. Typically used to monitor, log, or modify behavior of an object's method execution.

```go
type ProxyBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewProxyBase"></a>
### func [NewProxyBase](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L146>)

```go
func NewProxyBase[T any](target T, interceptors []MethodInterceptor) ProxyBase
```

NewProxyBase creates a ProxyBase instance with the provided target and a sorted list of method interceptors. Useful for setting up method interception on the target object.

<a name="ProxyBase.InvokeEnterMethodInterceptors"></a>
### func \(\*ProxyBase\) [InvokeEnterMethodInterceptors](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L113>)

```go
func (p *ProxyBase) InvokeEnterMethodInterceptors(callContext *MethodCallContext)
```

InvokeEnterMethodInterceptors triggers the Enter method on all registered interceptors before the target method executes.

<a name="ProxyBase.InvokeExitMethodInterceptors"></a>
### func \(\*ProxyBase\) [InvokeExitMethodInterceptors](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L130>)

```go
func (p *ProxyBase) InvokeExitMethodInterceptors(callContext *MethodCallContext)
```

InvokeExitMethodInterceptors triggers the Exit method of all registered interceptors after the target method completes.

<a name="ProxyBase.InvokeMethodErrorInterceptors"></a>
### func \(\*ProxyBase\) [InvokeMethodErrorInterceptors](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L99>)

```go
func (p *ProxyBase) InvokeMethodErrorInterceptors(callContext *MethodCallContext, returnValues ...interface{})
```

InvokeMethodErrorInterceptors intercepts the return values of a method, checks for errors, and triggers OnError for registered interceptors.

<a name="ReturnValueInfo"></a>
## type [ReturnValueInfo](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L69-L72>)

ReturnValueInfo represents the value and type information of a method's return value, used in method interception.

```go
type ReturnValueInfo struct {
    // contains filtered or unexported fields
}
```

<a name="ReturnValueInfo.String"></a>
### func \(ReturnValueInfo\) [String](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/proxies.go#L75>)

```go
func (r ReturnValueInfo) String() string
```

String returns a string representation of ReturnValueInfo, formatting the value and its type for debugging purposes.

<a name="TimesFunc"></a>
## type [TimesFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L25>)

TimesFunc is used to verify the number of times a mock function is called. It allows flexibility in call count assertions.

```go
type TimesFunc func(times int) bool
```

<a name="TimesAtLeastOnce"></a>
### func [TimesAtLeastOnce](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L35>)

```go
func TimesAtLeastOnce() TimesFunc
```

TimesAtLeastOnce returns a TimesFunc that verifies if a mock function is called at least once.

<a name="TimesExactly"></a>
### func [TimesExactly](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L42>)

```go
func TimesExactly(n int) TimesFunc
```

TimesExactly returns a TimesFunc that checks if the number of function calls is exactly equal to the specified value.

<a name="TimesNever"></a>
### func [TimesNever](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L49>)

```go
func TimesNever() TimesFunc
```

TimesNever returns a TimesFunc that ensures the function has never been called, providing a strict zero call condition.

<a name="TimesOnce"></a>
### func [TimesOnce](<https://github.com/matzefriedrich/parsley/blob/main/pkg/features/mock_verify.go#L28>)

```go
func TimesOnce() TimesFunc
```

TimesOnce returns a TimesFunc that checks if the number of function calls equals one. It is useful for verifying single call assertions.

