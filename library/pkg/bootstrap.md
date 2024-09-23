---
icon: package
label: bootstrap
tags: [ "packages", "bootstrap" ]
---
# bootstrap

```go
import "github.com/matzefriedrich/parsley/pkg/bootstrap"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func RunParsleyApplication\(cxt context.Context, appFactoryFunc any, configure ...types.ModuleFunc\) error](<#RunParsleyApplication>)
- [type Application](<#Application>)


## Constants

<a name="ErrorCannotRegisterAppFactory"></a>

```go
const (
    ErrorCannotRegisterAppFactory = "cannot register application factory"
)
```

## Variables

<a name="ErrCannotRegisterAppFactory"></a>

```go
var (
    ErrCannotRegisterAppFactory = errors.New(ErrorCannotRegisterAppFactory)
)
```

<a name="RunParsleyApplication"></a>
## func RunParsleyApplication

```go
func RunParsleyApplication(cxt context.Context, appFactoryFunc any, configure ...types.ModuleFunc) error
```



<a name="Application"></a>
## type Application



```go
type Application interface {
    Run(context.Context) error
}
```

