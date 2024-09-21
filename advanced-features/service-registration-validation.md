---
icon: file
label: Validating Service-Registrations
tags: [ registration, validation ]
---
# Validating Service-Registrations

Parsley provides a `Validator` service to improve service registration and dependency management by detecting two key issues: missing dependencies and circular dependencies. These configuration issues can be challenging, but the `Validator` ensures your application's dependency graph is valid. It is an essential tool for enhancing the robustness of service registration and resolution, especially in applications with complex service interactions.

## What are missing dependencies?

A **missing dependency** occurs when a service relies on a [constructor function](../registration/register-constructor-functions.md) that requires a service (the dependency) that hasn't been registered. In such cases, Parsley will fail to resolve the necessary object at runtime, leading to potential crashes or undefined behavior. Missing dependencies are often overlooked during development, especially in large, complex applications where services are registered in multiple places or based on runtime conditions.

For example, if a service `A` depends on service `B` but `B` isn't registered, this will cause an error during resolution:

```go
type A struct {
    b *B
}

type B struct {
}

func NewA(b *B) *A {
    return &A{b: b}
}
```

With the validator, such omissions are detected early, and a meaningful error message is returned before the application proceeds further.

## What are circular dependencies?

A **circular dependency** happens when two or more services depend on each other, directly or indirectly, creating an endless loop of dependencies. Without proper validation, this can lead to a **stack overflow** error during resolution, causing the application to crash. Circular dependencies are particularly problematic because they can easily occur unintentionally, especially in large systems where many services interact.

For instance, if service `A` depends on service `B`, and service `B` depends on service `A`, neither can be resolved because each requires the other to be instantiated first:

```go
type A struct {
    b *B
}

type B struct {
    a *A
}

func NewA(b *B) *A {
    return &A{b: b}
}

func NewB(a *A) *B {
    return &B{a: a}
}
```

Without a validator, such issues may be difficult to diagnose. Parsley's built-in **circular dependency check** ensures that the problem is caught early, preventing the application from entering a non-operable state. Although the resolver module itself checks for circular dependencies at runtime, the `Validator` services offers better error messaging and earlier detection.

### Example Usage

To validate your service registrations, use the following code:

```go
func main() {
  
    registry := registration.NewServiceRegistry()

    _ = registration.RegisterTransient(NewA)
    _ = registration.RegisterTransient(NewB)

    registryValidator := registration.NewServiceRegistrationsValidator()
    err := registryValidator.Validate(registry)
    if err != nil {
        panic(err)
    }
}
```

The validator inspects all registered services in the `ServiceRegistry`, ensuring that all dependencies are properly resolved and that there are no circular dependencies.

## When should the validator be used?

It is recommended that service registrations be validated after the registration of service types is complete. This is also the case after registering services to a **linked** or **scoped** service registry during runtime.

### Why This Matters

- **Prevent runtime errors:** Detect missing or invalid service registrations early to prevent application crashes.

- **Improve developer experience:** Get clear, actionable error messages, helping you resolve issues quickly.

- **Avoid circular dependency traps:** Protect your application from hard-to-diagnose circular dependencies that can lead to infinite recursion and stack overflow errors. For instance, you can organize application dependencies as [modules](../registration/register-module.md) and verify the modules using unit tests. Use the validator at runtime if dependencies get dynamically registered depending on configuration.