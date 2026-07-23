---
meta:
  title: Parsley - Context-Aware Factory Functions
description: Register context-aware factory functions in Parsley to resolve services on demand using a specific context.
icon: file
label: Context-Aware Factory Functions
tags: [ registration, features, factory ]
order: 89
category:
  - Registration
---
# Context-Aware Factory Functions

Parsley provides a specialized way to handle on-demand service resolution through context-aware factory functions. By using the `RegisterFactory[T]` function from the `pkg/features` package, you can register a factory that returns a function capable of resolving a service of type `T` whenever it is called.

## The FactoryFunc Type

A context-aware factory function is represented by the `FactoryFunc[T]` type:

```golang
type FactoryFunc[T any] func(ctx context.Context) (T, error)
```

This function takes a `context.Context` and returns an instance of type `T` (or an error). When invoked, it uses the underlying Parsley resolver to find and return the requested service.

## Registering a Factory

To register a context-aware factory, use the `features.RegisterFactory[T]` function. This function registers the `FactoryFunc[T]` type in the provided service registry.

```golang
func RegisterFactory[T any](registry types.ServiceRegistry, scope types.LifetimeScope) error
```

### Example

The following example demonstrates how to register and use a context-aware factory for a `Greeter` service.

```golang
package main

import (
    "context"
    "fmt"
    "github.com/matzefriedrich/parsley/pkg/features"
    "github.com/matzefriedrich/parsley/pkg/registration"
    "github.com/matzefriedrich/parsley/pkg/resolving"
    "github.com/matzefriedrich/parsley/pkg/types"
)

type Greeter interface {
    Greet() string
}

type greeter struct{}

func (g *greeter) Greet() string {
    return "Hello, Parsley!"
}

func NewGreeter() Greeter {
    return &greeter{}
}

func main() {
    registry := registration.NewServiceRegistry()
    
    // Register the actual service
    _ = registration.RegisterSingleton(registry, NewGreeter)
    
    // Register a FactoryFunc for the Greeter service
    _ = features.RegisterFactory[Greeter](registry, types.LifetimeSingleton)

    resolver := resolving.NewResolver(registry)
    ctx := context.Background()

    // Resolve the factory function
    factory, _ := resolving.ResolveRequiredService[features.FactoryFunc[Greeter]](ctx, resolver)

    // Use the factory to resolve the service on demand
    svc, _ := factory(ctx)
    fmt.Println(svc.Greet())
}
```

## Use Cases

Context-aware factory functions are particularly useful when you need to:

* **Deferred Resolution:** Resolve a service only when a specific condition is met at runtime.
* **Contextual Resolution:** Pass a specific context to the resolution process, which is useful for tracing, logging, or accessing request-scoped values.
* **Circular Dependencies:** Break circular dependencies by resolving one of the services on demand rather than at construction time.
