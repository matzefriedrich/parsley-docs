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



<a name="RegisterList"></a>
## func RegisterList

```go
func RegisterList[T any](registry types.ServiceRegistry) error
```



<a name="RegisterNamed"></a>
## func RegisterNamed

```go
func RegisterNamed[T any](registry types.ServiceRegistry, services ...registration.NamedServiceRegistrationFunc) error
```



<a name="ArgMatch"></a>
## type ArgMatch



```go
type ArgMatch func(actual any) bool
```

<a name="Exact"></a>
### func Exact

```go
func Exact[T comparable](expected T) ArgMatch
```



<a name="IsAny"></a>
### func IsAny

```go
func IsAny() ArgMatch
```



<a name="Interceptor"></a>
## type Interceptor



```go
type Interceptor interface {
    Name() string
    Position() int
}
```

<a name="InterceptorBase"></a>
## type InterceptorBase



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



<a name="InterceptorBase.Name"></a>
### func \(InterceptorBase\) Name

```go
func (i InterceptorBase) Name() string
```



<a name="InterceptorBase.Position"></a>
### func \(InterceptorBase\) Position

```go
func (i InterceptorBase) Position() int
```



<a name="Lazy"></a>
## type Lazy



```go
type Lazy[T any] interface {
    Value() T
}
```

<a name="MethodCallContext"></a>
## type MethodCallContext



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



<a name="MethodInterceptor"></a>
## type MethodInterceptor



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



<a name="MockBase.AddFunction"></a>
### func \(\*MockBase\) AddFunction

```go
func (m *MockBase) AddFunction(name string, signature string)
```



<a name="MockBase.TraceMethodCall"></a>
### func \(\*MockBase\) TraceMethodCall

```go
func (m *MockBase) TraceMethodCall(name string, arguments ...any)
```



<a name="MockBase.Verify"></a>
### func \(\*MockBase\) Verify

```go
func (m *MockBase) Verify(name string, times TimesFunc, matches ...ArgMatch) bool
```



<a name="MockFunction"></a>
## type MockFunction



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



<a name="ParameterInfo"></a>
## type ParameterInfo



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



<a name="ProxyBase"></a>
## type ProxyBase



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



<a name="ProxyBase.InvokeEnterMethodInterceptors"></a>
### func \(\*ProxyBase\) InvokeEnterMethodInterceptors

```go
func (p *ProxyBase) InvokeEnterMethodInterceptors(callContext *MethodCallContext)
```



<a name="ProxyBase.InvokeExitMethodInterceptors"></a>
### func \(\*ProxyBase\) InvokeExitMethodInterceptors

```go
func (p *ProxyBase) InvokeExitMethodInterceptors(callContext *MethodCallContext)
```



<a name="ProxyBase.InvokeMethodErrorInterceptors"></a>
### func \(\*ProxyBase\) InvokeMethodErrorInterceptors

```go
func (p *ProxyBase) InvokeMethodErrorInterceptors(callContext *MethodCallContext, returnValues ...interface{})
```



<a name="ReturnValueInfo"></a>
## type ReturnValueInfo



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



<a name="TimesFunc"></a>
## type TimesFunc



```go
type TimesFunc func(times int) bool
```

<a name="TimesAtLeastOnce"></a>
### func TimesAtLeastOnce

```go
func TimesAtLeastOnce() TimesFunc
```



<a name="TimesExactly"></a>
### func TimesExactly

```go
func TimesExactly(n int) TimesFunc
```



<a name="TimesNever"></a>
### func TimesNever

```go
func TimesNever() TimesFunc
```



<a name="TimesOnce"></a>
### func TimesOnce

```go
func TimesOnce() TimesFunc
```



