---
meta:
  title: Parsley - Registering Named Services to Manage Multiple Implementations
description: Learn how to use named service registration in Parsley for precise control over multiple service implementations. Discover how to associate unique names with services, configure different contexts, and resolve them dynamically to create a more flexible and adaptable application architecture.
icon: file
label: Register Named Services
tags: [ registration, named services, service factory ]
category:
  - Registration
  - Multiple Service Registrations
---
# Register Named Services

In Parsely, you can register services with specific names to support more granular control and flexible resolution of service instances. Named services are beneficial when you need to manage multiple implementations of the same interface or provide different configurations for different contexts.

When registering named services, you associate a unique name with each implementation. This allows you to resolve and retrieve the correct service based on its name, enhancing your application's flexibility and configurability.

Suppose you have an interface `DataService` with two implementations: `remoteDataService` and `localDataService`.

:::code language="golang" source="/examples/registration-concepts/internal/data_services.go" :::

## Register named services

Register each service with a unique name using `RegisterNamed`. This allows you to specify different names for each implementation and control their lifecycle.

:::code language="golang" source="/examples/registration-concepts/cmd/named-services/main.go" range="16-20" :::

In this example, `RegisterNamed` is used to register the constructor functions `NewRemoteDataService` and `NewLocalDataService` with the names `remote` and `local`, respectively. The `LifetimeTransient` parameter indicates that each service instance is transient and will be created anew each time it is resolved.

> **Note** Implementation types with names are automatically registered as an unnamed service with the desired interface type and lifetime scope. So, named services can be resolved (separately) by name or as a list of service instances per interface type.

## Resolve and use named services

To resolve a named service, you use a factory function that takes the name of the service (a `string` parameter) and returns an instance of the service along with any errors. This is done through the `ResolveRequiredService` function:

:::code language="golang" source="/examples/registration-concepts/cmd/named-services/main.go" range="25-32" :::

> **Note** You can combine the concept of factory functions and named services, but you must ensure that the signature of the factory function is the same for all named services.