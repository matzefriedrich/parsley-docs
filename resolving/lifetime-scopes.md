---
icon: file
label:  Lifetime Scopes
order: 100
tags: [ registration, resolving, instances, lifetime scope ]
---
# Lifetime Scopes in Parsley

Parsley allows you to control the lifetime of your service instances through different lifetime scopes. The lifetime setting determines how often the constructor or factory method of a service registration is called and how instances are managed. 

## Supported service lifetimes

The following lifetime scopes are supported:

| Name      | Value                                                                                                           |
|-----------|:----------------------------------------------------------------------------------------------------------------|
| Transient | A new instance is created every time the service is requested.                                                  |
| Scoped    | The same instance is reused within a scope.                                                                     |
| Singleton | The same instance is reused for all requests. Resolved instances remain valid for the lifetime of the resolver. |

> **Note:** Multiple registrations for the same service type can use different lifetime scopes. The same is true for named service registrations.

## Convenience Functions

Parsley provides several **convenience functions** to simplify the registration of services with different lifetimes. These functions allow you to register services with minimal boilerplate and ensure clarity in service management:

- **RegisterSingleton:** Registers a service that will have a single instance throughout the application’s lifetime.
- **RegisterScoped:** Registers services with a new instance per scope (such as a request or session).
- **RegisterTransient:** Registers services that will have a new instance every time they are requested.

Each of these functions accept **multiple activator functions**, allowing multiple services to be registered in one call.

## Example

The following demonstration is based on the greeter example code:

:::code language="golang" source="/examples/resolving-services/internal/greeter.go" range="5-25" :::

In this example, a factory method (instead of a constructor method) is used to intercept the creation of a service instance and trace an event each time Parsley activates a new `Greeter` service instance. The `traceResolveEventFor` method uses reflection to determine the type of the resolved service and traces the type name and pointer to the standard output.

```go
func resolveGreeter(salutation string) func(resolver types.Resolver) internal.Greeter {
	factory := internal.NewGreeterFactory(salutation)
	return func(resolver types.Resolver) internal.Greeter {
		greeter := factory()
		traceResolveEventFor(greeter)
		return greeter
	}
}

func traceResolveEventFor(service any) {
	value := reflect.ValueOf(service)
	instancePointer := value.Pointer()
	typeName := value.Elem().Type().String()
	fmt.Printf("New %s instance created: %d\n", typeName, instancePointer)
}
```

The `resolveGreeter` factory method is registered with Parsley. The lifetime for `Greeter` instances is set to `LifetimeScoped`, instructing the resolver to keep track of instances in the `Context` given when resolving the service.

```go
registry := registration.NewServiceRegistry()
err := registry.Register(resolveGreeter("Hi"), types.LifetimeScoped)
if err != nil {
    panic(err)
}
```

The following code attempts to resolve `Greeter` service instances repeatedly. Since the service factory is assigned with the scoped lifetime behavior and the same context is shared with all calls to the `ResolveRequiredService` method, the factory function is expected to be called only once.

```go
ctx := resolving.NewScopedContext(context.Background())
resolver := resolving.NewResolver(registry)

for i := 0; i < 3; i++ {
    greeter, _ := resolving.ResolveRequiredService[internal.Greeter](resolver, ctx)
    greeter.SayHello("John", false)
}
```

The example produces the following output:

```sh
New internal.greeter instance created: 824633868736
Hi, John
Hi, John
Hi, John
```

If a new context is passed to each call of the `ResolveRequiredService` method, a different behavior can be observed.

```go
resolver := resolving.NewResolver(registry)

for i := 0; i < 3; i++ {
    ctx := resolving.NewScopedContext(context.Background())
    greeter, _ := resolving.ResolveRequiredService[internal.Greeter](resolver, ctx)
    greeter.SayHello("John", false)
}
```

Now, for each iteration, a new context is created, thus requiring Parsley to activate a new service instance. The example produces the following output:

```sh
New internal.greeter instance created: 824633868736
Hi, John
New internal.greeter instance created: 824633868960
Hi, John
New internal.greeter instance created: 824633869200
Hi, John
```

If you change the example once again, setting the lifetime behavior for the `Greeter` service to `LifetimeSingleton` at registration, ...

```go
registry.Register(resolveGreeter("Hi"), types.LifetimeSingleton)
```

... Parsley does not store created instances in the given `Context` but in the instance cache attached to the resolver itself, resulting in the following output:

```sh
New internal.greeter instance created: 824633868736
Hi, John
Hi, John
Hi, John
```

Understanding and utilizing lifetime scopes in Parsley allows you to manage service instances effectively. Adjust the lifetime settings based on your application's requirements to optimize performance and resource usage.