# Registry Hierarchy

Parsley supports creating hierarchical relationships between service registries. This allows you to manage service registration inheritance and isolate services within specific parts of your application.

## Linked Registries

A linked registry is a new, empty service registry that shares the same internal state (such as service ID generation) as an existing registry. By itself, a linked registry does not inherit registrations, making it ideal for creating isolated registration sets that can later be combined.

To create a linked registry, use the `CreateLinkedRegistry` method:

```golang
func (s *serviceRegistry) CreateLinkedRegistry() types.ServiceRegistry
```

A common pattern is to use a linked registry in combination with [Multi-Registry Aggregation](multi-registry-aggregation.md) to provide fallback resolution logic.

## Scopes

A scope is a specialized service registry that inherits all current registrations from its parent registry at the time of creation. Scopes are primarily used for managing service lifetimes, especially the **Scoped** lifetime, where an instance is shared within a scope but not across different scopes.

To create a new scope, use the `CreateScope` method:

```golang
func (s *serviceRegistry) CreateScope() types.ServiceRegistry
```

### Inheritance and Isolation

* **Inheritance:** Scopes inherit all existing registrations from their parent. Linked registries start empty but share the same service identity context.
* **Isolation:** Registries created via `CreateLinkedRegistry` or `CreateScope` are isolated from each other. Registrations added to a child registry are not visible to the parent or to other sibling registries.

## Example

The following example demonstrates how to create a scope and use it to manage isolated service registrations.

```golang
package main

import (
    "github.com/matzefriedrich/parsley/pkg/registration"
)

func main() {
    parentRegistry := registration.NewServiceRegistry()
    _ = registration.RegisterSingleton(parentRegistry, NewGlobalService)

    // Create a new scope that inherits GlobalService
    scope := parentRegistry.CreateScope()
    
    // Register a service only within this scope
    _ = registration.RegisterScoped(scope, NewScopedService)
    
    // The scope can resolve both GlobalService and ScopedService.
    // The parentRegistry can only resolve GlobalService.
}
```

## Use Cases

* **Request Isolation:** In web applications, create a new scope for every incoming HTTP request to manage request-specific services.
* **Modular Configuration:** Use linked registries to isolate registrations for different modules, preventing naming conflicts or unintended service overrides when used with aggregation.
* **Testing:** Create a scope or linked registry to override certain services with mocks or stubs for a specific test case without affecting the global registry.
