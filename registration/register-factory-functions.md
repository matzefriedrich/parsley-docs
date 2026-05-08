# Register Factory Functions

In Parsley, you can register a factory function to provide more control and customization when creating service instances. A factory function differs from a constructor function because it allows you to pass additional parameters to configure the service creation process.

Instead of directly registering a constructor function that returns a service instance, you register a factory function that returns another function. This returned function is then used to create the actual service instances.

Here's a detailed look at how to register and use a factory function:

### Example

Suppose you want to create a `Greeter` service where the salutation can be customized. You can use a factory function to achieve this. Here's how you would define and register a factory function:

:::code language="golang" source="/examples/registration-concepts/internal/greeter.go" :::

### Supported Signatures

The function returned by the factory function supports various signatures, similar to constructor functions:

* `func(...) T`
* `func(context.Context, ...) T`
* `func(...) (T, error)`
* `func(context.Context, ...) (T, error)`

If the signature includes a `context.Context`, it must be the first parameter. Returning an `error` allows Parsley to handle and propagate errors during service resolution.

### Example Usage

:::code language="golang" source="/examples/registration-concepts/cmd/factory-functions/main.go" range="12-21" :::
