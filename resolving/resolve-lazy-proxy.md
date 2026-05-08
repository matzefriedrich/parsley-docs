# Lazy Proxies

Parsley supports lazy proxies, a powerful feature that allows for the delayed activation of services. A lazy proxy acts as a placeholder for a dependency, deferring its creation until it is needed. This is particularly useful for services that are resource-intensive to create or may not always be required immediately.

Once the service is activated, the lazy proxy retains the instance, ensuring that subsequent calls to the `Value()` method return the same instance. This balances performance and resource management, particularly in complex applications.

### Support for Dependencies

The `RegisterLazy[T]` method supports activator functions with dependencies. You can pass a constructor function that requires other services as the second argument, and Parsley will automatically resolve those dependencies when the lazy service is first accessed.

## Example

In this example, we register a `Greeter` service using a lazy proxy via the `RegisterLazy[T]` method. The `NewGreeterFactory` function is passed as the factory for creating `Greeter` instances, and the `LifetimeTransient` scope is used, meaning a new instance would typically be created each time.

:::code language="golang" source="/examples/resolving-services/cmd/resolve-lazy-proxy/main.go" :::

However, the `Greeter` instance is not created immediately when using a lazy proxy. Instead, a `Lazy[Greeter]` proxy is resolved, and the actual `Greeter` instance is only created when `lazy.Value(ctx)` is called for the first time. This instance is then cached within the proxy, ensuring the same `Greeter` object is returned on subsequent calls to `Value()`.

## Benefits and Use Cases

Lazy proxies are ideal for optimizing the performance of applications where certain dependencies are resource-intensive to create but may not always be needed immediately. By deferring the creation of such dependencies until they are required, you can improve startup times and reduce unnecessary resource consumption.

This feature is especially useful in scenarios involving complex object graphs, optional dependencies, or services that are conditionally used based on runtime factors.
