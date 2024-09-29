---
meta:
  title: Parsley Docs - "bootstrap" package
description: Official documentation of the Parsley´s "bootstrap" package
icon: package
label: bootstrap
tags: [ "packages", "bootstrap" ]
category:
  - Package Reference
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

<a name="ErrCannotRegisterAppFactory"></a>ErrCannotRegisterAppFactory is returned when the application factory cannot be registered, indicating an issue with the bootstrap process.

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

RunParsleyApplication initializes and runs the Parsley application lifecycle. It registers the application factory, configures additional modules, resolves the main application instance, and invokes its Run method.

<a name="Application"></a>
## type Application

Application provides an abstract interface for creating and running an application. It primarily facilitates the use of dependency injection for resolving services and the managing application lifecycle.

```go
type Application interface {
    Run(context.Context) error
}
```

