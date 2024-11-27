---
meta:
  title: Parsley Docs - "resolving" package
description: Official documentation of the Parsley´s "resolving" package
icon: package
label: resolving
tags: [ "packages", "resolving" ]
category:
  - Package Reference
---
# resolving

```go
import "github.com/matzefriedrich/parsley/pkg/resolving"
```

## Index

- [func Activate\[T any\]\(resolver types.Resolver, ctx context.Context, activatorFunc any, options ...types.ResolverOptionsFunc\) \(T, error\)](<#Activate>)
- [func NewResolver\(registry types.ServiceRegistry\) types.Resolver](<#NewResolver>)
- [func NewScopedContext\(ctx context.Context\) context.Context](<#NewScopedContext>)
- [func ResolveRequiredService\[T any\]\(resolver types.Resolver, ctx context.Context\) \(T, error\)](<#ResolveRequiredService>)
- [func ResolveRequiredServices\[T any\]\(resolver types.Resolver, ctx context.Context\) \(\[\]T, error\)](<#ResolveRequiredServices>)
- [func WithInstance\[T any\]\(instance T\) types.ResolverOptionsFunc](<#WithInstance>)
- [type NamedServiceResolverActivatorFunc](<#NamedServiceResolverActivatorFunc>)
  - [func CreateNamedServiceResolverActivatorFunc\[T any\]\(\) NamedServiceResolverActivatorFunc\[T\]](<#CreateNamedServiceResolverActivatorFunc>)


<a name="Activate"></a>
## func [Activate](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/activate.go#L11>)

```go
func Activate[T any](resolver types.Resolver, ctx context.Context, activatorFunc any, options ...types.ResolverOptionsFunc) (T, error)
```

Activate attempts to create and return an instance of the requested type using the provided resolver. Use this method to instantiate service objects of unregistered types. The specified activator function can have parameters to demand service instances for registered service types.

<a name="NewResolver"></a>
## func [NewResolver](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L58>)

```go
func NewResolver(registry types.ServiceRegistry) types.Resolver
```

NewResolver creates and returns a new Resolver instance based on the provided ServiceRegistry.

<a name="NewScopedContext"></a>
## func [NewScopedContext](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/scope.go#L9>)

```go
func NewScopedContext(ctx context.Context) context.Context
```

NewScopedContext creates a new context with an associated service instance map, useful for managing service lifetimes within scope.

<a name="ResolveRequiredService"></a>
## func [ResolveRequiredService](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L43>)

```go
func ResolveRequiredService[T any](resolver types.Resolver, ctx context.Context) (T, error)
```

ResolveRequiredService resolves a single service instance of the specified type using the given resolver and context. The method can return the following errors: ErrorCannotResolveService, ErrorAmbiguousServiceInstancesResolved.

<a name="ResolveRequiredServices"></a>
## func [ResolveRequiredServices](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L18>)

```go
func ResolveRequiredServices[T any](resolver types.Resolver, ctx context.Context) ([]T, error)
```

ResolveRequiredServices resolves all registered services of a specified type T using the given resolver and context.

<a name="WithInstance"></a>
## func [WithInstance](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver_options.go#L19>)

```go
func WithInstance[T any](instance T) types.ResolverOptionsFunc
```

WithInstance Creates a ResolverOptionsFunc that registers a specific instance of a type T with a service registry to be resolved as a singleton.

<a name="NamedServiceResolverActivatorFunc"></a>
## type [NamedServiceResolverActivatorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/named_service_resolver.go#L9>)

NamedServiceResolverActivatorFunc defines a function for resolving named services.

```go
type NamedServiceResolverActivatorFunc[T any] func(types.Resolver) func(string) (T, error)
```

<a name="CreateNamedServiceResolverActivatorFunc"></a>
### func [CreateNamedServiceResolverActivatorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/named_service_resolver.go#L12>)

```go
func CreateNamedServiceResolverActivatorFunc[T any]() NamedServiceResolverActivatorFunc[T]
```

CreateNamedServiceResolverActivatorFunc creates a NamedServiceResolverActivatorFunc for resolving named services.

