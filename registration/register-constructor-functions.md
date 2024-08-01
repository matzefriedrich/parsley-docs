---
label: Register Constructor Functions
tags: [ registration, constructor function ]
---
# Register Constructor Functions

In Parsely, the fundamental way to register services is by providing a constructor function to the service registry. This method enables you to create instances of your services in a controlled manner and inject them into your application as needed.

Here’s a basic example of registering a service using a constructor function:

:::code language="golang" source="/examples/quick-start/cmd/main.go" range="13-14" :::

In this example, the `NewGreeter` function is used to create instances of the `Greeter` interface. This function returns an implementation of `Greeter` while hiding the concrete type `*greeter`.

The `Greeter` interface and its implementation are defined as follows:

:::code language="golang" source="/examples/quick-start/internal/greeter.go" :::

While it’s possible for `NewGreeter` to return a *greeter directly, it is generally advisable to return the interface type `Greeter`. This approach maintains flexibility and abstraction, making it easier to change implementations without modifying dependent code.