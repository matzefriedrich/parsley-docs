# Multi-Registry Aggregation

In some scenarios, you may need to combine multiple service registries to resolve dependencies. Parsley provides the `NewMultiRegistryAccessor` function to aggregate several registries into a single accessor.

## Combining Registries

The `NewMultiRegistryAccessor` function takes one or more objects that implement the `types.ServiceRegistryAccessor` interface and returns a new accessor that can query all of them.

```golang
func NewMultiRegistryAccessor(registries ...types.ServiceRegistryAccessor) types.ServiceRegistryAccessor
```

When resolving a service through a multi-registry accessor, Parsley searches the registries in the order they were provided. The first registry that contains a registration for the requested service type is used for resolution.

## Example

The following example shows how to combine two separate registries:

```golang
package main

import (
    "github.com/matzefriedrich/parsley/pkg/registration"
    "github.com/matzefriedrich/parsley/pkg/resolving"
    "github.com/matzefriedrich/parsley/pkg/types"
)

func main() {
    registry1 := registration.NewServiceRegistry()
    _ = registration.RegisterInstance(registry1, "Service from Registry 1")

    registry2 := registration.NewServiceRegistry()
    _ = registration.RegisterInstance(registry2, 42)

    // Aggregate registries
    multiRegistry := registration.NewMultiRegistryAccessor(registry1, registry2)

    // Use the multi-registry accessor with a resolver
    resolver := resolving.NewResolver(multiRegistry)
    
    // Resolve services from both registries
    // ...
}
```

## Use Cases

Multi-registry aggregation is useful in modular applications where different components maintain their own service registries, but a central resolver needs to access all of them. It also facilitates scenarios where you want to override certain service registrations by placing a specialized registry before a general one in the aggregation list.
