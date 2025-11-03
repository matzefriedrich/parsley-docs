---
meta:
  title: Parsley Docs - "resolving" package
description: Official documentation of the Parsley´s "resolving" package
icon: '<svg viewBox="0 0 24 24"><path d="M10.68,10.16c-.28.08-.57.16-.9.25-.16.05-.2.06-.35-.13-.18-.22-.31-.36-.57-.48-.76-.39-1.5-.28-2.18.19-.82.56-1.24,1.38-1.23,2.41.01,1.02.68,1.86,1.63,1.99.82.11,1.51-.19,2.05-.84.11-.14.21-.29.33-.47h-2.33c-.25,0-.31-.17-.23-.38.16-.39.45-1.05.62-1.38.04-.08.12-.2.3-.2h3.88c.17-.58.46-1.13.83-1.65.88-1.22,1.94-1.86,3.38-2.12,1.23-.23,2.39-.1,3.44.65.95.69,1.54,1.61,1.7,2.83.21,1.72-.27,3.11-1.39,4.31-.8.85-1.77,1.39-2.89,1.63-.21.04-.43.06-.64.08-.11.01-.22.02-.33.03-1.1-.03-2.1-.36-2.94-1.12-.59-.54-1-1.21-1.21-1.98-.14.3-.31.59-.51.86-.87,1.21-2,1.96-3.44,2.16-1.18.17-2.28-.08-3.24-.84-.89-.71-1.4-1.65-1.53-2.82-.16-1.39.23-2.63,1.03-3.72.86-1.18,1.99-1.93,3.38-2.2,1.13-.22,2.22-.08,3.2.62.64.44,1.1,1.05,1.4,1.79.07.11.02.18-.12.22-.42.11-.77.21-1.13.31h-.01ZM18.66,11.61v.14c-.06,1.09-.58,1.91-1.53,2.43-.64.34-1.3.38-1.97.08-.87-.41-1.33-1.41-1.11-2.4.27-1.19.99-1.94,2.11-2.21,1.15-.28,2.24.43,2.46,1.69.02.09.02.18.03.28h.01Z" style="fill-rule: evenodd;"/></svg>'
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

- [func Activate\[T any\]\(ctx context.Context, resolver types.Resolver, activatorFunc any, options ...types.ResolverOptionsFunc\) \(T, error\)](<#Activate>)
- [func NewResolver\(registry types.ServiceRegistry\) types.Resolver](<#NewResolver>)
- [func NewScopedContext\(ctx context.Context\) context.Context](<#NewScopedContext>)
- [func ResolveRequiredService\[T any\]\(ctx context.Context, resolver types.Resolver\) \(T, error\)](<#ResolveRequiredService>)
- [func ResolveRequiredServices\[T any\]\(ctx context.Context, resolver types.Resolver\) \(\[\]T, error\)](<#ResolveRequiredServices>)
- [func WithInstance\[T any\]\(instance T\) types.ResolverOptionsFunc](<#WithInstance>)
- [type NamedServiceResolverActivatorFunc](<#NamedServiceResolverActivatorFunc>)
  - [func CreateNamedServiceResolverActivatorFunc\[T any\]\(ctx context.Context\) NamedServiceResolverActivatorFunc\[T\]](<#CreateNamedServiceResolverActivatorFunc>)


<a name="Activate"></a>
## func [Activate](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/activate.go#L12>)

```go
func Activate[T any](ctx context.Context, resolver types.Resolver, activatorFunc any, options ...types.ResolverOptionsFunc) (T, error)
```

Activate attempts to create and return an instance of the requested type using the provided resolver. Use this method to instantiate service objects of unregistered types. The specified activator function can have parameters to demand service instances for registered service types.

<a name="NewResolver"></a>
## func [NewResolver](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L59>)

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
## func [ResolveRequiredService](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L44>)

```go
func ResolveRequiredService[T any](ctx context.Context, resolver types.Resolver) (T, error)
```

ResolveRequiredService resolves a single service instance of the specified type using the given resolver and context. The method can return the following errors: ErrorCannotResolveService, ErrorAmbiguousServiceInstancesResolved.

<a name="ResolveRequiredServices"></a>
## func [ResolveRequiredServices](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver.go#L19>)

```go
func ResolveRequiredServices[T any](ctx context.Context, resolver types.Resolver) ([]T, error)
```

ResolveRequiredServices resolves all registered services of a specified type T using the given resolver and context.

<a name="WithInstance"></a>
## func [WithInstance](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/resolver_options.go#L19>)

```go
func WithInstance[T any](instance T) types.ResolverOptionsFunc
```

WithInstance Creates a ResolverOptionsFunc that registers a specific instance of a type T with a service registry to be resolved as a singleton.

<a name="NamedServiceResolverActivatorFunc"></a>
## type [NamedServiceResolverActivatorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/named_service_resolver.go#L10>)

NamedServiceResolverActivatorFunc defines a function for resolving named services.

```go
type NamedServiceResolverActivatorFunc[T any] func(types.Resolver) func(string) (T, error)
```

<a name="CreateNamedServiceResolverActivatorFunc"></a>
### func [CreateNamedServiceResolverActivatorFunc](<https://github.com/matzefriedrich/parsley/blob/main/pkg/resolving/named_service_resolver.go#L13>)

```go
func CreateNamedServiceResolverActivatorFunc[T any](ctx context.Context) NamedServiceResolverActivatorFunc[T]
```

CreateNamedServiceResolverActivatorFunc creates a NamedServiceResolverActivatorFunc for resolving named services.

