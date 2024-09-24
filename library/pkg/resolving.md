---
icon: package
label: resolving
tags: [ "packages", "resolving" ]
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
## func Activate

```go
func Activate[T any](resolver types.Resolver, ctx context.Context, activatorFunc any, options ...types.ResolverOptionsFunc) (T, error)
```

Activate attempts to create and return an instance of the requested type using the provided resolver. Use this method to instantiate service objects of unregistered types. The specified activator function can have parameters to demand service instances for registered service types.

<a name="NewResolver"></a>
## func NewResolver

```go
func NewResolver(registry types.ServiceRegistry) types.Resolver
```

NewResolver creates and returns a new Resolver instance based on the provided ServiceRegistry.

<a name="NewScopedContext"></a>
## func NewScopedContext

```go
func NewScopedContext(ctx context.Context) context.Context
```

NewScopedContext creates a new context with an associated service instance map, useful for managing service lifetimes within scope.

<a name="ResolveRequiredService"></a>
## func ResolveRequiredService

```go
func ResolveRequiredService[T any](resolver types.Resolver, ctx context.Context) (T, error)
```

ResolveRequiredService resolves a single service instance of the specified type using the given resolver and context. The method can return the following errors: ErrorCannotResolveService, ErrorAmbiguousServiceInstancesResolved.

<a name="ResolveRequiredServices"></a>
## func ResolveRequiredServices

```go
func ResolveRequiredServices[T any](resolver types.Resolver, ctx context.Context) ([]T, error)
```

ResolveRequiredServices resolves all registered services of a specified type T using the given resolver and context.

<a name="WithInstance"></a>
## func WithInstance

```go
func WithInstance[T any](instance T) types.ResolverOptionsFunc
```

WithInstance Creates a ResolverOptionsFunc that registers a specific instance of a type T with a service registry to be resolved as a singleton.

<a name="NamedServiceResolverActivatorFunc"></a>
## type NamedServiceResolverActivatorFunc

NamedServiceResolverActivatorFunc defines a function for resolving named services.

```go
type NamedServiceResolverActivatorFunc[T any] func(types.Resolver) func(string) (T, error)
```

<a name="CreateNamedServiceResolverActivatorFunc"></a>
### func CreateNamedServiceResolverActivatorFunc

```go
func CreateNamedServiceResolverActivatorFunc[T any]() NamedServiceResolverActivatorFunc[T]
```

CreateNamedServiceResolverActivatorFunc creates a NamedServiceResolverActivatorFunc for resolving named services.

