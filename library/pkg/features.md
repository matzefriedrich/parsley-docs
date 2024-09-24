---
icon: package
label: features
tags: [ "packages", "features" ]
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
## func RegisterLazy

```go
func RegisterLazy[T any](registry types.ServiceRegistry, activatorFunc func() T, _ types.LifetimeScope) error
```

RegisterLazy registers a lazily\-activated service in the service registry using the provided activator function.

<a name="RegisterList"></a>
## func RegisterList

```go
func RegisterList[T any](registry types.ServiceRegistry) error
```

RegisterList registers a function that resolves and returns a list of services of type T with the specified registry.

<a name="RegisterNamed"></a>
## func RegisterNamed

```go
func RegisterNamed[T any](registry types.ServiceRegistry, services ...registration.NamedServiceRegistrationFunc) error
```

RegisterNamed registers named services with their respective activator functions and lifetime scopes. It supports dependency injection by associating names with service instances.

<a name="ArgMatch"></a>
## type ArgMatch

ArgMatch is a function type used to match an argument against a certain condition during mock function verification.

```go
type ArgMatch func(actual any) bool
```

<a name="Exact"></a>
### func Exact

```go
func Exact[T comparable](expected T) ArgMatch
```

Exact returns an ArgMatch that checks if a given argument is exactly equal to the specified expected value.

<a name="IsAny"></a>
### func IsAny

```go
func IsAny() ArgMatch
```

IsAny always returns true, enabling it to match any given argument during mock function verification.

<a name="Interceptor"></a>
## type Interceptor

Interceptor is a base interface type for defining interceptors that can be used to monitor or alter the behavior of other components.

```go
type Interceptor interface {
    Name() string
    Position() int
}
```

<a name="InterceptorBase"></a>
## type InterceptorBase

InterceptorBase serves as a foundational structure for defining interceptors, managing essential data like name and position.

```go
type InterceptorBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewInterceptorBase"></a>
### func NewInterceptorBase

```go
func NewInterceptorBase(name string, position int) InterceptorBase
```

NewInterceptorBase creates a new instance of InterceptorBase with the specified name and position for managing interceptor metadata.

<a name="InterceptorBase.Name"></a>
### func \(InterceptorBase\) Name

```go
func (i InterceptorBase) Name() string
```

Name retrieves the name of the interceptor, which is useful for identification and debugging purposes.

<a name="InterceptorBase.Position"></a>
### func \(InterceptorBase\) Position

```go
func (i InterceptorBase) Position() int
```

Position returns the position of the interceptor, helping determine its order in processing flows within a system.

<a name="Lazy"></a>
## type Lazy

Lazy represents a type whose value is initialized lazily upon first access, typically to improve performance or manage resources.

```go
type Lazy[T any] interface {
    Value() T
}
```

<a name="MethodCallContext"></a>
## type MethodCallContext

MethodCallContext captures the context of a method call, including method name, parameters, and return values.

```go
type MethodCallContext struct {
    // contains filtered or unexported fields
}
```

<a name="NewMethodCallContext"></a>
### func NewMethodCallContext

```go
func NewMethodCallContext(methodName string, parameters map[string]interface{}) *MethodCallContext
```

NewMethodCallContext creates a new MethodCallContext instance with the provided method name and parameters.

<a name="MethodInterceptor"></a>
## type MethodInterceptor

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
## type MockBase

MockBase is used as a foundational struct to track and manage mocked functions and their call history. It helps in testing by allowing function signature tracking and call verification.

```go
type MockBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewMockBase"></a>
### func NewMockBase

```go
func NewMockBase() MockBase
```

NewMockBase initializes and returns an instance of MockBase, ideal for setting up and using mock functions in tests.

<a name="MockBase.AddFunction"></a>
### func \(\*MockBase\) AddFunction

```go
func (m *MockBase) AddFunction(name string, signature string)
```

AddFunction adds a new mock function with the specified name and signature to the MockBase instance.

<a name="MockBase.TraceMethodCall"></a>
### func \(\*MockBase\) TraceMethodCall

```go
func (m *MockBase) TraceMethodCall(name string, arguments ...any)
```

TraceMethodCall logs the invocation of a mocked function with specified arguments to facilitate function call tracking during testing. Before function calls can be tracked, the function must be registered with the MockBase instance; use AddFunction.

<a name="MockBase.Verify"></a>
### func \(\*MockBase\) Verify

```go
func (m *MockBase) Verify(name string, times TimesFunc, matches ...ArgMatch) bool
```

Verify checks if a mock function was called a specific number of times, optionally matching provided argument conditions.

<a name="MockFunction"></a>
## type MockFunction

MockFunction provides a structure to represent a mocked function in test scenarios. It allows tracking its calls and signature.

```go
type MockFunction struct {
    // contains filtered or unexported fields
}
```

<a name="MockFunction.String"></a>
### func \(MockFunction\) String

```go
func (m MockFunction) String() string
```

String returns the signature of the mocked function if it exists, otherwise it returns the function's name.

<a name="ParameterInfo"></a>
## type ParameterInfo

ParameterInfo represents information about a method parameter, including its value, type, and name. It is used in method interception where parameters need to be inspected or logged.

```go
type ParameterInfo struct {
    // contains filtered or unexported fields
}
```

<a name="ParameterInfo.String"></a>
### func \(ParameterInfo\) String

```go
func (p ParameterInfo) String() string
```

String returns a formatted string representation of the ParameterInfo, useful for logging and debugging purposes.

<a name="ProxyBase"></a>
## type ProxyBase

ProxyBase facilitates method interception by allowing the inclusion of multiple interceptors to target method calls. Typically used to monitor, log, or modify behavior of an object's method execution.

```go
type ProxyBase struct {
    // contains filtered or unexported fields
}
```

<a name="NewProxyBase"></a>
### func NewProxyBase

```go
func NewProxyBase[T any](target T, interceptors []MethodInterceptor) ProxyBase
```

NewProxyBase creates a ProxyBase instance with the provided target and a sorted list of method interceptors. Useful for setting up method interception on the target object.

<a name="ProxyBase.InvokeEnterMethodInterceptors"></a>
### func \(\*ProxyBase\) InvokeEnterMethodInterceptors

```go
func (p *ProxyBase) InvokeEnterMethodInterceptors(callContext *MethodCallContext)
```

InvokeEnterMethodInterceptors triggers the Enter method on all registered interceptors before the target method executes.

<a name="ProxyBase.InvokeExitMethodInterceptors"></a>
### func \(\*ProxyBase\) InvokeExitMethodInterceptors

```go
func (p *ProxyBase) InvokeExitMethodInterceptors(callContext *MethodCallContext)
```

InvokeExitMethodInterceptors triggers the Exit method of all registered interceptors after the target method completes.

<a name="ProxyBase.InvokeMethodErrorInterceptors"></a>
### func \(\*ProxyBase\) InvokeMethodErrorInterceptors

```go
func (p *ProxyBase) InvokeMethodErrorInterceptors(callContext *MethodCallContext, returnValues ...interface{})
```

InvokeMethodErrorInterceptors intercepts the return values of a method, checks for errors, and triggers OnError for registered interceptors.

<a name="ReturnValueInfo"></a>
## type ReturnValueInfo

ReturnValueInfo represents the value and type information of a method's return value, used in method interception.

```go
type ReturnValueInfo struct {
    // contains filtered or unexported fields
}
```

<a name="ReturnValueInfo.String"></a>
### func \(ReturnValueInfo\) String

```go
func (r ReturnValueInfo) String() string
```

String returns a string representation of ReturnValueInfo, formatting the value and its type for debugging purposes.

<a name="TimesFunc"></a>
## type TimesFunc

TimesFunc is used to verify the number of times a mock function is called. It allows flexibility in call count assertions.

```go
type TimesFunc func(times int) bool
```

<a name="TimesAtLeastOnce"></a>
### func TimesAtLeastOnce

```go
func TimesAtLeastOnce() TimesFunc
```

TimesAtLeastOnce returns a TimesFunc that verifies if a mock function is called at least once.

<a name="TimesExactly"></a>
### func TimesExactly

```go
func TimesExactly(n int) TimesFunc
```

TimesExactly returns a TimesFunc that checks if the number of function calls is exactly equal to the specified value.

<a name="TimesNever"></a>
### func TimesNever

```go
func TimesNever() TimesFunc
```

TimesNever returns a TimesFunc that ensures the function has never been called, providing a strict zero call condition.

<a name="TimesOnce"></a>
### func TimesOnce

```go
func TimesOnce() TimesFunc
```

TimesOnce returns a TimesFunc that checks if the number of function calls equals one. It is useful for verifying single call assertions.

