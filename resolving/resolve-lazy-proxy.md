---
label: Lazy Proxies
tags: [ registration, lazy proxy, resolve behavior, resolving, performance ]
---
# Lazy Proxies

Parsley supports lazy proxies, a powerful feature that allows for the delayed activation of services. A lazy proxy acts as a placeholder for a dependency, deferring its creation until it’s needed. This is particularly useful for services that are expensive to create or may not always be required immediately.

Once the service is activated, the lazy proxy retains the instance, ensuring that subsequent calls to the Value() method return the same instance. This balances performance and resource management, particularly in complex applications.

## Example

In this example, we register a `Greeter` service using a lazy proxy via the `RegisterLazy[T]` method. The `NewGreeterFactory` function is passed as the factory for creating `Greeter` instances, and the `LifetimeTransient` scope is used, meaning a new instance would typically be created each time.

:::code language="golang" source="/examples/resolving-services/cmd/resolve-lazy-proxy/main.go" :::

However, the `Greeter` instance is not created immediately using a lazy proxy. Instead, a `Lazy[Greeter]` proxy is resolved, and the actual `Greeter` instance is only created when `lazy.Value()` is called for the first time. This instance is then cached within the proxy, ensuring that the same `Greeter` object is returned on subsequent calls to `Value()`.

## Benefits and use cases

Lazy proxies are ideal for optimizing the performance of applications where certain dependencies are resource-intensive to create but may not always be needed immediately. By deferring the creation of such dependencies until they are required, you can improve startup times and reduce unnecessary resource consumption.

This feature is especially useful in scenarios involving complex object graphs, optional dependencies, or services that are conditionally used based on runtime factors.